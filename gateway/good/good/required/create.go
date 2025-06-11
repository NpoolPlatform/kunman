package required

import (
	"context"

	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
	requiredmw "github.com/NpoolPlatform/kunman/middleware/good/good/required"

	"github.com/google/uuid"
)

func (h *Handler) CreateRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithEntID(h.EntID, true),
		requiredmw.WithMainGoodID(h.MainGoodID, true),
		requiredmw.WithRequiredGoodID(h.RequiredGoodID, true),
		requiredmw.WithMust(h.Must, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateRequired(ctx); err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
