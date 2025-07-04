package notif

import (
	"context"
	"fmt"
	"time"

	basenotif "github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	retry1 "github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/types"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	notifbenefitmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
)

type handler struct{}

func NewNotif() basenotif.Notify {
	return &handler{}
}

func (p *handler) notifyGoodBenefit(good *types.PersistentPowerRental) error {
	return pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		now := uint32(time.Now().Unix())
		req := &notifbenefitmwpb.GoodBenefitReq{
			GoodID:      &good.GoodID,
			GoodType:    &good.GoodType,
			GoodName:    &good.Name,
			State:       &good.BenefitResult,
			BenefitDate: &now,
		}
		for _, reward := range good.Rewards {
			req.CoinTypeID = &reward.CoinTypeID
			req.Message = func() *string {
				for _, _reward := range good.CoinRewards {
					if reward.CoinTypeID == _reward.CoinTypeID {
						return &_reward.BenefitMessage
					}
				}
				if good.Error != nil {
					s := wlog.Unwrap(good.Error).Error()
					return &s
				}
				return nil
			}()
			if err := publisher.Update(
				basetypes.MsgID_CreateGoodBenefitReq.String(),
				nil,
				nil,
				nil,
				req,
			); err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *handler) Notify(ctx context.Context, good interface{}, retry chan interface{}) error {
	_good, ok := good.(*types.PersistentPowerRental)
	if !ok {
		return fmt.Errorf("invalid good")
	}
	if err := p.notifyGoodBenefit(_good); err != nil {
		retry1.Retry(_good.EntID, _good, retry)
		return err
	}
	return nil
}
