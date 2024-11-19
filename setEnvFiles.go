package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func getKeys() []string {
	file, err := os.Open("ganache-output.txt")
	handleError(err, "Error opening file")
	defer file.Close()

	var accounts []string
	var privateKeys []string
	var keys []string
	scanner := bufio.NewScanner(file)
	readAccounts := false
	readKeys := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Available Accounts") {
			readAccounts = true
			readKeys = false
			continue
		}

		if strings.Contains(line, "Private Keys") {
			readAccounts = false
			readKeys = true
			continue
		}

		if strings.Contains(line, "HD Wallet") {
			break
		}

		if readAccounts && strings.Contains(line, "0x") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				accounts = append(accounts, parts[1])
			}
		}

		if readKeys && strings.Contains(line, "0x") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				privateKeys = append(privateKeys, parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for i, account := range accounts {
		keys = append(keys, account)
		keys = append(keys, privateKeys[i][2:])
	}
	return keys
}

func testkeys(addresses []common.Address, keys []string, env string) {
	file, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	handleError(err, "Error opening file")
	defer file.Close()

	file.WriteString("BASECONTRACTADDRESS=")
	file.WriteString(addresses[0].String())
	file.WriteString("\n")
	file.WriteString("QUOTECONTRACTADDRESS=")
	file.WriteString(addresses[1].String())
	file.WriteString("\n")
	file.WriteString("SWAPCONTRACTADDRESS=")
	file.WriteString(addresses[2].String())
	file.WriteString("\n")
	file.WriteString("FAKEUSERPUBLICKEY=")
	file.WriteString(keys[8])
	file.WriteString("\n")
	file.WriteString("FAKEUSERPRIVATEKEY=")
	file.WriteString(keys[9])
	file.WriteString("\n")
	file.WriteString("REALUSERPUBLICKEY=")
	file.WriteString(keys[10])
	file.WriteString("\n")
	file.WriteString("REALUSERPRIVATEKEY=")
	file.WriteString(keys[11])
	file.WriteString("\n")
	file.WriteString("STARTPRIVATEKEY=")
	file.WriteString(keys[19])
	file.WriteString("\n")

	if env == "PROD" {
		file.WriteString("ETHCLIENT=https://eth-sepolia.g.alchemy.com/v2/S11GIEyZx8jm495NDjFMxJfsLEjF-sWJ")
	} else {
		file.WriteString("ETHCLIENT=http://localhost:8545")
	}
}
