package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	apppoolmwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/app/pool"
	rootusermwpb "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	apppoolmwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/app/pool"
	rootusermwcli "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"

	"github.com/google/uuid"
)

type CreateHander struct {
	*Handler
	powerRental *powerrentalmwpb.PowerRental
}

func (h *CreateHander) getPowerRental(ctx context.Context) (err error) {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	h.powerRental, err = powerrentalmwcli.GetPowerRental(ctx, *h.GoodID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *CreateHander) checkAppPoolAuth(ctx context.Context) error {
	if h.powerRental.StockMode != goodtypes.GoodStockMode_GoodStockByMiningPool {
		return nil
	}

	if h.powerRental == nil || h.AppID == nil {
		return wlog.Errorf("cannot check auth")
	}

	if h.powerRental.State == goodtypes.GoodState_DefaultGoodState ||
		h.powerRental.State == goodtypes.GoodState_GoodStatePreWait {
		return wlog.Errorf("cannot auth to app, wait goodstate allow")
	}

	if len(h.powerRental.MiningGoodStocks) == 0 {
		return nil
	}

	rootuserIDs := []string{}
	for _, miningGoodStock := range h.powerRental.MiningGoodStocks {
		rootuserIDs = append(rootuserIDs, miningGoodStock.PoolRootUserID)
	}

	rUsers, _, err := rootusermwcli.GetRootUsers(ctx,
		&rootusermwpb.Conds{
			EntIDs: &v1.StringSliceVal{
				Op:    cruder.IN,
				Value: rootuserIDs,
			},
		},
		0,
		int32(len(rootuserIDs)))
	if err != nil {
		return wlog.WrapError(err)
	}

	poolIDs := []string{}
	for _, rUser := range rUsers {
		poolIDs = append(poolIDs, rUser.PoolID)
	}

	for _, poolID := range poolIDs {
		appPools, _, err := apppoolmwcli.GetPools(ctx,
			&apppoolmwpb.Conds{
				AppID:  &v1.StringVal{Op: cruder.EQ, Value: *h.AppID},
				PoolID: &v1.StringVal{Op: cruder.EQ, Value: poolID},
			}, 0, 1)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(appPools) == 0 {
			return wlog.Errorf("have no permission for poolid: %v,appid: %v", poolID, *h.AppID)
		}
	}

	return nil
}

// TODO: check start mode with power rental start mode
func (h *Handler) CreatePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.FixedDuration == nil || *h.FixedDuration {
		if h.MaxOrderDurationSeconds != nil && *h.MinOrderDurationSeconds != *h.MaxOrderDurationSeconds {
			return nil, wlog.Errorf("invalid maxorderdurationseconds")
		}
		h.MaxOrderDurationSeconds = h.MinOrderDurationSeconds
	}

	createH := &CreateHander{Handler: h}

	if err := createH.getPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := createH.checkAppPoolAuth(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := apppowerrentalmwcli.CreatePowerRental(ctx, &apppowerrentalmwpb.PowerRentalReq{
		EntID:                        h.EntID,
		AppID:                        h.AppID,
		GoodID:                       h.GoodID,
		AppGoodID:                    h.AppGoodID,
		Purchasable:                  h.Purchasable,
		EnableProductPage:            h.EnableProductPage,
		ProductPage:                  h.ProductPage,
		Online:                       h.Online,
		Visible:                      h.Visible,
		Name:                         h.Name,
		DisplayIndex:                 h.DisplayIndex,
		Banner:                       h.Banner,
		ServiceStartAt:               h.ServiceStartAt,
		CancelMode:                   h.CancelMode,
		CancelableBeforeStartSeconds: h.CancelableBeforeStartSeconds,
		EnableSetCommission:          h.EnableSetCommission,
		MinOrderAmount:               h.MinOrderAmount,
		MaxOrderAmount:               h.MaxOrderAmount,
		MaxUserAmount:                h.MaxUserAmount,
		MinOrderDurationSeconds:      h.MinOrderDurationSeconds,
		MaxOrderDurationSeconds:      h.MaxOrderDurationSeconds,
		UnitPrice:                    h.UnitPrice,
		SaleStartAt:                  h.SaleStartAt,
		SaleEndAt:                    h.SaleEndAt,
		SaleMode:                     h.SaleMode,
		FixedDuration:                h.FixedDuration,
		PackageWithRequireds:         h.PackageWithRequireds,
		StartMode:                    h.StartMode,
	}); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
