package appsimulatepowerrental

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	apppowerrental1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/powerrental"
	devicetype1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device"
	manufacturer1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device/manufacturer"
	powerrental1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/powerrental"
	vendorbrand1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/vender/location"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Simulate{
	EntID:                uuid.NewString(),
	AppID:                uuid.NewString(),
	GoodID:               uuid.NewString(),
	GoodName:             uuid.NewString(),
	AppGoodID:            uuid.NewString(),
	AppGoodName:          uuid.NewString(),
	OrderUnits:           "10",
	OrderDurationSeconds: 86400,
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	manufacturerID := uuid.NewString()
	h1, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		manufacturer1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	deviceTypeID := uuid.NewString()
	h2, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&deviceTypeID, true),
		devicetype1.WithType(func() *string { s := uuid.NewString(); return &s }(), true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
		devicetype1.WithPowerConsumption(func() *uint32 { u := uint32(100); return &u }(), true),
		devicetype1.WithShipmentAt(func() *uint32 { u := uint32(100); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h2.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	brandID := uuid.NewString()
	h3, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithEntID(&brandID, true),
		vendorbrand1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorbrand1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h3.CreateBrand(context.Background())
	assert.Nil(t, err)

	vendorLocationID := uuid.NewString()
	h4, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&vendorLocationID, true),
		vendorlocation1.WithBrandID(&brandID, true),
		vendorlocation1.WithCountry(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithProvince(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithCity(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithAddress(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h4.CreateLocation(context.Background())
	assert.Nil(t, err)

	powerRentalID := uuid.NewString()
	h5, err := powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithEntID(&powerRentalID, true),
		powerrental1.WithGoodID(&ret.GoodID, true),
		powerrental1.WithDeviceTypeID(&deviceTypeID, true),
		powerrental1.WithVendorLocationID(&vendorLocationID, true),
		powerrental1.WithUnitPrice(func() *string { s := decimal.NewFromInt(100).String(); return &s }(), true),
		powerrental1.WithQuantityUnit(func() *string { s := "TiB"; return &s }(), true),
		powerrental1.WithQuantityUnitAmount(func() *string { s := decimal.NewFromInt(100).String(); return &s }(), true),
		powerrental1.WithDeliveryAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		powerrental1.WithUnitLockDeposit(func() *string { s := decimal.NewFromInt(90).String(); return &s }(), true),
		powerrental1.WithDurationDisplayType(func() *types.GoodDurationType { e := types.GoodDurationType_GoodDurationByDay; return &e }(), true),
		powerrental1.WithGoodType(func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(), true),
		powerrental1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		powerrental1.WithName(&ret.GoodName, true),
		powerrental1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		powerrental1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		powerrental1.WithStockMode(func() *types.GoodStockMode { e := types.GoodStockMode_GoodStockByUnique; return &e }(), true),
		powerrental1.WithTotal(func() *string { s := decimal.NewFromInt(120).String(); return &s }(), true),
		powerrental1.WithTestOnly(func() *bool { b := true; return &b }(), true),
		powerrental1.WithBenefitIntervalHours(func() *uint32 { n := uint32(24); return &n }(), true),
		powerrental1.WithPurchasable(func() *bool { b := true; return &b }(), true),
		powerrental1.WithOnline(func() *bool { b := true; return &b }(), true),
	)
	assert.Nil(t, err)

	err = h5.CreatePowerRental(context.Background())
	assert.Nil(t, err)

	h6, err := apppowerrental1.NewHandler(
		context.Background(),
		apppowerrental1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		apppowerrental1.WithAppID(&ret.AppID, true),
		apppowerrental1.WithGoodID(&ret.GoodID, true),
		apppowerrental1.WithAppGoodID(&ret.AppGoodID, true),
		apppowerrental1.WithPurchasable(func() *bool { b := true; return &b }(), true),
		apppowerrental1.WithEnableProductPage(func() *bool { b := true; return &b }(), true),
		apppowerrental1.WithProductPage(func() *string { s := uuid.NewString(); return &s }(), true),
		apppowerrental1.WithName(&ret.AppGoodName, true),
		apppowerrental1.WithOnline(func() *bool { b := true; return &b }(), true),
		apppowerrental1.WithVisible(func() *bool { b := true; return &b }(), true),
		apppowerrental1.WithDisplayIndex(func() *int32 { n := int32(24); return &n }(), true),
		apppowerrental1.WithBanner(func() *string { s := uuid.NewString(); return &s }(), true),
		apppowerrental1.WithServiceStartAt(func() *uint32 { n := uint32(24); return &n }(), true),
		apppowerrental1.WithCancelMode(func() *types.CancelMode { e := types.CancelMode_Uncancellable; return &e }(), true),
		apppowerrental1.WithCancelableBeforeStartSeconds(func() *uint32 { n := uint32(24); return &n }(), true),
		apppowerrental1.WithEnableSetCommission(func() *bool { b := true; return &b }(), true),
		apppowerrental1.WithMinOrderAmount(func() *string { s := decimal.NewFromInt(10).String(); return &s }(), true),
		apppowerrental1.WithMaxOrderAmount(func() *string { s := decimal.NewFromInt(90).String(); return &s }(), true),
		apppowerrental1.WithMaxUserAmount(func() *string { s := decimal.NewFromInt(90).String(); return &s }(), true),
		apppowerrental1.WithMinOrderDurationSeconds(func() *uint32 { n := uint32(86400); return &n }(), true),
		apppowerrental1.WithMaxOrderDurationSeconds(func() *uint32 { n := uint32(86400); return &n }(), true),
		apppowerrental1.WithUnitPrice(func() *string { s := decimal.NewFromInt(100).String(); return &s }(), true),
		apppowerrental1.WithSaleStartAt(func() *uint32 { n := uint32(24); return &n }(), true),
		apppowerrental1.WithSaleEndAt(func() *uint32 { n := uint32(24); return &n }(), true),
		apppowerrental1.WithSaleMode(func() *types.GoodSaleMode { e := types.GoodSaleMode_GoodSaleModeMainnetSpot; return &e }(), true),
		apppowerrental1.WithFixedDuration(func() *bool { b := false; return &b }(), true),
		apppowerrental1.WithPackageWithRequireds(func() *bool { b := false; return &b }(), true),
	)
	assert.Nil(t, err)

	err = h6.CreatePowerRental(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h6.DeletePowerRental(context.Background())
		_ = h5.DeletePowerRental(context.Background())
		_ = h4.DeleteLocation(context.Background())
		_ = h3.DeleteBrand(context.Background())
		_ = h2.DeleteDeviceType(context.Background())
		_ = h1.DeleteManufacturer(context.Background())
	}
}

func createSimulate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderUnits(&ret.OrderUnits, true),
		WithOrderDurationSeconds(&ret.OrderDurationSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateSimulate(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetSimulate(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateSimulate(t *testing.T) {
	ret.OrderUnits = "20"
	ret.OrderDurationSeconds = 86400
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderUnits(&ret.OrderUnits, true),
		WithOrderDurationSeconds(&ret.OrderDurationSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateSimulate(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetSimulate(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getSimulate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetSimulate(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getSimulates(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetSimulates(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteSimulate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteSimulate(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetSimulate(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestSimulate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createSimulate", createSimulate)
	t.Run("updateSimulate", updateSimulate)
	t.Run("getSimulate", getSimulate)
	t.Run("getSimulates", getSimulates)
	t.Run("deleteSimulate", deleteSimulate)
}
