package channel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/channel"
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

var ret = npool.Channel{
	AppID:        uuid.NewString(),
	EventType:    basetypes.UsedFor_KYCApproved,
	EventTypeStr: basetypes.UsedFor_KYCApproved.String(),
	Channel:      basetypes.NotifChannel_ChannelEmail,
	ChannelStr:   basetypes.NotifChannel_ChannelEmail.String(),
}

func createChannel(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&ret.AppID, true),
		WithEventType(&ret.EventType, true),
		WithChannel(&ret.Channel, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateChannel(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}
}

func getChannel(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetChannel(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getChannels(t *testing.T) {
	conds := &npool.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.EventType)},
		Channel:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.Channel)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetChannels(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteChannel(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteChannel(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetChannel(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestChannel(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createChannel", createChannel)
	t.Run("getChannel", getChannel)
	t.Run("getChannels", getChannels)
	t.Run("deleteChannel", deleteChannel)
}
