package crud

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/price"
	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
)

type CreateTransactionParam struct {
	CoinType         plugin.CoinType
	TransactionState proxy.TransactionState
	TransactionID    string
	Name             string
	From             string
	To               string
	Value            float64
	Memo             string
}

func CreateTransaction(ctx context.Context, t *CreateTransactionParam) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := cli.Transaction.Create().
			SetCoinType(int32(t.CoinType)).
			SetTransactionID(t.TransactionID).
			SetName(t.Name).
			SetFrom(t.From).
			SetTo(t.To).
			SetMemo(t.Memo).
			SetAmount(price.VisualPriceToDBPrice(t.Value)).
			SetState(uint8(t.TransactionState)).
			Save(ctx)
		return err
	})
}
