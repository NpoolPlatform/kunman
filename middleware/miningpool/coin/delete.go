//nolint:dupl
package coin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	crud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/coin"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	info *coin.Coin
}

func (h *deleteHandler) deleteCoinBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.Coin.UpdateOneID(h.info.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return wlog.WrapError(err)
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) DeleteCoin(ctx context.Context) error {
	handler := deleteHandler{Handler: h}
	var err error

	handler.info, err = handler.GetCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if handler.info == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCoinBase(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
