package powerrental

import (
	"context"

	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type sumHandler struct {
	*baseQueryHandler
}

func (h *Handler) SumPowerRentalUnits(ctx context.Context) (units string, err error) {
	handler := &sumHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	_units := []struct {
		AppID uuid.UUID       `json:"app_id"`
		Units decimal.Decimal `json:"total_units"`
	}{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		return wlog.WrapError(handler.stmSelect.
			GroupBy(entorderbase.FieldAppID).
			Aggregate(func(s *sql.Selector) string {
				return sql.As(sql.Sum("units"), "total_units")
			}).
			Scan(ctx, &_units))
	})
	if err != nil {
		return decimal.NewFromInt(0).String(), wlog.WrapError(err)
	}
	if len(_units) == 0 {
		return decimal.NewFromInt(0).String(), nil
	}
	if len(_units) > 1 {
		return decimal.NewFromInt(0).String(), wlog.Errorf("invalid units")
	}
	return _units[0].Units.String(), nil
}
