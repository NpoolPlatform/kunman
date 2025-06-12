package reward

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	rewardcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/user/reward"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteUserReward(ctx context.Context, cli *ent.Client) error {
	if _, err := rewardcrud.UpdateSet(
		cli.UserReward.UpdateOneID(*h.ID),
		&rewardcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteUserReward(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetUserReward(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteUserReward(_ctx, cli)
	})
}
