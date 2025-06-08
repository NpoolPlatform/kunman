package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	requiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
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

	if err := requiredmwcli.UpdateRequired(ctx, &requiredmwpb.RequiredReq{
		ID:   h.ID,
		Must: h.Must,
	}); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
