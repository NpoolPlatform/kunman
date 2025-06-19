package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/finish/types"
	payaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	payaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &payaccmwpb.Conds{
		Locked:   &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		LockedBy: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.AccountLockedBy_Collecting)},
	}

	for {
		handler, err := payaccmw.NewHandler(
			ctx,
			payaccmw.WithConds(conds),
			payaccmw.WithOffset(offset),
			payaccmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		accounts, _, err := handler.GetAccounts(ctx)
		if err != nil {
			return err
		}
		if len(accounts) == 0 {
			return nil
		}

		for _, account := range accounts {
			cancelablefeed.CancelableFeed(ctx, account, exec)
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
	if account, ok := ent.(*types.PersistentAccount); ok {
		return account.EntID
	}
	return ent.(*payaccmwpb.Account).EntID
}
