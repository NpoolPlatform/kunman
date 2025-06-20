package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/checkproportion/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	orderusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	appPowerRental      *powerrentalgoodmwpb.PowerRental
	powerRentalOrderReq *powerrentalordermwpb.PowerRentalOrderReq

	coinTypeIDs []string
	nextState   ordertypes.OrderState
	persistent  chan interface{}
	done        chan interface{}
	notif       chan interface{}
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

func (h *orderHandler) checkProportion(ctx context.Context) error {
	for _, coinTypeID := range h.coinTypeIDs {
		handler, err := orderusermw.NewHandler(
			ctx,
			orderusermw.WithEntID(h.PoolOrderUserID, true),
			orderusermw.WithCoinTypeID(&coinTypeID, true),
		)
		if err != nil {
			return err
		}

		proportion, err := handler.GetOrderUserProportion(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		proportionDec, err := decimal.NewFromString(proportion)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !proportionDec.IsZero() {
			return wlog.Errorf("invalid proportion: %v, orderid: %v poolorderuserid: %v, cointypeid: %v",
				proportionDec.String(),
				h.OrderID,
				h.PoolOrderUserID,
				coinTypeID,
			)
		}
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
		PowerRentalOrderReq: h.powerRentalOrderReq,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
	}
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	h.nextState = ordertypes.OrderState_OrderStateRestoreExpiredStock

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

	if err = h.validatePoolOrderUserID(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkProportion(ctx); err != nil {
		return wlog.WrapError(err)
	}

	h.constructPowerRentalOrderReq()

	return nil
}
