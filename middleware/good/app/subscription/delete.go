package subscription

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	appsubscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*appSubscriptionGoodQueryHandler
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

func (h *deleteHandler) deleteAppSubscription(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsubscriptioncrud.UpdateSet(
		tx.AppSubscription.UpdateOneID(h.appSubscription.ID),
		&appsubscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSubscription(ctx context.Context) error {
	handler := &deleteHandler{
		appSubscriptionGoodQueryHandler: &appSubscriptionGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppSubscriptionGood(ctx); err != nil {
		return err
	}
	if handler.appSubscription == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteAppSubscription(_ctx, tx)
	})
}
