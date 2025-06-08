package quota

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription/quota"
	quotamw "github.com/NpoolPlatform/kunman/middleware/agi/subscription/quota"
)

func (h *Handler) GetQuotas(ctx context.Context) ([]*npool.Quota, error) {
	handler, err := quotamw.NewHandler(
		ctx,
		quotamw.WithAppID(h.AppID, false),
		quotamw.WithUserID(h.UserID, false),
		quotamw.WithOffset(h.Offset),
		quotamw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetQuotas(ctx)
}
