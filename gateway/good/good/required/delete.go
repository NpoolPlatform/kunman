package required

import (
	"context"

	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
	requiredmw "github.com/NpoolPlatform/kunman/middleware/good/good/required"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkRequired(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetRequired(ctx)
	if err != nil {
		return nil, err
	}

	requiredHandler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithID(h.ID, true),
		requiredmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := requiredHandler.DeleteRequired(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
