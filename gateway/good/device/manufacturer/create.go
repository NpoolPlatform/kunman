package manufacturer

import (
	"context"

	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"

	"github.com/google/uuid"
)

func (h *Handler) CreateManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithEntID(h.EntID, true),
		manufacturermw.WithName(h.Name, true),
		manufacturermw.WithLogo(h.Logo, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateManufacturer(ctx); err != nil {
		return nil, err
	}
	return h.GetManufacturer(ctx)
}
