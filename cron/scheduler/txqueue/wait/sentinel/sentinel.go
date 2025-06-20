package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/wait/types"
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

func (h *handler) feedTx(ctx context.Context, tx *txmwpb.Tx, exec chan interface{}) error {
	if tx.State == basetypes.TxState_TxStateWait {
		state := basetypes.TxState_TxStateWaitCheck

		handler, err := txmw.NewHandler(
			ctx,
			txmw.WithID(&tx.ID, true),
			txmw.WithState(&state, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateTx(ctx); err != nil {
			return err
		}
	}
	cancelablefeed.CancelableFeed(ctx, tx, exec)
	return nil
}

func (h *handler) scanTxs(ctx context.Context, state basetypes.TxState, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &txmwpb.Conds{
		State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
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

		ignores := map[string]struct{}{}
		for _, tx := range txs {
			if _, ok := ignores[tx.FromAccountID]; ok {
				continue
			}
			if err := h.feedTx(ctx, tx, exec); err != nil {
				return err
			}
			ignores[tx.FromAccountID] = struct{}{}
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if err := h.scanTxs(ctx, basetypes.TxState_TxStateWait, exec); err != nil {
		return err
	}
	return h.scanTxs(ctx, basetypes.TxState_TxStateWaitCheck, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return h.scanTxs(ctx, basetypes.TxState_TxStateWaitCheck, exec)
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
