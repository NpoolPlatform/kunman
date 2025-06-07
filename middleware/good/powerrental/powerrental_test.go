package powerrental

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	devicetype1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device"
	manufacturer1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device/manufacturer"
	goodcoin1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/coin"
	vendorbrand1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/vender/location"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	stockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/stock"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.PowerRental{
	EntID:  uuid.NewString(),
	GoodID: uuid.NewString(),

	DeviceTypeID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturerName: uuid.NewString(),
	DeviceManufacturerLogo: uuid.NewString(),
	DevicePowerConsumption: 120,
	DeviceShipmentAt:       uint32(time.Now().Unix()),

	VendorLocationID: uuid.NewString(),
	VendorBrand:      uuid.NewString(),
	VendorLogo:       uuid.NewString(),
	VendorCountry:    uuid.NewString(),
	VendorProvince:   uuid.NewString(),

	UnitPrice:           decimal.NewFromInt(120).String(),
	QuantityUnit:        "TiB",
	QuantityUnitAmount:  decimal.NewFromInt(2).String(),
	DeliveryAt:          uint32(time.Now().Unix()),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,
	StockMode:           types.GoodStockMode_GoodStockByMiningPool,

	GoodType:             types.GoodType_PowerRental,
	BenefitType:          types.BenefitType_BenefitTypePool,
	Name:                 uuid.NewString(),
	ServiceStartAt:       uint32(time.Now().Unix()),
	StartMode:            types.GoodStartMode_GoodStartModeInstantly,
	BenefitIntervalHours: 20,
	UnitLockDeposit:      decimal.NewFromInt(1).String(),

	GoodStockID:      uuid.NewString(),
	GoodTotal:        decimal.NewFromInt(120).String(),
	GoodSpotQuantity: decimal.NewFromInt(120).String(),
	GoodLocked:       decimal.NewFromInt(0).String(),
	GoodInService:    decimal.NewFromInt(0).String(),
	GoodWaitStart:    decimal.NewFromInt(0).String(),
	GoodSold:         decimal.NewFromInt(0).String(),
	GoodAppReserved:  decimal.NewFromInt(0).String(),

	GoodCoins: []*goodcoinmwpb.GoodCoinInfo{
		{
			CoinTypeID: uuid.NewString(),
			Main:       true,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
	},

	MiningGoodStocks: []*stockmwpb.MiningGoodStock{
		{
			EntID:          uuid.NewString(),
			PoolRootUserID: uuid.NewString(),
			Total:          decimal.NewFromInt(70).String(),
		},
		{
			EntID:          uuid.NewString(),
			PoolRootUserID: uuid.NewString(),
			Total:          decimal.NewFromInt(50).String(),
		},
	},

	RewardState: types.BenefitState_BenefitWait,
	State:       types.GoodState_GoodStatePreWait,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.StartModeStr = ret.StartMode.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()
	ret.StartModeStr = ret.StartMode.String()
	ret.StockModeStr = ret.StockMode.String()
	ret.StateStr = ret.State.String()
	for _, stock := range ret.MiningGoodStocks {
		stock.GoodStockID = ret.GoodStockID
		stock.SpotQuantity = stock.Total
		stock.Locked = decimal.NewFromInt(0).String()
		stock.InService = decimal.NewFromInt(0).String()
		stock.WaitStart = decimal.NewFromInt(0).String()
		stock.Sold = decimal.NewFromInt(0).String()
	}
	ret.RewardStateStr = ret.RewardState.String()
	for _, goodCoin := range ret.GoodCoins {
		ret.Rewards = append(ret.Rewards, &goodcoinrewardmwpb.RewardInfo{
			GoodID:                ret.GoodID,
			CoinTypeID:            goodCoin.CoinTypeID,
			RewardTID:             uuid.Nil.String(),
			LastRewardAmount:      decimal.NewFromInt(0).String(),
			NextRewardStartAmount: decimal.NewFromInt(0).String(),
			LastUnitRewardAmount:  decimal.NewFromInt(0).String(),
			TotalRewardAmount:     decimal.NewFromInt(0).String(),
			MainCoin:              goodCoin.Main,
		})
	}

	manufacturerID := uuid.NewString()
	h1, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(&ret.DeviceManufacturerName, true),
		manufacturer1.WithLogo(&ret.DeviceManufacturerLogo, true),
	)
	assert.Nil(t, err)

	err = h1.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	h2, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&ret.DeviceTypeID, true),
		devicetype1.WithType(&ret.DeviceType, true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
		devicetype1.WithPowerConsumption(&ret.DevicePowerConsumption, true),
		devicetype1.WithShipmentAt(&ret.DeviceShipmentAt, true),
	)
	assert.Nil(t, err)

	err = h2.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	brandID := uuid.NewString()
	h3, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithEntID(&brandID, true),
		vendorbrand1.WithName(&ret.VendorBrand, true),
		vendorbrand1.WithLogo(&ret.VendorLogo, true),
	)
	assert.Nil(t, err)

	err = h3.CreateBrand(context.Background())
	assert.Nil(t, err)

	h4, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&ret.VendorLocationID, true),
		vendorlocation1.WithBrandID(&brandID, true),
		vendorlocation1.WithCountry(&ret.VendorCountry, true),
		vendorlocation1.WithProvince(&ret.VendorProvince, true),
		vendorlocation1.WithCity(&ret.VendorProvince, true),
		vendorlocation1.WithAddress(&ret.VendorProvince, true),
	)
	assert.Nil(t, err)

	err = h4.CreateLocation(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h4.DeleteLocation(context.Background())
		_ = h3.DeleteBrand(context.Background())
		_ = h2.DeleteDeviceType(context.Background())
		_ = h1.DeleteManufacturer(context.Background())
	}
}

