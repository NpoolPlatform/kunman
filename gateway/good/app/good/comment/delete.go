package comment

import (
	"context"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
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
	if err := commentmwcli.DeleteComment(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
