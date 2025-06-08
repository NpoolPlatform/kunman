package topmost

import (
	"context"

	topmostmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"

	"github.com/google/uuid"
)

func (h *Handler) CreateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := topmostmwcli.CreateTopMost(ctx, &topmostmwpb.TopMostReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		TopMostType: h.TopMostType,
		Title:       h.Title,
		Message:     h.Message,
		TargetUrl:   h.TargetURL,
		StartAt:     h.StartAt,
		EndAt:       h.EndAt,
	}); err != nil {
		return nil, err
	}
	return h.GetTopMost(ctx)
}
