package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// function to get users orders based on the address
func Swap(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := r.URL.Query()
	amount := params.Get("amount")
	token := params.Get("token")

	token1, token2, err := swapcall(amount, token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, err)
		return
	}
	response := map[string]interface{}{
		"token1": token1,
		"token2": token2,
	}

	writeJSON(w, response, nil)

}

func swapcall(amount string, token string) (string, string, error) {

	// Connect to client
	client, err := ethclient.Dial(GlobalEthClient)
	handleError(err, "Failed to connect to the Ethereum client")

	// Get contract instances
	parsedSwapABI, _ := getAbidataBytecode("../../build/MiniSwap.abi", "../../build/MiniSwap.bin")
	parsedBaseTokenABI, _ := getAbidataBytecode("../../build/fakeToken.abi", "../../build/fakeToken.bin")
	parsedQuoteTokenABI, _ := getAbidataBytecode("../../build/fakeToken.abi", "../../build/fakeToken.bin")

	swapAddress := common.HexToAddress(swapAddress)
	baseTokenAddress := common.HexToAddress(baseTokenAddress)
	quoteTokenAddress := common.HexToAddress(quoteTokenAddress)
	userKey := common.HexToAddress(userKey)

	swapInstance := bind.NewBoundContract(swapAddress, parsedSwapABI, client, client, client)
	baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
	quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)

	// Setup the fake user's private key for transactions
	fakeUserPrivateKey := os.Getenv("FAKEUSERPRIVATEKEY")
	privateKey, err := crypto.HexToECDSA(fakeUserPrivateKey)
	handleError(err, "Failed to load private key")

	// Create auth
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	handleError(err, "Failed to create transactor")

	// Get initial nonce
	nonce, err := client.PendingNonceAt(context.Background(), userKey)
	handleError(err, "Failed to get nonce")

	// Setup auth parameters
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(1000000000)
	auth.From = userKey

	var token1 string
	var token2 string

	bToken1, bToken2 := getBalances(baseTokenInstance, quoteTokenInstance, userKey, "Initial")

	fmt.Println("WWWWWW", bToken1, bToken2)
	rawAmount := new(big.Int)
	_, ok := rawAmount.SetString(amount, 10)
	if !ok {
		return "", "", errors.New("failed to parse amount")
	}

	// Amount to swap (10 tokens)
	amountToSwap := new(big.Int).Mul(rawAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	if token == "base" {

		rawToken1 := new(big.Float)
		_, ok = rawToken1.SetString(bToken1)
		if !ok {
			return "0", "0", errors.New("failed to parse amount")
		}
		rawAmountFloat := new(big.Float).SetInt(rawAmount)
		if rawToken1.Cmp(rawAmountFloat) == -1 {
			return "0", "0", errors.New("insufficient balance")
		}

		fmt.Printf("Approving %v Token1 for swap...\n", new(big.Float).Quo(new(big.Float).SetInt(amountToSwap), new(big.Float).SetInt64(1e18)))
		tx, err := baseTokenInstance.Transact(auth, "approve", swapAddress, amountToSwap)
		handleError(err, "Failed to approve Token1")
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		handleError(err, "Failed waiting for Token1 approval")
		if receipt.Status == 0 {
			log.Fatal("Token1 approval failed")
		}
		fmt.Println("✅ Token1 approved")

		auth.Nonce = big.NewInt(int64(nonce + 1))

		tx, err = swapInstance.Transact(auth, "swap", baseTokenAddress, amountToSwap)
		handleError(err, "Failed to swap Token1")
		receipt, err = bind.WaitMined(context.Background(), client, tx)
		handleError(err, "Failed waiting for Token1 swap")
		if receipt.Status == 0 {
			log.Fatal("Token1 swap failed")
		}
		fmt.Println("swap completed")

		token1, token2 = getBalances(baseTokenInstance, quoteTokenInstance, userKey, "Final")

	} else if token == "quote" {

		rawToken2 := new(big.Float)
		_, ok = rawToken2.SetString(bToken2)
		if !ok {
			return "0", "0", errors.New("failed to parse amount")
		}
		rawAmountFloat := new(big.Float).SetInt(rawAmount)
		if rawToken2.Cmp(rawAmountFloat) == -1 {
			return "0", "0", errors.New("insufficient balance")
		}

		fmt.Printf("Approving %v Token2 for swap...\n", new(big.Float).Quo(new(big.Float).SetInt(amountToSwap), new(big.Float).SetInt64(1e18)))
		tx, err := quoteTokenInstance.Transact(auth, "approve", swapAddress, amountToSwap)
		handleError(err, "Failed to approve Token2")
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		handleError(err, "Failed waiting for Token2 approval")
		if receipt.Status == 0 {
			log.Fatal("Token2 approval failed")
		}
		fmt.Println("✅ Token2 approved")

		auth.Nonce = big.NewInt(int64(nonce + 1))

		tx, err = swapInstance.Transact(auth, "swap", quoteTokenAddress, amountToSwap)
		handleError(err, "Failed to swap Token2")
		receipt, err = bind.WaitMined(context.Background(), client, tx)
		handleError(err, "Failed waiting for Token2 swap")
		if receipt.Status == 0 {
			log.Fatal("Token2 swap failed")
		}
		fmt.Println("swap completed")

		token1, token2 = getBalances(baseTokenInstance, quoteTokenInstance, userKey, "Final")

	}

	return token1, token2, nil
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	token1, token2, err := Balances()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		writeJSON(w, nil, errors.New("record not found"))
		return
	}

	response := map[string]interface{}{
		"token1": token1,
		"amount": token2,
	}

	writeJSON(w, response, nil)

}

func Balances() (string, string, error) {

	// Connect to client
	client, err := ethclient.Dial(GlobalEthClient)
	handleError(err, "Failed to connect to the Ethereum client")

	parsedBaseTokenABI, _ := getAbidataBytecode("../../build/fakeToken.abi", "../../build/fakeToken.bin")
	parsedQuoteTokenABI, _ := getAbidataBytecode("../../build/fakeToken.abi", "../../build/fakeToken.bin")

	baseTokenAddress := common.HexToAddress(baseTokenAddress)
	quoteTokenAddress := common.HexToAddress(quoteTokenAddress)
	userKey := common.HexToAddress(userKey)

	baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
	quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)

	token1, token2 := getBalances(baseTokenInstance, quoteTokenInstance, userKey, "Final")

	return token1, token2, nil
}
