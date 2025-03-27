package models

import (
	"math/big"
)

type BlockInfo struct {
	LatestBlockNum int64
}

type SwapEvent struct {
	User      string // Changed from UserAddress
	TokenIn   string
	TokenOut  string
	AmountIn  *big.Int
	AmountOut *big.Int
	Fee       *big.Int
}
