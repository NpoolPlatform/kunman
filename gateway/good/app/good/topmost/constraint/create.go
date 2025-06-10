package constraint

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
	constraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/constraint"

	"github.com/google/uuid"
)

func (h *Handler) CreateConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	if err := h.CheckTopMost(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := constraintmw.NewHandler(
		ctx,
		constraintmw.WithEntID(h.EntID, true),
		constraintmw.WithTopMostID(h.TopMostID, true),
		constraintmw.WithConstraint(h.Constraint, true),
		constraintmw.WithTargetValue(h.TargetValue, true),
		constraintmw.WithIndex(h.Index, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateConstraint(ctx); err != nil {
		return nil, err
	}

	return h.GetConstraint(ctx)
}
