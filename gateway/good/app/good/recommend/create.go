package recommend

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
	recommendmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) rewardWriteRecommend() {
	// TODO: reward write recommend
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

	recommendHandler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithEntID(h.EntID, true),
		recommendmw.WithRecommenderID(h.RecommenderID, true),
		recommendmw.WithAppGoodID(h.AppGoodID, true),
		recommendmw.WithRecommendIndex(h.RecommendIndex, true),
		recommendmw.WithMessage(h.Message, true),
	)
	if err != nil {
		return nil, err
	}

	if err := recommendHandler.CreateRecommend(ctx); err != nil {
		return nil, err
	}

	handler.rewardWriteRecommend()

	return h.GetRecommend(ctx)
}
