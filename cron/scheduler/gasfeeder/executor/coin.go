//nolint:dupl
package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/gasfeeder/types"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/logger"
	accountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	payaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	accountmw "github.com/NpoolPlatform/kunman/middleware/account/account"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
	payaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	txmw "github.com/NpoolPlatform/kunman/middleware/chain/tx"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type coinHandler struct {
	*coinmwpb.Coin
	persistent         chan interface{}
	notif              chan interface{}
	done               chan interface{}
	gasProviderAccount *accountmwpb.Account
}

func (h *coinHandler) getPlatformAccount(ctx context.Context, coinTypeID string, usedFor basetypes.AccountUsedFor) (*pltfaccmwpb.Account, error) {
	conds := &pltfaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: coinTypeID},
		UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(usedFor)},
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
		return nil, err
	}

	account, err := handler.GetAccountOnly(ctx)
	if err != nil {
		return nil, err
	}
	if account == nil || account.Address == "" {
		return nil, fmt.Errorf("invalid ")
	}
	return account, nil
}

func (h *coinHandler) getGasProvider(ctx context.Context) error {
	account, err := h.getPlatformAccount(ctx, h.FeeCoinTypeID, basetypes.AccountUsedFor_GasProvider)
	if err != nil {
		return err
	}
	if account == nil {
		return fmt.Errorf("invalid gasprovider")
	}

	handler, err := accountmw.NewHandler(
		ctx,
		accountmw.WithEntID(&account.AccountID, true),
	)
	if err != nil {
		return err
	}

	_account, err := handler.GetAccount(ctx)
	if err != nil {
		return err
	}
	if _account == nil {
		return fmt.Errorf("invalid gasprovider")
	}

	h.gasProviderAccount = _account
	return nil
}

func (h *coinHandler) feeding(ctx context.Context, account *accountmwpb.Account) (bool, error) {
	conds := &txmwpb.Conds{
		AccountID: &basetypes.StringVal{Op: cruder.EQ, Value: account.EntID},
		Type:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.TxType_TxFeedGas)},
		States: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(basetypes.TxState_TxStateCreated),
			uint32(basetypes.TxState_TxStateCreatedCheck),
			uint32(basetypes.TxState_TxStateWait),
			uint32(basetypes.TxState_TxStateWaitCheck),
			uint32(basetypes.TxState_TxStateTransferring),
			uint32(basetypes.TxState_TxStateSuccessful),
		}},
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
		logger.Sugar().Debugw(
			"feeding",
			"Account", account,
			"Txs", txs,
			"State", "Feeding",
		)
		return true, nil
	}

	const coolDown = uint32(10 * timedef.SecondsPerMinute)
	if txs[0].UpdatedAt+coolDown > uint32(time.Now().Unix()) {
		logger.Sugar().Debugw(
			"feeding",
			"Account", account,
			"Txs", txs,
			"State", "Feeding",
		)
		return true, nil
	}

	return false, nil
}

func (h *coinHandler) enough(ctx context.Context, account *accountmwpb.Account, amount decimal.Decimal) (bool, error) {
	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.FeeCoinName,
		Address: account.Address,
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
	logger.Sugar().Debugw(
		"enough",
		"Account", account,
		"Amount", amount,
		"Balance", bal,
	)
	return bal.Cmp(amount) >= 0, nil
}

