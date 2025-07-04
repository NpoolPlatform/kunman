package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/setproportion/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	orderusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	appPowerRental *powerrentalgoodmwpb.PowerRental

	powerRentalOrderReq *powerrentalordermwpb.PowerRentalOrderReq
	orderUserReqs       []*orderusermwpb.OrderUserReq
	nextState           ordertypes.OrderState

	coinTypeIDs []string
	proportion  string
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

func (h *orderHandler) getProportion() error {
	if h.appPowerRental == nil {
		return wlog.Errorf("invalid powerrental")
	}

	var miningGoodStockID *string
	var total *string
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
			total = &miningGoodStock.Total
			break
		}
	}

	if total == nil {
		return wlog.Errorf("cannot find mininggoodstock, mininggoodstockid: %v", miningGoodStockID)
	}

	unitsDec, err := decimal.NewFromString(h.Units)
	if err != nil {
		return wlog.WrapError(err)
	}

	totalDec, err := decimal.NewFromString(*total)
	if err != nil {
		return wlog.WrapError(err)
	}

	percentDec, err := decimal.NewFromString("100")
	if err != nil {
		return wlog.WrapError(err)
	}

	precision := 2
	h.proportion = unitsDec.Mul(percentDec).Div(totalDec).Truncate(int32(precision)).String()
	return nil
}

func (h *orderHandler) constructOrderUserReq() {
	for _, coinTypeID := range h.coinTypeIDs {
		h.orderUserReqs = append(h.orderUserReqs, &orderusermwpb.OrderUserReq{
			EntID:      h.PoolOrderUserID,
			CoinTypeID: &coinTypeID,
			Proportion: &h.proportion,
		})
	}
}

func (h *orderHandler) constructPowerRentalOrderReq() {
	h.powerRentalOrderReq = &powerrentalordermwpb.PowerRentalOrderReq{
		ID:         &h.ID,
		OrderState: &h.nextState,
	}
}

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
		OrderUserReqs:       h.orderUserReqs,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, h.PowerRentalOrder, h.done)
	}
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.nextState = ordertypes.OrderState_OrderStateSetRevenueAddress

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
	if err = h.getProportion(); err != nil {
		return wlog.WrapError(err)
	}

	if err = h.validatePoolOrderUserID(ctx); err != nil {
		return wlog.WrapError(err)
	}

	h.constructOrderUserReq()
	h.constructPowerRentalOrderReq()

	return nil
}
