package appstock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	devicetype1 "github.com/NpoolPlatform/kunman/middleware/good/device"
	manufacturer1 "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
	powerrental1 "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
	vendorbrand1 "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/stock"
	appmininggoodstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/stock/mining"
	stockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"

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

var ret = npool.Stock{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	GoodID:      uuid.NewString(),
	GoodName:    uuid.NewString(),
	AppGoodID:   uuid.NewString(),
	AppGoodName: uuid.NewString(),
	Reserved:    "100",
	Locked:      decimal.NewFromInt(0).String(),
	InService:   decimal.NewFromInt(0).String(),
	WaitStart:   decimal.NewFromInt(0).String(),
	Sold:        decimal.NewFromInt(0).String(),
	StockMode:   types.GoodStockMode_GoodStockByMiningPool,
	AppMiningGoodStocks: []*appmininggoodstockmwpb.StockInfo{
		{
			EntID:             uuid.NewString(),
			MiningGoodStockID: uuid.NewString(),
			Reserved:          decimal.NewFromInt(0).String(),
		},
		{
			EntID:             uuid.NewString(),
			MiningGoodStockID: uuid.NewString(),
			Reserved:          decimal.NewFromInt(0).String(),
		},
		{
			EntID:             uuid.NewString(),
			MiningGoodStockID: uuid.NewString(),
			Reserved:          decimal.NewFromInt(0).String(),
		},
	},
}

var lockID = uuid.NewString()

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	ret.StockModeStr = ret.StockMode.String()
	for _, stock := range ret.AppMiningGoodStocks {
		stock.AppGoodStockID = ret.EntID
		stock.SpotQuantity = stock.Reserved
		stock.Locked = decimal.NewFromInt(0).String()
		stock.InService = decimal.NewFromInt(0).String()
		stock.WaitStart = decimal.NewFromInt(0).String()
		stock.Sold = decimal.NewFromInt(0).String()
	}

	brandID := uuid.NewString()
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithEntID(&brandID, true),
		vendorbrand1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorbrand1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	locationID := uuid.NewString()
	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&locationID, true),
		vendorlocation1.WithCountry(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithProvince(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithCity(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithAddress(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithBrandID(&brandID, true),
	)
	assert.Nil(t, err)

	err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	manufacturerID := uuid.NewString()
	h31, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		manufacturer1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h31.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	deviceTypeID := uuid.NewString()
	h3, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&deviceTypeID, true),
		devicetype1.WithType(func() *string { s := uuid.NewString(); return &s }(), true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
		devicetype1.WithPowerConsumption(func() *uint32 { u := uint32(100); return &u }(), true),
		devicetype1.WithShipmentAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h3.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	miningGoodStocks := func() (reqs []*stockmwpb.MiningGoodStockReq) {
		for _, stock := range ret.AppMiningGoodStocks {
			reqs = append(reqs, &stockmwpb.MiningGoodStockReq{
				EntID:          &stock.MiningGoodStockID,
				PoolRootUserID: func() *string { s := uuid.NewString(); return &s }(),
				PoolGoodUserID: func() *string { s := uuid.NewString(); return &s }(),
				Total:          func() *string { s := decimal.NewFromInt(3333).String(); return &s }(),
			})
		}
		return
	}()
	h4, err := powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		powerrental1.WithGoodID(&ret.GoodID, true),
		powerrental1.WithDeviceTypeID(&deviceTypeID, true),
		powerrental1.WithVendorLocationID(&locationID, true),
		powerrental1.WithUnitPrice(func() *string { s := decimal.NewFromInt(999).String(); return &s }(), true),
		powerrental1.WithQuantityUnit(func() *string { s := "TiB"; return &s }(), true),
		powerrental1.WithQuantityUnitAmount(func() *string { s := decimal.NewFromInt(9).String(); return &s }(), true),
		powerrental1.WithDeliveryAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		powerrental1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		powerrental1.WithGoodType(func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(), true),
		powerrental1.WithTestOnly(func() *bool { b := true; return &b }(), false),
		powerrental1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
		powerrental1.WithName(&ret.GoodName, true),
		powerrental1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		powerrental1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		powerrental1.WithStockMode(&ret.StockMode, true),
		powerrental1.WithTotal(func() *string { s := decimal.NewFromInt(9999).String(); return &s }(), true),
		powerrental1.WithStocks(miningGoodStocks, true),
		powerrental1.WithUnitLockDeposit(func() *string { s := decimal.NewFromInt(9).String(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h4.CreatePowerRental(context.Background())
	assert.Nil(t, err)

	h5, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&ret.GoodID, true),
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h5.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	appMiningGoodStocks := func() (reqs []*appmininggoodstockmwpb.StockReq) {
		for _, stock := range ret.AppMiningGoodStocks {
			reqs = append(reqs, &appmininggoodstockmwpb.StockReq{
				EntID:             &stock.EntID,
				MiningGoodStockID: &stock.MiningGoodStockID,
				Reserved:          func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
			})
		}
		return
	}()
	h6, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithAppMiningGoodStocks(appMiningGoodStocks, true),
	)
	assert.Nil(t, err)

	err = h6.createStock(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h6.deleteStock(context.Background())
		_ = h5.DeleteGoodBase(context.Background())
		_ = h4.DeletePowerRental(context.Background())
		_ = h3.DeleteDeviceType(context.Background())
		_ = h31.DeleteManufacturer(context.Background())
		_ = h2.DeleteLocation(context.Background())
		_ = h1.DeleteBrand(context.Background())
	}
}

func reserveStock(t *testing.T) {
	ret.SpotQuantity = ret.Reserved
	ret.AppMiningGoodStocks[0].Reserved = ret.Reserved
	ret.AppMiningGoodStocks[0].SpotQuantity = ret.Reserved

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.AppMiningGoodStocks[0].EntID, true), // For mining pool stock, here is the ent_id of mining pool stock
		WithAppGoodID(&ret.AppGoodID, true),
		WithReserved(&ret.Reserved, true),
	)
	if assert.Nil(t, err) {
		err = handler.ReserveStock(context.Background())
		if assert.Nil(t, err) {
			handler, err := NewHandler(
				context.Background(),
				WithEntID(&ret.EntID, true),
			)
			if assert.Nil(t, err) {
				info, err := handler.GetStock(context.Background())
				if assert.Nil(t, err) {
					ret.CreatedAt = info.CreatedAt
					ret.UpdatedAt = info.UpdatedAt
					ret.ID = info.ID
					for i, stock := range info.AppMiningGoodStocks {
						ret.AppMiningGoodStocks[i].ID = stock.ID
					}
					assert.Equal(t, &ret, info)
				}
			}
		}
	}
}

func lockStock(t *testing.T) {
	ret.Locked = decimal.NewFromInt(10).String()
	ret.SpotQuantity = decimal.NewFromInt(90).String()
	ret.AppMiningGoodStocks[0].Locked = ret.Locked
	ret.AppMiningGoodStocks[0].SpotQuantity = ret.SpotQuantity

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.AppMiningGoodStocks[0].EntID, true), // For mining pool stock, here is the ent_id of mining pool stock
		WithAppGoodID(&ret.AppGoodID, true),
		WithLocked(&ret.Locked, true),
		WithLockID(&lockID, true),
		WithAppSpotLocked(&ret.Locked, true),
	)
	if assert.Nil(t, err) {
		err = handler.LockStock(context.Background())
		if assert.Nil(t, err) {
			handler, err := NewHandler(
				context.Background(),
				WithEntID(&ret.EntID, true),
			)
			if assert.Nil(t, err) {
				info, err := handler.GetStock(context.Background())
				if assert.Nil(t, err) {
					ret.UpdatedAt = info.UpdatedAt
					assert.Equal(t, &ret, info)
				}
			}
		}
	}
}

