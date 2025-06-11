package manufacturer

import (
	"context"

	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
)

func (h *Handler) GetManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	handler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetManufacturer(ctx)
}

func (h *Handler) GetManufacturers(ctx context.Context) ([]*manufacturermwpb.Manufacturer, uint32, error) {
	handler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithConds(&manufacturermwpb.Conds{}),
		manufacturermw.WithOffset(h.Offset),
		manufacturermw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetManufacturers(ctx)
}
