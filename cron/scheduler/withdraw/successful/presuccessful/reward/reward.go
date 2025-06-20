package reward

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basereward "github.com/NpoolPlatform/kunman/cron/scheduler/base/reward"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/types"
)

type handler struct{}

func NewReward() basereward.Rewarder {
	return &handler{}
}

func (p *handler) rewardWithdraw(_withdraw *types.PersistentWithdraw) {
	// TODO
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	p.rewardWithdraw(_withdraw)

	return nil
}
