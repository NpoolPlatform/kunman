package malfunction

import (
	"context"

	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/malfunction"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkMalfunction(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetMalfunction(ctx)
	if err != nil {
		return nil, err
	}
	if err := malfunctionmwcli.DeleteMalfunction(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
