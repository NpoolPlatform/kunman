package like

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
	likemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/like"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateLike(ctx context.Context) (*npool.Like, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUser(ctx); err != nil {
		return nil, err
	}
	if err := handler.checkUserLike(ctx); err != nil {
		return nil, err
	}

	likeHandler, err := likemw.NewHandler(
		ctx,
		likemw.WithID(h.ID, true),
		likemw.WithEntID(h.EntID, true),
		likemw.WithLike(h.Like, true),
	)
	if err != nil {
		return nil, err
	}

	if err := likeHandler.UpdateLike(ctx); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
