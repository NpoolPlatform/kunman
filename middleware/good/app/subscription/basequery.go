package subscription

import (
	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscription"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscription"

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
			sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_Subscription.String()),
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

func (h *baseQueryHandler) queryJoinSubscription(s *sql.Selector) error {
	t1 := sql.Table(entsubscription.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entsubscription.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entsubscription.FieldDeletedAt), 0),
		)
	if h.SubscriptionConds.GoodID != nil {
		uid, ok := h.SubscriptionConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entsubscription.FieldGoodID), uid))
	}
	if h.SubscriptionConds.GoodIDs != nil {
		uids, ok := h.SubscriptionConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(
				t1.C(entsubscription.FieldGoodID),
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
		t1.C(entsubscription.FieldDurationUnits),
		t1.C(entsubscription.FieldDurationQuota),
		t1.C(entsubscription.FieldDailyBonusQuota),
		t1.C(entsubscription.FieldDurationDisplayType),
	)
	return nil
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinAppSubscription(s *sql.Selector) error {
	t1 := sql.Table(entappsubscription.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entappsubscription.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t1.C(entappsubscription.FieldDeletedAt), 0),
		)
	if h.ID != nil {
		s.OnP(sql.EQ(t1.C(entappsubscription.FieldID), *h.ID))
	}
	if h.AppSubscriptionConds.ID != nil {
		u, ok := h.AppSubscriptionConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entappsubscription.FieldID), u))
	}
	if h.AppSubscriptionConds.IDs != nil {
		ids, ok := h.AppSubscriptionConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(
			sql.In(
				t1.C(entappsubscription.FieldID),
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
		s.OnP(sql.EQ(t1.C(entappsubscription.FieldEntID), *h.EntID))
	}
	if h.AppSubscriptionConds.EntID != nil {
		uid, ok := h.AppSubscriptionConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entappsubscription.FieldEntID), uid))
	}
	if h.AppSubscriptionConds.EntIDs != nil {
		uids, ok := h.AppSubscriptionConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(
			sql.In(
				t1.C(entappsubscription.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, uid)
					}
					return _uids
				}()...,
			),
		)
	}
	if h.AppSubscriptionConds.AppGoodID != nil {
		uid, ok := h.AppSubscriptionConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(sql.EQ(t1.C(entappsubscription.FieldAppGoodID), uid))
	}
	if h.AppSubscriptionConds.AppGoodIDs != nil {
		uids, ok := h.AppSubscriptionConds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(
				t1.C(entappsubscription.FieldAppGoodID),
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
		t1.C(entappsubscription.FieldID),
		t1.C(entappsubscription.FieldEntID),
		t1.C(entappsubscription.FieldUsdPrice),
		t1.C(entappsubscription.FieldProductID),
		t1.C(entappsubscription.FieldPlanID),
		t1.C(entappsubscription.FieldTrialUsdPrice),
		t1.C(entappsubscription.FieldTrialUnits),
		t1.C(entappsubscription.FieldPriceFiatID),
		t1.C(entappsubscription.FieldFiatPrice),
		t1.C(entappsubscription.FieldTrialFiatPrice),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppSubscription(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppSubscription", "Error", err)
			return
		}
		if err := h.queryJoinGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinGood", "Error", err)
			return
		}
		if err := h.queryJoinSubscription(s); err != nil {
			logger.Sugar().Errorw("queryJoinSubscription", "Error", err)
		}
	})
}
