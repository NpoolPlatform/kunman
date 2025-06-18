package appcountry

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/appcountry"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	country "github.com/NpoolPlatform/kunman/middleware/g11n/country"
	"github.com/NpoolPlatform/kunman/middleware/g11n/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.Country{
		EntID:     uuid.NewString(),
		AppID:     uuid.NewString(),
		CountryID: uuid.NewString(),
		Country:   "test country" + uuid.NewString(),
		Flag:      "test flag" + uuid.NewString(),
		Code:      "test code" + uuid.NewString(),
		Short:     "test short" + uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	ch, err := country.NewHandler(
		context.Background(),
		country.WithEntID(&ret.CountryID, true),
		country.WithCountry(&ret.Country, true),
		country.WithFlag(&ret.Flag, true),
		country.WithCode(&ret.Code, true),
		country.WithShort(&ret.Short, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ch)
	country1, err := ch.CreateCountry(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, country1)

	return func(*testing.T) {
		_, _ = ch.DeleteCountry(context.Background())
	}
}

func createCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCountryID(&ret.CountryID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCountry(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCountry(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCountries(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCountries(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCountry(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCountry(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCountry(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCountry", createCountry)
	t.Run("getCountry", getCountry)
	t.Run("getCountries", getCountries)
	t.Run("deleteCountry", deleteCountry)
}
