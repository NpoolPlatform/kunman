package comment

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	commentmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/comment"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
	commentmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/comment"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
	purchasedUser bool
	trialUser     bool
}

func (h *createHandler) rewardWriteComment() {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       *h.AppID,
			UserID:      *h.CommentUserID,
			EventType:   basetypes.UsedFor_WriteComment,
			Consecutive: 1,
		}
		return publisher.Update(
			basetypes.MsgID_CalculateEventRewardReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"rewardWriteComment",
			"AppID", *h.AppID,
			"UserID", h.UserID,
			"Error", err,
		)
	}
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

	if err := commentmwcli.CreateComment(ctx, &commentmwpb.CommentReq{
		EntID:         h.EntID,
		UserID:        h.CommentUserID,
		AppGoodID:     h.AppGoodID,
		OrderID:       h.OrderID,
		Content:       h.Content,
		ReplyToID:     h.ReplyToID,
		Anonymous:     h.Anonymous,
		PurchasedUser: &handler.purchasedUser,
		TrialUser:     &handler.trialUser,
		Score:         h.Score,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.rewardWriteComment()

	return h.GetComment(ctx)
}
