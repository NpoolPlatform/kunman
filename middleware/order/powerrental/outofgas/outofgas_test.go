package outofgas

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	outofgasmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/outofgas"
	outofgas1 "github.com/NpoolPlatform/kunman/middleware/order/outofgas"
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

const outOfGasSeconds = uint32(1)

var ret = outofgasmiddlewarepb.OutOfGas{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
	GoodID:    uuid.NewString(),
	GoodType:  goodtypes.GoodType_PowerRental,
	AppGoodID: uuid.NewString(),
	OrderID:   uuid.NewString(),
	StartAt:   uint32(time.Now().Unix()),
}

func setup(t *testing.T) func(*testing.T) {
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

func getOutOfGas() (*outofgasmiddlewarepb.OutOfGas, error) {
	handler, err := outofgas1.NewHandler(
		context.Background(),
		outofgas1.WithEntID(&ret.EntID, true),
	)
	if err != nil {
		return nil, err
	}
	return handler.GetOutOfGas(context.Background())
}

func getPowerRentalOutOfGasSeconds() (uint32, error) {
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
	return info.OutOfGasSeconds, nil
}

func createOutOfGas(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
		WithStartAt(&ret.StartAt, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateOutOfGas(context.Background())
		if assert.Nil(t, err) {
			info, err := getOutOfGas()
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
			seconds, err := getPowerRentalOutOfGasSeconds()
			if assert.Nil(t, err) {
				assert.Equal(t, uint32(0), seconds)
			}
		}
	}

	err = handler.CreateOutOfGas(context.Background())
	assert.NotNil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID, true),
		WithStartAt(&ret.StartAt, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateOutOfGas(context.Background())
		assert.NotNil(t, err)
	}
}

func updateOutOfGas(t *testing.T) {
	ret.EndAt = ret.StartAt + outOfGasSeconds

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
		WithEndAt(&ret.EndAt, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateOutOfGas(context.Background())
		if assert.Nil(t, err) {
			info, err := getOutOfGas()
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
			seconds, err := getPowerRentalOutOfGasSeconds()
			if assert.Nil(t, err) {
				assert.Equal(t, outOfGasSeconds, seconds)
			}
		}
	}
}

func getOutOfGas1(t *testing.T) {
	info, err := getOutOfGas()
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getOutOfGases(t *testing.T) {
	handler, err := outofgas1.NewHandler(
		context.Background(),
		outofgas1.WithConds(&outofgasmiddlewarepb.Conds{
			OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
			OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
		}),
		outofgas1.WithOffset(0),
		outofgas1.WithLimit(0),
	)
	if assert.Nil(t, err) {
		_, total, err := handler.GetOutOfGases(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(1), total)
		}
	}
}

func deleteOutOfGas(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteOutOfGas(context.Background())
		assert.Nil(t, err)

		info, err := getOutOfGas()
		assert.Nil(t, err)
		assert.Nil(t, info)

		seconds, err := getPowerRentalOutOfGasSeconds()
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(0), seconds)
		}
	}
}

func TestOutOfGas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createOutOfGas", createOutOfGas)
	t.Run("updateOutOfGas", updateOutOfGas)
	t.Run("getOutOfGas", getOutOfGas1)
	t.Run("getOutOfGases", getOutOfGases)
	t.Run("deleteOutOfGas", deleteOutOfGas)
}
