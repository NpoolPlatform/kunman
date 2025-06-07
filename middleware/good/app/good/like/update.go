package like

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	likecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/like"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entlike "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/like"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	like int
}

func (h *updateHandler) updateLike(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Like.
		Query().
		Where(
			entlike.ID(*h.ID),
			entlike.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if *h.Like && !info.Like {
		h.like = 1
	} else if !*h.Like && info.Like {
		h.like = -1
	}

	if _, err := likecrud.UpdateSet(
		info.Update(),
		&likecrud.Req{
			Like: h.Like,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodLike(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	info.Likes += uint32(h.like)
	info.Dislikes -= uint32(h.like)

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Likes:    &info.Likes,
			Dislikes: &info.Dislikes,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateLike(ctx context.Context) error {
	if h.Like == nil {
		return wlog.Errorf("%v", cruder.ErrUpdateNothing)
	}

	info, err := h.GetLike(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return wlog.Errorf("invalid like")
	}

	h.AppGoodID = func() *uuid.UUID { uid, _ := uuid.Parse(info.AppGoodID); return &uid }()
	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateLike(ctx, tx); err != nil {
			return err
		}
		return handler.updateGoodLike(ctx, tx)
	})
}
