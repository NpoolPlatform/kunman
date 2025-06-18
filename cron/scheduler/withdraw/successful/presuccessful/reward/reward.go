package reward

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basereward "github.com/NpoolPlatform/kunman/cron/scheduler/base/reward"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/types"
)

type handler struct{}

func NewReward() basereward.Rewarder {
	return &handler{}
}

func (p *handler) rewardWithdraw(_withdraw *types.PersistentWithdraw) {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       _withdraw.AppID,
			UserID:      _withdraw.UserID,
			EventType:   basetypes.UsedFor_WithdrawalCompleted,
			Consecutive: 1,
		}
		return publisher.Update(
			basetypes.MsgID_CalculateEventRewardReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"rewardWithdraw",
			"AppID", _withdraw.AppID,
			"UserID", _withdraw.UserID,
			"Error", err,
		)
	}
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
