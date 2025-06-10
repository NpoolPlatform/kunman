package topmost

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"

	"github.com/google/uuid"
)

func (h *Handler) CreateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithEntID(h.EntID, true),
		topmostmw.WithAppID(h.AppID, true),
		topmostmw.WithTopMostType(h.TopMostType, true),
		topmostmw.WithTitle(h.Title, true),
		topmostmw.WithMessage(h.Message, true),
		topmostmw.WithTargetURL(h.TargetURL, true),
		topmostmw.WithStartAt(h.StartAt, true),
		topmostmw.WithEndAt(h.EndAt, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateTopMost(ctx); err != nil {
		return nil, err
	}
	return h.GetTopMost(ctx)
}
