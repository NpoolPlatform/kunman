package device

import (
	"context"
	"time"

	devicecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteDeviceType(ctx context.Context, cli *ent.Client) error {
	if _, err := devicecrud.UpdateSet(
		cli.DeviceInfo.UpdateOneID(*h.ID),
		&devicecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteDeviceType(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetDeviceType(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteDeviceType(_ctx, cli)
	})
}
