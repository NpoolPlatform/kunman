package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/setrevenueaddress/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderbenefitmw "github.com/NpoolPlatform/kunman/middleware/account/orderbenefit"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	orderusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder

	appPowerRental       *powerrentalgoodmwpb.PowerRental
	orderbenefitAccounts map[string]*orderbenefit.Account
	powerRentalOrderReq  *powerrentalordermwpb.PowerRentalOrderReq
	nextState            ordertypes.OrderState

	coinTypeIDs   []string
	orderUserReqs []*orderusermwpb.OrderUserReq
	persistent    chan interface{}
	done          chan interface{}
	notif         chan interface{}
}

func (h *orderHandler) getAppPowerRental(ctx context.Context) error {
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithAppGoodID(&h.AppGoodID, true),
	)
	if err != nil {
		return err
	}

	good, err := handler.GetPowerRental(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if good == nil {
		return wlog.Errorf("invalid powerrental")
	}
	h.appPowerRental = good
	return nil
}

func (h *orderHandler) getOrderBenefits(ctx context.Context) error {
	conds := &orderbenefit.Conds{
		OrderID: &v1.StringVal{Op: cruder.EQ, Value: h.OrderID},
	}
	handler, err := orderbenefitmw.NewHandler(
		ctx,
		orderbenefitmw.WithConds(conds),
		orderbenefitmw.WithOffset(0),
		orderbenefitmw.WithLimit(0),
	)
	if err != nil {
		return err
	}

	accounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.orderbenefitAccounts = make(map[string]*orderbenefit.Account)
	for _, acc := range accounts {
		h.orderbenefitAccounts[acc.CoinTypeID] = acc
	}
	return nil
}

func (h *orderHandler) checkAppPowerRental() error {
	if h.appPowerRental == nil {
		return wlog.Errorf("invalid powerrental")
	}
	if h.appPowerRental.State != goodtypes.GoodState_GoodStateReady {
		return wlog.Errorf("powerrental good not ready")
	}
	return nil
}

func (h *orderHandler) getCoinTypeIDs() error {
	for _, goodCoin := range h.appPowerRental.GoodCoins {
		h.coinTypeIDs = append(h.coinTypeIDs, goodCoin.CoinTypeID)
	}

	if len(h.coinTypeIDs) == 0 {
		return wlog.Errorf("have no goodcoins")
	}
	return nil
}

func (h *orderHandler) validatePoolOrderUserID(ctx context.Context) error {
	if h.PoolOrderUserID == nil {
		return wlog.Errorf("invalid poolorderuserid")
	}

	handler, err := orderusermw.NewHandler(
		ctx,
		orderusermw.WithEntID(h.PoolOrderUserID, true),
	)
	if err != nil {
		return err
	}

	info, err := handler.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poolorderuserid")
	}
	return nil
}

func (h *orderHandler) constructOrderUserReqs() error {
	h.orderUserReqs = []*orderusermwpb.OrderUserReq{}
	autoPay := true
	for _, coinTypeID := range h.coinTypeIDs {
		acc, ok := h.orderbenefitAccounts[coinTypeID]
		if !ok {
			return wlog.Errorf("cannot find orderbenefit account for cointypeid: %v", coinTypeID)
		}
		h.orderUserReqs = append(h.orderUserReqs, &orderusermwpb.OrderUserReq{
			EntID:          h.PoolOrderUserID,
			RevenueAddress: &acc.Address,
			CoinTypeID:     &coinTypeID,
			AutoPay:        &autoPay,
		})
	}
	return nil
}

func (h *orderHandler) constructPowerRentalOrderReq() {
	h.powerRentalOrderReq = &powerrentalordermwpb.PowerRentalOrderReq{
		ID:         &h.ID,
		OrderState: &h.nextState,
	}
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"AdminSetCanceled", h.AdminSetCanceled,
			"UserSetCanceled", h.UserSetCanceled,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder:    h.PowerRentalOrder,
		OrderUserReqs:       h.orderUserReqs,
		PowerRentalOrderReq: h.powerRentalOrderReq,
		AppGoodStockLockID:  &h.AppGoodStockLockID,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
	}
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.nextState = ordertypes.OrderState_OrderStateInService

	var err error
	defer h.final(ctx, &err)

	if h.GoodStockMode != goodtypes.GoodStockMode_GoodStockByMiningPool {
		h.constructPowerRentalOrderReq()
		return nil
	}

	if err = h.getAppPowerRental(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkAppPowerRental(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.getCoinTypeIDs(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.getOrderBenefits(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.validatePoolOrderUserID(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.constructOrderUserReqs(); err != nil {
		return wlog.WrapError(err)
	}

	h.constructPowerRentalOrderReq()

	return nil
}
