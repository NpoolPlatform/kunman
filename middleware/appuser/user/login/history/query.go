package history

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	historycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/login/history"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	entappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuser"
	entloginhistory "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/loginhistory"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/login/history"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
)

type queryHandler struct {
	*Handler
	stm   *ent.LoginHistorySelect
	infos []*npool.History
	total uint32
}

func (h *queryHandler) selectHistory(stm *ent.LoginHistoryQuery) {
	h.stm = stm.Select(
		entloginhistory.FieldID,
		entloginhistory.FieldEntID,
		entloginhistory.FieldAppID,
		entloginhistory.FieldUserID,
		entloginhistory.FieldClientIP,
		entloginhistory.FieldUserAgent,
		entloginhistory.FieldLocation,
		entloginhistory.FieldLoginType,
		entloginhistory.FieldCreatedAt,
	)
}

func (h *queryHandler) queryHistory(cli *ent.Client) {
	stm := cli.LoginHistory.
		Query().
		Where(entloginhistory.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entloginhistory.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entloginhistory.EntID(*h.EntID))
	}
	h.selectHistory(stm)
}

func (h *queryHandler) queryLoginHistories(ctx context.Context, cli *ent.Client) error {
	stm, err := historycrud.SetQueryConds(cli.LoginHistory.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectHistory(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entloginhistory.FieldAppID),
			t.C(entapp.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entloginhistory.FieldAppID),
			t.C(entappuser.FieldAppID),
		).
		On(
			s.C(entloginhistory.FieldUserID),
			t.C(entappuser.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			sql.As(t.C(entappuser.FieldPhoneNo), "phone_no"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.LoginType = basetypes.LoginType(basetypes.LoginType_value[info.LoginTypeStr])
	}
}

func (h *Handler) GetHistory(ctx context.Context) (*npool.History, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryHistory(cli)
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetHistories(ctx context.Context) ([]*npool.History, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLoginHistories(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Desc(entloginhistory.FieldUpdatedAt)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
