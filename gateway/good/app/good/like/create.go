package like

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
	likemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/like"

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

	likeHandler, err := likemw.NewHandler(
		ctx,
		likemw.WithEntID(h.EntID, true),
		likemw.WithUserID(h.UserID, true),
		likemw.WithAppGoodID(h.AppGoodID, true),
		likemw.WithLike(h.Like, true),
	)
	if err != nil {
		return nil, err
	}

	if err := likeHandler.CreateLike(ctx); err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
