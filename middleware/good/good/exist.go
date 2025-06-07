package good

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistGoodConds(ctx context.Context) (exist bool, err error) {
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	}); err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
