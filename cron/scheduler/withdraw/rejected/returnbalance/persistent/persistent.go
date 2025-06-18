package persistent

import (
	"context"
	"fmt"

	withdrawmwcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/withdraw"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/rejected/returnbalance/types"
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

	state := ledgertypes.WithdrawState_Rejected
	if _, err := withdrawmwcli.UpdateWithdraw(ctx, &withdrawmwpb.WithdrawReq{
		ID:    &_withdraw.ID,
		State: &state,
	}); err != nil {
		return err
	}

	return nil
}
