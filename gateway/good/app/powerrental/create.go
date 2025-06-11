package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	apppoolmwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/app/pool"
	rootusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	apppoolmw "github.com/NpoolPlatform/kunman/middleware/miningpool/app/pool"
	rootusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/rootuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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

	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return err
	}

	h.powerRental, err = handler.GetPowerRental(ctx)
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

	rootUserIDs := []string{}
	for _, miningGoodStock := range h.powerRental.MiningGoodStocks {
		rootUserIDs = append(rootUserIDs, miningGoodStock.PoolRootUserID)
	}

	handler, err := rootusermw.NewHandler(
		ctx,
		rootusermw.WithConds(
			&rootusermwpb.Conds{
				EntIDs: &v1.StringSliceVal{
					Op:    cruder.IN,
					Value: rootUserIDs,
				},
			},
		),
		rootusermw.WithOffset(0),
		rootusermw.WithLimit(int32(len(rootUserIDs))),
	)
	if err != nil {
		return err
	}
	rootUsers, _, err := handler.GetRootUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	poolIDs := []string{}
	for _, rootUser := range rootUsers {
		poolIDs = append(poolIDs, rootUser.PoolID)
	}

	conds := &apppoolmwpb.Conds{
		AppID:   &v1.StringVal{Op: cruder.EQ, Value: *h.AppID},
		PoolIDs: &v1.StringSliceVal{Op: cruder.EQ, Value: poolIDs},
	}
	poolHandler, err := apppoolmw.NewHandler(
		ctx,
		apppoolmw.WithConds(conds),
		apppoolmw.WithOffset(0),
		apppoolmw.WithLimit(int32(len(poolIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	appPools, _, err := poolHandler.GetPools(ctx)
	if len(appPools) == 0 {
		return wlog.Errorf("Permission denied")
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

	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithEntID(h.EntID, true),
		apppowerrentalmw.WithAppID(h.AppID, true),
		apppowerrentalmw.WithGoodID(h.GoodID, true),
		apppowerrentalmw.WithAppGoodID(h.AppGoodID, true),
		apppowerrentalmw.WithPurchasable(h.Purchasable, true),
		apppowerrentalmw.WithEnableProductPage(h.EnableProductPage, true),
		apppowerrentalmw.WithProductPage(h.ProductPage, true),
		apppowerrentalmw.WithOnline(h.Online, true),
		apppowerrentalmw.WithVisible(h.Visible, true),
		apppowerrentalmw.WithName(h.Name, true),
		apppowerrentalmw.WithDisplayIndex(h.DisplayIndex, true),
		apppowerrentalmw.WithBanner(h.Banner, true),
		apppowerrentalmw.WithServiceStartAt(h.ServiceStartAt, true),
		apppowerrentalmw.WithCancelMode(h.CancelMode, true),
		apppowerrentalmw.WithCancelableBeforeStartSeconds(h.CancelableBeforeStartSeconds, true),
		apppowerrentalmw.WithEnableSetCommission(h.EnableSetCommission, true),
		apppowerrentalmw.WithMinOrderAmount(h.MinOrderAmount, true),
		apppowerrentalmw.WithMaxOrderAmount(h.MaxOrderAmount, true),
		apppowerrentalmw.WithMaxUserAmount(h.MaxUserAmount, true),
		apppowerrentalmw.WithMinOrderDurationSeconds(h.MinOrderDurationSeconds, true),
		apppowerrentalmw.WithMaxOrderDurationSeconds(h.MaxOrderDurationSeconds, true),
		apppowerrentalmw.WithUnitPrice(h.UnitPrice, true),
		apppowerrentalmw.WithSaleStartAt(h.SaleStartAt, true),
		apppowerrentalmw.WithSaleEndAt(h.SaleEndAt, true),
		apppowerrentalmw.WithSaleMode(h.SaleMode, true),
		apppowerrentalmw.WithFixedDuration(h.FixedDuration, true),
		apppowerrentalmw.WithPackageWithRequireds(h.PackageWithRequireds, true),
		apppowerrentalmw.WithStartMode(h.StartMode, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreatePowerRental(ctx); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
