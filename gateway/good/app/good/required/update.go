package required

import (
	"context"

	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateRequired(ctx context.Context) (*requiredappgoodmwpb.Required, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkRequired(ctx); err != nil {
		return nil, err
	}

	requiredHandler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithID(h.ID, true),
		requiredappgoodmw.WithEntID(h.EntID, true),
		requiredappgoodmw.WithMust(h.Must, true),
	)
	if err != nil {
		return nil, err
	}

	if err := requiredHandler.UpdateRequired(ctx); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
