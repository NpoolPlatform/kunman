package goodbase

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
)

func (h *Handler) ExistGoodBase(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppGoodBase.Query()
		if h.ID != nil {
			stm.Where(entappgoodbase.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappgoodbase.EntID(*h.EntID))
		}
		count, err := stm.Limit(1).Count(_ctx)
		if err != nil {
			return err
		}

		exist = count > 0

		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
