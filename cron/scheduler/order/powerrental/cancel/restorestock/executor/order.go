package executor

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/logger"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/restorestock/types"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent     chan interface{}
	notif          chan interface{}
	done           chan interface{}
	appPowerRental *apppowerrentalmwpb.PowerRental
}

func (h *orderHandler) getAppPowerRental(ctx context.Context) error {
	good, err := apppowerrentalmwcli.GetPowerRental(ctx, h.AppGoodID)
	if err != nil {
		return err
	}
	if good == nil {
		return fmt.Errorf("invalid powerrental")
	}
	h.appPowerRental = good
	return nil
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"AppPowerRental", h.appPowerRental,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder:   h.PowerRentalOrder,
		AppGoodStockLockID: h.AppGoodStockLockID,
	}
	if h.appPowerRental != nil {
		persistentOrder.AppGoodStockID = h.appPowerRental.AppGoodStockID
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

	if err = h.getAppPowerRental(ctx); err != nil {
		return err
	}

	return nil
}
