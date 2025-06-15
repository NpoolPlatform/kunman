package oneshot

import (
	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entapponeshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscriptiononeshot"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entoneshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscriptiononeshot"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldCreatedAt)
}

func (h *baseQueryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectAppGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		OnP(
			sql.EQ(t1.C(entappgoodbase.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodbase.FieldEntID), "app_good_id"),
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldProductPage),
			t1.C(entappgoodbase.FieldBanner),
			t1.C(entappgoodbase.FieldCreatedAt),
			t1.C(entappgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinGood(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_OneShot.String()),
		).
		OnP(
			sql.EQ(t1.C(entgoodbase.FieldDeletedAt), 0),
		)
	if h.GoodBaseConds.EntID != nil {
		uid, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldEntID), uid))
	}
	if h.GoodBaseConds.EntIDs != nil {
		uids, ok := h.GoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(
				t1.C(entgoodbase.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.GoodBaseConds.GoodType != nil {
		_type, ok := h.GoodBaseConds.GoodType.Val.(types.GoodType)
		if !ok {
			return wlog.Errorf("invalid goodtype")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldGoodType), _type.String()))
	}
	s.AppendSelect(
		t1.C(entgoodbase.FieldGoodType),
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinOneShot(s *sql.Selector) error {
	t1 := sql.Table(entoneshot.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entoneshot.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entoneshot.FieldDeletedAt), 0),
		)
	if h.OneShotConds.GoodID != nil {
		uid, ok := h.OneShotConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entoneshot.FieldGoodID), uid))
	}
	if h.OneShotConds.GoodIDs != nil {
		uids, ok := h.OneShotConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(
				t1.C(entoneshot.FieldGoodID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	s.AppendSelect(
		t1.C(entoneshot.FieldQuota),
	)
	return nil
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinAppOneShot(s *sql.Selector) error {
	t1 := sql.Table(entapponeshot.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entapponeshot.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t1.C(entapponeshot.FieldDeletedAt), 0),
		)
	if h.ID != nil {
		s.OnP(sql.EQ(t1.C(entapponeshot.FieldID), *h.ID))
	}
	if h.AppOneShotConds.ID != nil {
		u, ok := h.AppOneShotConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entapponeshot.FieldID), u))
	}
	if h.AppOneShotConds.IDs != nil {
		ids, ok := h.AppOneShotConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(
			sql.In(
				t1.C(entapponeshot.FieldID),
				func() (_ids []interface{}) {
					for _, id := range ids {
						_ids = append(_ids, id)
					}
					return
				}()...,
			),
		)
	}
	if h.EntID != nil {
		s.OnP(sql.EQ(t1.C(entapponeshot.FieldEntID), *h.EntID))
	}
	if h.AppOneShotConds.EntID != nil {
		uid, ok := h.AppOneShotConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entapponeshot.FieldEntID), uid))
	}
	if h.AppOneShotConds.EntIDs != nil {
		uids, ok := h.AppOneShotConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(
			sql.In(
				t1.C(entapponeshot.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.AppOneShotConds.AppGoodID != nil {
		uid, ok := h.AppOneShotConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(sql.EQ(t1.C(entapponeshot.FieldAppGoodID), uid))
	}
	if h.AppOneShotConds.AppGoodIDs != nil {
		uids, ok := h.AppOneShotConds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(
				t1.C(entapponeshot.FieldAppGoodID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	s.AppendSelect(
		t1.C(entapponeshot.FieldID),
		t1.C(entapponeshot.FieldEntID),
		t1.C(entapponeshot.FieldUsdPrice),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppOneShot(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppOneShot", "Error", err)
			return
		}
		if err := h.queryJoinGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinGood", "Error", err)
			return
		}
		if err := h.queryJoinOneShot(s); err != nil {
			logger.Sugar().Errorw("queryJoinOneShot", "Error", err)
		}
	})
}
