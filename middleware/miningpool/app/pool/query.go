package apppool

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/app/pool"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	entapppool "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
	entpool "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/pool"

	apppoolcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/app/pool"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppPoolSelect
	infos []*npool.Pool
	total uint32
}

func (h *queryHandler) selectPool(stm *ent.AppPoolQuery) {
	h.stm = stm.Select(
		entapppool.FieldID,
		entapppool.FieldEntID,
		entapppool.FieldAppID,
		entapppool.FieldPoolID,
		entapppool.FieldCreatedAt,
		entapppool.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(
		h.queryJoinPool,
	)
}

func (h *queryHandler) queryJoinPool(s *sql.Selector) {
	poolT := sql.Table(entpool.Table)
	s.Join(poolT).On(
		s.C(entapppool.FieldPoolID),
		poolT.C(entpool.FieldEntID),
	).OnP(
		sql.EQ(poolT.C(entpool.FieldDeletedAt), 0),
	)
}

func (h *queryHandler) queryPool(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppPool.Query().Where(entapppool.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapppool.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entapppool.EntID(*h.EntID))
	}
	h.selectPool(stm)
	return nil
}

func (h *queryHandler) queryPools(ctx context.Context, cli *ent.Client) error {
	stm, err := apppoolcrud.SetQueryConds(cli.AppPool.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := apppoolcrud.SetQueryConds(cli.AppPool.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(
		h.queryJoinPool,
	)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectPool(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
}

func (h *Handler) GetPool(ctx context.Context) (*npool.Pool, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPool(cli); err != nil {
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

func (h *Handler) GetPools(ctx context.Context) ([]*npool.Pool, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPools(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entapppool.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
