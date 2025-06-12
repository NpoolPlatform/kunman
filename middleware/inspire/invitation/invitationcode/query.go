package invitationcode

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/invitationcode"
	invitationcodecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/invitationcode"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entinvitationcode "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/invitationcode"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.InvitationCodeSelect
	stmSelect *ent.InvitationCodeSelect
	infos     []*npool.InvitationCode
	total     uint32
}

func (h *queryHandler) selectInvitationCode(stm *ent.InvitationCodeQuery) *ent.InvitationCodeSelect {
	return stm.Select(
		entinvitationcode.FieldID,
	)
}

func (h *queryHandler) queryInvitationCode(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.InvitationCode.Query().Where(entinvitationcode.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entinvitationcode.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entinvitationcode.EntID(*h.EntID))
	}
	h.stmSelect = h.selectInvitationCode(stm)
	return nil
}

func (h *queryHandler) queryInvitationCodes(cli *ent.Client) (*ent.InvitationCodeSelect, error) {
	stm, err := invitationcodecrud.SetQueryConds(cli.InvitationCode.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectInvitationCode(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entinvitationcode.Table)
	s.LeftJoin(t).
		On(
			s.C(entinvitationcode.FieldID),
			t.C(entinvitationcode.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entinvitationcode.FieldEntID), "ent_id"),
			sql.As(t.C(entinvitationcode.FieldAppID), "app_id"),
			sql.As(t.C(entinvitationcode.FieldUserID), "user_id"),
			sql.As(t.C(entinvitationcode.FieldInvitationCode), "invitation_code"),
			sql.As(t.C(entinvitationcode.FieldDisabled), "disabled"),
			sql.As(t.C(entinvitationcode.FieldCreatedAt), "created_at"),
			sql.As(t.C(entinvitationcode.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return wlog.WrapError(err)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetInvitationCode(ctx context.Context) (*npool.InvitationCode, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.InvitationCode{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryInvitationCode(cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetInvitationCodes(ctx context.Context) ([]*npool.InvitationCode, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.InvitationCode{},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryInvitationCodes(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryInvitationCodes(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryJoin(); err != nil {
			return wlog.WrapError(err)
		}
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	return handler.infos, handler.total, nil
}
