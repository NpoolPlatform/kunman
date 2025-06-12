//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAppPowerRentals(ctx context.Context, appGoodIDs []string) (map[string]*apppowerrentalmwpb.PowerRental, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &apppowerrentalmwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithConds(conds),
		apppowerrentalmw.WithOffset(0),
		apppowerrentalmw.WithLimit(int32(len(appGoodIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appPowerRentals, _, err := handler.GetPowerRentals(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appPowerRentalMap := map[string]*apppowerrentalmwpb.PowerRental{}
	for _, appPowerRental := range appPowerRentals {
		appPowerRentalMap[appPowerRental.AppGoodID] = appPowerRental
	}
	return appPowerRentalMap, nil
}
