package required

import (
	"context"

	requiredappgoodmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"

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

	if err := requiredappgoodmwcli.CreateRequired(ctx, &requiredappgoodmwpb.RequiredReq{
		EntID:             h.EntID,
		MainAppGoodID:     h.MainAppGoodID,
		RequiredAppGoodID: h.RequiredAppGoodID,
		Must:              h.Must,
	}); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
