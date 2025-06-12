package goodachievement

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodachievementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/good"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entgoodachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodachievement"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodAchievementSelect
}

func (h *baseQueryHandler) selectGoodAchievement(stm *ent.GoodAchievementQuery) *ent.GoodAchievementSelect {
	return stm.Select(entgoodachievement.FieldID)
}

func (h *baseQueryHandler) queryGoodAchievement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.GoodAchievement.Query().Where(entgoodachievement.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entgoodachievement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodachievement.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entgoodachievement.AppGoodID(*h.AppGoodID))
	}
	h.stmSelect = h.selectGoodAchievement(stm)
	return nil
}

func (h *baseQueryHandler) queryGoodAchievements(cli *ent.Client) (*ent.GoodAchievementSelect, error) {
	stm, err := goodachievementcrud.SetQueryConds(cli.GoodAchievement.Query(), h.GoodAchievementConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGoodAchievement(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodachievement.Table)
	s.AppendSelect(
		t.C(entgoodachievement.FieldID),
		t.C(entgoodachievement.FieldEntID),
		t.C(entgoodachievement.FieldAppID),
		t.C(entgoodachievement.FieldUserID),
		t.C(entgoodachievement.FieldGoodID),
		t.C(entgoodachievement.FieldAppGoodID),
		t.C(entgoodachievement.FieldTotalUnits),
		t.C(entgoodachievement.FieldSelfUnits),
		t.C(entgoodachievement.FieldTotalAmountUsd),
		t.C(entgoodachievement.FieldSelfAmountUsd),
		t.C(entgoodachievement.FieldTotalCommissionUsd),
		t.C(entgoodachievement.FieldSelfCommissionUsd),
		t.C(entgoodachievement.FieldCreatedAt),
		t.C(entgoodachievement.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
