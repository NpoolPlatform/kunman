package poster

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	devicetype1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device"
	manufacturer1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/device/manufacturer"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"

	"github.com/google/uuid"
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

var ret = npool.Poster{
	EntID:        uuid.NewString(),
	DeviceTypeID: uuid.NewString(),
	DeviceType:   uuid.NewString(),
	Manufacturer: uuid.NewString(),
	Poster:       uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	manufacturerID := uuid.NewString()
	h10, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(&ret.Manufacturer, true),
	)
	assert.Nil(t, err)

	err = h10.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	h1, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&ret.DeviceTypeID, true),
		devicetype1.WithType(&ret.DeviceType, true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
	)
	assert.Nil(t, err)

	err = h1.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteDeviceType(context.Background())
		_ = h10.DeleteManufacturer(context.Background())
	}
}

func createPoster(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithPoster(&ret.Poster, true),
	)
	assert.Nil(t, err)

	err = h1.CreatePoster(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetPoster(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithPoster(&ret.Poster, true),
	)
	assert.Nil(t, err)

	err = h2.CreatePoster(context.Background())
	assert.Nil(t, err)
}

func updatePoster(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithPoster(&ret.Poster, true),
	)
	assert.Nil(t, err)

	err = handler.UpdatePoster(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetPoster(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getPoster(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetPoster(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getPosters(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			DeviceTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.DeviceTypeID},
			DeviceTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.DeviceTypeID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetPosters(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deletePoster(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeletePoster(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetPoster(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestPoster(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPoster", createPoster)
	t.Run("updatePoster", updatePoster)
	t.Run("getPoster", getPoster)
	t.Run("getPosters", getPosters)
	t.Run("deletePoster", deletePoster)
}
