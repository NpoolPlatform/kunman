package notif

import (
	"context"
	"fmt"

	basenotif "github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	retry1 "github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/review/notify/types"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	withdrawreviewnotifypb "github.com/NpoolPlatform/kunman/message/scheduler/middleware/v1/withdraw/review/notify"

	"github.com/google/uuid"
)

type handler struct{}

func NewNotif() basenotif.Notify {
	return &handler{}
}

func (p *handler) notifyWithdrawReview(notify *types.PersistentWithdrawReviewNotify) error {
	return pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		return publisher.Update(
			basetypes.MsgID_WithdrawReviewNotifyReq.String(),
			nil,
			nil,
			nil,
			&withdrawreviewnotifypb.MsgWithdrawReviewNotifyReq{
				AppWithdraws: notify.AppWithdraws,
			},
		)
	})
}

func (p *handler) Notify(ctx context.Context, notify interface{}, retry chan interface{}) error {
	_notify, ok := notify.(*types.PersistentWithdrawReviewNotify)
	if !ok {
		return fmt.Errorf("invalid notify")
	}
	if err := p.notifyWithdrawReview(_notify); err != nil {
		retry1.Retry(uuid.Nil.String(), _notify, retry)
		return err
	}
	return nil
}
