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
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/types"
)

type handler struct{}

func NewReward() basereward.Rewarder {
	return &handler{}
}

func (p *handler) rewardDeposit(_account *types.PersistentAccount) {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       _account.AppID,
			UserID:      _account.UserID,
			EventType:   basetypes.UsedFor_DepositReceived,
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
			"rewardDeposit",
			"AppID", _account.AppID,
			"UserID", _account.UserID,
			"Error", err,
		)
	}
}

func (p *handler) Update(ctx context.Context, account interface{}, notif, done chan interface{}) error {
	_account, ok := account.(*types.PersistentAccount)
	if !ok {
		return fmt.Errorf("invalid account")
	}

	defer asyncfeed.AsyncFeed(ctx, _account, done)

	p.rewardDeposit(_account)

	return nil
}
