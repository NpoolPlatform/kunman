package user

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	usercrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/task/user"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteTaskUser(ctx context.Context, cli *ent.Client) error {
	if _, err := usercrud.UpdateSet(
		cli.TaskUser.UpdateOneID(*h.ID),
		&usercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteTaskUser(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetTaskUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteTaskUser(_ctx, cli)
	})
}
