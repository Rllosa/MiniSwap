package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ENV string

func main() {

	keys := getKeys()
	_, addresses, _ := deployContracts(keys)
	testkeys(addresses, keys, ENV)

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	CheckSwapContract(client, addresses)
}