func waitStartStock(t *testing.T) {
	ret.WaitStart = ret.Locked
	ret.AppMiningGoodStocks[0].WaitStart = ret.Locked
	ret.Sold = ret.Locked
	ret.AppMiningGoodStocks[0].Sold = ret.Locked
	ret.Locked = decimal.NewFromInt(0).String()
	ret.AppMiningGoodStocks[0].Locked = ret.Locked

	handler, err := NewHandler(
		context.Background(),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.WaitStartStock(context.Background())
		if assert.Nil(t, err) {
			handler, err := NewHandler(
				context.Background(),
				WithEntID(&ret.EntID, true),
			)
			if assert.Nil(t, err) {
				info, err := handler.GetStock(context.Background())
				if assert.Nil(t, err) {
					ret.UpdatedAt = info.UpdatedAt
					assert.Equal(t, &ret, info)
				}
			}
		}
	}
}

func lockFailStock(t *testing.T) {
	locked := decimal.NewFromInt(1999).String()
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.AppMiningGoodStocks[0].EntID, true), // For mining pool stock, here is the ent_id of mining pool stock
		WithAppGoodID(&ret.AppGoodID, true),
		WithLocked(&locked, true),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.LockStock(context.Background())
		assert.NotNil(t, err)
	}
}

func chargeBackStock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.ChargeBackStock(context.Background())
		if assert.Nil(t, err) {
			handler, err := NewHandler(
				context.Background(),
				WithEntID(&ret.EntID, true),
			)
			if assert.Nil(t, err) {
				info, err := handler.GetStock(context.Background())
				if assert.Nil(t, err) {
					ret.WaitStart = decimal.NewFromInt(0).String()
					ret.Sold = decimal.NewFromInt(0).String()
					ret.AppMiningGoodStocks[0].WaitStart = decimal.NewFromInt(0).String()
					ret.AppMiningGoodStocks[0].Sold = decimal.NewFromInt(0).String()
					ret.UpdatedAt = info.UpdatedAt
					assert.Equal(t, &ret, info)
				}
			}
		}
	}
}

func TestStock(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("reserveStock", reserveStock)
	t.Run("lockStock", lockStock)
	t.Run("waitStartStock", waitStartStock)
	t.Run("lockFailStock", lockFailStock)
	t.Run("chargeBackStock", chargeBackStock)
}
