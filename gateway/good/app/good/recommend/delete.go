package recommend

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
	recommendmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkUserRecommend(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}

	recommendHandler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithID(h.ID, true),
		recommendmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := recommendHandler.DeleteRecommend(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
