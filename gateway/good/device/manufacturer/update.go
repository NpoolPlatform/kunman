package manufacturer

import (
	"context"

	manufacturermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/manufacturer"
	manufacturermw "github.com/NpoolPlatform/kunman/middleware/good/device/manufacturer"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateManufacturer(ctx context.Context) (*manufacturermwpb.Manufacturer, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkManufacturer(ctx); err != nil {
		return nil, err
	}

	manufacturerHandler, err := manufacturermw.NewHandler(
		ctx,
		manufacturermw.WithID(h.ID, true),
		manufacturermw.WithEntID(h.EntID, true),
		manufacturermw.WithName(h.Name, false),
		manufacturermw.WithLogo(h.Logo, false),
	)
	if err != nil {
		return nil, err
	}

	if err := manufacturerHandler.UpdateManufacturer(ctx); err != nil {
		return nil, err
	}
	return h.GetManufacturer(ctx)
}
