package topmost

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	enttopmost "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmost"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostSelect
}

func (h *baseQueryHandler) selectTopMost(stm *ent.TopMostQuery) *ent.TopMostSelect {
	return stm.Select(enttopmost.FieldID)
}

func (h *baseQueryHandler) queryTopMost(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.TopMost.Query().Where(enttopmost.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmost.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmost.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMost(stm)
	return nil
}

func (h *baseQueryHandler) queryTopMosts(cli *ent.Client) (*ent.TopMostSelect, error) {
	stm, err := topmostcrud.SetQueryConds(cli.TopMost.Query(), h.Conds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectTopMost(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmost.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmost.FieldID),
			t.C(enttopmost.FieldID),
		).
		AppendSelect(
			t.C(enttopmost.FieldEntID),
			t.C(enttopmost.FieldAppID),
			t.C(enttopmost.FieldTopMostType),
			t.C(enttopmost.FieldTitle),
			t.C(enttopmost.FieldMessage),
			t.C(enttopmost.FieldTargetURL),
			t.C(enttopmost.FieldStartAt),
			t.C(enttopmost.FieldEndAt),
			t.C(enttopmost.FieldCreatedAt),
			t.C(enttopmost.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
