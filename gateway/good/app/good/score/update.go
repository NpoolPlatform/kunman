package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
	scoremw "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"
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

	scoreHandler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithID(h.ID, true),
		scoremw.WithEntID(h.EntID, true),
		scoremw.WithScore(h.Score, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := scoreHandler.UpdateScore(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetScore(ctx)
}
