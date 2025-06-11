package fee

import (
	"context"

	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	feemw "github.com/NpoolPlatform/kunman/middleware/good/fee"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateFee(ctx context.Context) (*feemwpb.Fee, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkFee(ctx); err != nil {
		return nil, err
	}

	feeHandler, err := feemw.NewHandler(
		ctx,
		feemw.WithID(h.ID, true),
		feemw.WithEntID(h.EntID, true),
		feemw.WithGoodID(h.GoodID, true),
		feemw.WithGoodType(h.GoodType, true),
		feemw.WithName(h.Name, true),
		feemw.WithSettlementType(h.SettlementType, true),
		feemw.WithUnitValue(h.UnitValue, true),
		feemw.WithDurationDisplayType(h.DurationDisplayType, true),
	)
	if err != nil {
		return nil, err
	}

	if err := feeHandler.UpdateFee(ctx); err != nil {
		return nil, err
	}
	return h.GetFee(ctx)
}
