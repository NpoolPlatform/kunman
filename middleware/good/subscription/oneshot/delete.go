package oneshot

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	oneshotcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*oneShotGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.goodBase == nil {
		return wlog.Errorf("invalid goodbase")
	}
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteOneShot(ctx context.Context, tx *ent.Tx) error {
	if h.oneShot == nil {
		return wlog.Errorf("invalid oneShot")
	}
	if _, err := oneshotcrud.UpdateSet(
		tx.SubscriptionOneShot.UpdateOneID(h.oneShot.ID),
		&oneshotcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOneShot(ctx context.Context) error {
	handler := &deleteHandler{
		oneShotGoodQueryHandler: &oneShotGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getOneShotGood(ctx); err != nil {
		return err
	}
	if handler.oneShot == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteOneShot(_ctx, tx)
	})
}
