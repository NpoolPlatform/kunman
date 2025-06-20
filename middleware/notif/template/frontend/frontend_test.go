package frontend

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template/frontend"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/notif/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.FrontendTemplate{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	LangID:     uuid.NewString(),
	UsedFor:    basetypes.UsedFor_KYCApproved,
	UsedForStr: basetypes.UsedFor_KYCApproved.String(),
	Title:      "title " + uuid.NewString(),
	Content:    "content " + uuid.NewString(),
}

func createFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithAppID(&ret.AppID, true),
		WithLangID(&ret.LangID, true),
		WithUsedFor(&ret.UsedFor, true),
		WithTitle(&ret.Title, false),
		WithContent(&ret.Content, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateFrontendTemplate(t *testing.T) {
	ret.Title = "change Title " + uuid.NewString()
	ret.Content = "change Content " + uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, false),
		WithLangID(&ret.LangID, false),
		WithUsedFor(&ret.UsedFor, false),
		WithTitle(&ret.Title, false),
		WithContent(&ret.Content, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFrontendTemplates(t *testing.T) {
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

	infos, _, err := handler.GetFrontendTemplates(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteFrontendTemplate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteFrontendTemplate(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetFrontendTemplate(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFrontendTemplate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createFrontendTemplate", createFrontendTemplate)
	t.Run("updateFrontendTemplate", updateFrontendTemplate)
	t.Run("getFrontendTemplate", getFrontendTemplate)
	t.Run("getFrontendTemplates", getFrontendTemplates)
	t.Run("deleteFrontendTemplate", deleteFrontendTemplate)
}
