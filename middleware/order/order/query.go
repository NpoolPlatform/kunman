package order

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderBaseSelect
	infos    []*npool.Order
	total    uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.GoodType = goodtypes.GoodType(goodtypes.GoodType_value[info.GoodTypeStr])
		info.OrderType = types.OrderType(types.OrderType_value[info.OrderTypeStr])
		info.PaymentType = types.PaymentType(types.PaymentType_value[info.PaymentTypeStr])
		info.OrderState = types.OrderState(types.OrderState_value[info.OrderStateStr])
		info.CreateMethod = types.OrderCreateMethod(types.OrderCreateMethod_value[info.CreateMethodStr])
	}
}

func (h *Handler) GetOrder(ctx context.Context) (*npool.Order, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.Offset(int(0)).Limit(int(2))
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

func (h *Handler) GetOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entorderbase.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
