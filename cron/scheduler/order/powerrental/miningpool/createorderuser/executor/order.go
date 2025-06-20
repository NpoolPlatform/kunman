package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/createorderuser/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	goodusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	"github.com/google/uuid"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	appPowerRental *powerrentalgoodmwpb.PowerRental

	powerRentalOrderReq *powerrentalordermwpb.PowerRentalOrderReq
	orderUserReq        *orderusermwpb.OrderUserReq
	nextState           *ordertypes.OrderState
	poolGoodUserID      *string
	persistent          chan interface{}
	done                chan interface{}
	notif               chan interface{}
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

func (h *orderHandler) getPoolGoodUserID() error {
	if h.appPowerRental == nil {
		return wlog.Errorf("invalid powerrental")
	}

	var miningGoodStockID *string
	var poolGoodOrderID *string
	for _, appMiningGoodStock := range h.appPowerRental.AppMiningGoodStocks {
		if appMiningGoodStock.EntID == h.AppGoodStockID {
			miningGoodStockID = &appMiningGoodStock.MiningGoodStockID
			break
		}
	}

	if miningGoodStockID == nil {
		return wlog.Errorf("cannot find appmininggoodstock, appgoodstockid: %v", h.AppGoodStockID)
	}

	for _, miningGoodStock := range h.appPowerRental.MiningGoodStocks {
		if miningGoodStock.EntID == *miningGoodStockID {
			poolGoodOrderID = &miningGoodStock.PoolGoodUserID
			break
		}
	}

	if poolGoodOrderID == nil {
		return wlog.Errorf("cannot find mininggoodstock, mininggoodstockid: %v", miningGoodStockID)
	}

	h.poolGoodUserID = poolGoodOrderID
	return nil
}

func (h *orderHandler) validatePoolGoodUserID(ctx context.Context) error {
	if h.poolGoodUserID == nil {
		return wlog.Errorf("invalid poolgooduserid")
	}

	handler, err := goodusermw.NewHandler(
		ctx,
		goodusermw.WithEntID(h.poolGoodUserID, true),
	)
	if err != nil {
		return err
	}

	info, err := handler.GetGoodUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poolgooduserid")
	}
	return nil
}

func (h *orderHandler) constructCreateOrderUserReq() {
	h.orderUserReq = &orderusermwpb.OrderUserReq{
		EntID:      func() *string { entID := uuid.NewString(); return &entID }(),
		AppID:      &h.AppID,
		UserID:     &h.UserID,
		GoodUserID: h.poolGoodUserID,
	}
}

func (h *orderHandler) constructUpdatePowerrentalOrderReq() {
	h.powerRentalOrderReq = &powerrentalordermwpb.PowerRentalOrderReq{
		ID:              &h.ID,
		EntID:           &h.EntID,
		PoolOrderUserID: h.orderUserReq.EntID,
		OrderState:      h.nextState,
	}
}

func (h *orderHandler) constructUpdatePowerrentalOrderReqForSkip() {
	h.powerRentalOrderReq = &powerrentalordermwpb.PowerRentalOrderReq{
		ID:         &h.ID,
		EntID:      &h.EntID,
		OrderState: h.nextState,
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
		OrderUserReq:        h.orderUserReq,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
	}
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	h.nextState = ordertypes.OrderState_OrderStateSetProportion.Enum()

	var err error
	defer h.final(ctx, &err)

	if h.GoodStockMode != goodtypes.GoodStockMode_GoodStockByMiningPool {
		h.constructUpdatePowerrentalOrderReqForSkip()
		return nil
	}

	if err = h.getAppPowerRental(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.checkAppPowerRental(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.getPoolGoodUserID(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.validatePoolGoodUserID(ctx); err != nil {
		return wlog.WrapError(err)
	}

	h.constructCreateOrderUserReq()
	h.constructUpdatePowerrentalOrderReq()

	return nil
}
