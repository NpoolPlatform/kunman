package cashcontrol

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	cashcontrolcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/app/cashcontrol"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/cashcontrol"
)

func (h *Handler) UpdateCashControl(ctx context.Context) (*npool.CashControl, error) {
	info, err := h.GetCashControl(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid cashcontrol")
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := cashcontrolcrud.UpdateSet(
			cli.CashControl.UpdateOneID(*h.ID),
			&cashcontrolcrud.Req{
				Value: h.Value,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCashControl(ctx)
}
