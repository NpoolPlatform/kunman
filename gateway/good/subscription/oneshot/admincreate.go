package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	oneshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	oneshotmw "github.com/NpoolPlatform/kunman/middleware/good/subscription/oneshot"

	"github.com/google/uuid"
)

func (h *Handler) AdminCreateSubscrption(ctx context.Context) (*oneshotmwpb.OneShot, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithEntID(h.EntID, true),
		oneshotmw.WithGoodID(h.GoodID, true),
		oneshotmw.WithName(h.Name, true),
		oneshotmw.WithQuota(h.Quota, true),
		oneshotmw.WithUSDPrice(h.USDPrice, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.CreateOneShot(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetOneShot(ctx)
}
