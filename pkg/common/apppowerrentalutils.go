//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"

	"github.com/google/uuid"
)

func GetAppPowerRentals(ctx context.Context, appGoodIDs []string) (map[string]*apppowerrentalmwpb.PowerRental, error) {
	for _, appGoodID := range appGoodIDs {
		if _, err := uuid.Parse(appGoodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	appPowerRentals, _, err := apppowerrentalmwcli.GetPowerRentals(ctx, &apppowerrentalmwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appGoodIDs},
	}, int32(0), int32(len(appGoodIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	appPowerRentalMap := map[string]*apppowerrentalmwpb.PowerRental{}
	for _, appPowerRental := range appPowerRentals {
		appPowerRentalMap[appPowerRental.AppGoodID] = appPowerRental
	}
	return appPowerRentalMap, nil
}
