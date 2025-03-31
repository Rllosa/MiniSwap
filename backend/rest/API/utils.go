package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	userKey           string
	swapAddress       string
	baseTokenAddress  string
	quoteTokenAddress string
	GlobalEthClient   string
)

func init() {
	// Load environment variables from the .env file (optional)
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get addresses from env
	userKey = os.Getenv("FAKEUSERPUBLICKEY")
	swapAddress = os.Getenv("SWAPCONTRACTADDRESS")
	baseTokenAddress = os.Getenv("BASECONTRACTADDRESS")
	quoteTokenAddress = os.Getenv("QUOTECONTRACTADDRESS")

	GlobalEthClient = os.Getenv("ETHCLIENT")
	if GlobalEthClient == "" {
		log.Fatal("What env?")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type errResp struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func writeJSON(w http.ResponseWriter, v interface{}, err error) {
	var respVal interface{}

	// Set the Content-Type header for JSON response
	w.Header().Set("Content-Type", "application/json")

	// Check for an error and set the appropriate status code and response body
	if err != nil {
		// Only call WriteHeader once before writing the response
		w.WriteHeader(http.StatusBadRequest)

		var e errResp
		e.Error.Message = err.Error()
		respVal = e
	} else {
		respVal = v
		// Status 200 OK is implicit, no need to call WriteHeader here
	}

	// Encode the response value to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(respVal); err != nil {
		// Handle encoding errors by logging and setting a 500 status code
		log.Printf("json.NewEncoder.Encode: %v", err)
		w.WriteHeader(http.StatusInternalServerError)                      // Send 500 error if JSON encoding fails
		w.Write([]byte(`{"error": {"message": "Internal Server Error"}}`)) // Manually write the error message
		return
	}

	// Write the encoded response to the client
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// Generic utility function to format values with a specified number of extra decimals
func formatValue(value string, extraDecimals int) string {
	valueInt := new(big.Int)
	valueInt.SetString(value, 10)

	divisor := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(extraDecimals)), nil))
	formattedValue := new(big.Float).SetInt(valueInt)
	formattedValue.Quo(formattedValue, divisor)

	return formattedValue.Text('f', extraDecimals)
}

// Utility function to calculate amount_eth as formattedAmount / formattedPrice
func calculateAmountEth(formattedAmount string, formattedPrice string) string {
	amount := new(big.Float)
	price := new(big.Float)

	amount.SetString(formattedAmount)
	price.SetString(formattedPrice)

	amountEth := new(big.Float).Quo(amount, price)

	return amountEth.Text('f', 18) // Returning amount_eth with 18 decimal places
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func getAbidataBytecode(abipath string, bytecodePath string) (abi.ABI, []byte) {

	abiData, err := os.ReadFile(abipath)
	handleError(err, "Failed to read contract ABI")

	bytecode, err := os.ReadFile(bytecodePath)
	handleError(err, "Failed to read contract bytecode")

	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	handleError(err, "Failed to parse contract ABI")

	return parsedABI, bytecode
}

func waitMined(client *ethclient.Client, tx *types.Transaction) *types.Receipt {
	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed to wait for transaction to be mined")

	// Check if the transaction failed
	if receipt.Status != 1 {
		fmt.Println("Transaction failed. Analyzing failure reasons...")

		// Inspect the receipt details
		fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("Block Number: %v\n", receipt.BlockNumber)
		fmt.Printf("Gas Used: %v\n", receipt.GasUsed)
		fmt.Printf("Cumulative Gas Used: %v\n", receipt.CumulativeGasUsed)

		// Check if the gas used is close to the gas limit, indicating a possible out-of-gas error
		if receipt.GasUsed >= tx.Gas() {
			fmt.Println("Possible cause: Transaction ran out of gas.")
		}

		// Decode any revert reason from the transaction receipt logs if available
		reason, err := getRevertReason(client, tx, receipt)
		if err != nil {
			fmt.Printf("Failed to fetch revert reason: %v\n", err)
		} else {
			fmt.Printf("Revert Reason: %s\n", reason)
		}

		// Display any logs emitted during the transaction
		for i, logEntry := range receipt.Logs {
			fmt.Printf("Log %d: %v\n", i+1, logEntry)
		}

		log.Fatalf("Transaction failed: %s", tx.Hash().Hex())
	}

	return receipt
}

func getRevertReason(client *ethclient.Client, tx *types.Transaction, receipt *types.Receipt) (string, error) {
	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
	}

	// Simulate the transaction to get revert reason
	output, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
	if err != nil {
		return "", fmt.Errorf("error calling contract: %v", err)
	}

	// Revert reasons follow a standard format: Error(string)
	if len(output) < 4 || string(output[:4]) != "\x08\xc3y\xa0" {
		return "Unknown or no revert reason", nil
	}

	// Decode the revert reason
	reason, err := abi.UnpackRevert(output[4:])
	if err != nil {
		return "Unable to decode revert reason", err
	}

	return reason, nil
}

func getBalances(token1Instance, token2Instance *bind.BoundContract, userKey common.Address, label string) (string, string) {
	var token1Balance []interface{}
	var token2Balance []interface{}

	err := token1Instance.Call(&bind.CallOpts{}, &token1Balance, "balanceOf", userKey)
	handleError(err, "Failed to get Token1 balance")

	err = token2Instance.Call(&bind.CallOpts{}, &token2Balance, "balanceOf", userKey)
	handleError(err, "Failed to get Token2 balance")

	fmt.Printf("\n%s Balances:\n", label)
	token1 := new(big.Float).Quo(new(big.Float).SetInt(token1Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18))
	token2 := new(big.Float).Quo(new(big.Float).SetInt(token2Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18))

	fmt.Printf("Token1: %v\n", new(big.Float).Quo(new(big.Float).SetInt(token1Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
	fmt.Printf("Token2: %v\n", new(big.Float).Quo(new(big.Float).SetInt(token2Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))

	token1Str := strings.TrimRight(strings.TrimRight(token1.Text('f', 18), "0"), ".")
	token2Str := strings.TrimRight(strings.TrimRight(token2.Text('f', 18), "0"), ".")
	return token1Str, token2Str
}
