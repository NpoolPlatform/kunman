package malfunction

import (
	"context"

	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkMalfunction(ctx); err != nil {
		return nil, err
	}

	if err := malfunctionmwcli.UpdateMalfunction(ctx, &malfunctionmwpb.MalfunctionReq{
		ID:                h.ID,
		EntID:             h.EntID,
		Title:             h.Title,
		Message:           h.Message,
		StartAt:           h.StartAt,
		DurationSeconds:   h.DurationSeconds,
		CompensateSeconds: h.CompensateSeconds,
	}); err != nil {
		return nil, err
	}
	return h.GetMalfunction(ctx)
}
