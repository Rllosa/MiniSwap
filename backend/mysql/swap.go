package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Rllosa/miniSwap/backend/mysql/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertSwapEvent(db *sql.DB, swapEvent models.SwapEvent) error {

	fmt.Println("InsertSwapEvent")

	query := `
	INSERT INTO swap_events (
		user_address,
		token_in,
		token_out,
		amount_in,
		amount_out,
		fee,
		created_at
	) VALUES (?, ?, ?, ?, ?, ?, FROM_UNIXTIME(?))
`
	unixTime := time.Now().Unix()

	_, err := db.Exec(
		query,
		swapEvent.User,
		swapEvent.TokenIn,
		swapEvent.TokenOut,
		swapEvent.AmountIn.String(),  // Convert big.Int to decimal string
		swapEvent.AmountOut.String(), // Convert big.Int to string for DB storage
		swapEvent.Fee.String(),
		unixTime,
	)

	if err != nil {
		fmt.Println("Error inserting swap event:", err)
	}

	fmt.Println("Swap event inserted successfully")

	return nil
}
