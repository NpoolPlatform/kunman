package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/limitation/types"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/logger"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type coinHandler struct {
	*coinmwpb.Coin
	persistent             chan interface{}
	notif                  chan interface{}
	done                   chan interface{}
	userBenefitHotAccount  *pltfaccmwpb.Account
	userBenefitColdAccount *pltfaccmwpb.Account
	amount                 decimal.Decimal
}

func (h *coinHandler) getPlatformAccount(ctx context.Context, usedFor basetypes.AccountUsedFor) (*pltfaccmwpb.Account, error) {
	conds := &pltfaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
		UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(usedFor)},
		Backup:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := pltfaccmw.NewHandler(
		ctx,
		pltfaccmw.WithConds(conds),
	)
	if err != nil {
		return nil, err
	}

	account, err := handler.GetAccountOnly(ctx)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("invalid account")
	}
	return account, nil
}

func (h *coinHandler) checkBalanceLimitation(ctx context.Context) (bool, error) {
	limit, err := decimal.NewFromString(h.HotWalletAccountAmount)
	if err != nil {
		return false, err
	}

	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.Name,
		Address: h.userBenefitHotAccount.Address,
	})
	if err != nil {
		return false, err
	}
	if balance == nil {
		return false, fmt.Errorf("invalid balance")
	}

	bal, err := decimal.NewFromString(balance.BalanceStr)
	if err != nil {
		return false, err
	}
	if bal.Cmp(limit.Mul(decimal.NewFromInt(2))) < 0 {
		return false, nil
	}

	h.amount = bal.Sub(limit)

	return true, nil
}

func (h *coinHandler) checkTransferring(ctx context.Context) (bool, error) {
	conds := &txmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
		AccountIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{
			h.userBenefitHotAccount.AccountID,
			h.userBenefitColdAccount.AccountID,
		}},
		States: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(basetypes.TxState_TxStateCreated),
			uint32(basetypes.TxState_TxStateCreatedCheck),
			uint32(basetypes.TxState_TxStateWait),
			uint32(basetypes.TxState_TxStateWaitCheck),
			uint32(basetypes.TxState_TxStateTransferring),
			uint32(basetypes.TxState_TxStateSuccessful),
		}},
		Type: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.TxType_TxLimitation)},
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
		return true, nil
	}

	if txs[0].State != basetypes.TxState_TxStateSuccessful {
		return true, nil
	}

	const coolDown = timedef.SecondsPerHour
	if txs[0].CreatedAt+coolDown > uint32(time.Now().Unix()) {
		return true, nil
	}
	return false, nil
}

func (h *coinHandler) checkAccountCoin() error {
	if h.userBenefitHotAccount.CoinTypeID != h.EntID {
		return fmt.Errorf("invalid hot account")
	}
	if h.userBenefitColdAccount.CoinTypeID != h.EntID {
		return fmt.Errorf("invalid hot account")
	}
	return nil
}

func (h *coinHandler) checkFeeBalance(ctx context.Context) error {
	if h.EntID == h.FeeCoinTypeID {
		return nil
	}
	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.FeeCoinName,
		Address: h.userBenefitHotAccount.Address,
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

//nolint:gocritic
func (h *coinHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Coin", h,
			"HotAccount", h.userBenefitHotAccount,
			"Amount", h.amount,
			"Error", *err,
		)
	}

	persistentCoin := &types.PersistentCoin{
		Coin:      h.Coin,
		Amount:    h.amount.String(),
		FeeAmount: decimal.NewFromInt(0).String(),
		Error:     *err,
	}

	if h.userBenefitHotAccount != nil {
		persistentCoin.FromAccountID = h.userBenefitHotAccount.AccountID
		persistentCoin.FromAddress = h.userBenefitHotAccount.Address
	}
	if h.userBenefitColdAccount != nil {
		persistentCoin.ToAccountID = h.userBenefitColdAccount.AccountID
		persistentCoin.ToAddress = h.userBenefitColdAccount.Address
	}

	if h.amount.Cmp(decimal.NewFromInt(0)) <= 0 && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.notif)
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentCoin, h.done)
}

//nolint:gocritic
func (h *coinHandler) exec(ctx context.Context) error {
	var err error
	var yes bool
	defer h.final(ctx, &err)

	h.userBenefitHotAccount, err = h.getPlatformAccount(ctx, basetypes.AccountUsedFor_UserBenefitHot)
	if err != nil {
		return err
	}
	h.userBenefitColdAccount, err = h.getPlatformAccount(ctx, basetypes.AccountUsedFor_UserBenefitCold)
	if err != nil {
		return err
	}
	if err = h.checkAccountCoin(); err != nil {
		return err
	}
	if err = h.checkFeeBalance(ctx); err != nil {
		return err
	}
	if yes, err = h.checkTransferring(ctx); err != nil || yes {
		return err
	}
	if yes, err = h.checkBalanceLimitation(ctx); err != nil || !yes {
		return err
	}

	return nil
}
