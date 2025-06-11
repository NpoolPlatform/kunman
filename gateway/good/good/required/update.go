package required

import (
	"context"

	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
	requiredmw "github.com/NpoolPlatform/kunman/middleware/good/good/required"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkRequired(ctx); err != nil {
		return nil, err
	}

	requiredHandler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithID(h.ID, true),
		requiredmw.WithEntID(h.EntID, true),
		requiredmw.WithMust(h.Must, false),
	)
	if err != nil {
		return nil, err
	}

	if err := requiredHandler.UpdateRequired(ctx); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
