package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
