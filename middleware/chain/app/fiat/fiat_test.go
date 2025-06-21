package appfiat

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/fiat"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	fiat1 "github.com/NpoolPlatform/kunman/middleware/chain/fiat"

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

var (
	fiatUnit = "BTC"
)

var ret = &npool.Fiat{
	AppID:        uuid.NewString(),
	Name:         "My BTC1",
	DisplayNames: []string{"123123", "2323"},
	Logo:         uuid.NewString(),
	Unit:         fiatUnit,
	Display:      true,
}

var req = &npool.FiatReq{
	AppID:        &ret.AppID,
	Name:         &ret.Name,
	DisplayNames: ret.DisplayNames,
	Logo:         &ret.Logo,
}

func setupFiat(t *testing.T) func(*testing.T) {
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID
	ret.FiatName = uuid.NewString()

	h1, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithEntID(&ret.FiatID, true),
		fiat1.WithName(&ret.FiatName, true),
		fiat1.WithUnit(&fiatUnit, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateFiat(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteFiat(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, false),
		WithAppID(req.AppID, true),
		WithFiatID(req.FiatID, true),
		WithName(req.Name, true),
		WithDisplayNames(req.DisplayNames, false),
		WithLogo(req.Logo, true),
		WithDisplay(req.Display, false),
		WithDisplayIndex(req.DisplayIndex, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.DisplayNamesStr = info.DisplayNamesStr
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	logo := uuid.NewString()

	ret.Logo = logo

	req.ID = &ret.ID
	req.Logo = &logo

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithName(req.Name, false),
		WithDisplayNames(req.DisplayNames, false),
		WithLogo(req.Logo, false),
		WithDisplay(req.Display, false),
		WithDisplayIndex(req.DisplayIndex, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func _delete(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	_, err = handler.DeleteFiat(context.Background())
	assert.Nil(t, err)
}

func TestFiat(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupFiat(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", _delete)
}
