package fee

import (
	"context"

	feemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	feemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"

	"github.com/google/uuid"
)

func (h *Handler) CreateFee(ctx context.Context) (*feemwpb.Fee, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := feemwcli.CreateFee(ctx, &feemwpb.FeeReq{
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
