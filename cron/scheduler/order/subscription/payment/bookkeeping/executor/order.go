package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/bookkeeping/types"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*subscriptionordermwpb.SubscriptionOrder
	done                 chan interface{}
	persistent           chan interface{}
	notif                chan interface{}
	paymentTransfers     []*types.XPaymentTransfer
	paymentTransferCoins map[string]*coinmwpb.Coin
	paymentFiats         map[string]*fiatmwpb.Fiat
	xPaymentFiats        []*types.XPaymentFiat
	paymentAccounts      map[string]*paymentaccountmwpb.Account
}

func (h *orderHandler) getPaymentCoins(ctx context.Context) (err error) {
	if len(h.paymentTransfers) == 0 {
		return nil
	}

	h.paymentTransferCoins, err = schedcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, paymentTransfer := range h.PaymentTransfers {
			coinTypeIDs = append(coinTypeIDs, paymentTransfer.CoinTypeID)
		}
		return
	}())
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, paymentTransfer := range h.PaymentTransfers {
		if _, ok := h.paymentTransferCoins[paymentTransfer.CoinTypeID]; !ok {
			return wlog.Errorf("invalid paymenttransfercoin")
		}
	}
	for _, paymentCoin := range h.paymentTransferCoins {
		if !paymentCoin.ForPay {
			return wlog.Errorf("invalid paymenttransfercoin")
		}
	}
	return nil
}

func (h *orderHandler) getPaymentFiats(ctx context.Context) (err error) {
	if len(h.PaymentFiats) == 0 {
		return nil
	}

	h.paymentFiats, err = schedcommon.GetFiats(ctx, func() (fiatIDs []string) {
		for _, paymentFiat := range h.PaymentFiats {
			fiatIDs = append(fiatIDs, paymentFiat.FiatID)
		}
		return
	}())
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, paymentFiat := range h.PaymentFiats {
		if _, ok := h.paymentFiats[paymentFiat.FiatID]; !ok {
			return wlog.Errorf("invalid paymentfiat")
		}
	}
	return nil
}

func (h *orderHandler) getPaymentAccounts(ctx context.Context) (err error) {
	h.paymentAccounts, err = schedcommon.GetPaymentAccounts(ctx, func() (accountIDs []string) {
		for _, paymentTransfer := range h.PaymentTransfers {
			accountIDs = append(accountIDs, paymentTransfer.AccountID)
		}
		return
	}())
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, paymentTransfer := range h.PaymentTransfers {
		if _, ok := h.paymentAccounts[paymentTransfer.AccountID]; !ok {
			return wlog.Errorf("invalid paymentaccount")
		}
	}
	return nil
}

func (h *orderHandler) updatePaymentTransfers(ctx context.Context) error {
	for _, _paymentTransfer := range h.PaymentTransfers {
		paymentTransfer := &types.XPaymentTransfer{
			PaymentTransferID: _paymentTransfer.EntID,
			CoinTypeID:        _paymentTransfer.CoinTypeID,
			AccountID:         _paymentTransfer.AccountID,
			Amount:            _paymentTransfer.Amount,
			StartAmount:       _paymentTransfer.StartAmount,
		}

		paymentCoin, ok := h.paymentTransferCoins[paymentTransfer.CoinTypeID]
		if !ok {
			return wlog.Errorf("invalid paymentcoin")
		}
		paymentAccount, ok := h.paymentAccounts[paymentTransfer.AccountID]
		if !ok {
			return wlog.Errorf("invalid paymentaccount")
		}

		balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    paymentCoin.Name,
			Address: paymentAccount.Address,
		})
		if err != nil {
			return wlog.WrapError(err)
		}
		if balance == nil {
			return wlog.Errorf("invalid balance")
		}

		bal, err := decimal.NewFromString(balance.BalanceStr)
		if err != nil {
			return wlog.WrapError(err)
		}
		paymentTransfer.PaymentAccountBalance = bal.String()
		startAmount, err := decimal.NewFromString(paymentTransfer.StartAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		paymentTransfer.IncomingAmount = func() *string {
			amount := bal.Sub(startAmount)
			if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
				return nil
			}
			s := amount.String()
			return &s
		}()
		// If we could be here, we should have enough balance
		if paymentTransfer.IncomingAmount != nil {
			paymentTransfer.IncomingExtra = fmt.Sprintf(
				`{"PaymentID": "%v","OrderID":"%v","PaymentState":"%v","GoodID":"%v","AppGoodID":"%v"}`,
				h.PaymentID,
				h.OrderID,
				h.PaymentState,
				h.GoodID,
				h.AppGoodID,
			)
			paymentTransfer.OutcomingExtra = fmt.Sprintf(
				`{"PaymentID":"%v","OrderID": "%v","FromTransfer":true,"GoodID":"%v","AppGoodID":"%v","PaymentType":"%v"}`,
				h.PaymentID,
				h.OrderID,
				h.GoodID,
				h.AppGoodID,
				h.PaymentType,
			)
		}
		paymentTransfer.FinishAmount = balance.BalanceStr

		h.paymentTransfers = append(h.paymentTransfers, paymentTransfer)
	}
	return nil
}

func (h *orderHandler) updatePaymentFiats() {
	for _, paymentFiat := range h.PaymentFiats {
		h.xPaymentFiats = append(h.xPaymentFiats, &types.XPaymentFiat{
			PaymentFiatID: paymentFiat.EntID,
			FiatID:        paymentFiat.FiatID,
			Amount:        paymentFiat.Amount,
			Extra: fmt.Sprintf(
				`{"PaymentID":"%v","OrderID":"%v","GoodID":"%v","AppGoodID":"%v","PaymentType":"%v"}`,
				h.PaymentID,
				h.OrderID,
				h.GoodID,
				h.AppGoodID,
				h.PaymentType,
			),
		})
	}
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"SubscriptionOrder", h.SubscriptionOrder,
			"PaymentTransfers", h.paymentTransfers,
			"Error", *err,
		)
	}

	persistentOrder := &types.PersistentOrder{
		SubscriptionOrder: h.SubscriptionOrder,
		XPaymentTransfers: h.paymentTransfers,
		XPaymentFiats:     h.xPaymentFiats,
		Error:             *err,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getPaymentCoins(ctx); err != nil {
		return err
	}
	if err = h.getPaymentFiats(ctx); err != nil {
		return err
	}
	if err = h.getPaymentAccounts(ctx); err != nil {
		return err
	}
	if err = h.updatePaymentTransfers(ctx); err != nil {
		return err
	}
	h.updatePaymentFiats()
	return nil
}
