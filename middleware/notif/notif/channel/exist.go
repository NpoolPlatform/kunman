package channel

import (
	"context"

	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/channel"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) ExistChannelConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(cli.NotifChannel.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(ctx)
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
