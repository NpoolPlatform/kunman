package poster

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostpostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/poster"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	enttopmost "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmost"
	enttopmostposter "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostposter"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostPosterSelect
}

func (h *baseQueryHandler) selectPoster(stm *ent.TopMostPosterQuery) *ent.TopMostPosterSelect {
	return stm.Select(enttopmostposter.FieldID)
}

func (h *baseQueryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.TopMostPoster.Query().Where(enttopmostposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *baseQueryHandler) queryPosters(cli *ent.Client) (*ent.TopMostPosterSelect, error) {
	stm, err := topmostpostercrud.SetQueryConds(cli.TopMostPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectPoster(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostposter.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostposter.FieldID),
			t.C(enttopmostposter.FieldID),
		).
		AppendSelect(
			t.C(enttopmostposter.FieldEntID),
			t.C(enttopmostposter.FieldTopMostID),
			t.C(enttopmostposter.FieldPoster),
			t.C(enttopmostposter.FieldIndex),
			t.C(enttopmostposter.FieldCreatedAt),
			t.C(enttopmostposter.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinTopMost(s *sql.Selector) error {
	t1 := sql.Table(enttopmost.Table)
	s.Join(t1).
		On(
			s.C(enttopmostposter.FieldTopMostID),
			t1.C(enttopmost.FieldEntID),
		)
	if h.TopMostConds.AppID != nil {
		id, ok := h.TopMostConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(enttopmost.FieldAppID), id),
		)
	}
	if h.TopMostConds.EntID != nil {
		id, ok := h.TopMostConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t1.C(enttopmost.FieldEntID), id),
		)
	}
	if h.TopMostConds.EntIDs != nil {
		ids, ok := h.TopMostConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(
			sql.In(t1.C(enttopmost.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t1.C(enttopmost.FieldAppID),
		t1.C(enttopmost.FieldTopMostType),
		sql.As(t1.C(enttopmost.FieldTitle), "top_most_title"),
		sql.As(t1.C(enttopmost.FieldMessage), "top_most_message"),
		sql.As(t1.C(enttopmost.FieldTargetURL), "top_most_target_url"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinTopMost(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMost", "Error", err)
		}
	})
}
