package malfunction

import (
	"context"

	malfunctionmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/malfunction"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
	malfunctionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"

	"github.com/google/uuid"
)

func (h *Handler) CreateMalfunction(ctx context.Context) (*npool.Malfunction, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := malfunctionmwcli.CreateMalfunction(ctx, &malfunctionmwpb.MalfunctionReq{
		EntID:             h.EntID,
		GoodID:            h.GoodID,
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
