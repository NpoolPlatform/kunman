package persistent

import (
	"context"
	"fmt"

	powerrentalmwcli "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodpowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/wait/types"
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

	return powerrentalmwcli.UpdatePowerRental(ctx, &goodpowerrentalmwpb.PowerRentalReq{
		ID:               &_good.ID,
		EntID:            &_good.EntID,
		GoodID:           &_good.GoodID,
		State:            v1.GoodState_GoodStateCreateGoodUser.Enum(),
		MiningGoodStocks: _good.MiningGoodStockReqs,
		Rollback:         func() *bool { rollback := true; return &rollback }(),
	})
}
