package location

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	brand1 "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
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

var ret = npool.Location{
	EntID:     uuid.NewString(),
	Country:   uuid.NewString(),
	Province:  uuid.NewString(),
	City:      uuid.NewString(),
	Address:   uuid.NewString(),
	BrandID:   uuid.NewString(),
	BrandName: uuid.NewString(),
	BrandLogo: uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := brand1.NewHandler(
		context.Background(),
		brand1.WithEntID(&ret.BrandID, true),
		brand1.WithName(&ret.BrandName, true),
		brand1.WithLogo(&ret.BrandLogo, true),
	)
	assert.Nil(t, err)

	err = h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteBrand(context.Background())
	}
}

func createLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithCountry(&ret.Country, true),
		WithProvince(&ret.Province, true),
		WithCity(&ret.City, true),
		WithAddress(&ret.Address, true),
		WithBrandID(&ret.BrandID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateLocation(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetLocation(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateLocation(t *testing.T) {
	ret.Address = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCountry(&ret.Country, true),
		WithProvince(&ret.Province, true),
		WithCity(&ret.City, true),
		WithAddress(&ret.Address, true),
		WithBrandID(&ret.BrandID, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateLocation(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetLocation(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetLocation(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLocations(t *testing.T) {
	conds := &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Country:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.Country},
		Province: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Province},
		BrandID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.BrandID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetLocations(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteLocation(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetLocation(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestLocation(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createLocation", createLocation)
	t.Run("updateLocation", updateLocation)
	t.Run("getLocation", getLocation)
	t.Run("getLocations", getLocations)
	t.Run("deleteLocation", deleteLocation)
}
