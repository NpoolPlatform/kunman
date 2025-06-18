package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/unlockaccount/types"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent      chan interface{}
	notif           chan interface{}
	done            chan interface{}
	paymentAccounts map[string]*paymentaccountmwpb.Account
}

func (h *orderHandler) payWithTransfer() bool {
	return len(h.PaymentTransfers) > 0
}

func (h *orderHandler) checkUnlockable() bool {
	if !h.payWithTransfer() {
		return false
	}
	switch h.CancelState {
	case ordertypes.OrderState_OrderStateWaitPayment:
	case ordertypes.OrderState_OrderStatePaymentTimeout:
	default:
		return false
	}
	return true
}

func (h *orderHandler) getPaymentAccounts(ctx context.Context) (err error) {
	h.paymentAccounts, err = schedcommon.GetPaymentAccounts(ctx, func() (accountIDs []string) {
		for _, paymentTransfer := range h.PaymentTransfers {
			accountIDs = append(accountIDs, paymentTransfer.AccountID)
		}
		return
	}())
	return wlog.WrapError(err)
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"PaymentAccounts", h.paymentAccounts,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
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
