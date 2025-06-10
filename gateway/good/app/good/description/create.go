package description

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
	appgooddescriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/description"

	"github.com/google/uuid"
)

func (h *Handler) CreateDescription(ctx context.Context) (*npool.Description, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appgooddescriptionmw.NewHandler(
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

	if err := handler.CreateDescription(ctx); err != nil {
		return nil, err
	}
	return h.GetDescription(ctx)
}
