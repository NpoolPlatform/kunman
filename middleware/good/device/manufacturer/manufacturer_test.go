package manufacturer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
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

var ret = npool.Manufacturer{
	EntID: uuid.NewString(),
	Name:  uuid.NewString(),
	Logo:  uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createManufacturer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
	)
	assert.Nil(t, err)

	err = handler.CreateManufacturer(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetManufacturer(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateManufacturer(t *testing.T) {
	ret.Logo = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateManufacturer(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetManufacturer(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getManufacturer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetManufacturer(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getManufacturers(t *testing.T) {
	conds := &npool.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Name:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetManufacturers(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteManufacturer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteManufacturer(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetManufacturer(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestManufacturer(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createManufacturer", createManufacturer)
	t.Run("updateManufacturer", updateManufacturer)
	t.Run("getManufacturer", getManufacturer)
	t.Run("getManufacturers", getManufacturers)
	t.Run("deleteManufacturer", deleteManufacturer)
}
