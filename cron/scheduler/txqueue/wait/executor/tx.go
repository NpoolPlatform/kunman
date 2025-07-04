package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/wait/types"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	accountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	useraccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	accountmw "github.com/NpoolPlatform/kunman/middleware/account/account"
	useraccmw "github.com/NpoolPlatform/kunman/middleware/account/user"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type txHandler struct {
	*txmwpb.Tx
	persistent       chan interface{}
	notif            chan interface{}
	done             chan interface{}
	newState         basetypes.TxState
	transactionExist bool
	fromAccount      *accountmwpb.Account
	toAccount        *accountmwpb.Account
	transferAmount   decimal.Decimal
	txCoin           *coinmwpb.Coin
	toAccountCoin    *coinmwpb.Coin
	memo             *string
}

func (h *txHandler) checkTransfer(ctx context.Context) (bool, error) {
	tx, err := sphinxproxycli.GetTransaction(ctx, h.EntID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false, nil
		}
		return false, err
	}
	if tx == nil {
		return false, nil
	}
	h.newState = basetypes.TxState_TxStateTransferring
	h.transactionExist = true
	return true, nil
}

func (h *txHandler) getAccount(ctx context.Context, accountID string) (*accountmwpb.Account, error) {
	handler, err := accountmw.NewHandler(
		ctx,
		accountmw.WithEntID(&accountID, true),
	)
	if err != nil {
		return nil, err
	}

	account, err := handler.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	if account == nil {
		h.newState = basetypes.TxState_TxStateFail
		return nil, fmt.Errorf("invalid account")
	}
	return account, nil
}

func (h *txHandler) getCoin(ctx context.Context, coinTypeID string) (*coinmwpb.Coin, error) {
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
		h.newState = basetypes.TxState_TxStateFail
		return nil, fmt.Errorf("invalid coin")
	}
	return coin, nil
}

func (h *txHandler) checkTransferAmount(ctx context.Context) error {
	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.txCoin.Name,
		Address: h.fromAccount.Address,
	})
	if err != nil {
		return fmt.Errorf("fail check transfer amount (%v)", err)
	}
	if bal == nil {
		return fmt.Errorf("invalid balance")
	}

	amount, err := decimal.NewFromString(h.Amount)
	if err != nil {
		return err
	}
	feeAmount, err := decimal.NewFromString(h.FeeAmount)
	if err != nil {
		return err
	}
	balance, err := decimal.NewFromString(bal.BalanceStr)
	if err != nil {
		return err
	}
	reserved, err := decimal.NewFromString(h.txCoin.ReservedAmount)
	if err != nil {
		return err
	}
	if amount.Cmp(feeAmount) <= 0 {
		h.newState = basetypes.TxState_TxStateFail
		return fmt.Errorf("invalid amount")
	}

	h.transferAmount = amount.Sub(feeAmount)
	if h.transferAmount.Add(reserved).Cmp(balance) > 0 {
		return fmt.Errorf("insufficient funds")
	}

	return nil
}

func (h *txHandler) checkFeeAmount(ctx context.Context) error {
	if h.txCoin.EntID == h.txCoin.FeeCoinTypeID {
		return nil
	}

	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.txCoin.FeeCoinName,
		Address: h.fromAccount.Address,
	})
	if err != nil {
		return fmt.Errorf("fail check fee amount (%v)", err)
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

func (h *txHandler) getMemo(ctx context.Context) error {
	if h.Type != basetypes.TxType_TxWithdraw {
		return nil
	}

	conds := &useraccmwpb.Conds{
		AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.ToAccountID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.UsedFor_Withdraw)},
	}
	handler, err := useraccmw.NewHandler(
		ctx,
		useraccmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	account, err := handler.GetAccountOnly(ctx)
	if err != nil {
		return err
	}
	if account == nil {
		h.newState = basetypes.TxState_TxStateFail
		return fmt.Errorf("invalid useraccount")
	}
	if account.Memo == "" {
		return nil
	}
	h.memo = &account.Memo
	return nil
}

func (h *txHandler) checkAccountCoin() error {
	switch h.fromAccount.CoinTypeID {
	case h.CoinTypeID:
	default:
		h.newState = basetypes.TxState_TxStateFail
		return fmt.Errorf("invalid from account coin")
	}
	switch h.CoinTypeID {
	case h.toAccount.CoinTypeID:
	case h.toAccountCoin.FeeCoinTypeID:
	default:
		h.newState = basetypes.TxState_TxStateFail
		return fmt.Errorf("invalid to account coin")
	}
	return nil
}

//nolint:gocritic
func (h *txHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Tx", h,
			"NewTxState", h.newState,
			"TransactionExist", h.transactionExist,
			"Coin", h.txCoin,
			"FromAccount", h.fromAccount,
			"ToAccount", h.toAccount,
			"Error", *err,
		)
	}

	persistentTx := &types.PersistentTx{
		Tx:               h.Tx,
		TransactionExist: h.transactionExist,
		Amount:           h.transferAmount.String(),
		FloatAmount:      h.transferAmount.InexactFloat64(),
		AccountMemo:      h.memo,
		NewTxState:       h.newState,
		Error:            *err,
	}
	if h.txCoin != nil {
		persistentTx.CoinName = h.txCoin.Name
	}
	if h.fromAccount != nil {
		persistentTx.FromAddress = h.fromAccount.Address
	}
	if h.toAccount != nil {
		persistentTx.ToAddress = h.toAccount.Address
	}

	if h.newState == h.State && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.notif)
	}
	if h.newState != h.State {
		asyncfeed.AsyncFeed(ctx, persistentTx, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentTx, h.done)
}

//nolint:gocritic
func (h *txHandler) exec(ctx context.Context) error {
	h.newState = h.State

	var err error
	defer h.final(ctx, &err)
	if exist, err := h.checkTransfer(ctx); err != nil || exist {
		return err
	}
	h.txCoin, err = h.getCoin(ctx, h.CoinTypeID)
	if err != nil {
		return err
	}
	h.fromAccount, err = h.getAccount(ctx, h.FromAccountID)
	if err != nil {
		return err
	}
	h.toAccount, err = h.getAccount(ctx, h.ToAccountID)
	if err != nil {
		return err
	}
	h.toAccountCoin, err = h.getCoin(ctx, h.toAccount.CoinTypeID)
	if err != nil {
		return err
	}
	if err = h.checkAccountCoin(); err != nil {
		return err
	}
	if err = h.getMemo(ctx); err != nil {
		return err
	}
	if err = h.checkTransferAmount(ctx); err != nil {
		return err
	}
	if err = h.checkFeeAmount(ctx); err != nil {
		return err
	}

	h.newState = basetypes.TxState_TxStateTransferring

	return nil
}
