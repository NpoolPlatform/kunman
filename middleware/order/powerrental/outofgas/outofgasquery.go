package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entoutofgas "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/outofgas"
)

type outOfGasQueryHandler struct {
	*Handler
	_ent outOfGas
}

func (h *outOfGasQueryHandler) getOutOfGasEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.OutOfGas.Query()
	if h.ID != nil {
		stm.Where(entoutofgas.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entoutofgas.EntID(*h.EntID))
	}
	if h._ent.entOutOfGas, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *outOfGasQueryHandler) _getOutOfGas(ctx context.Context, must bool) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h.getOutOfGasEnt(ctx, cli, must)
	})
}

func (h *outOfGasQueryHandler) requireOutOfGas(ctx context.Context) error {
	return h._getOutOfGas(ctx, true)
}
