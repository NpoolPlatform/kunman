package transfer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/transfer"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/account/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Transfer{
	EntID:        uuid.NewString(),
	AppID:        uuid.NewString(),
	UserID:       uuid.NewString(),
	TargetUserID: uuid.NewString(),
}

func createTransfer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, false),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithTargetUserID(&ret.TargetUserID, true),
	)
	assert.Nil(t, err)
	info, err := handler.CreateTransfer(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func getTransfer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)
	info, err := handler.GetTransfer(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getTransfers(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			TargetUserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.TargetUserID},
		}),
		WithOffset(0),
		WithLimit(2),
	)
	assert.Nil(t, err)
	infos, _, err := handler.GetTransfers(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteTransfer(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)
	info, err := handler.DeleteTransfer(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetTransfer(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createTransfer", createTransfer)
	t.Run("getTransfer", getTransfer)
	t.Run("getTransfers", getTransfers)
	t.Run("deleteTransfer", deleteTransfer)
}
