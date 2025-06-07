package displayname

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaynamecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/name"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappgooddisplayname "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddisplayname"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodDisplayNameSelect
}

func (h *baseQueryHandler) selectDisplayName(stm *ent.AppGoodDisplayNameQuery) *ent.AppGoodDisplayNameSelect {
	return stm.Select(entappgooddisplayname.FieldID)
}

func (h *baseQueryHandler) queryDisplayName(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppGoodDisplayName.Query().Where(entappgooddisplayname.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgooddisplayname.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgooddisplayname.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDisplayName(stm)
	return nil
}

func (h *baseQueryHandler) queryDisplayNames(cli *ent.Client) (*ent.AppGoodDisplayNameSelect, error) {
	stm, err := appgooddisplaynamecrud.SetQueryConds(cli.AppGoodDisplayName.Query(), h.DisplayNameConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectDisplayName(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgooddisplayname.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgooddisplayname.FieldID),
			t.C(entappgooddisplayname.FieldID),
		).
		AppendSelect(
			t.C(entappgooddisplayname.FieldEntID),
			t.C(entappgooddisplayname.FieldAppGoodID),
			t.C(entappgooddisplayname.FieldName),
			t.C(entappgooddisplayname.FieldIndex),
			t.C(entappgooddisplayname.FieldCreatedAt),
			t.C(entappgooddisplayname.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgooddisplayname.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)

	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	if h.AppGoodBaseConds.GoodID != nil {
		id, ok := h.AppGoodBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldGoodID), id),
		)
	}
	if h.AppGoodBaseConds.GoodIDs != nil {
		ids, ok := h.AppGoodBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldGoodID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
		sql.As(t2.C(entgoodbase.FieldEntID), "good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "good_name"),
		t2.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGood", "Error", err)
		}
	})
}
