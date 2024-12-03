package main

import (
	"context"
	"crypto/ecdsa"
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

func addLiquidity() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get addresses from env
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

	// Setup the real user's private key (for minting and adding liquidity)
	realUserPrivateKey := os.Getenv("REALUSERPRIVATEKEY")
	privateKey, err := crypto.HexToECDSA(realUserPrivateKey)
	handleError(err, "Failed to load private key")

	// Get the real user's address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Create auth
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	handleError(err, "Failed to create transactor")

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	handleError(err, "Failed to get nonce")

	// Setup auth parameters
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(1000000000)
	auth.From = fromAddress

	fmt.Printf("\nAdding liquidity to swap contract...\n")

	// Amount of liquidity to add (1000 of each token)
	liquidityAmount := new(big.Int).Mul(big.NewInt(1000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// First add WETH liquidity
	fmt.Println("Adding WETH liquidity...")

	// Mint WETH
	tx, err := baseTokenInstance.Transact(auth, "mint", fromAddress, liquidityAmount)
	handleError(err, "Failed to mint WETH for liquidity")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for WETH mint")
	if receipt.Status == 0 {
		log.Fatal("WETH mint failed")
	}

	// Approve WETH
	auth.Nonce = big.NewInt(int64(nonce + 1))
	tx, err = baseTokenInstance.Transact(auth, "approve", swapAddress, liquidityAmount)
	handleError(err, "Failed to approve WETH")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for WETH approval")
	if receipt.Status == 0 {
		log.Fatal("WETH approval failed")
	}

	// Add WETH liquidity
	auth.Nonce = big.NewInt(int64(nonce + 2))
	tx, err = swapInstance.Transact(auth, "addLiquidity", baseTokenAddress, liquidityAmount)
	handleError(err, "Failed to add WETH liquidity")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for WETH liquidity addition")
	if receipt.Status == 0 {
		log.Fatal("WETH liquidity addition failed")
	}
	fmt.Println("✅ WETH liquidity added")

	// Then add USDT liquidity
	fmt.Println("\nAdding USDT liquidity...")

	// Mint USDT
	auth.Nonce = big.NewInt(int64(nonce + 3))
	tx, err = quoteTokenInstance.Transact(auth, "mint", fromAddress, liquidityAmount)
	handleError(err, "Failed to mint USDT for liquidity")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for USDT mint")
	if receipt.Status == 0 {
		log.Fatal("USDT mint failed")
	}

	// Approve USDT
	auth.Nonce = big.NewInt(int64(nonce + 4))
	tx, err = quoteTokenInstance.Transact(auth, "approve", swapAddress, liquidityAmount)
	handleError(err, "Failed to approve USDT")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for USDT approval")
	if receipt.Status == 0 {
		log.Fatal("USDT approval failed")
	}

	// Add USDT liquidity
	auth.Nonce = big.NewInt(int64(nonce + 5))
	tx, err = swapInstance.Transact(auth, "addLiquidity", quoteTokenAddress, liquidityAmount)
	handleError(err, "Failed to add USDT liquidity")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for USDT liquidity addition")
	if receipt.Status == 0 {
		log.Fatal("USDT liquidity addition failed")
	}
	fmt.Println("✅ USDT liquidity added")

	fmt.Printf("\nLiquidity added successfully!\n")
	fmt.Printf("Added %v WETH and %v USDT to the swap contract\n",
		new(big.Float).Quo(new(big.Float).SetInt(liquidityAmount), new(big.Float).SetInt64(1e18)),
		new(big.Float).Quo(new(big.Float).SetInt(liquidityAmount), new(big.Float).SetInt64(1e18)))
}
