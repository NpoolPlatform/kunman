package reward

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basereward "github.com/NpoolPlatform/kunman/cron/scheduler/base/reward"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/unlockaccount/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

type handler struct{}

func NewReward() basereward.Rewarder {
	return &handler{}
}

func (p *handler) rewardOrderCompleted(order *types.PersistentOrder) {
	// TODO
}

func (p *handler) rewardFirstOrderCompleted(order *types.PersistentOrder) {
	if order.ExistOrderCompletedHistory {
		return
	}

	// TODO
}

func (p *handler) Update(ctx context.Context, order interface{}, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if _order.OrderType == ordertypes.OrderType_Normal {
		p.rewardFirstOrderCompleted(_order)
		p.rewardOrderCompleted(_order)
	}

	return nil
}
