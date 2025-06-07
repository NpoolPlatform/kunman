package goodbase

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entextrainfo "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/extrainfo"
)

type queryHandler struct {
	*Handler
	goodBase  *ent.AppGoodBase
	extraInfo *ent.ExtraInfo
}

func (h *queryHandler) getGoodBase(ctx context.Context, cli *ent.Client) (err error) {
	stm := cli.AppGoodBase.Query()
	if h.ID != nil {
		stm.Where(entappgoodbase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodbase.EntID(*h.EntID))
	}
	h.goodBase, err = stm.Only(ctx)
	return err
}

func (h *queryHandler) getExtraInfo(ctx context.Context, cli *ent.Client) (err error) {
	h.extraInfo, err = cli.
		ExtraInfo.
		Query().
		Where(entextrainfo.AppGoodID(*h.EntID)).
		Only(ctx)
	return err
}

func (h *Handler) GetGoodBase(ctx context.Context) (_goodBase AppGoodBase, err error) {
	handler := &queryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.getGoodBase(ctx, cli); err != nil {
			return err
		}
		return handler.getExtraInfo(ctx, cli)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return &goodBase{
		_ent:      handler.goodBase,
		_extraEnt: handler.extraInfo,
	}, nil
}
