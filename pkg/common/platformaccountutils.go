//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	platformaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	platformaccountmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetCoinPlatformAccounts(ctx context.Context, usedFor basetypes.AccountUsedFor, coinTypeIDs []string) (map[string]*platformaccountmwpb.Account, error) {
	for _, coinTypeID := range coinTypeIDs {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &platformaccountmwpb.Conds{
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
		UsedFor:     &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(usedFor)},
		Backup:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := platformaccountmw.NewHandler(
		ctx,
		platformaccountmw.WithConds(conds),
		platformaccountmw.WithOffset(0),
		platformaccountmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	platformAccounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	platformAccountMap := map[string]*platformaccountmwpb.Account{}
	for _, platformAccount := range platformAccounts {
		platformAccountMap[platformAccount.CoinTypeID] = platformAccount
	}
	return platformAccountMap, nil
}
