package required

import (
	"context"

	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"

	"github.com/google/uuid"
)

func (h *Handler) CreateRequired(ctx context.Context) (*requiredappgoodmwpb.Required, error) {
	if err := h.CheckAppGoodWithAppGoodID(ctx, *h.MainAppGoodID); err != nil {
		return nil, err
	}
	if err := h.CheckAppGoodWithAppGoodID(ctx, *h.RequiredAppGoodID); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithEntID(h.EntID, true),
		requiredappgoodmw.WithMainAppGoodID(h.MainAppGoodID, true),
		requiredappgoodmw.WithRequiredAppGoodID(h.RequiredAppGoodID, true),
		requiredappgoodmw.WithMust(h.Must, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateRequired(ctx); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
