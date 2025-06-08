package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
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
	if err := requiredmwcli.DeleteRequired(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
