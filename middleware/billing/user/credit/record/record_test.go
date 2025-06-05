package record

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/user/credit/record"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/billing/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Record{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	UserID:           uuid.NewString(),
	OperationType:    types.OperationType_Exchange,
	OperationTypeStr: types.OperationType_Exchange.String(),
	CreditsChange:    int32(10),
	Extra:            uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createRecord(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithOperationType(&ret.OperationType, true),
		WithCreditsChange(&ret.CreditsChange, true),
		WithExtra(&ret.Extra, true),
	)
	assert.Nil(t, err)

	err = handler.CreateRecord(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetRecord(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func getRecord(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetRecord(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRecords(t *testing.T) {
	conds := &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		OperationType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.OperationType)},
		IDs:           &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetRecords(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteRecord(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteRecord(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetRecord(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRecord(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRecord", createRecord)
	t.Run("getRecord", getRecord)
	t.Run("getRecords", getRecords)
	t.Run("deleteRecord", deleteRecord)
}
