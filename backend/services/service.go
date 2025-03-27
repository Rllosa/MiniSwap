package services

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/Rllosa/miniSwap/backend/mysql"
	"github.com/Rllosa/miniSwap/backend/mysql/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var SwapContractAdd string
var GlobalClient string

func init() {
	// Load environment variables from the .env file (optional)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load the environment variable
	SwapContractAdd = os.Getenv("SWAPCONTRACTADDRESS")
	GlobalClient = os.Getenv("ETHCLIENT")
	if GlobalClient == "" {
		log.Fatal("What env?")
	}

	fmt.Println("ENV?: ", GlobalClient)
	fmt.Println("Contract?: ", SwapContractAdd)
	if SwapContractAdd == "" {
		log.Fatal("PAIRCONTRACTADDRESS is not set")
	}
}

func SwapingEvent(vLog types.Log, client *ethclient.Client, db *sql.DB) {
	fmt.Println("SwapingEvent")

	// Topics[0] is the event signature
	// Topics[1] is the indexed user address
	// Topics[2] is the indexed tokenIn address
	// Topics[3] is the indexed tokenOut address
	user := common.HexToAddress(vLog.Topics[1].Hex())
	tokenIn := common.HexToAddress(vLog.Topics[2].Hex())
	tokenOut := common.HexToAddress(vLog.Topics[3].Hex())

	// Create a struct for the non-indexed parameters
	event := struct {
		AmountIn  *big.Int
		AmountOut *big.Int
		Fee       *big.Int
	}{}

	abiFile, err := os.ReadFile("Contract.json")
	if err != nil {
		log.Fatal("error reading ABI file", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatal(err)
	}

	// Unpack only the non-indexed parameters from data
	err = contractAbi.UnpackIntoInterface(&event, "Swap", vLog.Data)
	if err != nil {
		log.Fatal("Failed to unpack:", err)
	}

	swap := models.SwapEvent{
		User:      user.Hex(),
		TokenIn:   tokenIn.Hex(),
		TokenOut:  tokenOut.Hex(),
		AmountIn:  event.AmountIn,
		AmountOut: event.AmountOut,
		Fee:       event.Fee,
	}

	fmt.Printf("Full swap event: %+v\n", swap)
	mysql.InsertSwapEvent(db, swap)
}

func BurningEvent(vLog types.Log, client *ethclient.Client, db *sql.DB) {
	fmt.Println("BurningEvent")
}
