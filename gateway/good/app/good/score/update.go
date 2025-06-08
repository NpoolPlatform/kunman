package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	scoremwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/score"
	scoremwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateScore(ctx context.Context) (*npool.Score, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.CheckUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.checkScore(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if h.Score != nil {
		if err := handler.validateScore(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	if err := scoremwcli.UpdateScore(ctx, &scoremwpb.ScoreReq{
		ID:    h.ID,
		Score: h.Score,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetScore(ctx)
}
