package comment

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/comment"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateComment(ctx context.Context) (*npool.Comment, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if h.UserID != nil {
		if err := handler.CheckUserWithUserID(ctx, *h.UserID); err != nil {
			return nil, err
		}
	}
	if err := handler.CheckUserWithUserID(ctx, *h.CommentUserID); err != nil {
		return nil, err
	}
	if err := handler.checkUserComment(ctx); err != nil {
		return nil, err
	}

	commentHandler, err := commentmw.NewHandler(
		ctx,
		commentmw.WithID(h.ID, true),
		commentmw.WithEntID(h.EntID, true),
		commentmw.WithAnonymous(h.Anonymous, true),
		commentmw.WithContent(h.Content, true),
		commentmw.WithHide(h.Hide, true),
		commentmw.WithHideReason(h.HideReason, true),
	)
	if err != nil {
		return nil, err
	}

	if err := commentHandler.UpdateComment(ctx); err != nil {
		return nil, err
	}
	return h.GetComment(ctx)
}
