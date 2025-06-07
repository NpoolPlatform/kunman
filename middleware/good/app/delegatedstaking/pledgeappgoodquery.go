package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdelegatedstaking"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
)

type delegatedstakingAppGoodQueryHandler struct {
	*Handler
	_ent delegatedstaking
}

func (h *delegatedstakingAppGoodQueryHandler) getDelegatedStaking(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.delegatedstaking, err = cli.
		DelegatedStaking.
		Query().
		Where(
			entdelegatedstaking.GoodID(*h.AppGoodBaseReq.GoodID),
			entdelegatedstaking.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingAppGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h._ent.delegatedstaking.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingAppGoodQueryHandler) _getDelegatedStakingGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getDelegatedStaking(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.delegatedstaking == nil {
			return nil
		}
		if err := h.getGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *delegatedstakingAppGoodQueryHandler) getAppDelegatedStaking(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.AppDelegatedStaking.Query().Where(entappdelegatedstaking.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappdelegatedstaking.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappdelegatedstaking.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entappdelegatedstaking.AppGoodID(*h.AppGoodID))
	}
	h._ent.appDelegatedStaking, err = stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingAppGoodQueryHandler) getAppGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.appGoodBase, err = cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntID(h._ent.appDelegatedStaking.AppGoodID),
			entappgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingAppGoodQueryHandler) _getAppDelegatedStakingAppGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppDelegatedStaking(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.appDelegatedStaking == nil {
			return nil
		}
		if err := h.getAppGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}

		return nil
	}); err != nil {
		return wlog.WrapError(err)
	}

	if h._ent.appGoodBase == nil {
		if !must {
			return nil
		}
		return wlog.Errorf("invalid appgoodbase")
	}

	h.AppGoodBaseReq.GoodID = &h._ent.appGoodBase.GoodID
	return h._getDelegatedStakingGood(ctx, must)
}

//nolint:unused
func (h *delegatedstakingAppGoodQueryHandler) getDelegatedStakingGood(ctx context.Context) error {
	return h._getDelegatedStakingGood(ctx, false)
}

func (h *delegatedstakingAppGoodQueryHandler) requireDelegatedStakingGood(ctx context.Context) error {
	return h._getDelegatedStakingGood(ctx, true)
}

func (h *delegatedstakingAppGoodQueryHandler) getAppDelegatedStakingAppGood(ctx context.Context) error {
	return h._getAppDelegatedStakingAppGood(ctx, false)
}

func (h *delegatedstakingAppGoodQueryHandler) requireAppDelegatedStakingAppGood(ctx context.Context) error {
	return h._getAppDelegatedStakingAppGood(ctx, true)
}

func (h *Handler) QueryDelegatedStakingEnt(ctx context.Context) (DelegatedStaking, error) {
	handler := &delegatedstakingAppGoodQueryHandler{
		Handler: h,
	}
	if err := handler.requireAppDelegatedStakingAppGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return &handler._ent, nil
}
