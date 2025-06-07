package comment

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	commentcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/comment"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entcomment "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/comment"
	entscore "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/score"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.CommentSelect
}

func (h *baseQueryHandler) selectComment(stm *ent.CommentQuery) *ent.CommentSelect {
	return stm.Select(entcomment.FieldID)
}

func (h *baseQueryHandler) queryComment(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Comment.Query().Where(entcomment.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcomment.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcomment.EntID(*h.EntID))
	}
	h.stmSelect = h.selectComment(stm)
	return nil
}

func (h *baseQueryHandler) queryComments(cli *ent.Client) (*ent.CommentSelect, error) {
	stm, err := commentcrud.SetQueryConds(cli.Comment.Query(), h.CommentConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectComment(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcomment.Table)
	s.LeftJoin(t).
		On(
			s.C(entcomment.FieldID),
			t.C(entcomment.FieldID),
		).
		AppendSelect(
			t.C(entcomment.FieldEntID),
			t.C(entcomment.FieldUserID),
			t.C(entcomment.FieldAppGoodID),
			t.C(entcomment.FieldOrderID),
			t.C(entcomment.FieldContent),
			t.C(entcomment.FieldReplyToID),
			t.C(entcomment.FieldAnonymous),
			t.C(entcomment.FieldPurchasedUser),
			t.C(entcomment.FieldTrialUser),
			t.C(entcomment.FieldHide),
			t.C(entcomment.FieldHideReason),
			t.C(entcomment.FieldCreatedAt),
			t.C(entcomment.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGoodBase(s *sql.Selector) error {
	t := sql.Table(entappgoodbase.Table)
	s.Join(t).
		On(
			s.C(entcomment.FieldAppGoodID),
			t.C(entappgoodbase.FieldEntID),
		)
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
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
			sql.EQ(t.C(entappgoodbase.FieldGoodID), id),
		)
	}
	if h.AppGoodBaseConds.GoodIDs != nil {
		ids, ok := h.AppGoodBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t.C(entappgoodbase.FieldGoodID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t.C(entappgoodbase.FieldAppID),
		t.C(entappgoodbase.FieldGoodID),
		sql.As(t.C(entappgoodbase.FieldName), "good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinScore(s *sql.Selector) {
	t := sql.Table(entscore.Table)
	s.LeftJoin(t).
		On(
			s.C(entcomment.FieldEntID),
			t.C(entscore.FieldCommentID),
		).
		AppendSelect(
			t.C(entscore.FieldScore),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGoodBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGoodBase", "Error", err)
		}
		h.queryJoinScore(s)
	})
}
