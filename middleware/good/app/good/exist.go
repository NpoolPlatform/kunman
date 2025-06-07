package good

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistGoodConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryGoods(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	}); err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
