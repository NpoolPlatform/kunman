package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/transfer/types"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/logger"
	payaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	payaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type accountHandler struct {
	*payaccmwpb.Account
	persistent     chan interface{}
	notif          chan interface{}
	done           chan interface{}
	amount         decimal.Decimal
	coin           *coinmwpb.Coin
	collectAccount *pltfaccmwpb.Account
}

func (h *accountHandler) getCoin(ctx context.Context, coinTypeID string) (*coinmwpb.Coin, error) {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(&coinTypeID, true),
	)
	if err != nil {
		return nil, err
	}

	coin, err := handler.GetCoin(ctx)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, fmt.Errorf("invalid coin")
	}
	return coin, nil
}

func (h *accountHandler) checkAccountCoin() error {
	if h.collectAccount.CoinTypeID != h.CoinTypeID {
		return fmt.Errorf("invalid collect account coin")
	}
	return nil
}

func (h *accountHandler) recheckAccount(ctx context.Context) (bool, error) {
	handler, err := payaccmw.NewHandler(
		ctx,
		payaccmw.WithEntID(&h.EntID, true),
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
	if account.Locked || account.Blocked || !account.Active {
		return false, nil
	}
	if account.AvailableAt >= uint32(time.Now().Unix()) {
		return false, nil
	}
	return true, nil
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

	limit, err := decimal.NewFromString(h.coin.PaymentAccountCollectAmount)
	if err != nil {
		return err
	}
	reserved, err := decimal.NewFromString(h.coin.ReservedAmount)
	if err != nil {
		return err
	}
	if balance.Cmp(limit) < 0 || balance.Cmp(reserved) <= 0 {
		return nil
	}
	h.amount = balance.Sub(reserved)
	return nil
}

func (h *accountHandler) checkFeeBalance(ctx context.Context) error {
	if h.coin.EntID == h.coin.FeeCoinTypeID {
		return nil
	}

	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.coin.FeeCoinName,
		Address: h.Address,
	})
	if err != nil {
		return err
	}
	if balance == nil {
		return fmt.Errorf("invalid balance")
	}
	bal, err := decimal.NewFromString(balance.BalanceStr)
	if err != nil {
		return err
	}
	if bal.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("insufficient gas")
	}

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

func (h *accountHandler) checkTransferring(ctx context.Context) (bool, error) {
	conds := &txmwpb.Conds{
		AccountID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AccountID},
		States: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(basetypes.TxState_TxStateCreated),
			uint32(basetypes.TxState_TxStateCreatedCheck),
			uint32(basetypes.TxState_TxStateWait),
			uint32(basetypes.TxState_TxStateWaitCheck),
			uint32(basetypes.TxState_TxStateTransferring),
			uint32(basetypes.TxState_TxStateSuccessful),
		}},
		Type: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.TxType_TxPaymentCollect)},
	}
	handler, err := txmw.NewHandler(
		ctx,
		txmw.WithConds(conds),
		txmw.WithOffset(0),
		txmw.WithLimit(1),
	)
	if err != nil {
		return false, err
	}

	txs, _, err := handler.GetTxs(ctx)
	if err != nil {
		return false, err
	}
	if len(txs) == 0 {
		return false, nil
	}

	if txs[0].State != basetypes.TxState_TxStateSuccessful {
		return false, nil
	}

	const coolDown = timedef.SecondsPerHour
	if txs[0].CreatedAt+coolDown > uint32(time.Now().Unix()) {
		return true, nil
	}
	return false, nil
}

//nolint:gocritic
func (h *accountHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Account", h,
			"Coin", h.coin,
			"CollectAccount", h.collectAccount,
			"Amount", h.amount,
			"Error", *err,
		)
	}

	persistentAccount := &types.PersistentAccount{
		Account:          h.Account,
		CollectAmount:    h.amount.String(),
		FeeAmount:        decimal.NewFromInt(0).String(),
		PaymentAccountID: h.AccountID,
		PaymentAddress:   h.Address,
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
		return nil
	}

	var err error
	var executable bool
	var yes bool

	defer h.final(ctx, &err)

	h.coin, err = h.getCoin(ctx, h.CoinTypeID)
	if err != nil {
		return err
	}
	if err = h.getCollectAccount(ctx); err != nil {
		return err
	}
	if err = h.checkAccountCoin(); err != nil {
		return err
	}
	if executable, err = h.recheckAccount(ctx); err != nil || !executable {
		return err
	}
	if yes, err = h.checkTransferring(ctx); err != nil || yes {
		return err
	}
	if err = h.checkFeeBalance(ctx); err != nil {
		return err
	}
	if err = h.checkBalance(ctx); err != nil {
		return err
	}
	return nil
}
