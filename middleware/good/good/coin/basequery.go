package coin

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoincrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodCoinSelect
}

func (h *baseQueryHandler) selectGoodCoin(stm *ent.GoodCoinQuery) *ent.GoodCoinSelect {
	return stm.Select(entgoodcoin.FieldID)
}

func (h *baseQueryHandler) queryGoodCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid goodcoinid")
	}
	stm := cli.GoodCoin.
		Query().
		Where(
			entgoodcoin.DeletedAt(0),
		)
	if h.ID != nil {
		stm.Where(entgoodcoin.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodcoin.EntID(*h.EntID))
	}
	h.stmSelect = h.selectGoodCoin(stm)
	return nil
}

func (h *baseQueryHandler) queryGoodCoins(cli *ent.Client) (*ent.GoodCoinSelect, error) {
	stm, err := goodcoincrud.SetQueryConds(cli.GoodCoin.Query(), h.GoodCoinConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGoodCoin(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodcoin.FieldID),
			t1.C(entgoodcoin.FieldID),
		).
		AppendSelect(
			t1.C(entgoodcoin.FieldEntID),
			t1.C(entgoodcoin.FieldGoodID),
			t1.C(entgoodcoin.FieldCoinTypeID),
			t1.C(entgoodcoin.FieldMain),
			t1.C(entgoodcoin.FieldIndex),
			t1.C(entgoodcoin.FieldCreatedAt),
			t1.C(entgoodcoin.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodcoin.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		)
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntID != nil {
		uid, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldEntID), uid))
	}
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntIDs != nil {
		uids, ok := h.GoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t1.C(entgoodbase.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	s.AppendSelect(
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
		t1.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodBase", "Error", err)
		}
	})
}
