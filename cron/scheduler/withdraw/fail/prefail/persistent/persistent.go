package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/fail/prefail/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, withdraw interface{}, reward, notif, done chan interface{}) error {
	_withdraw, ok := withdraw.(*types.PersistentWithdraw)
	if !ok {
		return fmt.Errorf("invalid withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _withdraw, done)

	state := ledgertypes.WithdrawState_ReturnFailBalance

	handler, err := withdrawmw.NewHandler(
		ctx,
		withdrawmw.WithID(&_withdraw.ID, true),
		withdrawmw.WithState(&state, true),
	)
	if err != nil {
		return err
	}

	if _, err := handler.UpdateWithdraw(ctx); err != nil {
		return err
	}

	return nil
}
