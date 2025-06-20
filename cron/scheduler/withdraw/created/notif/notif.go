package notif

import (
	"context"
	"fmt"

	basenotif "github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	retry1 "github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/created/types"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type handler struct{}

func NewNotif() basenotif.Notify {
	return &handler{}
}

func (p *handler) notifyWithdraw(withdraw *types.PersistentWithdraw) error {
	return pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		return publisher.Update(
			basetypes.MsgID_WithdrawRequestReq.String(),
			nil,
			nil,
			nil,
			withdraw.Withdraw,
		)
	})
}

func (p *handler) Notify(ctx context.Context, withdraw interface{}, retry chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}
	if err := p.notifyWithdraw(_withdraw); err != nil {
		retry1.Retry(_withdraw.EntID, _withdraw, retry)
		return err
	}
	return nil
}
