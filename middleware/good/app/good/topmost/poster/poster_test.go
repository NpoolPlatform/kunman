package poster

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/poster"
	topmost1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	TopMostID:        uuid.NewString(),
	TopMostType:      types.GoodTopMostType_TopMostBestOffer,
	TopMostTitle:     uuid.NewString(),
	TopMostMessage:   uuid.NewString(),
	TopMostTargetUrl: uuid.NewString(),
	Poster:           uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.TopMostTypeStr = ret.TopMostType.String()

	h1, err := topmost1.NewHandler(
		context.Background(),
		topmost1.WithEntID(&ret.TopMostID, true),
		topmost1.WithAppID(&ret.AppID, true),
		topmost1.WithTopMostType(&ret.TopMostType, true),
		topmost1.WithTitle(&ret.TopMostTitle, true),
		topmost1.WithMessage(&ret.TopMostMessage, true),
		topmost1.WithTargetURL(&ret.TopMostTargetUrl, true),
		topmost1.WithStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		topmost1.WithEndAt(func() *uint32 { u := uint32(time.Now().Unix() + 10000); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateTopMost(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteTopMost(context.Background())
	}
}

func createPoster(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithTopMostID(&ret.TopMostID, true),
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
		WithTopMostID(&ret.TopMostID, true),
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
			ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			TopMostID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
			TopMostIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.TopMostID}},
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
