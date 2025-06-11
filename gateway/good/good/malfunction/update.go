package malfunction

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
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

	malfunctionHandler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithID(h.ID, true),
		malfunctionmw.WithEntID(h.EntID, true),
		malfunctionmw.WithTitle(h.Title, false),
		malfunctionmw.WithMessage(h.Message, false),
		malfunctionmw.WithStartAt(h.StartAt, false),
		malfunctionmw.WithDurationSeconds(h.DurationSeconds, false),
		malfunctionmw.WithCompensateSeconds(h.CompensateSeconds, false),
	)
	if err != nil {
		return nil, err
	}

	if err := malfunctionHandler.UpdateMalfunction(ctx); err != nil {
		return nil, err
	}
	return h.GetMalfunction(ctx)
}
