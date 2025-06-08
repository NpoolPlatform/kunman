package recommend

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	recommendmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/recommend"
	recommendmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (h *Handler) UpdateRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &deleteHandler{
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

	if err := recommendmwcli.UpdateRecommend(ctx, &recommendmwpb.RecommendReq{
		ID:             h.ID,
		RecommendIndex: h.RecommendIndex,
		Message:        h.Message,
		Hide:           h.Hide,
		HideReason:     h.HideReason,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetRecommend(ctx)
}
