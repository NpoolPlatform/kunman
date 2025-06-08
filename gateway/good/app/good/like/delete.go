package like

import (
	"context"

	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/like"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteLike(ctx context.Context) (*npool.Like, error) {
	handler := &deleteHandler{
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

	info, err := h.GetLike(ctx)
	if err != nil {
		return nil, err
	}

	if err := likemwcli.DeleteLike(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}

	return info, nil
}
