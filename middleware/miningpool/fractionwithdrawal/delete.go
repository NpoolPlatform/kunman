package fractionwithdrawal

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"
	crud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawal"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	info *fractionwithdrawal.FractionWithdrawal
}

func (h *deleteHandler) deleteFractionWithdrawalBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	updateOne, err := crud.UpdateSet(tx.FractionWithdrawal.UpdateOneID(h.info.ID), &crud.Req{DeletedAt: &now})
	if err != nil {
		return wlog.WrapError(err)
	}
	_, err = updateOne.Save(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) DeleteFractionWithdrawal(ctx context.Context) error {
	handler := deleteHandler{Handler: h}
	var err error

	handler.info, err = handler.GetFractionWithdrawal(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if handler.info == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteFractionWithdrawalBase(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
