package types

import (
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
)

type PersistentCoin struct {
	*coinmwpb.Coin
	FromAccountID string
	FromAddress   string
	ToAccountID   string
	ToAddress     string
	Amount        string
	FeeAmount     string
	Error         error
}
