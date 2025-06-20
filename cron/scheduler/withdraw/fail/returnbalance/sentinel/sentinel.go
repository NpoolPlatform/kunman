package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/fail/returnbalance/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	withdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanWithdraws(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &withdrawmwpb.Conds{
		State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.WithdrawState_ReturnFailBalance)},
	}

	for {
		handler, err := withdrawmw.NewHandler(
			ctx,
			withdrawmw.WithConds(conds),
			withdrawmw.WithOffset(offset),
			withdrawmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		withdraws, _, err := handler.GetWithdraws(ctx)
		if err != nil {
			return err
		}
		if len(withdraws) == 0 {
			return nil
		}

		for _, withdraw := range withdraws {
			cancelablefeed.CancelableFeed(ctx, withdraw, exec)
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	return h.scanWithdraws(ctx, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if withdraw, ok := ent.(*withdrawmwpb.Withdraw); ok {
		return withdraw.EntID
	}
	return ent.(*types.PersistentWithdraw).EntID
}
