//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbenefitaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/goodbenefit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodbenefitaccountmw "github.com/NpoolPlatform/kunman/middleware/account/goodbenefit"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetGoodCoinBenefitAccounts(ctx context.Context, goodID string, coinTypeIDs []string) (map[string]*goodbenefitaccountmwpb.Account, error) {
	for _, coinTypeID := range coinTypeIDs {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &goodbenefitaccountmwpb.Conds{
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: goodID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
		Backup:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := goodbenefitaccountmw.NewHandler(
		ctx,
		goodbenefitaccountmw.WithConds(conds),
		goodbenefitaccountmw.WithOffset(0),
		goodbenefitaccountmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	goodBenefitAccounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	goodBenefitAccountMap := map[string]*goodbenefitaccountmwpb.Account{}
	for _, goodBenefitAccount := range goodBenefitAccounts {
		goodBenefitAccountMap[goodBenefitAccount.CoinTypeID] = goodBenefitAccount
	}
	return goodBenefitAccountMap, nil
}
