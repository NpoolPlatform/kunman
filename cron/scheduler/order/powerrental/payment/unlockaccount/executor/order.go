package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/unlockaccount/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	taskconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/task/config"
	taskusermwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/task/user"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	eventmw "github.com/NpoolPlatform/kunman/middleware/inspire/event"
	taskconfigmw "github.com/NpoolPlatform/kunman/middleware/inspire/task/config"
	taskusermw "github.com/NpoolPlatform/kunman/middleware/inspire/task/user"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent                      chan interface{}
	notif                           chan interface{}
	done                            chan interface{}
	paymentAccounts                 map[string]*paymentaccountmwpb.Account
	existOrderCompletedHistory      bool
	existFirstOrderCompletedHistory bool
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

//nolint:dupl
func (h *orderHandler) checkFirstOrderComplatedHistory(ctx context.Context) error {
	eventType := basetypes.UsedFor_FirstOrderCompleted
	eventConds := &eventmwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(eventType)},
	}
	eventHandler, err := eventmw.NewHandler(
		ctx,
		eventmw.WithConds(eventConds),
	)
	if err != nil {
		return err
	}

	ev, err := eventHandler.GetEventOnly(ctx)
	if err != nil {
		return err
	}
	if ev == nil {
		return nil
	}

	tcConds := &taskconfigmwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ev.EntID},
	}
	tcHandler, err := taskconfigmw.NewHandler(
		ctx,
		taskconfigmw.WithConds(tcConds),
	)
	if err != nil {
		return err
	}

	taskConfig, err := tcHandler.GetTaskConfigOnly(ctx)
	if err != nil {
		return err
	}
	if taskConfig == nil {
		return nil
	}

	tuConds := &taskusermwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ev.EntID},
		TaskID:  &basetypes.StringVal{Op: cruder.EQ, Value: taskConfig.EntID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID},
	}
	tuHandler, err := taskusermw.NewHandler(
		ctx,
		taskusermw.WithConds(tuConds),
	)
	if err != nil {
		return err
	}

	existTaskUser, err := tuHandler.ExistTaskUserConds(ctx)
	if err != nil {
		return err
	}
	if existTaskUser {
		h.existFirstOrderCompletedHistory = true
	}

	return nil
}

//nolint:dupl
func (h *orderHandler) checkOrderComplatedHistory(ctx context.Context) error {
	eventType := basetypes.UsedFor_OrderCompleted
	eventConds := &eventmwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(eventType)},
	}
	eventHandler, err := eventmw.NewHandler(
		ctx,
		eventmw.WithConds(eventConds),
	)
	if err != nil {
		return err
	}

	ev, err := eventHandler.GetEventOnly(ctx)
	if err != nil {
		return err
	}
	if ev == nil {
		return nil
	}

	tcConds := &taskconfigmwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ev.EntID},
	}
	tcHandler, err := taskconfigmw.NewHandler(
		ctx,
		taskconfigmw.WithConds(tcConds),
	)
	if err != nil {
		return err
	}

	taskConfig, err := tcHandler.GetTaskConfigOnly(ctx)
	if err != nil {
		return err
	}
	if taskConfig == nil {
		return nil
	}

	tuConds := &taskusermwpb.Conds{
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		EventID: &basetypes.StringVal{Op: cruder.EQ, Value: ev.EntID},
		TaskID:  &basetypes.StringVal{Op: cruder.EQ, Value: taskConfig.EntID},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID},
	}
	tuHandler, err := taskusermw.NewHandler(
		ctx,
		taskusermw.WithConds(tuConds),
	)
	if err != nil {
		return err
	}

	existTaskUser, err := tuHandler.ExistTaskUserConds(ctx)
	if err != nil {
		return err
	}
	if existTaskUser {
		h.existOrderCompletedHistory = true
	}

	return nil
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
	existOrderCompletedHistory := h.existFirstOrderCompletedHistory || h.existOrderCompletedHistory
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		PaymentAccountIDs: func() (ids []uint32) {
			for _, paymentAccount := range h.paymentAccounts {
				ids = append(ids, paymentAccount.ID)
			}
			return
		}(),
		ExistOrderCompletedHistory: existOrderCompletedHistory,
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

	h.existFirstOrderCompletedHistory = false
	h.existOrderCompletedHistory = false
	if err = h.checkFirstOrderComplatedHistory(ctx); err != nil {
		return err
	}
	if err = h.checkOrderComplatedHistory(ctx); err != nil {
		return err
	}

	return nil
}
