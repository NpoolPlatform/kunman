package fee

import (
	"context"

	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	feemw "github.com/NpoolPlatform/kunman/middleware/good/fee"

	"github.com/google/uuid"
)

func (h *Handler) CreateFee(ctx context.Context) (*feemwpb.Fee, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := feemw.NewHandler(
		ctx,
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

	if err := handler.CreateFee(ctx); err != nil {
		return nil, err
	}
	return h.GetFee(ctx)
}
