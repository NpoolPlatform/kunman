package like

import (
	"context"

	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/like"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/like"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
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

	if err := likemwcli.UpdateLike(ctx, &likemwpb.LikeReq{
		ID:    h.ID,
		EntID: h.EntID,
		Like:  h.Like,
	}); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
