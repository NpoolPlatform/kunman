package lang

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/lang"
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

var ret = npool.Lang{
	EntID: uuid.NewString(),
	Lang:  uuid.NewString(),
	Logo:  uuid.NewString(),
	Name:  uuid.NewString(),
	Short: uuid.NewString(),
}

func createLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithLang(&ret.Lang, true),
		WithLogo(&ret.Logo, true),
		WithName(&ret.Name, true),
		WithShort(&ret.Short, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateLang(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateLang(t *testing.T) {
	ret.Lang = uuid.NewString()
	ret.Logo = uuid.NewString()
	ret.Name = uuid.NewString()
	ret.Short = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithLang(&ret.Lang, false),
		WithLogo(&ret.Logo, false),
		WithName(&ret.Name, false),
		WithShort(&ret.Short, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateLang(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetLang(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLangs(t *testing.T) {
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

	infos, _, err := handler.GetLangs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteLang(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteLang(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetLang(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestLang(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createLang", createLang)
	t.Run("updateLang", updateLang)
	t.Run("getLang", getLang)
	t.Run("getLangs", getLangs)
	t.Run("deleteLang", deleteLang)
}
