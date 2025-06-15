package crud

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated/transaction"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         proxy.TransactionState
	NextState     proxy.TransactionState
	Payload       []byte
	Cid           string
	ExitCode      int64
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t *UpdateTransactionParams) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stmt := cli.
			Transaction.
			Update().
			Where(
				transaction.TransactionIDEQ(t.TransactionID),
				transaction.StateEQ(uint8(t.State)),
			).
			SetPayload(t.Payload).
			SetState(uint8(t.NextState)).
			SetExitCode(t.ExitCode)

		if t.Cid != "" {
			stmt.
				SetCid(t.Cid)
		}

		return stmt.Exec(ctx)
	})
}
