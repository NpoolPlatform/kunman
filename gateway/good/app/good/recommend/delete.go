package recommend

import (
	"context"

	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
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

	if err := recommendmwcli.DeleteRecommend(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
