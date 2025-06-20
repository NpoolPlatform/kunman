package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/transferring/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
)

type withdrawHandler struct {
	*withdrawmwpb.Withdraw
	persistent       chan interface{}
	notif            chan interface{}
	done             chan interface{}
	newWithdrawState ledgertypes.WithdrawState
	chainTxID        string
}

func (h *withdrawHandler) checkTransfer(ctx context.Context) error {
	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithEntID(&h.PlatformTransactionID, true),
	)
	if err != nil {
		return err
	}

	tx, err := handler.GetTx(ctx)
	if err != nil {
		return err
	}
	if tx == nil {
		h.newWithdrawState = ledgertypes.WithdrawState_PreFail
		return fmt.Errorf("invalid tx")
	}
	switch tx.State {
	case basetypes.TxState_TxStateSuccessful:
		h.newWithdrawState = ledgertypes.WithdrawState_PreSuccessful
	case basetypes.TxState_TxStateFail:
		h.newWithdrawState = ledgertypes.WithdrawState_PreFail
	}
	h.chainTxID = tx.ChainTxID
	return nil
}

//nolint:gocritic
func (h *withdrawHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Withdraw", h.Withdraw,
			"NewWithdrawState", h.newWithdrawState,
			"ChainTxID", h.chainTxID,
			"Error", *err,
		)
	}
	persistentWithdraw := &types.PersistentWithdraw{
		Withdraw:         h.Withdraw,
		NewWithdrawState: h.newWithdrawState,
		ChainTxID:        h.chainTxID,
		Error:            *err,
	}
	if h.newWithdrawState == h.State && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.notif)
	}
	if h.newWithdrawState != h.State {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.done)
}

//nolint:gocritic
func (h *withdrawHandler) exec(ctx context.Context) error {
	h.newWithdrawState = h.State

	var err error
	defer h.final(ctx, &err)

	if err = h.checkTransfer(ctx); err != nil {
		return err
	}

	return nil
}
