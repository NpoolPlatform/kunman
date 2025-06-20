package device

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device"
	manufacturer1 "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

var ret = npool.DeviceType{
	EntID:            uuid.NewString(),
	Type:             uuid.NewString(),
	ManufacturerID:   uuid.NewString(),
	ManufacturerName: uuid.NewString(),
	ManufacturerLogo: uuid.NewString(),
	PowerConsumption: 120,
	ShipmentAt:       uint32(time.Now().Unix()),
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&ret.ManufacturerID, true),
		manufacturer1.WithName(&ret.ManufacturerName, true),
		manufacturer1.WithLogo(&ret.ManufacturerLogo, true),
	)
	assert.Nil(t, err)

	err = h1.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteManufacturer(context.Background())
	}
}

func createDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithType(&ret.Type, true),
		WithManufacturerID(&ret.ManufacturerID, true),
		WithPowerConsumption(&ret.PowerConsumption, true),
		WithShipmentAt(&ret.ShipmentAt, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, handler)

	err = handler.CreateDeviceType(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetDeviceType(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateDeviceType(t *testing.T) {
	ret.PowerConsumption = 1000
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithType(&ret.Type, false),
		WithManufacturerID(&ret.ManufacturerID, false),
		WithPowerConsumption(&ret.PowerConsumption, false),
		WithShipmentAt(&ret.ShipmentAt, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateDeviceType(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetDeviceType(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetDeviceType(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getDeviceTypes(t *testing.T) {
	conds := &npool.Conds{
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Type:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		ManufacturerID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ManufacturerID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetDeviceTypes(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteDeviceType(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetDeviceType(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDeviceType(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createDeviceType", createDeviceType)
	t.Run("updateDeviceType", updateDeviceType)
	t.Run("getDeviceType", getDeviceType)
	t.Run("getDeviceTypes", getDeviceTypes)
	t.Run("deleteDeviceType", deleteDeviceType)
}
