package apppool

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/app/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/registetestinfo"
	testinit "github.com/NpoolPlatform/kunman/middleware/miningpool/testinit"
	"github.com/google/uuid"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"

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

var ret = &npool.Pool{
	EntID:  uuid.NewString(),
	AppID:  uuid.NewString(),
	PoolID: uuid.NewString(),
}

var req = &npool.PoolReq{
	EntID:  &ret.EntID,
	AppID:  &ret.AppID,
	PoolID: &ret.PoolID,
}

func create(t *testing.T) {
	poolH, err := pool.NewHandler(context.Background(), pool.WithConds(nil), pool.WithOffset(0), pool.WithLimit(2))
	assert.Nil(t, err)

	infos, _, err := poolH.GetPools(context.Background())
	assert.Nil(t, err)
	if !assert.Equal(t, true, len(infos) > 0) {
		return
	}

	ret.PoolID = infos[0].EntID
	req.PoolID = &infos[0].EntID

	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, true),
		WithAppID(req.AppID, true),
		WithPoolID(req.PoolID, true),
	)
	assert.Nil(t, err)

	err = handler.CreatePool(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetPool(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}

	entID := uuid.New()
	handler.EntID = &entID
	err = handler.CreatePool(context.Background())
	assert.NotNil(t, err)
}

func update(t *testing.T) {
	ret.AppID = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, false),
	)
	assert.Nil(t, err)

	err = handler.UpdatePool(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetPool(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	poolID := uuid.NewString()
	_, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithPoolID(&poolID, false),
	)
	assert.NotNil(t, err)
}

func deleteRow(t *testing.T) {
	conds := &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetPools(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], ret)
	}

	ret.EntID = infos[0].EntID
	handler, err = NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistPool(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	err = handler.DeletePool(context.Background())
	assert.Nil(t, err)

	handler, err = NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: ret.EntID,
			},
		}),
	)
	assert.Nil(t, err)

	exist, err = handler.ExistPoolConds(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}

func TestPool(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	registetestinfo.InitTestInfo(context.Background())
	t.Run("create", create)
	t.Run("update", update)
	t.Run("deleteRow", deleteRow)
	registetestinfo.CleanTestInfo(context.Background())
}
