package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/transfer/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type accountHandler struct {
	*depositaccmwpb.Account
	persistent     chan interface{}
	notif          chan interface{}
	done           chan interface{}
	incoming       decimal.Decimal
	outcoming      decimal.Decimal
	amount         decimal.Decimal
	coin           *coinmwpb.Coin
	collectAccount *pltfaccmwpb.Account
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
	reserved, err := decimal.NewFromString(h.coin.ReservedAmount)
	if err != nil {
		return err
	}
	if balance.Cmp(reserved) <= 0 {
		return nil
	}
	collectAmount, err := decimal.NewFromString(h.coin.PaymentAccountCollectAmount)
	if err != nil {
		return err
	}
	if collectAmount.GreaterThan(balance) {
		return nil
	}
	h.amount = balance.Sub(reserved)
	return nil
}

func (h *accountHandler) getCollectAccount(ctx context.Context) error {
	conds := &pltfaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.coin.EntID},
		UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.AccountUsedFor_PaymentCollector)},
		Backup:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := pltfaccmw.NewHandler(
		ctx,
		pltfaccmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	account, err := handler.GetAccountOnly(ctx)
	if err != nil {
		return err
	}
	if account == nil {
		return fmt.Errorf("invalid collect account")
	}
	h.collectAccount = account
	return nil
}

func (h *accountHandler) checkAccountCoin() error {
	if h.collectAccount.CoinTypeID != h.CoinTypeID {
		return fmt.Errorf("invalid collect account coin")
	}
	return nil
}

func (h *accountHandler) checkFeeBalance(ctx context.Context) error {
	if h.CoinTypeID == h.coin.FeeCoinTypeID {
		return nil
	}

	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.coin.FeeCoinName,
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
	if balance.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("insufficient gas")
	}
	return nil
}

//nolint:gocritic
func (h *accountHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Infow(
			"final",
			"Account", h,
			"Incoming", h.incoming,
			"Outcoming", h.outcoming,
			"Amount", h.amount,
			"Coin", h.coin,
			"CollectAccount", h.collectAccount,
			"Error", *err,
		)
	}

	persistentAccount := &types.PersistentAccount{
		Account:          h.Account,
		CollectAmount:    h.amount.String(),
		FeeAmount:        decimal.NewFromInt(0).String(),
		DepositAccountID: h.AccountID,
		DepositAddress:   h.Address,
		Error:            *err,
	}
	if h.collectAccount != nil {
		persistentAccount.CollectAccountID = h.collectAccount.AccountID
		persistentAccount.CollectAddress = h.collectAccount.Address
	}

	if h.amount.Cmp(decimal.NewFromInt(0)) <= 0 && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentAccount, h.notif)
	}
	if h.amount.Cmp(decimal.NewFromInt(0)) > 0 {
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
	if err = h.getCollectAccount(ctx); err != nil {
		return err
	}
	if err = h.checkFeeBalance(ctx); err != nil {
		return err
	}
	if err = h.checkAccountCoin(); err != nil {
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
