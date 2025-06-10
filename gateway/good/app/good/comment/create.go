package comment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/comment"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
	purchasedUser bool
	trialUser     bool
}

func (h *createHandler) rewardWriteComment() {
	// TODO: publish reward event
}

func (h *Handler) CreateComment(ctx context.Context) (*npool.Comment, error) {
	handler := &createHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUserWithUserID(ctx, *h.CommentUserID); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.CheckAppGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.checkOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	// TODO: check if trial user

	commentHandler, err := commentmw.NewHandler(
		ctx,
		commentmw.WithEntID(h.EntID, true),
		commentmw.WithUserID(h.CommentUserID, true),
		commentmw.WithAppGoodID(h.AppGoodID, true),
		commentmw.WithOrderID(h.OrderID, true),
		commentmw.WithContent(h.Content, true),
		commentmw.WithReplyToID(h.ReplyToID, true),
		commentmw.WithAnonymous(h.Anonymous, true),
		commentmw.WithPurchasedUser(&handler.purchasedUser, true),
		commentmw.WithTrialUser(&handler.trialUser, true),
		commentmw.WithScore(h.Score, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := commentHandler.CreateComment(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.rewardWriteComment()

	return h.GetComment(ctx)
}
