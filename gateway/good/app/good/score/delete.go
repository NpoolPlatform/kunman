package score

import (
	"context"

	scoremwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
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
	if err := scoremwcli.DeleteScore(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
