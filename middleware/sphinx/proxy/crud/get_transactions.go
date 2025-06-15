package crud

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/sphinx/plugin"
	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated/transaction"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
)

type GetTransactionsParam struct {
	CoinType         plugin.CoinType
	TransactionState proxy.TransactionState
}

// GetTransactions ..
func GetTransactions(ctx context.Context, params GetTransactionsParam) (rows []*ent.Transaction, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stmt := cli.
			Transaction.
			Query().
			Where(
				transaction.StateEQ(uint8(params.TransactionState)),
			)

		if params.CoinType != plugin.CoinType_CoinTypeUnKnow {
			stmt = stmt.Where(transaction.CoinTypeEQ(int32(params.CoinType)))
		}

		rows, err = stmt.Order(ent.Asc(transaction.FieldCreatedAt)).
			Limit(int(constant.DefaultRowLimit)).
			All(ctx)
		return err
	})
	return
}
