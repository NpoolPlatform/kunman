package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/finish/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"

	"github.com/shopspring/decimal"
)

type accountHandler struct {
	*depositaccmwpb.Account
	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
	coin       *coinmwpb.Coin
	outcoming  decimal.Decimal
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
		h.outcoming = decimal.RequireFromString(tx.Amount)
		fallthrough
	case basetypes.TxState_TxStateFail:
		h.txFinished = true
	default:
		return nil
	}

	return nil
}

func (h *accountHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Account", h,
			"Outcoming", h.outcoming,
			"Coin", h.coin,
			"TxFinished", h.txFinished,
			"Error", *err,
		)
	}

	persistentAccount := &types.PersistentAccount{
		Account: h.Account,
		Error:   *err,
	}
	if h.outcoming.Cmp(decimal.NewFromInt(0)) > 0 {
		outcoming := h.outcoming.String()
		persistentAccount.CollectOutcoming = &outcoming
	}
	if !h.txFinished && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.notif)
	}
	if h.txFinished {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
}

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
