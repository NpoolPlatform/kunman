package poolorderuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	poolorderusermiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental/poolorderuser"
	poolorderusercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/poolorderuser"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	poolorderuserent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/poolorderuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.PoolOrderUserSelect
	infos []poolorderusermiddlewarepb.PoolOrderUser
	total uint32
}

func (h *queryHandler) selectPoolOrderUser(stm *ent.PoolOrderUserQuery) {
	h.stm = stm.Select(
		poolorderuserent.FieldID,
		poolorderuserent.FieldEntID,
		poolorderuserent.FieldOrderID,
		poolorderuserent.FieldPoolOrderUserID,
	)
}

func (h *queryHandler) queryPoolOrderUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.PoolOrderUser.Query().Where(poolorderuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(poolorderuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(poolorderuserent.EntID(*h.EntID))
	}
	h.selectPoolOrderUser(stm)
	return nil
}

func (h *queryHandler) queryPoolOrderUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := poolorderusercrud.SetQueryConds(cli.PoolOrderUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectPoolOrderUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetPoolOrderUser(ctx context.Context) (*poolorderusermiddlewarepb.PoolOrderUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoolOrderUser(cli); err != nil {
			return err
		}
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

	return &handler.infos[0], nil
}

func (h *Handler) GetPoolOrderUsers(ctx context.Context) ([]poolorderusermiddlewarepb.PoolOrderUser, uint32, error) {
	if h.PoolOrderUserID == nil {
		return nil, 0, nil
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoolOrderUsers(_ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	return handler.infos, handler.total, nil
}
