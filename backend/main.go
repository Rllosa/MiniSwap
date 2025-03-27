package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/Rllosa/miniSwap/backend/mysql"
	service "github.com/Rllosa/miniSwap/backend/services"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
)

const (
	swapEvent     = "0xd6d34547c69c5ee3d2667625c188acf1006abb93e0ee7cf03925c67cf7760413"
	burningEvent  = ""
	startBlock    = 0 // start block - 1 (since event log will add 1)
	blockInterval = 100
)

func main() {
	fmt.Println("Event Listener started")
	runCronJobs()
}

func runCronJobs() {
	client, err := ethclient.Dial(service.GlobalClient)
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := gocron.NewScheduler(time.UTC)
	s.Every(20).Seconds().Do(func() {
		cronService(client, db)
	})
	s.StartBlocking()
}

func cronService(client *ethclient.Client, db *sql.DB) {

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockInfo, err := mysql.ReadBlockInfo(db)
	if err != nil {
		log.Println("Db error", err)
	}

	if blockInfo.LatestBlockNum == 0 {
		mysql.WriteBlockInfo(db, startBlock)
		blockInfo.LatestBlockNum = startBlock
	}
	endBlock := header.Number.Int64()

	if endBlock > (blockInfo.LatestBlockNum + blockInterval) {
		endBlock = blockInfo.LatestBlockNum + blockInterval
	}

	fmt.Println(blockInfo.LatestBlockNum+1, endBlock)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockInfo.LatestBlockNum + 1),
		ToBlock:   big.NewInt(endBlock),
		Addresses: []common.Address{
			common.HexToAddress(service.SwapContractAdd),
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AYA")
	for _, vLog := range logs {

		fmt.Println(vLog.Topics[0].Hex())

		switch vLog.Topics[0].Hex() {

		case swapEvent:
			service.SwapingEvent(vLog, client, db)
		case burningEvent:
			service.BurningEvent(vLog, client, db)
		}
	}

	mysql.UpdateBlockInfo(db, endBlock)
}
