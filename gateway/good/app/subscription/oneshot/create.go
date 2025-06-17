package oneshot

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
	apponeshotmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription/oneshot"

	"github.com/google/uuid"
)

// TODO: check start mode with power rental start mode
func (h *Handler) CreateOneShot(ctx context.Context) (*npool.AppOneShot, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := apponeshotmw.NewHandler(
		ctx,
		apponeshotmw.WithEntID(h.EntID, true),
		apponeshotmw.WithAppID(h.AppID, true),
		apponeshotmw.WithGoodID(h.GoodID, true),
		apponeshotmw.WithAppGoodID(h.AppGoodID, true),
		apponeshotmw.WithName(h.Name, true),
		apponeshotmw.WithBanner(h.Banner, true),
		// apponeshotmw.WithEnableSetCommission(h.EnableSetCommission, true),
		apponeshotmw.WithUSDPrice(h.USDPrice, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateOneShot(ctx); err != nil {
		return nil, err
	}
	return h.GetOneShot(ctx)
}
