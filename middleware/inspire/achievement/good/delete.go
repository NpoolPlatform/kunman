package goodachievement

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodachievementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/good"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodAchievement(ctx context.Context, cli *ent.Client) error {
	_, err := goodachievementcrud.UpdateSet(
		cli.GoodAchievement.UpdateOneID(*h.ID),
		&goodachievementcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteAchievement(ctx context.Context) error {
	info, err := h.GetAchievement(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteGoodAchievement(_ctx, cli)
	})
}
