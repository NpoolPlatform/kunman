package score

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	scoremwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
	scoremwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/score"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
}

func (h *createHandler) rewardGoodScoring() {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       *h.AppID,
			UserID:      *h.UserID,
			EventType:   basetypes.UsedFor_GoodScoring,
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
			"rewardGoodScoring",
			"AppID", *h.AppID,
			"UserID", h.UserID,
			"Error", err,
		)
	}
}

func (h *Handler) CreateScore(ctx context.Context) (*npool.Score, error) {
	handler := &createHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}

	if err := h.CheckUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if h.Score != nil {
		if err := handler.validateScore(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	if err := scoremwcli.CreateScore(ctx, &scoremwpb.ScoreReq{
		EntID:     h.EntID,
		UserID:    h.UserID,
		AppGoodID: h.AppGoodID,
		Score:     h.Score,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.rewardGoodScoring()

	return h.GetScore(ctx)
}
