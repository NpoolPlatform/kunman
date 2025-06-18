package persistent

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/checkhashrate/types"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, good interface{}, reward, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentGoodPowerRental)
	if !ok {
		return fmt.Errorf("invalid feeorder")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(&_good.ID, true),
		powerrentalmw.WithEntID(&_good.EntID, true),
		powerrentalmw.WithGoodID(&_good.GoodID, true),
		powerrentalmw.WithState(v1.GoodState_GoodStateReady.Enum(), true),
		powerrentalmw.WithStocks(_good.MiningGoodStockReqs, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}
