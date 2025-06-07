package delegatedstaking

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appdelegatedstakingcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/delegatedstaking"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*delegatedstakingAppGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h._ent.appGoodBase.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *deleteHandler) deleteDelegatedStaking(ctx context.Context, tx *ent.Tx) error {
	if _, err := appdelegatedstakingcrud.UpdateSet(
		tx.AppDelegatedStaking.UpdateOneID(h._ent.appDelegatedStaking.ID),
		&appdelegatedstakingcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteDelegatedStaking(ctx context.Context) error {
	handler := &deleteHandler{
		delegatedstakingAppGoodQueryHandler: &delegatedstakingAppGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppDelegatedStakingAppGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if handler._ent.appDelegatedStaking == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.deleteDelegatedStaking(_ctx, tx)
	})
}
