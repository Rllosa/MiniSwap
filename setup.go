package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
)

type variables struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	nonce      uint64
	chainID    *big.Int
}

func setup(input string) *variables {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	handleError(err, "Failed to connect to the Ethereum client")

	privateKey, err := crypto.HexToECDSA(input)
	handleError(err, "Failed to load private key")

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	handleError(err, "Failed to get nonce")

	chainID := big.NewInt(1337)

	return &variables{
		client:     client,
		privateKey: privateKey,
		nonce:      nonce,
		chainID:    chainID,
	}

}

func createTransactor(v *variables) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(v.privateKey, v.chainID)
	handleError(err, "Failed to create transactor")

	gasPrice, err := v.client.SuggestGasPrice(context.Background())
	handleError(err, "Failed to suggest gas price")

	auth.Nonce = big.NewInt(int64(v.nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(80000000)
	auth.GasPrice = gasPrice

	return auth
}
