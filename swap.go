package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func performSwap() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get addresses from env
	userKey := common.HexToAddress(os.Getenv("FAKEUSERPUBLICKEY"))
	swapAddress := common.HexToAddress(os.Getenv("SWAPCONTRACTADDRESS"))
	baseTokenAddress := common.HexToAddress(os.Getenv("BASECONTRACTADDRESS"))
	quoteTokenAddress := common.HexToAddress(os.Getenv("QUOTECONTRACTADDRESS"))

	// Connect to client
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	handleError(err, "Failed to connect to the Ethereum client")

	// Get contract instances
	parsedSwapABI, _ := getAbidataBytecode("./build/MiniSwap.abi", "./build/MiniSwap.bin")
	parsedBaseTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")
	parsedQuoteTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")

	swapInstance := bind.NewBoundContract(swapAddress, parsedSwapABI, client, client, client)
	baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
	quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)

	// Print initial balances
	var wethBalance []interface{}
	var usdtBalance []interface{}

	err = baseTokenInstance.Call(&bind.CallOpts{}, &wethBalance, "balanceOf", userKey)
	handleError(err, "Failed to get WETH balance")

	err = quoteTokenInstance.Call(&bind.CallOpts{}, &usdtBalance, "balanceOf", userKey)
	handleError(err, "Failed to get USDT balance")

	fmt.Printf("\nInitial Balances:\n")
	fmt.Printf("WETH: %v\n", new(big.Float).Quo(new(big.Float).SetInt(wethBalance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
	fmt.Printf("USDT: %v\n", new(big.Float).Quo(new(big.Float).SetInt(usdtBalance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))

	// Setup the fake user's private key for transactions
	fakeUserPrivateKey := os.Getenv("FAKEUSERPRIVATEKEY")
	privateKey, err := crypto.HexToECDSA(fakeUserPrivateKey)
	handleError(err, "Failed to load private key")

	// Create auth
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	handleError(err, "Failed to create transactor")

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), userKey)
	handleError(err, "Failed to get nonce")

	// Setup auth parameters
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(1000000000)
	auth.From = userKey

	// Approve tokens for swap
	amountToSwap := new(big.Int).Mul(big.NewInt(1), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)) // 1 WETH
	fmt.Printf("\nApproving %v WETH for swap...\n", new(big.Float).Quo(new(big.Float).SetInt(amountToSwap), new(big.Float).SetInt64(1e18)))

	tx, err := baseTokenInstance.Transact(auth, "approve", swapAddress, amountToSwap)
	handleError(err, "Failed to approve WETH")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for WETH approval")
	if receipt.Status == 0 {
		log.Fatal("WETH approval failed")
	}
	fmt.Println("✅ WETH approved")

	// Update nonce
	auth.Nonce = big.NewInt(int64(nonce + 1))

	// Perform swap
	fmt.Printf("\nPerforming swap...\n")
	tx, err = swapInstance.Transact(auth, "swap", baseTokenAddress, amountToSwap)
	handleError(err, "Failed to swap")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for swap")
	if receipt.Status == 0 {
		log.Fatal("Swap failed")
	}
	fmt.Println("✅ Swap completed")
	fmt.Printf("Swap Event ID: %s\n", getSwapEventID())

	// Print final balances
	var weth2Balance []interface{}
	var usdt2Balance []interface{}

	err = baseTokenInstance.Call(&bind.CallOpts{}, &weth2Balance, "balanceOf", userKey)
	handleError(err, "Failed to get final WETH balance")

	err = quoteTokenInstance.Call(&bind.CallOpts{}, &usdt2Balance, "balanceOf", userKey)
	handleError(err, "Failed to get final USDT balance")

	fmt.Printf("\nFinal Balances:\n")
	fmt.Printf("WETH: %v\n", new(big.Float).Quo(new(big.Float).SetInt(weth2Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
	fmt.Printf("USDT: %v\n", new(big.Float).Quo(new(big.Float).SetInt(usdt2Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
}

func getSwapEventID() string {
	eventSignature := []byte("Swap(address,address,address,uint256,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	return hash.Hex()
}
