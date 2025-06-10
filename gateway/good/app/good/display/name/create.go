package displayname

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
	appgooddisplaynamemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/name"

	"github.com/google/uuid"
)

func (h *Handler) CreateDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appgooddisplaynamemw.NewHandler(
		ctx,
		appgooddisplaynamemw.WithEntID(h.EntID, true),
		appgooddisplaynamemw.WithAppGoodID(h.AppGoodID, true),
		appgooddisplaynamemw.WithName(h.Name, true),
		appgooddisplaynamemw.WithIndex(func() *uint8 {
			if h.Index == nil {
				return nil
			}
			index := uint8(*h.Index)
			return &index
		}(), true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateDisplayName(ctx); err != nil {
		return nil, err
	}
	return h.GetDisplayName(ctx)
}
