package malfunction

import (
	"context"

	malfunctionmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/malfunction"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/good/malfunction"
	malfunctionmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"

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
