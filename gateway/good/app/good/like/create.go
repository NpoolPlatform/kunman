package like

import (
	"context"

	likemwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/like"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/like"
	likemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
}

func (h *Handler) CreateLike(ctx context.Context) (*npool.Like, error) {
	handler := &createHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUser(ctx); err != nil {
		return nil, err
	}
	if err := handler.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := likemwcli.CreateLike(ctx, &likemwpb.LikeReq{
		EntID:     h.EntID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Like:      h.Like,
	}); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
