package manufacturer

import (
	"context"

	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkManufacturer(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetManufacturer(ctx)
	if err != nil {
		return nil, err
	}

	manufacturerHandler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithID(h.ID, true),
		manufacturermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := manufacturerHandler.DeleteManufacturer(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
