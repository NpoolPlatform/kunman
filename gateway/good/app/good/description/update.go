package description

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
	appgooddescriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/description"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDescription(ctx context.Context) (*npool.Description, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDescription(ctx); err != nil {
		return nil, err
	}

	descriptionHandler, err := appgooddescriptionmw.NewHandler(
		ctx,
		appgooddescriptionmw.WithEntID(h.EntID, true),
		appgooddescriptionmw.WithAppGoodID(h.AppGoodID, true),
		appgooddescriptionmw.WithDescription(h.Description, true),
		appgooddescriptionmw.WithIndex(func() *uint8 {
			if h.Index == nil {
				return nil
			}
			u := uint8(*h.Index)
			return &u
		}(), true),
	)
	if err != nil {
		return nil, err
	}

	if err := descriptionHandler.UpdateDescription(ctx); err != nil {
		return nil, err
	}
	return h.GetDescription(ctx)
}
