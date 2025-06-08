package constraint

import (
	"context"

	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/constraint"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/constraint"
	topmostconstraintmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/constraint"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkConstraint(ctx); err != nil {
		return nil, err
	}

	if err := topmostconstraintmwcli.UpdateTopMostConstraint(ctx, &topmostconstraintmwpb.TopMostConstraintReq{
		ID:          h.ID,
		EntID:       h.EntID,
		TargetValue: h.TargetValue,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetConstraint(ctx)
}
