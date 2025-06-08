package fee

import (
	"context"

	feemwcli "github.com/NpoolPlatform/kunman/middleware/good/fee"
	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
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
	if err := feemwcli.UpdateFee(ctx, &feemwpb.FeeReq{
		ID:                  h.ID,
		EntID:               h.EntID,
		GoodID:              h.GoodID,
		GoodType:            h.GoodType,
		Name:                h.Name,
		SettlementType:      h.SettlementType,
		UnitValue:           h.UnitValue,
		DurationDisplayType: h.DurationDisplayType,
	}); err != nil {
		return nil, err
	}
	return h.GetFee(ctx)
}
