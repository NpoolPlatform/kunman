package persistent

import (
	"context"
	"fmt"

	withdrawmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/transferring/types"
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

	if _, err := withdrawmwcli.UpdateWithdraw(ctx, &withdrawmwpb.WithdrawReq{
		ID:                 &_withdraw.ID,
		State:              &_withdraw.NewWithdrawState,
		ChainTransactionID: &_withdraw.ChainTxID,
	}); err != nil {
		return err
	}

	return nil
}
