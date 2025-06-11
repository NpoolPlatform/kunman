package malfunction

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmw "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"

	"github.com/google/uuid"
)

func (h *Handler) CreateMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := malfunctionmw.NewHandler(
		ctx,
		malfunctionmw.WithEntID(h.EntID, true),
		malfunctionmw.WithGoodID(h.GoodID, true),
		malfunctionmw.WithTitle(h.Title, true),
		malfunctionmw.WithMessage(h.Message, true),
		malfunctionmw.WithStartAt(h.StartAt, true),
		malfunctionmw.WithDurationSeconds(h.DurationSeconds, true),
		malfunctionmw.WithCompensateSeconds(h.CompensateSeconds, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateMalfunction(ctx); err != nil {
		return nil, err
	}
	return h.GetMalfunction(ctx)
}