func createPowerRental(t *testing.T) {
	miningGoodStocks := func() (reqs []*stockmwpb.MiningGoodStockReq) {
		for _, stock := range ret.MiningGoodStocks {
			reqs = append(reqs, &stockmwpb.MiningGoodStockReq{
				EntID:          &stock.EntID,
				PoolRootUserID: &stock.PoolRootUserID,
				Total:          &stock.Total,
			})
		}
		return
	}()

	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
		WithState(&ret.State, true),
		WithStockID(&ret.GoodStockID, true),
		WithTotal(&ret.GoodTotal, true),
		WithStockMode(&ret.StockMode, true),
		WithStocks(miningGoodStocks, true),
	)
	assert.Nil(t, err)

	err = h1.CreatePowerRental(context.Background())
	if assert.Nil(t, err) {
		for _, goodCoin := range ret.GoodCoins {
			goodCoin.GoodID = ret.GoodID
			h5, err := goodcoin1.NewHandler(
				context.Background(),
				goodcoin1.WithGoodID(&ret.GoodID, true),
				goodcoin1.WithCoinTypeID(&goodCoin.CoinTypeID, true),
				goodcoin1.WithMain(&goodCoin.Main, true),
				goodcoin1.WithIndex(&goodCoin.Index, true),
			)
			assert.Nil(t, err)

			err = h5.CreateGoodCoin(context.Background())
			assert.Nil(t, err)
		}
		handler, _ := NewHandler(
			context.Background(),
			WithGoodID(&ret.GoodID, true),
		)
		info, err := handler.GetPowerRental(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			ret.State = info.State
			ret.StateStr = info.StateStr
			ret.MiningGoodStocks = info.MiningGoodStocks
			assert.Equal(t, &ret, info)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
		WithState(&ret.State, true),
		WithStockID(&ret.GoodStockID, true),
		WithTotal(&ret.GoodTotal, true),
	)
	assert.Nil(t, err)

	err = h2.CreatePowerRental(context.Background())
	assert.NotNil(t, err)
}

func updatePowerRental(t *testing.T) {
	ret.GoodTotal = decimal.NewFromInt(10000).String()
	ret.GoodSpotQuantity = ret.GoodTotal
	ret.Name = uuid.NewString()
	ret.DeliveryAt = uint32(time.Now().Unix() + 10)

	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)
	info, err := h1.GetPowerRental(context.Background())
	assert.Nil(t, err)

	miningGoodStocks := func() (reqs []*stockmwpb.MiningGoodStockReq) {
		remain := decimal.NewFromInt(0)
		for i, stock := range info.MiningGoodStocks {
			if i == 0 {
				remain = decimal.RequireFromString(ret.GoodTotal).Sub(decimal.RequireFromString(stock.Total))
				continue
			}
			stock.Total = remain.String()
			stock.SpotQuantity = stock.Total
			reqs = append(reqs, &stockmwpb.MiningGoodStockReq{
				EntID:          &stock.EntID,
				PoolRootUserID: &stock.PoolRootUserID,
				Total:          &stock.Total,
			})
			return
		}
		return
	}()

	h2, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
		WithTotal(&ret.GoodTotal, true),
		WithStocks(miningGoodStocks, true),
	)
	assert.Nil(t, err)

	err = h2.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	info, err = h2.GetPowerRental(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.MiningGoodStocks = info.MiningGoodStocks
		assert.Equal(t, info, &ret)
	}
}

func getPowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetPowerRental(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getPowerRentals(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetPowerRentals(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deletePowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeletePowerRental(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetPowerRental(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPowerRental(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPowerRental", createPowerRental)
	t.Run("updatePowerRental", updatePowerRental)
	t.Run("getPowerRental", getPowerRental)
	t.Run("getPowerRentals", getPowerRentals)
	t.Run("deletePowerRental", deletePowerRental)
}
