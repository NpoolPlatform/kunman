package required

import (
	"context"

	requiredappgoodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/required"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
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

	if err := requiredappgoodmwcli.UpdateRequired(ctx, &requiredappgoodmwpb.RequiredReq{
		ID:    h.ID,
		EntID: h.EntID,
		Must:  h.Must,
	}); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
