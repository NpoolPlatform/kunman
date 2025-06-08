package required

import (
	"context"

	requiredappgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/required"
	requiredappgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
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
	if err := requiredappgoodmwcli.DeleteRequired(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
