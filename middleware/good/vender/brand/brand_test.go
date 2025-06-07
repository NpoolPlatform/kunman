package brand

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
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

var ret = npool.Brand{
	EntID: uuid.NewString(),
	Name:  uuid.NewString(),
	Logo:  uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createBrand(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
	)
	assert.Nil(t, err)

	err = handler.CreateBrand(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetBrand(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateBrand(t *testing.T) {
	ret.Name = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateBrand(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetBrand(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getBrand(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetBrand(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getBrands(t *testing.T) {
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

	infos, _, err := handler.GetBrands(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteBrand(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteBrand(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetBrand(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestBrand(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createBrand", createBrand)
	t.Run("updateBrand", updateBrand)
	t.Run("getBrand", getBrand)
	t.Run("getBrands", getBrands)
	t.Run("deleteBrand", deleteBrand)
}
