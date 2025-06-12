//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
)

func GetAppGoods(ctx context.Context, appGoodIDs []string) (map[string]*appgoodmwpb.Good, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	appGoods, _, err := appgoodmwcli.GetGoods(ctx, &appgoodmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}, int32(0), int32(len(appGoodIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appGoodMap := map[string]*appgoodmwpb.Good{}
	for _, appGood := range appGoods {
		appGoodMap[appGood.EntID] = appGood
	}
	return appGoodMap, nil
}
