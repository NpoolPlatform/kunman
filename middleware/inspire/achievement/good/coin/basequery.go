package goodcoinachievement

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoinachievementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/good/coin"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entgoodcoinachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodcoinachievement"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodCoinAchievementSelect
}

func (h *baseQueryHandler) selectGoodCoinAchievement(stm *ent.GoodCoinAchievementQuery) *ent.GoodCoinAchievementSelect {
	return stm.Select(entgoodcoinachievement.FieldID)
}

func (h *baseQueryHandler) queryGoodCoinAchievement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.GoodCoinTypeID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.GoodCoinAchievement.Query().Where(entgoodcoinachievement.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entgoodcoinachievement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodcoinachievement.EntID(*h.EntID))
	}
	if h.GoodCoinTypeID != nil {
		stm.Where(entgoodcoinachievement.GoodCoinTypeID(*h.GoodCoinTypeID))
	}
	h.stmSelect = h.selectGoodCoinAchievement(stm)
	return nil
}

func (h *baseQueryHandler) queryGoodCoinAchievements(cli *ent.Client) (*ent.GoodCoinAchievementSelect, error) {
	stm, err := goodcoinachievementcrud.SetQueryConds(cli.GoodCoinAchievement.Query(), h.GoodCoinAchievementConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGoodCoinAchievement(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodcoinachievement.Table)
	s.AppendSelect(
		t.C(entgoodcoinachievement.FieldID),
		t.C(entgoodcoinachievement.FieldEntID),
		t.C(entgoodcoinachievement.FieldAppID),
		t.C(entgoodcoinachievement.FieldUserID),
		t.C(entgoodcoinachievement.FieldGoodCoinTypeID),
		t.C(entgoodcoinachievement.FieldTotalUnits),
		t.C(entgoodcoinachievement.FieldSelfUnits),
		t.C(entgoodcoinachievement.FieldTotalAmountUsd),
		t.C(entgoodcoinachievement.FieldSelfAmountUsd),
		t.C(entgoodcoinachievement.FieldTotalCommissionUsd),
		t.C(entgoodcoinachievement.FieldSelfAmountUsd),
		t.C(entgoodcoinachievement.FieldCreatedAt),
		t.C(entgoodcoinachievement.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