func (h *coinHandler) feedable(ctx context.Context, account *accountmwpb.Account, amount, low decimal.Decimal) (bool, error) {
	if enough, err := h.enough(ctx, account, low); err != nil || enough {
		logger.Sugar().Debugw(
			"feedable",
			"Account", account,
			"Amount", amount,
			"Low", low,
			"Enough", enough,
			"Error", err,
		)
		return false, err
	}
	if feeding, err := h.feeding(ctx, account); err != nil || feeding {
		logger.Sugar().Debugw(
			"feedable",
			"Account", account,
			"Amount", amount,
			"Low", low,
			"Feeding", feeding,
			"Error", err,
		)
		return false, err
	}

	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.Name,
		Address: account.Address,
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
	reserved, err := decimal.NewFromString(h.ReservedAmount)
	if err != nil {
		return false, err
	}
	if bal.Cmp(reserved) <= 0 {
		logger.Sugar().Debugw(
			"feedable",
			"Account", account,
			"Amount", amount,
			"Low", low,
			"Reserved", reserved,
			"Balance", bal,
			"Error", err,
		)
		return false, nil
	}

	enough, err := h.enough(ctx, h.gasProviderAccount, amount)
	if err != nil {
		return false, err
	}
	if !enough {
		return false, fmt.Errorf("insufficient funds")
	}
	return true, nil
}

func (h *coinHandler) checkUserBenefitHot(ctx context.Context) (bool, *accountmwpb.Account, decimal.Decimal, error) {
	account, err := h.getPlatformAccount(ctx, h.EntID, basetypes.AccountUsedFor_UserBenefitHot)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}
	if account == nil {
		return false, nil, decimal.NewFromInt(0), fmt.Errorf("invalid account")
	}

	handler, err := accountmw.NewHandler(
		ctx,
		accountmw.WithEntID(&account.AccountID, true),
	)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}

	_account, err := handler.GetAccount(ctx)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}
	if _account == nil {
		return false, nil, decimal.NewFromInt(0), fmt.Errorf("invalid account")
	}

	amount, err := decimal.NewFromString(h.HotWalletFeeAmount)
	if err != nil {
		return false, _account, decimal.NewFromInt(0), err
	}
	lowFeeAmount, err := decimal.NewFromString(h.HotLowFeeAmount)
	if err != nil {
		return false, _account, decimal.NewFromInt(0), err
	}

	feedable, err := h.feedable(ctx, _account, amount, lowFeeAmount)
	if err != nil {
		return false, _account, decimal.NewFromInt(0), err
	}
	return feedable, _account, amount, nil
}

func (h *coinHandler) checkPaymentAccount(ctx context.Context) (bool, *accountmwpb.Account, decimal.Decimal, error) {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	amount, err := decimal.NewFromString(h.CollectFeeAmount)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}
	lowFeeAmount, err := decimal.NewFromString(h.LowFeeAmount)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}

	conds := &payaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}

	for {
		paymentHandler, err := payaccmw.NewHandler(
			ctx,
			payaccmw.WithConds(conds),
			payaccmw.WithOffset(offset),
			payaccmw.WithLimit(limit),
		)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		accounts, _, err := paymentHandler.GetAccounts(ctx)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}
		if len(accounts) == 0 {
			return false, nil, decimal.NewFromInt(0), nil
		}

		ids := []string{}
		for _, account := range accounts {
			ids = append(ids, account.AccountID)
		}

		_conds := &accountmwpb.Conds{
			EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
		}
		accountHandler, err := accountmw.NewHandler(
			ctx,
			accountmw.WithConds(_conds),
			accountmw.WithOffset(0),
			accountmw.WithLimit(int32(len(ids))),
		)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		_accounts, _, err := accountHandler.GetAccounts(ctx)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		for _, account := range _accounts {
			if feedable, err := h.feedable(ctx, account, amount, lowFeeAmount); err != nil || feedable {
				return feedable, account, amount, err
			}
		}

		offset += limit
	}
}

