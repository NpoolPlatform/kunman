package country

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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
		EntID:   uuid.NewString(),
		Country: uuid.NewString(),
		Flag:    uuid.NewString(),
		Code:    uuid.NewString(),
		Short:   uuid.NewString(),
	}
)

func createCountry(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithCountry(&ret.Country, true),
		WithFlag(&ret.Flag, true),
		WithCode(&ret.Code, true),
		WithShort(&ret.Short, true),
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

func updateCountry(t *testing.T) {
	ret.Country = uuid.NewString()
	ret.Flag = uuid.NewString()
	ret.Code = uuid.NewString()
	ret.Short = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCountry(&ret.Country, false),
		WithFlag(&ret.Flag, false),
		WithCode(&ret.Code, false),
		WithShort(&ret.Short, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCountry(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
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
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
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

	t.Run("createCountry", createCountry)
	t.Run("updateCountry", updateCountry)
	t.Run("getCountry", getCountry)
	t.Run("getCountries", getCountries)
	t.Run("deleteCountry", deleteCountry)
}
