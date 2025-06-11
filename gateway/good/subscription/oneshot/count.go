package oneshot

import (
	"context"

	oneshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	oneshotmw "github.com/NpoolPlatform/kunman/middleware/good/subscription/oneshot"
)

func (h *Handler) CountOneShots(ctx context.Context) (uint32, error) {
	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithConds(&oneshotmwpb.Conds{}),
	)
	if err != nil {
		return 0, err
	}

	return handler.CountOneShots(ctx)
}
