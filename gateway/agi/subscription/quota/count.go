package quota

import (
	"context"

	quotamw "github.com/NpoolPlatform/kunman/middleware/agi/subscription/quota"
)

func (h *Handler) CountQuotas(ctx context.Context) (uint32, error) {
	handler, err := quotamw.NewHandler(
		ctx,
		quotamw.WithAppID(h.AppID, false),
		quotamw.WithUserID(h.UserID, false),
	)
	if err != nil {
		return 0, err
	}

	return handler.CountQuotas(ctx)
}
