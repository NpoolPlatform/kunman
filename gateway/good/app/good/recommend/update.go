package recommend

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
	recommendmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/recommend"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkUserRecommend(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if h.Hide != nil && *h.Hide && h.HideReason == nil {
		return nil, wlog.Errorf("invalid hidereason")
	}

	recommendHandler, err := recommendmw.NewHandler(
		ctx,
		recommendmw.WithID(h.ID, false),
		recommendmw.WithEntID(h.EntID, false),
		recommendmw.WithRecommendIndex(h.RecommendIndex, false),
		recommendmw.WithMessage(h.Message, false),
		recommendmw.WithHide(h.Hide, false),
		recommendmw.WithHideReason(h.HideReason, false),
	)
	if err != nil {
		return nil, err
	}

	if err := recommendHandler.UpdateRecommend(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetRecommend(ctx)
}
