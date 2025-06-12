//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	appgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAppGoods(ctx context.Context, appGoodIDs []string) (map[string]*appgoodmwpb.Good, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &appgoodmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}
	handler, err := appgoodmw.NewHandler(
		ctx,
		appgoodmw.WithConds(conds),
		appgoodmw.WithOffset(0),
		appgoodmw.WithLimit(int32(len(appGoodIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appGoods, _, err := handler.GetGoods(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appGoodMap := map[string]*appgoodmwpb.Good{}
	for _, appGood := range appGoods {
		appGoodMap[appGood.EntID] = appGood
	}
	return appGoodMap, nil
}
