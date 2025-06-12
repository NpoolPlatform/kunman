package config

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	configcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/task/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteTaskConfig(ctx context.Context, cli *ent.Client) error {
	if _, err := configcrud.UpdateSet(
		cli.TaskConfig.UpdateOneID(*h.ID),
		&configcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteTaskConfig(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetTaskConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteTaskConfig(_ctx, cli)
	})
}
