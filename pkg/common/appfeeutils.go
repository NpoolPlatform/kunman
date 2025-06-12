//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appfeemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"

	"github.com/google/uuid"
)

func GetAppFees(ctx context.Context, appGoodIDs []string) (map[string]*appfeemwpb.Fee, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	appFees, _, err := appfeemwcli.GetFees(ctx, &appfeemwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}, int32(0), int32(len(appGoodIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appFeeMap := map[string]*appfeemwpb.Fee{}
	for _, appFee := range appFees {
		appFeeMap[appFee.AppGoodID] = appFee
	}
	return appFeeMap, nil
}
