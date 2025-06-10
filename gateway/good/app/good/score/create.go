package score

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"
	scoremw "github.com/NpoolPlatform/kunman/middleware/good/app/good/score"

	"github.com/google/uuid"
)

type createHandler struct {
	*checkHandler
}

func (h *createHandler) rewardGoodScoring() {
	// TODO: reward good scoring
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

	scoreHandler, err := scoremw.NewHandler(
		ctx,
		scoremw.WithEntID(h.EntID, true),
		scoremw.WithUserID(h.UserID, true),
		scoremw.WithAppGoodID(h.AppGoodID, true),
		scoremw.WithScore(h.Score, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := scoreHandler.CreateScore(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.rewardGoodScoring()

	return h.GetScore(ctx)
}
