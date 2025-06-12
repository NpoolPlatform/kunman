package registration

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	achievementusercrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/user"
	registrationcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/registration"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createOrAddInvites(ctx context.Context, tx *ent.Tx, req *registrationcrud.Req) error {
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

	invitees := uint32(1)
	if h.InviterID == req.InviterID {
		_req.DirectInvitees = &invitees
	} else {
		_req.IndirectInvitees = &invitees
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
	if h.InviterID == req.InviterID {
		invitees += directInvitees
		_req.DirectInvitees = &invitees
	} else {
		invitees += indirectInvitees
		_req.IndirectInvitees = &invitees
	}

	if _, err := achievementusercrud.UpdateSet(
		tx.AchievementUser.UpdateOneID(info.ID),
		_req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *createHandler) addInvites(ctx context.Context, tx *ent.Tx) error {
	handler, err := NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	handler.AppID = h.AppID
	handler.InviteeID = h.InviterID

	inviters, _, err := handler.GetSortedInvitersWithClient(ctx, tx.Client())
	if err != nil {
		return wlog.WrapError(err)
	}

	for _, inviter := range inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}
		inviterID := uuid.MustParse(inviter.InviterID)
		req := &registrationcrud.Req{
			AppID:     h.AppID,
			InviterID: &inviterID,
		}
		if err := h.createOrAddInvites(ctx, tx, req); err != nil {
			return wlog.WrapError(err)
		}
	}
	req := &registrationcrud.Req{
		AppID:     h.AppID,
		InviterID: h.InviterID,
	}
	if err := h.createOrAddInvites(ctx, tx, req); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) CreateRegistration(ctx context.Context) (*npool.Registration, error) {
	if err := h.validateInvitationCode(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	h.Conds = &registrationcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		InviteeID: &cruder.Cond{Op: cruder.EQ, Val: *h.InviteeID},
	}
	exist, err := h.ExistRegistrationConds(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if exist {
		return nil, wlog.Errorf("already exists")
	}

	handler := &createHandler{
		Handler: h,
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := registrationcrud.CreateSet(
			tx.Registration.Create(),
			&registrationcrud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				InviterID: h.InviterID,
				InviteeID: h.InviteeID,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}

		if err := handler.addInvites(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetRegistration(ctx)
}
