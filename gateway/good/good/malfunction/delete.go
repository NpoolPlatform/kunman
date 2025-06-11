package malfunction

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
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

	malfunctionHandler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithID(h.ID, true),
		malfunctionmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := malfunctionHandler.DeleteMalfunction(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
