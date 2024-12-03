package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func deploySwapContract(auth *bind.TransactOpts, v *variables, baseTokenAddress common.Address, quoteTokenAddress common.Address) common.Address {
	auth.GasLimit = uint64(8000000)

	parsedABI, bytecode := getAbidataBytecode("./build/MiniSwap.abi", "./build/MiniSwap.bin")

	address, tx, _, err := bind.DeployContract(auth, parsedABI, common.FromHex(string(bytecode)), v.client, baseTokenAddress, quoteTokenAddress)
	handleError(err, "Failed to deploy contract")
	waitMined(v.client, tx)

	fmt.Print("\n-----> MiniSwap - Deployment <-----\n\n")

	fmt.Printf("MiniSwap Contract deployed to address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	fmt.Print("\n---------------------------------------\n\n")

	v.nonce++
	return address
}

func deployQuoteTokenContract(auth *bind.TransactOpts, v *variables, governanceTreasuryAddress string) common.Address {
	auth.GasLimit = uint64(8000000)

	parsedABI, bytecode := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")

	name := "USDT"
	symbol := "USDT"
	decimal := uint8(18)
	tokenHolder := common.HexToAddress(governanceTreasuryAddress) // Replace with actual address
	initialValue := new(big.Int).Mul(big.NewInt(30000000000000), big.NewInt(1e6))
	address, tx, _, err := bind.DeployContract(auth, parsedABI, common.FromHex(string(bytecode)), v.client, name, symbol, tokenHolder, decimal, initialValue)
	handleError(err, "Failed to deploy contract")
	waitMined(v.client, tx)

	fmt.Print("\n-----> USDT - Deployment <-----\n\n")

	fmt.Printf("USDT Contract deployed to address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	v.nonce++
	return address
}

func deployBaseTokenContract(auth *bind.TransactOpts, v *variables, governanceTreasuryAddress string) common.Address {
	auth.GasLimit = uint64(8000000)

	parsedABI, bytecode := getAbidataBytecode("./build/fakeToken.abi", "./build/fakeToken.bin")

	name := "WETH"
	symbol := "WETH"
	decimal := uint8(18)
	tokenHolder := common.HexToAddress(governanceTreasuryAddress) // Replace with actual address
	initialValue := new(big.Int).Mul(big.NewInt(100000000000000), big.NewInt(1e18))

	address, tx, _, err := bind.DeployContract(auth, parsedABI, common.FromHex(string(bytecode)), v.client, name, symbol, tokenHolder, decimal, initialValue)
	handleError(err, "Failed to deploy contract")
	waitMined(v.client, tx)

	fmt.Print("\n---------------------------------------\n")

	fmt.Print("\n-----> WETH - Deployment <-----\n\n")

	fmt.Printf("WETH Contract deployed to address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

	v.nonce++

	return address

}

func deploymentSetup(privateKey string) (*bind.TransactOpts, *variables) {
	v := setup(privateKey)
	auth := createTransactor(v)

	return auth, v
}

func deployContracts(keys []string) (*variables, []common.Address, *bind.TransactOpts) {

	auth, v := deploymentSetup(keys[1])
	baseTokenAddress := deployBaseTokenContract(auth, v, keys[18])

	auth, v = deploymentSetup(keys[3])
	quoteTokenAddress := deployQuoteTokenContract(auth, v, keys[18])

	auth, v = deploymentSetup(keys[5])
	swapAddress := deploySwapContract(auth, v, baseTokenAddress, quoteTokenAddress)

	var addresses = []common.Address{baseTokenAddress, quoteTokenAddress, swapAddress}

	return v, addresses, auth
}
