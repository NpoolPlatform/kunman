package crud

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db"
	ent "github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/sphinx/proxy/db/ent/generated/transaction"
)

// GetTransaction ..
func GetTransaction(ctx context.Context, transactionID string) (row *ent.Transaction, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		row, err = cli.
			Transaction.
			Query().
			Select(
				transaction.FieldTransactionID,
				transaction.FieldCoinType,
				transaction.FieldName,
				transaction.FieldFrom,
				transaction.FieldTo,
				transaction.FieldMemo,
				transaction.FieldAmount,
				transaction.FieldCid,
				transaction.FieldExitCode,
				transaction.FieldPayload,
				transaction.FieldState,
				transaction.FieldCreatedAt,
				transaction.FieldUpdatedAt,
			).
			Where(
				transaction.TransactionIDEQ(transactionID),
			).
			Only(ctx)
		return err
	})
	return
}

type GetTransactionExistParam struct {
	TransactionID    string
	TransactionState proxy.TransactionState
}

func GetTransactionExist(ctx context.Context, params GetTransactionExistParam) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Transaction.
			Query().
			Select(transaction.FieldID).
			Where(
				transaction.TransactionIDEQ(params.TransactionID),
			)

		if params.TransactionState != proxy.TransactionState_TransactionStateUnKnow {
			stm.Where(transaction.StateEQ(uint8(params.TransactionState)))
		}

		exist, err = stm.Exist(ctx)
		return err
	})
	return
}
