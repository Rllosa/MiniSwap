package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "server:password@tcp(127.0.0.1:3306)/MiniSwap")

	if err != nil {
		return nil, err
	}

	return db, nil
}
