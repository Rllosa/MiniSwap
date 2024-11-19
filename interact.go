package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckSwapContract(client *ethclient.Client, addresses []common.Address) {

	// Contract addresses from deployment
	swapAddress := addresses[2]
	baseTokenAddress := addresses[0]
	quoteTokenAddress := addresses[1]

	// Wait for contract deployment to be confirmed
	fmt.Println("Waiting for contract deployment confirmation...")
	time.Sleep(2 * time.Second)

	// Verify contract code exists
	code, err := client.CodeAt(context.Background(), swapAddress, nil)
	if err != nil {
		log.Fatalf("Failed to get contract code: %v", err)
	}
	if len(code) == 0 {
		log.Fatal("No contract code found at swap address")
	}

	// Get contract ABIs and bytecode
	parsedSwapABI, _ := getAbidataBytecode("./build/MiniSwap.abi", "./build/MiniSwap.bin")
	// parsedBaseTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")
	// parsedQuoteTokenABI, _ := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")

	// Create contract instances
	// baseTokenInstance := bind.NewBoundContract(baseTokenAddress, parsedBaseTokenABI, client, client, client)
	// quoteTokenInstance := bind.NewBoundContract(quoteTokenAddress, parsedQuoteTokenABI, client, client, client)
	swapInstance := bind.NewBoundContract(swapAddress, parsedSwapABI, client, client, client)

	// Get token1 and token2 addresses from swap contract
	var token1Address common.Address
	var token2Address common.Address

	var result []interface{}
	err = swapInstance.Call(&bind.CallOpts{}, &result, "token1")
	if err != nil {
		log.Fatalf("Failed to get token1 address: %v", err)
	}
	token1Address = result[0].(common.Address)

	result = nil
	err = swapInstance.Call(&bind.CallOpts{}, &result, "token2")
	if err != nil {
		log.Fatalf("Failed to get token2 address: %v", err)
	}
	token2Address = result[0].(common.Address)

	fmt.Println("\n-----> MiniSwap Contract Details <-----")
	fmt.Printf("Contract Address: %s\n", swapAddress.Hex())
	fmt.Printf("\nToken1 Address: %s\n", token1Address.Hex())
	fmt.Printf("Token2 Address: %s\n", token2Address.Hex())
	fmt.Printf("WETH Address: %s\n", baseTokenAddress.Hex())
	fmt.Printf("USDT Address: %s\n", quoteTokenAddress.Hex())

	// Verify addresses match
	if token1Address == baseTokenAddress {
		fmt.Println("\n✅ Token1 matches WETH address")
	} else {
		fmt.Println("\n❌ Token1 does not match WETH address")
	}

	if token2Address == quoteTokenAddress {
		fmt.Println("✅ Token2 matches USDT address")
	} else {
		fmt.Println("❌ Token2 does not match USDT address")
	}
}
