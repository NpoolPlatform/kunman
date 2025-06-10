package comment

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/comment"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteComment(ctx context.Context) (*npool.Comment, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUserWithUserID(ctx, *h.CommentUserID); err != nil {
		return nil, err
	}
	if err := handler.checkUserComment(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetComment(ctx)
	if err != nil {
		return nil, err
	}

	commentHandler, err := commentmw.NewHandler(
		ctx,
		commentmw.WithID(h.ID, true),
		commentmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := commentHandler.DeleteComment(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
