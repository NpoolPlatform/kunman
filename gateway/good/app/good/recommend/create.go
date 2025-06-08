package recommend

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	recommendmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/recommend"
	eventmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) rewardWriteRecommend() {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventmwpb.CalcluateEventRewardsRequest{
			AppID:       *h.AppID,
			UserID:      *h.RecommenderID,
			EventType:   basetypes.UsedFor_WriteRecommend,
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
			"rewardWriteRecommend",
			"AppID", *h.AppID,
			"UserID", h.UserID,
			"Error", err,
		)
	}
}

func (h *Handler) CreateRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := h.CheckUserWithUserID(ctx, *h.RecommenderID); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if h.RecommendIndex != nil {
		maxRecommendIndex := decimal.RequireFromString("10.0")
		score, err := decimal.NewFromString(*h.RecommendIndex)
		if err != nil {
			return nil, err
		}
		if score.GreaterThan(maxRecommendIndex) || score.LessThan(decimal.NewFromInt(0)) {
			return nil, wlog.Errorf("invalid recommendindex")
		}
	}

	if err := recommendmwcli.CreateRecommend(ctx, &recommendmwpb.RecommendReq{
		EntID:          h.EntID,
		RecommenderID:  h.RecommenderID,
		AppGoodID:      h.AppGoodID,
		RecommendIndex: h.RecommendIndex,
		Message:        h.Message,
	}); err != nil {
		return nil, err
	}

	handler.rewardWriteRecommend()

	return h.GetRecommend(ctx)
}
