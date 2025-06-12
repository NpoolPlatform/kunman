//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAppFees(ctx context.Context, appGoodIDs []string) (map[string]*appfeemwpb.Fee, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &appfeemwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}
	handler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithConds(conds),
		appfeemw.WithOffset(0),
		appfeemw.WithLimit(int32(len(appGoodIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appFees, _, err := handler.GetFees(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appFeeMap := map[string]*appfeemwpb.Fee{}
	for _, appFee := range appFees {
		appFeeMap[appFee.AppGoodID] = appFee
	}
	return appFeeMap, nil
}
