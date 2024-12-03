// package main

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"
// 	"os"
// 	"time"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/joho/godotenv"
// )

// func CheckSwapContract(client *ethclient.Client, addresses []common.Address, auth *bind.TransactOpts) {

// 	// Contract addresses from deployment
// 	swapAddress := addresses[2]
// 	baseTokenAddress := addresses[0]
// 	quoteTokenAddress := addresses[1]

// 	// Wait for contract deployment to be confirmed
// 	fmt.Println("Waiting for contract deployment confirmation...")
// 	time.Sleep(2 * time.Second)

// 	// Verify contract code exists
// 	code, err := client.CodeAt(context.Background(), swapAddress, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to get contract code: %v", err)
// 	}
// 	if len(code) == 0 {
// 		log.Fatal("No contract code found at swap address")
// 	}

// 	// Get contract ABIs and bytecode
// 	parsedSwapABI, _ := getAbidataBytecode("./build/MiniSwap.abi", "./build/MiniSwap.bin")
// 	// parsedBaseTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")
// 	// parsedQuoteTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")

// 	// Create contract instances
// 	// baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
// 	// quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)
// 	swapInstance := bind.NewBoundContract(swapAddress, parsedSwapABI, client, client, client)

// 	// Get token1 and token2 addresses from swap contract
// 	var token1Address common.Address
// 	var token2Address common.Address

// 	var result []interface{}
// 	err = swapInstance.Call(&bind.CallOpts{}, &result, "token1")
// 	if err != nil {
// 		log.Fatalf("Failed to get token1 address: %v", err)
// 	}
// 	token1Address = result[0].(common.Address)

// 	result = nil
// 	err = swapInstance.Call(&bind.CallOpts{}, &result, "token2")
// 	if err != nil {
// 		log.Fatalf("Failed to get token2 address: %v", err)
// 	}
// 	token2Address = result[0].(common.Address)

// 	fmt.Println("\n-----> MiniSwap Contract Details <-----")
// 	fmt.Printf("Contract Address: %s\n", swapAddress.Hex())
// 	fmt.Printf("\nToken1 Address: %s\n", token1Address.Hex())
// 	fmt.Printf("Token2 Address: %s\n", token2Address.Hex())
// 	fmt.Printf("WETH Address: %s\n", baseTokenAddress.Hex())
// 	fmt.Printf("USDT Address: %s\n", quoteTokenAddress.Hex())

// 	// Verify addresses match
// 	if token1Address == baseTokenAddress {
// 		fmt.Println("\n✅ Token1 matches WETH address")
// 	} else {
// 		fmt.Println("\n❌ Token1 does not match WETH address")
// 	}

// 	if token2Address == quoteTokenAddress {
// 		fmt.Println("✅ Token2 matches USDT address")
// 	} else {
// 		fmt.Println("❌ Token2 does not match USDT address")
// 	}
// }

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

func setUserWallet() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get addresses and keys from env
	userKey := common.HexToAddress(os.Getenv("FAKEUSERPUBLICKEY"))
	realUserPrivateKey := os.Getenv("REALUSERPRIVATEKEY")
	BTA := os.Getenv("BASECONTRACTADDRESS")
	QTA := os.Getenv("QUOTECONTRACTADDRESS")

	// Connect to client
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	handleError(err, "Failed to connect to the Ethereum client")

	// Setup the private key
	privateKey, err := crypto.HexToECDSA(realUserPrivateKey)
	handleError(err, "Failed to load private key")

	// Get the public address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("\nAddresses:\n")
	fmt.Printf("From (real user): %s\n", fromAddress.Hex())
	fmt.Printf("To (fake user): %s\n", userKey.Hex())

	// Check ETH balances before minting
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	handleError(err, "Failed to get balance")
	fmt.Printf("\nBalances:\n")
	fmt.Printf("Real user ETH balance: %v wei (%v ETH)\n",
		balance,
		new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt64(1e18)))

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	handleError(err, "Failed to get gas price")
	fmt.Printf("Suggested gas price: %v wei\n", gasPrice)

	// Create auth
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	handleError(err, "Failed to create transactor")

	// Setup auth parameters with lower gas values
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	handleError(err, "Failed to get nonce")

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(100000)         // Lower gas limit
	auth.GasPrice = big.NewInt(1000000000) // Set to 1 gwei
	auth.From = fromAddress

	fmt.Printf("\nTransaction settings:\n")
	fmt.Printf("Gas limit: %v\n", auth.GasLimit)
	fmt.Printf("Gas price: %v wei\n", auth.GasPrice)
	fmt.Printf("Max transaction cost: %v wei\n", new(big.Int).Mul(big.NewInt(int64(auth.GasLimit)), auth.GasPrice))

	// Get contract instances
	parsedBaseTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")
	parsedQuoteTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")
	baseTokenAddress := common.HexToAddress(BTA)
	quoteTokenAddress := common.HexToAddress(QTA)
	baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
	quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)

	fmt.Printf("\nStarting token minting...\n")

	// Mint USDT (18 decimals)
	fmt.Println("Minting USDT...")
	usdtAmount := new(big.Int).Mul(big.NewInt(650), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	tx, err := quoteTokenInstance.Transact(auth, "mint", userKey, usdtAmount)
	handleError(err, "Failed to mint USDT")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for USDT mint")
	if receipt.Status == 0 {
		log.Fatal("USDT mint transaction failed")
	}
	fmt.Println("✅ USDT minted")

	// Update nonce for next transaction
	auth.Nonce = big.NewInt(int64(nonce + 1))

	// Mint WETH (18 decimals)
	fmt.Println("Minting WETH...")
	wethAmount := new(big.Int).Mul(big.NewInt(92), new(big.Int).Exp(big.NewInt(10), big.NewInt(17), nil))
	tx, err = baseTokenInstance.Transact(auth, "mint", userKey, wethAmount)
	handleError(err, "Failed to mint WETH")
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	handleError(err, "Failed waiting for WETH mint")
	if receipt.Status == 0 {
		log.Fatal("WETH mint transaction failed")
	}
	fmt.Println("✅ WETH minted")

	fmt.Printf("\nTokens minted successfully!\n")
	fmt.Printf("Amounts minted to %s:\n", userKey.Hex())
	fmt.Printf("- WETH: 9.2 (18 decimals)\n")
	fmt.Printf("- USDT: 650 (18 decimals)\n")

	// Print final balances
	var wethBalance []interface{}
	var usdtBalance []interface{}

	err = baseTokenInstance.Call(&bind.CallOpts{}, &wethBalance, "balanceOf", userKey)
	handleError(err, "Failed to get WETH balance")

	err = quoteTokenInstance.Call(&bind.CallOpts{}, &usdtBalance, "balanceOf", userKey)
	handleError(err, "Failed to get USDT balance")

	fmt.Printf("\nFinal Token Balances:\n")
	fmt.Printf("WETH: %v\n", new(big.Float).Quo(new(big.Float).SetInt(wethBalance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
	fmt.Printf("USDT: %v\n", new(big.Float).Quo(new(big.Float).SetInt(usdtBalance[0].(*big.Int)), new(big.Float).SetInt64(1e18)))
}

// func main() {
// 	setUserWallet()
// }
