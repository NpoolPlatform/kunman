package sentinel

import (
	"context"
	"time"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/finish/types"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
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

	conds := &depositaccmwpb.Conds{
		Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		LockedBy:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.AccountLockedBy_Collecting)},
		ScannableAt: &basetypes.Uint32Val{Op: cruder.LT, Value: uint32(time.Now().Unix())},
	}

	for {
		handler, err := depositaccmw.NewHandler(
			ctx,
			depositaccmw.WithConds(conds),
			depositaccmw.WithOffset(offset),
			depositaccmw.WithLimit(limit),
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
	return ent.(*depositaccmwpb.Account).EntID
}
