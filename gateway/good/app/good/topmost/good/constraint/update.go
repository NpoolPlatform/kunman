package constraint

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
	topmostconstraintmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/constraint"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateConstraint(ctx context.Context) (*npool.TopMostGoodConstraint, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkConstraint(ctx); err != nil {
		return nil, err
	}

	constraintHandler, err := topmostconstraintmw.NewHandler(
		ctx,
		topmostconstraintmw.WithID(h.ID, true),
		topmostconstraintmw.WithEntID(h.EntID, true),
		topmostconstraintmw.WithTargetValue(h.TargetValue, false),
		topmostconstraintmw.WithIndex(h.Index, false),
	)
	if err != nil {
		return nil, err
	}

	if err := constraintHandler.UpdateConstraint(ctx); err != nil {
		return nil, err
	}
	return h.GetConstraint(ctx)
}
