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
	printBalances(baseTokenInstance, quoteTokenInstance, userKey, "Initial")

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

	// Amount to swap (10 tokens)
	amountToSwap := new(big.Int).Mul(big.NewInt(10), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// First Swap: Token1 -> Token2
	fmt.Println("\n=== Performing Token1 -> Token2 Swap ===")

	// Approve Token1
	fmt.Printf("Approving %v Token1 for swap...\n", new(big.Float).Quo(new(big.Float).SetInt(amountToSwap), new(big.Float).SetInt64(1e18)))
	tx, err := baseTokenInstance.Transact(auth, "approve", swapAddress, amountToSwap)
	handleError(err, "Failed to approve Token1")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for Token1 approval")
	if receipt.Status == 0 {
		log.Fatal("Token1 approval failed")
	}
	fmt.Println("✅ Token1 approved")

	// Update nonce
	auth.Nonce = big.NewInt(int64(nonce + 1))

	// Perform first swap
	tx, err = swapInstance.Transact(auth, "swap", baseTokenAddress, amountToSwap)
	handleError(err, "Failed to swap Token1")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for Token1 swap")
	if receipt.Status == 0 {
		log.Fatal("Token1 swap failed")
	}
	fmt.Println("✅ First swap completed")

	// Print intermediate balances
	printBalances(baseTokenInstance, quoteTokenInstance, userKey, "After first swap")

	// Second Swap: Token2 -> Token1
	fmt.Println("\n=== Performing Token2 -> Token1 Swap ===")

	// Update nonce
	auth.Nonce = big.NewInt(int64(nonce + 2))

	// Approve Token2
	fmt.Printf("Approving %v Token2 for swap...\n", new(big.Float).Quo(new(big.Float).SetInt(amountToSwap), new(big.Float).SetInt64(1e18)))
	tx, err = quoteTokenInstance.Transact(auth, "approve", swapAddress, amountToSwap)
	handleError(err, "Failed to approve Token2")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for Token2 approval")
	if receipt.Status == 0 {
		log.Fatal("Token2 approval failed")
	}
	fmt.Println("✅ Token2 approved")

	// Update nonce
	auth.Nonce = big.NewInt(int64(nonce + 3))

	// Perform second swap
	tx, err = swapInstance.Transact(auth, "swap", quoteTokenAddress, amountToSwap)
	handleError(err, "Failed to swap Token2")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for Token2 swap")
	if receipt.Status == 0 {
		log.Fatal("Token2 swap failed")
	}
	fmt.Println("✅ Second swap completed")

	// Print final balances
	printBalances(baseTokenInstance, quoteTokenInstance, userKey, "Final")
}

func printBalances(token1Instance, token2Instance *bind.BoundContract, userKey common.Address, label string) {
	var token1Balance []interface{}
	var token2Balance []interface{}

	err := token1Instance.Call(&bind.CallOpts{}, &token1Balance, "balanceOf", userKey)
	handleError(err, "Failed to get Token1 balance")

	err = token2Instance.Call(&bind.CallOpts{}, &token2Balance, "balanceOf", userKey)
	handleError(err, "Failed to get Token2 balance")

	fmt.Printf("\n%s Balances:\n", label)
	fmt.Printf("Token1: %v\n", new(big.Float).Quo(new(big.Float).SetInt(token1Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
	fmt.Printf("Token2: %v\n", new(big.Float).Quo(new(big.Float).SetInt(token2Balance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
}

func getSwapEventID() string {
	eventSignature := []byte("Swap(address,address,address,uint256,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	return hash.Hex()
}
