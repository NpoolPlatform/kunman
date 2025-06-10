package score

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
	scoremw "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteScore(ctx context.Context) (*npool.Score, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUser(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkScore(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetScore(ctx)
	if err != nil {
		return nil, err
	}

	scoreHandler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithID(h.ID, true),
		scoremw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := scoreHandler.DeleteScore(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