func (h *coinHandler) checkDepositAccount(ctx context.Context) (bool, *accountmwpb.Account, decimal.Decimal, error) {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	amount, err := decimal.NewFromString(h.CollectFeeAmount)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}
	lowFeeAmount, err := decimal.NewFromString(h.LowFeeAmount)
	if err != nil {
		return false, nil, decimal.NewFromInt(0), err
	}

	conds := &depositaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.EntID},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Locked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}

	for {
		depositHandler, err := depositaccmw.NewHandler(
			ctx,
			depositaccmw.WithConds(conds),
			depositaccmw.WithOffset(offset),
			depositaccmw.WithLimit(limit),
		)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		accounts, _, err := depositHandler.GetAccounts(ctx)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}
		if len(accounts) == 0 {
			return false, nil, decimal.NewFromInt(0), nil
		}

		ids := []string{}
		for _, account := range accounts {
			ids = append(ids, account.AccountID)
		}

		_conds := &accountmwpb.Conds{
			EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
		}
		accountHandler, err := accountmw.NewHandler(
			ctx,
			accountmw.WithConds(_conds),
			accountmw.WithOffset(0),
			accountmw.WithLimit(int32(len(ids))),
		)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		_accounts, _, err := accountHandler.GetAccounts(ctx)
		if err != nil {
			return false, nil, decimal.NewFromInt(0), err
		}

		for _, account := range _accounts {
			if feedable, err := h.feedable(ctx, account, amount, lowFeeAmount); err != nil || feedable {
				return feedable, account, amount, err
			}
		}

		offset += limit
	}
}

// TODO: in case some mining product get rewards other than it native coin (native coin is fee coin)
func (h *coinHandler) checkGoodBenefit(ctx context.Context) (bool, *accountmwpb.Account, decimal.Decimal, error) { //nolint
	return false, nil, decimal.NewFromInt(0), nil
}

//nolint:gocritic,interfacer
func (h *coinHandler) final(ctx context.Context, account **accountmwpb.Account, usedFor *basetypes.AccountUsedFor, amount *decimal.Decimal, feedable *bool, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Coin", h.Coin,
			"GasProvider", h.gasProviderAccount,
			"Amount", amount,
			"UsedFor", *usedFor,
			"Account", *account,
			"Feedable", *feedable,
			"Error", *err,
		)
	}

	persistentCoin := &types.PersistentCoin{
		Coin:      h.Coin,
		Amount:    amount.String(),
		FeeAmount: decimal.NewFromInt(0).String(),
		UsedFor:   *usedFor,
		Extra:     fmt.Sprintf(`{"Coin":"%v","FeeCoin":"%v","Type":"%v"}`, h.Name, h.FeeCoinName, *usedFor),
		Error:     *err,
	}
	if *account != nil {
		persistentCoin.ToAccountID = (*account).EntID
		persistentCoin.ToAddress = (*account).Address
	}
	if h.gasProviderAccount != nil {
		persistentCoin.FromAccountID = h.gasProviderAccount.EntID
		persistentCoin.FromAddress = h.gasProviderAccount.Address
	}

	if !*feedable && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.notif)
	}
	if *feedable {
		asyncfeed.AsyncFeed(ctx, persistentCoin, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentCoin, h.done)
}

//nolint:gocritic
func (h *coinHandler) exec(ctx context.Context) error {
	var err error
	var account *accountmwpb.Account
	var amount decimal.Decimal
	var feedable bool
	var usedFor basetypes.AccountUsedFor
	var feeding bool

	defer h.final(ctx, &account, &usedFor, &amount, &feedable, &err)

	if err = h.getGasProvider(ctx); err != nil {
		return err
	}
	if feeding, err = h.feeding(ctx, h.gasProviderAccount); err != nil || feeding {
		return err
	}
	if feedable, account, amount, err = h.checkUserBenefitHot(ctx); err != nil || feedable {
		usedFor = basetypes.AccountUsedFor_UserBenefitHot
		return err
	}
	if feedable, account, amount, err = h.checkPaymentAccount(ctx); err != nil || feedable {
		usedFor = basetypes.AccountUsedFor_GoodPayment
		return err
	}
	if feedable, account, amount, err = h.checkDepositAccount(ctx); err != nil || feedable {
		usedFor = basetypes.AccountUsedFor_UserDeposit
		return err
	}
	if feedable, account, amount, err = h.checkGoodBenefit(ctx); err != nil || feedable {
		usedFor = basetypes.AccountUsedFor_GoodBenefit
		return err
	}

	return nil
}
