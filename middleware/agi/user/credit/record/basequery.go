package record

import (
	"entgo.io/ent/dialect/sql"

	recordcrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/user/credit/record"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entrecord "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/usercreditrecord"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.UserCreditRecordSelect
}

func (h *queryHandler) queryJoin() {
	if h.stmSelect != nil {
		h.baseQueryHandler.queryJoin()
	}
}

func (h *baseQueryHandler) selectRecord(stm *ent.UserCreditRecordQuery) *ent.UserCreditRecordSelect {
	return stm.Select(entrecord.FieldID)
}

func (h *baseQueryHandler) queryRecord(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.UserCreditRecord.Query().Where(entrecord.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrecord.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrecord.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRecord(stm)
	return nil
}

func (h *baseQueryHandler) queryRecords(cli *ent.Client) error {
	stm, err := recordcrud.SetQueryConds(cli.UserCreditRecord.Query(), h.RecordConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectRecord(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entrecord.Table)
	s.Join(t1).
		On(
			s.C(entrecord.FieldID),
			t1.C(entrecord.FieldID),
		).
		AppendSelect(
			t1.C(entrecord.FieldEntID),
			t1.C(entrecord.FieldAppID),
			t1.C(entrecord.FieldUserID),
			t1.C(entrecord.FieldOperationType),
			t1.C(entrecord.FieldCreditsChange),
			t1.C(entrecord.FieldExtra),
			t1.C(entrecord.FieldCreatedAt),
			t1.C(entrecord.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
