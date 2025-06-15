package oneshot

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	apponeshotcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*appOneShotGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.appGoodBase == nil {
		return wlog.Errorf("invalid appgoodbase")
	}
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h.appGoodBase.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppOneShot(ctx context.Context, tx *ent.Tx) error {
	if _, err := apponeshotcrud.UpdateSet(
		tx.AppSubscriptionOneShot.UpdateOneID(h.appOneShot.ID),
		&apponeshotcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOneShot(ctx context.Context) error {
	handler := &deleteHandler{
		appOneShotGoodQueryHandler: &appOneShotGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppOneShotGood(ctx); err != nil {
		return err
	}
	if handler.appOneShot == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteAppOneShot(_ctx, tx)
	})
}
