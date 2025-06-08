package fiat

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.Fiat{
	Name: uuid.NewString(),
	Logo: uuid.NewString(),
	Unit: "USD",
}

var req = &npool.FiatReq{
	Name: &ret.Name,
	Logo: &ret.Logo,
	Unit: &ret.Unit,
}

func setupFiat(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func create(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithName(&ret.Name, true),
		WithLogo(&ret.Logo, true),
		WithUnit(&ret.Unit, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info.String(), ret.String())
	}
}

func update(t *testing.T) {
	ret.Logo = "ABCED"
	ret.Name = uuid.NewString()

	req.Logo = &ret.Logo
	req.Name = &ret.Name

	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithName(req.Name, true),
		WithUnit(req.Unit, true),
		WithLogo(req.Logo, true),
	)
	assert.Nil(t, err)

	info, err := h1.UpdateFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func TestFiat(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupFiat(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
}
