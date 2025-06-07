package label

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodlabelcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/label"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappgoodlabel "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodlabel"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodLabelSelect
}

func (h *baseQueryHandler) selectLabel(stm *ent.AppGoodLabelQuery) *ent.AppGoodLabelSelect {
	return stm.Select(entappgoodlabel.FieldID)
}

func (h *baseQueryHandler) queryLabel(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppGoodLabel.Query().Where(entappgoodlabel.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgoodlabel.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodlabel.EntID(*h.EntID))
	}
	h.stmSelect = h.selectLabel(stm)
	return nil
}

func (h *baseQueryHandler) queryLabels(cli *ent.Client) (*ent.AppGoodLabelSelect, error) {
	stm, err := appgoodlabelcrud.SetQueryConds(cli.AppGoodLabel.Query(), h.LabelConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectLabel(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgoodlabel.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodlabel.FieldID),
			t.C(entappgoodlabel.FieldID),
		).
		AppendSelect(
			t.C(entappgoodlabel.FieldEntID),
			t.C(entappgoodlabel.FieldAppGoodID),
			t.C(entappgoodlabel.FieldIcon),
			t.C(entappgoodlabel.FieldIconBgColor),
			t.C(entappgoodlabel.FieldLabel),
			t.C(entappgoodlabel.FieldLabelBgColor),
			t.C(entappgoodlabel.FieldIndex),
			t.C(entappgoodlabel.FieldCreatedAt),
			t.C(entappgoodlabel.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodlabel.FieldAppGoodID),
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
