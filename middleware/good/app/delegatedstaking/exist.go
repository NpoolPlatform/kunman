package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistDelegatedStaking(ctx context.Context) (exist bool, err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return false, wlog.Errorf("invalid appgoodid")
	}
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppGoodBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistDelegatedStakingConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryAppGoodBases(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return wlog.WrapError(err)
	}); err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
