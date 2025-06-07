package orderlock

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/lock"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderLockSelect
	infos    []*npool.OrderLock
	total    uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinOrderBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderBase", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.LockType = types.OrderLockType(types.OrderLockType_value[info.LockTypeStr])
		info.GoodType = goodtypes.GoodType(goodtypes.GoodType_value[info.GoodTypeStr])
	}
}

func (h *queryHandler) getOrderLockWithClient(ctx context.Context, cli *ent.Client) error {
	if err := h.queryOrderLock(cli); err != nil {
		return wlog.WrapError(err)
	}
	h.queryJoin()
	if err := h.scan(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if len(h.infos) > 1 {
		return wlog.Errorf("too many records")
	}

	h.formalize()
	return nil
}

func (h *Handler) GetOrderLockWithTx(ctx context.Context, tx *ent.Tx) (*npool.OrderLock, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := handler.getOrderLockWithClient(ctx, tx.Client()); err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	return handler.infos[0], nil
}

func (h *Handler) GetOrderLock(ctx context.Context) (*npool.OrderLock, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.getOrderLockWithClient(_ctx, cli)
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	return handler.infos[0], nil
}

func (h *Handler) GetOrderLocks(ctx context.Context) (infos []*npool.OrderLock, total uint32, err error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryOrderLocks(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryOrderLocks(cli); err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.
			stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entorderlock.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	}); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
