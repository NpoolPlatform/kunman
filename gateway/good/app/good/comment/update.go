package comment

import (
	"context"

	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"
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
	if err := commentmwcli.UpdateComment(ctx, &commentmwpb.CommentReq{
		ID:         h.ID,
		EntID:      h.EntID,
		Anonymous:  h.Anonymous,
		Content:    h.Content,
		Hide:       h.Hide,
		HideReason: h.HideReason,
	}); err != nil {
		return nil, err
	}
	return h.GetComment(ctx)
}
