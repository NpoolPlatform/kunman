package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/unlockaccount/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent      chan interface{}
	notif           chan interface{}
	done            chan interface{}
	paymentAccounts map[string]*paymentaccountmwpb.Account
}

func (h *orderHandler) payWithTransfer() bool {
	return len(h.PaymentTransfers) > 0
}

func (h *orderHandler) checkUnlockable() bool {
	return h.payWithTransfer()
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

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"FeeOrder", h.FeeOrder,
			"PaymentAccounts", h.paymentAccounts,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		FeeOrder: h.FeeOrder,
		PaymentAccountIDs: func() (ids []uint32) {
			for _, paymentAccount := range h.paymentAccounts {
				ids = append(ids, paymentAccount.ID)
			}
			return
		}(),
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

	if able := h.checkUnlockable(); !able {
		return nil
	}
	if err = h.getPaymentAccounts(ctx); err != nil {
		return err
	}

	return nil
}
