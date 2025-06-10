package topmost

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMost(ctx); err != nil {
		return nil, err
	}

	topMostHandler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithID(h.ID, true),
		topmostmw.WithEntID(h.EntID, true),
		topmostmw.WithTitle(h.Title, true),
		topmostmw.WithMessage(h.Message, true),
		topmostmw.WithTargetURL(h.TargetURL, true),
		topmostmw.WithStartAt(h.StartAt, true),
		topmostmw.WithEndAt(h.EndAt, true),
	)
	if err != nil {
		return nil, err
	}

	if err := topMostHandler.UpdateTopMost(ctx); err != nil {
		return nil, err
	}
	return h.GetTopMost(ctx)
}
