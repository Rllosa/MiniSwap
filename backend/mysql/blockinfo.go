package mysql

import (
	"database/sql"

	"github.com/Rllosa/miniSwap/backend/mysql/models"
)

func ReadBlockInfo(db *sql.DB) (*models.BlockInfo, error) {
	var blockInfo models.BlockInfo
	err := db.QueryRow("SELECT * FROM blockInfo").Scan(&blockInfo.LatestBlockNum)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &blockInfo, nil
}

func WriteBlockInfo(db *sql.DB, latestBlockNum int64) error {
	_, err := db.Exec("INSERT INTO blockInfo (latestBlockNum) VALUES (?)", latestBlockNum)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlockInfo(db *sql.DB, latestBlockNum int64) error {
	_, err := db.Exec("UPDATE blockInfo SET latestBlockNum = ?", latestBlockNum)
	if err != nil {
		return err
	}
	return nil
}
