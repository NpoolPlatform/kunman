package message

import (
	"context"

	messagecrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/message"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
)

func (h *Handler) ExistMessageCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := messagecrud.SetQueryConds(cli.Message.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	return stm.Exist(ctx)
}

func (h *Handler) ExistMessageConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistMessageCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
