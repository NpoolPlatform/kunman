package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/transferring/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &txmwpb.Conds{
		State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.TxState_TxStateTransferring)},
	}

	for {
		handler, err := txmw.NewHandler(
			ctx,
			txmw.WithConds(conds),
			txmw.WithOffset(offset),
			txmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		txs, _, err := handler.GetTxs(ctx)
		if err != nil {
			return err
		}
		if len(txs) == 0 {
			return nil
		}

		for _, tx := range txs {
			cancelablefeed.CancelableFeed(ctx, tx, exec)
		}

		offset += limit
	}
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if tx, ok := ent.(*types.PersistentTx); ok {
		return tx.EntID
	}
	return ent.(*txmwpb.Tx).EntID
}
