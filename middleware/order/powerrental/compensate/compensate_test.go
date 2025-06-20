package compensate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	compensatemiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/compensate"
	compensate1 "github.com/NpoolPlatform/kunman/middleware/order/compensate"
	powerrental1 "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/order/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = compensatemiddlewarepb.Compensate{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UserID:            uuid.NewString(),
	GoodID:            uuid.NewString(),
	GoodType:          goodtypes.GoodType_PowerRental,
	AppGoodID:         uuid.NewString(),
	OrderID:           uuid.NewString(),
	CompensateFromID:  uuid.NewString(),
	CompensateType:    types.CompensateType_CompensateMalfunction,
	CompensateSeconds: 10000,
}

func setup(t *testing.T) func(*testing.T) {
	ret.CompensateTypeStr = ret.CompensateType.String()
	ret.GoodTypeStr = ret.GoodType.String()

	h1, err := powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		powerrental1.WithAppID(&ret.AppID, true),
		powerrental1.WithUserID(&ret.UserID, true),
		powerrental1.WithGoodID(&ret.GoodID, true),
		powerrental1.WithGoodType(&ret.GoodType, true),
		powerrental1.WithAppGoodID(&ret.AppGoodID, true),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderType(func() *types.OrderType { e := types.OrderType_Offline; return &e }(), true),
		powerrental1.WithAppGoodStockID(func() *string { s := uuid.NewString(); return &s }(), true),
		powerrental1.WithUnits(func() *string { s := decimal.NewFromInt(10).String(); return &s }(), true),
		powerrental1.WithPaymentType(func() *types.PaymentType { e := types.PaymentType_PayWithOffline; return &e }(), true),
		powerrental1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
		powerrental1.WithGoodValueUSD(func() *string { s := decimal.NewFromInt(100).String(); return &s }(), true),
		powerrental1.WithDurationSeconds(func() *uint32 { u := uint32(10000); return &u }(), true),
		powerrental1.WithAppGoodStockLockID(func() *string { s := uuid.NewString(); return &s }(), true),
		powerrental1.WithStartMode(func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateWaitPayment.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStatePaymentTransferReceived.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStatePaymentTransferBookKeeping.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStatePaymentSpendBalance.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateTransferGoodStockLocked.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateAchievementBookKeeping.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateAddCommission.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStatePaymentUnlockAccount.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStatePaid.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateTransferGoodStockWaitStart.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateCreateOrderUser.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateSetProportion.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateSetRevenueAddress.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	h1, err = powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
		powerrental1.WithOrderState(types.OrderState_OrderStateInService.Enum(), true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)

	return func(t *testing.T) {
		_ = h1.DeletePowerRental(context.Background())
	}
}

func getCompensate() (*compensatemiddlewarepb.Compensate, error) {
	handler, err := compensate1.NewHandler(
		context.Background(),
		compensate1.WithEntID(&ret.EntID, true),
	)
	if err != nil {
		return nil, err
	}
	return handler.GetCompensate(context.Background())
}

func getPowerRentalCompensateSeconds() (uint32, error) {
	handler, err := powerrental1.NewHandler(
		context.Background(),
		powerrental1.WithOrderID(&ret.OrderID, true),
	)
	if err != nil {
		return 0, err
	}
	info, err := handler.GetPowerRental(context.Background())
	if err != nil {
		return 0, err
	}
	if info == nil {
		return 0, fmt.Errorf("invalid powerrental")
	}
	return info.CompensateSeconds, nil
}

func createCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
		WithCompensateFromID(&ret.CompensateFromID, true),
		WithCompensateType(&ret.CompensateType, true),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateCompensate(context.Background())
		if assert.Nil(t, err) {
			info, err := getCompensate()
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
			seconds, err := getPowerRentalCompensateSeconds()
			if assert.Nil(t, err) {
				assert.Equal(t, ret.CompensateSeconds, seconds)
			}
		}
	}

	handler, err = NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithCompensateFromID(func() *string { s := uuid.NewString(); return &s }(), true),
		WithCompensateType(&ret.CompensateType, true),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateCompensate(context.Background())
		if assert.Nil(t, err) {
			seconds, err := getPowerRentalCompensateSeconds()
			if assert.Nil(t, err) {
				assert.Equal(t, ret.CompensateSeconds*2, seconds)
			}
		}
	}
}

func getCompensate1(t *testing.T) {
	info, err := getCompensate()
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getCompensates(t *testing.T) {
	handler, err := compensate1.NewHandler(
		context.Background(),
		compensate1.WithConds(&compensatemiddlewarepb.Conds{
			OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
			OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
		}),
		compensate1.WithOffset(0),
		compensate1.WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetCompensates(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(2), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteCompensate(context.Background())
		assert.Nil(t, err)

		info, err := getCompensate()
		assert.Nil(t, err)
		assert.Nil(t, info)

		seconds, err := getPowerRentalCompensateSeconds()
		if assert.Nil(t, err) {
			assert.Equal(t, ret.CompensateSeconds, seconds)
		}
	}
}

func TestCompensate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCompensate", createCompensate)
	t.Run("getCompensate", getCompensate1)
	t.Run("getCompensates", getCompensates)
	t.Run("deleteCompensate", deleteCompensate)
}
