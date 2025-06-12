package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	achievementusermwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/user"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	achievementuser1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user"
	common1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
	achievementusercrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/user"
	registrationcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/registration"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	registration   *npool.Registration
	subInviterIDs  []string
	addInviterIDs  []string
	inviteeInvites uint32
}

func (h *updateHandler) subAchievementInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
	stm, err := achievementusercrud.SetQueryConds(
		tx.AchievementUser.Query(),
		&achievementusercrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *req.InviterID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}
	if info == nil {
		return nil
	}

	_req := &achievementusercrud.Req{}

	invites := uint32(1)
	directInvitees := info.DirectInvitees
	indirectInvitees := info.IndirectInvitees
	oldInviterID := uuid.MustParse(h.registration.InviterID)
	if oldInviterID == *req.InviterID {
		if directInvitees != uint32(0) {
			directInvitees -= invites
			_req.DirectInvitees = &directInvitees
		}
		if indirectInvitees != uint32(0) {
			indirectInvitees -= h.inviteeInvites
			_req.IndirectInvitees = &indirectInvitees
		}
	}
	if oldInviterID != *req.InviterID && indirectInvitees != uint32(0) {
		indirectInvitees = indirectInvitees - invites - h.inviteeInvites
		_req.IndirectInvitees = &indirectInvitees
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *updateHandler) getTotalInvites(ctx context.Context) error {
	handler, err := achievementuser1.NewHandler(
		ctx,
		common1.WithConds(&achievementusermwpb.Conds{
			AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.registration.AppID},
			UserID: &basetypes.StringVal{Op: cruder.EQ, Value: h.registration.InviteeID},
		}),
		common1.WithLimit(int32(1)),
	)
	if err != nil {
		return nil
	}
	achivmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(achivmentUsers) == 0 {
		return nil
	}
	h.inviteeInvites = achivmentUsers[0].DirectInvitees + achivmentUsers[0].IndirectInvitees
	return nil
}

func (h *updateHandler) createOrAddInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
	stm, err := achievementusercrud.SetQueryConds(
		tx.AchievementUser.Query(),
		&achievementusercrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *req.InviterID},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}

	_req := &achievementusercrud.Req{
		AppID:  req.AppID,
		UserID: req.InviterID,
	}

	invites := uint32(1)
	if h.InviterID == req.InviterID {
		_req.DirectInvitees = &invites
	} else {
		_req.IndirectInvitees = &invites
	}

	if info == nil {
		if _, err = achievementusercrud.CreateSet(
			tx.AchievementUser.Create(),
			_req,
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	}

	directInvitees := info.DirectInvitees
	indirectInvitees := info.IndirectInvitees

	if h.InviterID.String() == req.InviterID.String() {
		directInvitees += invites
		_req.DirectInvitees = &directInvitees
		indirectInvitees += h.inviteeInvites
		_req.IndirectInvitees = &indirectInvitees
	} else {
		indirectInvitees = indirectInvitees + invites + h.inviteeInvites
		_req.IndirectInvitees = &indirectInvitees
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *updateHandler) addInvites(ctx context.Context, tx *ent.Tx) error {
	for _, inviter := range h.addInviterIDs {
		if inviter == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.createOrAddInvites(ctx, tx, req); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) subInvites(ctx context.Context, tx *ent.Tx) error {
	for _, inviter := range h.subInviterIDs {
		if inviter == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.subAchievementInvites(ctx, tx, req); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *updateHandler) getInviters(ctx context.Context, inviterID *uuid.UUID) ([]string, error) {
	handler, err := NewHandler(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.AppID = h.AppID
	handler.InviteeID = inviterID

	_, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return inviterIDs, nil
}

func (h *Handler) UpdateRegistration(ctx context.Context) (*npool.Registration, error) {
	info, err := h.GetRegistration(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("registration not found")
	}
	if info.InviterID == h.InviterID.String() || info.InviteeID == h.InviterID.String() {
		return nil, wlog.Errorf("invalid inviterid")
	}

	if err := h.validateInvitationCode(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &updateHandler{
		Handler:        h,
		registration:   info,
		subInviterIDs:  []string{},
		addInviterIDs:  []string{},
		inviteeInvites: uint32(0),
	}

	if err := handler.getTotalInvites(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	inviteeID := uuid.MustParse(handler.registration.InviterID)
	subInviters, err := handler.getInviters(ctx, &inviteeID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	addInviters, err := handler.getInviters(ctx, h.InviterID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.subInviterIDs = subInviters
	handler.addInviterIDs = addInviters

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.subInvites(ctx, tx); err != nil {
			return nil
		}

		if _, err := registrationcrud.UpdateSet(
			tx.Registration.UpdateOneID(*h.ID),
			&registrationcrud.Req{
				InviterID: h.InviterID,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.addInvites(ctx, tx); err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetRegistration(ctx)
}
