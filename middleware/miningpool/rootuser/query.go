package rootuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/rootuser"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"
	rootuserent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/rootuser"

	rootusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/rootuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.RootUserSelect
	infos []*npool.RootUser
	total uint32
}

func (h *queryHandler) selectRootUser(stm *ent.RootUserQuery) {
	h.stm = stm.Select(
		rootuserent.FieldID,
		rootuserent.FieldEntID,
		rootuserent.FieldName,
		rootuserent.FieldPoolID,
		rootuserent.FieldEmail,
		rootuserent.FieldAuthToken,
		rootuserent.FieldAuthed,
		rootuserent.FieldRemark,
		rootuserent.FieldCreatedAt,
		rootuserent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryRootUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.RootUser.Query().Where(rootuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(rootuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(rootuserent.EntID(*h.EntID))
	}
	h.selectRootUser(stm)
	return nil
}

func (h *queryHandler) queryRootUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := rootusercrud.SetQueryConds(cli.RootUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := rootusercrud.SetQueryConds(cli.RootUser.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectRootUser(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(
		h.queryJoinPool,
	)
}

func (h *queryHandler) queryJoinPool(s *sql.Selector) {
	poolT := sql.Table(pool.Table)
	s.Join(poolT).On(
		s.C(rootuserent.FieldPoolID),
		poolT.C(pool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(pool.FieldDeletedAt), 0),
	).AppendSelect(
		poolT.C(pool.FieldMiningPoolType),
	).Distinct()
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningPoolType = basetypes.MiningPoolType(basetypes.MiningPoolType_value[info.MiningPoolTypeStr])
	}
}

func (h *Handler) GetRootUser(ctx context.Context) (*npool.RootUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRootUser(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetRootUsers(ctx context.Context) ([]*npool.RootUser, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRootUsers(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(rootuserent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
