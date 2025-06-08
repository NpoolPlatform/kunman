package required

import (
	"context"

	requiredmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/required"
	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"

	"github.com/google/uuid"
)

func (h *Handler) CreateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := requiredmwcli.CreateRequired(ctx, &requiredmwpb.RequiredReq{
		EntID:          h.EntID,
		MainGoodID:     h.MainGoodID,
		RequiredGoodID: h.RequiredGoodID,
		Must:           h.Must,
	}); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
