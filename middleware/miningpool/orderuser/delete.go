//nolint:dupl
package orderuser

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	crud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/orderuser"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	info *orderuser.OrderUser
}

func (h *deleteHandler) deleteOrderUserBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.OrderUser.UpdateOneID(h.info.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return wlog.WrapError(err)
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) DeleteOrderUser(ctx context.Context) error {
	handler := deleteHandler{Handler: h}
	var err error

	handler.info, err = handler.GetOrderUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if handler.info == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOrderUserBase(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
