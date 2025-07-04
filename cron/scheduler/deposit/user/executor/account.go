package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type accountHandler struct {
	*depositaccmwpb.Account
	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
	incoming   decimal.Decimal
	outcoming  decimal.Decimal
	amount     decimal.Decimal
	coin       *coinmwpb.Coin
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

func (h *accountHandler) recheckAccountLock(ctx context.Context) (bool, error) {
	handler, err := depositaccmw.NewHandler(
		ctx,
		depositaccmw.WithEntID(&h.EntID, true),
	)
	if err != nil {
		return false, err
	}

	account, err := handler.GetAccount(ctx)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, fmt.Errorf("invalid account")
	}
	return account.Locked, nil
}

func (h *accountHandler) checkBalance(ctx context.Context) error {
	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.coin.Name,
		Address: h.Address,
	})
	if err != nil {
		return err
	}
	if bal == nil {
		return fmt.Errorf("invalid balance")
	}

	balance, err := decimal.NewFromString(bal.BalanceStr)
	if err != nil {
		return err
	}
	if balance.Cmp(h.incoming.Sub(h.outcoming)) <= 0 {
		return nil
	}
	h.amount = balance.Sub(h.incoming.Sub(h.outcoming))
	return nil
}

//nolint:gocritic
func (h *accountHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Account", h,
			"Incoming", h.incoming,
			"Outcoming", h.outcoming,
			"Amount", h.amount,
			"Coin", h.coin,
		)
	}

	persistentAccount := &types.PersistentAccount{
		Account:       h.Account,
		DepositAmount: h.amount.String(),
		Error:         *err,
	}

	if h.amount.Cmp(decimal.NewFromInt(0)) <= 0 && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.notif)
	}
	if h.amount.Cmp(decimal.NewFromInt(0)) > 0 {
		ioExtra := fmt.Sprintf(
			`{"AppID":"%v","UserID":"%v","AccountID":"%v","CoinName":"%v","Address":"%v","Date":"%v"}`,
			h.AppID,
			h.UserID,
			h.AccountID,
			h.coin.Name,
			h.Address,
			time.Now(),
		)
		persistentAccount.Extra = ioExtra
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
}

//nolint:gocritic
func (h *accountHandler) exec(ctx context.Context) error {
	if h.Locked {
		asyncfeed.AsyncFeed(ctx, h.Account, h.done)
		return nil
	}

	var err error
	var locked bool

	defer h.final(ctx, &err)

	h.incoming, err = decimal.NewFromString(h.Incoming)
	if err != nil {
		return err
	}
	h.outcoming, err = decimal.NewFromString(h.Outcoming)
	if err != nil {
		return err
	}

	if err = h.getCoin(ctx); err != nil {
		return err
	}
	if locked, err = h.recheckAccountLock(ctx); err != nil || locked {
		return err
	}
	if err = h.checkBalance(ctx); err != nil {
		return err
	}
	return nil
}
