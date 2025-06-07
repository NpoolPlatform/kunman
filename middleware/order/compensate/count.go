package compensate

import (
	"context"

	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/compensate"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entcompensate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/compensate"
)

type countHandler struct {
	*baseQueryHandler
}

func (h *Handler) CountCompensates(ctx context.Context) (total uint32, err error) {
	handler := &countHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCompensates(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_total, err := handler.stmSelect.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		total = uint32(_total)
		return nil
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return total, nil
}

func (h *Handler) CountCompensateOrders(ctx context.Context) (infos []*npool.CompensateOrderNumber, err error) {
	handler := &countHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryCompensates(cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.stmSelect.
			GroupBy(entcompensate.FieldCompensateFromID).
			Aggregate(func(s *sql.Selector) string {
				return sql.As(sql.Count("*"), "orders")
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return infos, nil
}
