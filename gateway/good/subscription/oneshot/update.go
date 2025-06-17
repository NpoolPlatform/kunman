package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	oneshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	oneshotmw "github.com/NpoolPlatform/kunman/middleware/good/subscription/oneshot"
)

func (h *Handler) UpdateOneShot(ctx context.Context) (*oneshotmwpb.OneShot, error) {
	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithID(h.ID, true),
		oneshotmw.WithEntID(h.EntID, true),
		oneshotmw.WithGoodID(h.GoodID, true),
		oneshotmw.WithName(h.Name, false),
		oneshotmw.WithQuota(h.Quota, false),
		oneshotmw.WithUSDPrice(h.USDPrice, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.UpdateOneShot(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return h.GetOneShot(ctx)
}
