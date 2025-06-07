package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type countHandler struct {
	*baseQueryHandler
	total uint32
}

func (h *Handler) CountSubscriptions(ctx context.Context) (uint32, error) {
	handler := &countHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.queryGoodBases(cli); err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		count, err := handler.stmSelect.Count(_ctx)

		handler.total = uint32(count)

		return err
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}

	return handler.total, nil
}
