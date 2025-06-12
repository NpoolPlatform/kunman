package registration

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	invitationcode1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/invitationcode"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var regs1 = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _regs1 = npool.RegistrationReq{
	EntID:     &regs1.EntID,
	AppID:     &regs1.AppID,
	InviterID: &regs1.InviterID,
	InviteeID: &regs1.InviteeID,
}

var regs2 = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     regs1.AppID,
	InviterID: regs1.InviteeID,
	InviteeID: uuid.NewString(),
}

var _regs2 = npool.RegistrationReq{
	EntID:     &regs2.EntID,
	AppID:     &regs2.AppID,
	InviterID: &regs2.InviterID,
	InviteeID: &regs2.InviteeID,
}

var regs4 = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     regs1.AppID,
	InviterID: uuid.NewString(),
	InviteeID: uuid.NewString(),
}

var _regs4 = npool.RegistrationReq{
	EntID:     &regs4.EntID,
	AppID:     &regs4.AppID,
	InviterID: &regs4.InviterID,
	InviteeID: &regs4.InviteeID,
}

var regs5 = npool.Registration{
	EntID:     uuid.NewString(),
	AppID:     regs1.AppID,
	InviterID: regs4.InviteeID,
	InviteeID: uuid.NewString(),
}

var _regs5 = npool.RegistrationReq{
	EntID:     &regs5.EntID,
	AppID:     &regs5.AppID,
	InviterID: &regs5.InviterID,
	InviteeID: &regs5.InviteeID,
}

func setupRegs(t *testing.T) func(*testing.T) { //nolint
	_h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_regs1.AppID, true),
		invitationcode1.WithUserID(_regs1.InviterID, true),
	)
	assert.Nil(t, err)

	_info1, err := _h1.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h1.ID = &_info1.ID
	}

	h1, err := NewHandler(
		context.Background(),
		WithEntID(_regs1.EntID, true),
		WithAppID(_regs1.AppID, true),
		WithInviterID(_regs1.InviterID, true),
		WithInviteeID(_regs1.InviteeID, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateRegistration(context.Background())
	assert.Nil(t, err)
	h1.ID = &info1.ID
	regs1.ID = info1.ID
	regs1.CreatedAt = info1.CreatedAt
	regs1.UpdatedAt = info1.UpdatedAt

	_h2, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_regs2.AppID, true),
		invitationcode1.WithUserID(_regs2.InviterID, true),
	)
	assert.Nil(t, err)

	_info2, err := _h2.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h2.ID = &_info2.ID
	}

	assert.Nil(t, err)
	assert.NotNil(t, _info2)

	h2, err := NewHandler(
		context.Background(),
		WithEntID(_regs2.EntID, true),
		WithAppID(_regs2.AppID, true),
		WithInviterID(_regs2.InviterID, true),
		WithInviteeID(_regs2.InviteeID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateRegistration(context.Background())
	assert.Nil(t, err)
	h2.ID = &info2.ID
	regs2.ID = info2.ID
	regs2.CreatedAt = info2.CreatedAt
	regs2.UpdatedAt = info2.UpdatedAt

	_h4, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_regs4.AppID, true),
		invitationcode1.WithUserID(_regs4.InviterID, true),
	)
	assert.Nil(t, err)

	_info4, err := _h4.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h4.ID = &_info4.ID
	}

	h4, err := NewHandler(
		context.Background(),
		WithEntID(_regs4.EntID, true),
		WithAppID(_regs4.AppID, true),
		WithInviterID(_regs4.InviterID, true),
		WithInviteeID(_regs4.InviteeID, true),
	)
	assert.Nil(t, err)

	info4, err := h4.CreateRegistration(context.Background())
	assert.Nil(t, err)
	h4.ID = &info4.ID
	regs4.ID = info4.ID
	regs4.CreatedAt = info4.CreatedAt
	regs4.UpdatedAt = info4.UpdatedAt

	_h5, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(_regs5.AppID, true),
		invitationcode1.WithUserID(_regs5.InviterID, true),
	)
	assert.Nil(t, err)

	_info5, err := _h5.CreateInvitationCode(context.Background())
	if assert.Nil(t, err) {
		_h5.ID = &_info5.ID
	}

	h5, err := NewHandler(
		context.Background(),
		WithEntID(_regs5.EntID, true),
		WithAppID(_regs5.AppID, true),
		WithInviterID(_regs5.InviterID, true),
		WithInviteeID(_regs5.InviteeID, true),
	)
	assert.Nil(t, err)

	info5, err := h5.CreateRegistration(context.Background())
	assert.Nil(t, err)
	h5.ID = &info5.ID
	regs5.ID = info5.ID
	regs5.CreatedAt = info5.CreatedAt
	regs5.UpdatedAt = info5.UpdatedAt

	return func(*testing.T) {
		_, _ = _h1.DeleteInvitationCode(context.Background())
		_, _ = _h2.DeleteInvitationCode(context.Background())
		_, _ = _h4.DeleteInvitationCode(context.Background())
		_, _ = _h5.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteRegistration(context.Background())
		_, _ = h2.DeleteRegistration(context.Background())
		_, _ = h4.DeleteRegistration(context.Background())
		_, _ = h5.DeleteRegistration(context.Background())
	}
}

func updateLevelRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&regs4.ID, true),
		WithAppID(&regs4.AppID, true),
		WithInviterID(&regs2.InviterID, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateRegistration(context.Background())
	if assert.Nil(t, err) {
		regs4.InviterID = regs2.InviterID
		regs4.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &regs4)
	}
}

var updateInviterID = uuid.NewString()

func setup(t *testing.T) func(*testing.T) {
	h, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(&ret.AppID, true),
		invitationcode1.WithUserID(&ret.InviterID, true),
	)
	assert.Nil(t, err)

	info, err := h.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
	h.ID = &info.ID

	h1, err := invitationcode1.NewHandler(
		context.Background(),
		invitationcode1.WithAppID(&ret.AppID, true),
		invitationcode1.WithUserID(&updateInviterID, true),
	)
	assert.Nil(t, err)

	info, err = h1.CreateInvitationCode(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, info)
	h1.ID = &info.ID

	return func(*testing.T) {
		_, _ = h.DeleteInvitationCode(context.Background())
		_, _ = h1.DeleteInvitationCode(context.Background())
	}
}

func createRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithInviterID(&ret.InviterID, true),
		WithInviteeID(&ret.InviteeID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateRegistration(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateRegistration(t *testing.T) {
	ret.InviterID = updateInviterID
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithInviterID(&ret.InviterID, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateRegistration(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetRegistration(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRegistrations(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		InviterID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.InviterID},
		InviteeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.InviteeID},
		InviterIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.InviterID}},
		InviteeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.InviteeID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetRegistrations(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteRegistration(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteRegistration(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetRegistration(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRegistration(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	return

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRegistration", createRegistration)
	t.Run("updateRegistration", updateRegistration)
	t.Run("getRegistration", getRegistration)
	t.Run("getRegistrations", getRegistrations)
	t.Run("deleteRegistration", deleteRegistration)

	testLevelReg := setupRegs(t)
	defer testLevelReg(t)

	t.Run("updateLevelRegistration", updateLevelRegistration)
}
