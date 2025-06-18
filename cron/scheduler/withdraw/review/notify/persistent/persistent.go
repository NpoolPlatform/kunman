package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/review/notify/types"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, notify interface{}, reward, notif, done chan interface{}) error {
	_notify, ok := notify.(*types.PersistentWithdrawReviewNotify)
	if !ok {
		return fmt.Errorf("invalid notify")
	}

	asyncfeed.AsyncFeed(ctx, _notify, done)

	return nil
}
