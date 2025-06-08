package appsubscribe

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"

	appsubscribecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/subscriber/app/subscribe"
)

func (h *Handler) ExistAppSubscribeConds(ctx context.Context) (bool, error) {
	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsubscribecrud.SetQueryConds(cli.AppSubscribe.Query(), h.Conds)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
