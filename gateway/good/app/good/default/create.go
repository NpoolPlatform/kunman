package default1

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"
	defaultmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/default"

	"github.com/google/uuid"
)

func (h *Handler) CreateDefault(ctx context.Context) (*npool.Default, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := defaultmw.NewHandler(
		ctx,
		defaultmw.WithEntID(h.EntID, true),
		defaultmw.WithCoinTypeID(h.CoinTypeID, true),
		defaultmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateDefault(ctx); err != nil {
		return nil, err
	}
	return h.GetDefault(ctx)
}
