package sentinel

import (
	"context"

	withdrawmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/transferring/types"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		withdraws, _, err := withdrawmwcli.GetWithdraws(ctx, &withdrawmwpb.Conds{
			State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.WithdrawState_Transferring)},
		}, offset, limit)
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

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if withdraw, ok := ent.(*types.PersistentWithdraw); ok {
		return withdraw.EntID
	}
	return ent.(*withdrawmwpb.Withdraw).EntID
}
