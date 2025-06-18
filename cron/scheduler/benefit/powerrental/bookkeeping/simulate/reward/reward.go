package reward

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basereward "github.com/NpoolPlatform/kunman/cron/scheduler/base/reward"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/simulate/types"
)

type handler struct{}

func NewReward() basereward.Rewarder {
	return &handler{}
}

func (p *handler) rewardProfit(good *types.PersistentGood) {
	// TODO
}

func (p *handler) Update(ctx context.Context, good interface{}, notif, done chan interface{}) error {
	_good, ok := good.(*types.PersistentGood)
	if !ok {
		return fmt.Errorf("invalid good")
	}

	defer asyncfeed.AsyncFeed(ctx, _good, done)

	p.rewardProfit(_good)

	return nil
}
