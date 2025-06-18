package executor

import (
	"context"

	txmwcli "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/created/types"
)

type txHandler struct {
	*txmwpb.Tx
	persistent chan interface{}
	done       chan interface{}
	newState   basetypes.TxState
}

func (h *txHandler) checkWait(ctx context.Context) error {
	exist, err := txmwcli.ExistTxConds(ctx, &txmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
		AccountIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{h.FromAccountID, h.ToAccountID}},
		States: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(basetypes.TxState_TxStateWaitCheck),
			uint32(basetypes.TxState_TxStateWait),
			uint32(basetypes.TxState_TxStateTransferring),
		}},
	})
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	h.newState = basetypes.TxState_TxStateWait
	return nil
}

//nolint:gocritic
func (h *txHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Tx", h,
			"NewTxState", h.newState,
			"Error", *err,
		)
	}
	persistentTx := &types.PersistentTx{
		Tx: h.Tx,
	}
	if h.newState == h.State && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
		return
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
}

//nolint:gocritic
func (h *txHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)

	if err = h.checkWait(ctx); err != nil {
		return err
	}
	return nil
}
