package executor

import (
	"context"

	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/fail/prefail/types"
)

type withdrawHandler struct {
	*withdrawmwpb.Withdraw
	persistent chan interface{}
}

func (h *withdrawHandler) final(ctx context.Context) {
	persistentWithdraw := &types.PersistentWithdraw{
		Withdraw: h.Withdraw,
	}
	asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.persistent)
}

func (h *withdrawHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
