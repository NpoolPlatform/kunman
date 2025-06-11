package oneshot

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entoneshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscriptiononeshot"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.GoodType(types.GoodType_OneShot.String()),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) error {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectGoodBase(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.Join(t).
		On(
			s.C(entgoodbase.FieldID),
			t.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t.C(entgoodbase.FieldGoodType),
			t.C(entgoodbase.FieldName),
			t.C(entgoodbase.FieldCreatedAt),
			t.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinOneShot(s *sql.Selector) error {
	t1 := sql.Table(entoneshot.Table)
	s.Join(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entoneshot.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entoneshot.FieldDeletedAt), 0),
		)
	if h.ID != nil {
		s.OnP(sql.EQ(t1.C(entoneshot.FieldID), *h.ID))
	}
	if h.OneShotConds.ID != nil {
		u, ok := h.OneShotConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entoneshot.FieldID), u))
	}
	if h.OneShotConds.IDs != nil {
		ids, ok := h.OneShotConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entoneshot.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, id)
			}
			return
		}()...))
	}
	if h.EntID != nil {
		s.OnP(sql.EQ(t1.C(entoneshot.FieldEntID), *h.EntID))
	}
	if h.OneShotConds.EntID != nil {
		uid, ok := h.OneShotConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entoneshot.FieldEntID), uid))
	}
	if h.OneShotConds.EntIDs != nil {
		uids, ok := h.OneShotConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entoneshot.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, uid)
			}
			return
		}()...))
	}
	s.AppendSelect(
		t1.C(entoneshot.FieldID),
		t1.C(entoneshot.FieldEntID),
		t1.C(entoneshot.FieldGoodID),
		t1.C(entoneshot.FieldQuota),
		t1.C(entoneshot.FieldUsdPrice),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOneShot(s); err != nil {
			logger.Sugar().Errorw("queryJoinOneShot", "Error", err)
		}
	})
}
