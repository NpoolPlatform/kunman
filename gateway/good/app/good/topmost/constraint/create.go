package constraint

import (
	"context"

	constraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/constraint"
	constraintmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"

	"github.com/google/uuid"
)

func (h *Handler) CreateConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	if err := h.CheckTopMost(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := constraintmwcli.CreateTopMostConstraint(ctx, &constraintmwpb.TopMostConstraintReq{
		EntID:       h.EntID,
		TopMostID:   h.TopMostID,
		Constraint:  h.Constraint,
		TargetValue: h.TargetValue,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}

	return h.GetConstraint(ctx)
}
