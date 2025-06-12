package user

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/task/user"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enttaskuser "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/taskuser"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/task/user"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TaskUserSelect
	stmCount  *ent.TaskUserSelect
	infos     []*npool.TaskUser
	total     uint32
}

func (h *queryHandler) selectTaskUser(stm *ent.TaskUserQuery) {
	h.stmSelect = stm.Select(enttaskuser.FieldID)
}

func (h *queryHandler) queryTaskUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.TaskUser.Query().Where(enttaskuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttaskuser.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttaskuser.EntID(*h.EntID))
	}
	h.selectTaskUser(stm)
	return nil
}

func (h *queryHandler) queryTaskUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.TaskUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectTaskUser(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(enttaskuser.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttaskuser.FieldEntID),
			t1.C(enttaskuser.FieldEntID),
		).
		AppendSelect(
			t1.C(enttaskuser.FieldEntID),
			t1.C(enttaskuser.FieldAppID),
			t1.C(enttaskuser.FieldUserID),
			t1.C(enttaskuser.FieldTaskID),
			t1.C(enttaskuser.FieldEventID),
			t1.C(enttaskuser.FieldTaskState),
			t1.C(enttaskuser.FieldRewardState),
			t1.C(enttaskuser.FieldCreatedAt),
			t1.C(enttaskuser.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TaskState = basetypes.TaskState(basetypes.TaskState_value[info.TaskStateStr])
		info.RewardState = basetypes.RewardState(basetypes.RewardState_value[info.RewardStateStr])
	}
}

func (h *Handler) GetTaskUser(ctx context.Context) (*npool.TaskUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTaskUser(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetTaskUsers(ctx context.Context) ([]*npool.TaskUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTaskUsers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
