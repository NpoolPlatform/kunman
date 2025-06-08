package constraint

import (
	"context"

	topmostconstraintmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/constraint"
	topmostconstraintmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
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

	if err := topmostconstraintmwcli.UpdateTopMostGoodConstraint(ctx, &topmostconstraintmwpb.TopMostGoodConstraintReq{
		ID:          h.ID,
		EntID:       h.EntID,
		TargetValue: h.TargetValue,
		Index:       h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetConstraint(ctx)
}
