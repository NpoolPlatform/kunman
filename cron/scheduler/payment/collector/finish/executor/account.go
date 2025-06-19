package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/finish/types"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"

	"github.com/google/uuid"
)

type accountHandler struct {
	*paymentaccountmwpb.Account
	persistent chan interface{}
	done       chan interface{}
	coin       *coinmwpb.Coin
	txFinished bool
}

func (h *accountHandler) getCoin(ctx context.Context) error {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(&h.CoinTypeID, true),
	)
	if err != nil {
		return err
	}

	coin, err := handler.GetCoin(ctx)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}
	h.coin = coin
	return nil
}

func (h *accountHandler) checkTransfer(ctx context.Context) error {
	if _, err := uuid.Parse(h.CollectingTID); err != nil {
		return err
	}
	if h.CollectingTID == uuid.Nil.String() {
		return nil
	}

	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithEntID(&h.CollectingTID, true),
	)
	if err != nil {
		return err
	}

	tx, err := handler.GetTx(ctx)
	if err != nil {
		return err
	}
	if tx == nil {
		h.txFinished = true
		return nil
	}

	switch tx.State {
	case basetypes.TxState_TxStateSuccessful:
		fallthrough //nolint
	case basetypes.TxState_TxStateFail:
		h.txFinished = true
	default:
		return nil
	}

	return nil
}

//nolint:gocritic
func (h *accountHandler) final(ctx context.Context, err *error) {
	persistentAccount := &types.PersistentAccount{
		Account: h.Account,
		Error:   *err,
	}

	if !h.txFinished && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
		return
	}
	if h.txFinished {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
}

//nolint:gocritic
func (h *accountHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getCoin(ctx); err != nil {
		return err
	}
	if err := h.checkTransfer(ctx); err != nil {
		return err
	}
	return nil
}
