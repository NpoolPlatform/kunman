package required

import (
	"context"

	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteRequired(ctx context.Context) (*requiredappgoodmwpb.Required, error) {
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

	requiredHandler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithID(h.ID, true),
		requiredappgoodmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := requiredHandler.DeleteRequired(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
