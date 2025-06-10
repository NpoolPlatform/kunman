package displayname

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
	appgooddisplaynamemw "github.com/NpoolPlatform/kunman/middleware/good/app/good/display/name"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDisplayName(ctx); err != nil {
		return nil, err
	}

	displayNameHandler, err := appgooddisplaynamemw.NewHandler(
		ctx,
		appgooddisplaynamemw.WithEntID(h.EntID, true),
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

	if err := displayNameHandler.UpdateDisplayName(ctx); err != nil {
		return nil, err
	}
	return h.GetDisplayName(ctx)
}
