package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	oneshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	oneshotmw "github.com/NpoolPlatform/kunman/middleware/good/subscription/oneshot"
)

func (h *Handler) AdminDeleteSubscrption(ctx context.Context) (*oneshotmwpb.OneShot, error) {
	info, err := h.GetOneShot(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithID(h.ID, true),
		oneshotmw.WithEntID(h.EntID, true),
		oneshotmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.DeleteOneShot(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
