package sentinel

import (
	"context"
	"time"

	depositaccmwcli "github.com/NpoolPlatform/account-middleware/pkg/client/deposit"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/transfer/types"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		accounts, _, err := depositaccmwcli.GetAccounts(ctx, &depositaccmwpb.Conds{
			Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
			ScannableAt: &basetypes.Uint32Val{Op: cruder.LT, Value: uint32(time.Now().Unix())},
		}, offset, limit)
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
