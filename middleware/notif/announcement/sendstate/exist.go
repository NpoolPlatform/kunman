package sendstate

import (
	"context"

	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/announcement/sendstate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db"
	ent "github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated"
)

func (h *Handler) ExistSendStateCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := crud.SetQueryConds(cli.SendAnnouncement.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistSendStateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistSendStateCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
