package constraint

import (
	"context"

	constraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"
	constraintmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"

	"github.com/google/uuid"
)

func (h *Handler) CreateConstraint(ctx context.Context) (*npool.TopMostGoodConstraint, error) {
	if err := h.CheckTopMostGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := constraintmwcli.CreateTopMostGoodConstraint(ctx, &constraintmwpb.TopMostGoodConstraintReq{
		EntID:         h.EntID,
		TopMostGoodID: h.TopMostGoodID,
		Constraint:    h.Constraint,
		TargetValue:   h.TargetValue,
		Index:         h.Index,
	}); err != nil {
		return nil, err
	}

	return h.GetConstraint(ctx)
}
