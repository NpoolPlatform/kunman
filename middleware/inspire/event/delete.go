package event

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	eventcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event"
	eventcoincrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coin"
	eventcouponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enteventcoin "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoin"
	enteventcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoupon"
	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

//nolint:dupl
func (h *deleteHandler) deleteCoupons(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.
		EventCoupon.
		Query().
		Where(
			enteventcoupon.EventID(*h.EntID),
			enteventcoupon.AppID(*h.AppID),
			enteventcoupon.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if _, err := eventcouponcrud.UpdateSet(
			info.Update(),
			&eventcouponcrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deleteCoins(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.
		EventCoin.
		Query().
		Where(
			enteventcoin.EventID(*h.EntID),
			enteventcoin.AppID(*h.AppID),
			enteventcoin.DeletedAt(0),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if _, err := eventcoincrud.UpdateSet(
			info.Update(),
			&eventcoincrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *Handler) DeleteEvent(ctx context.Context) error {
	info, err := h.GetEvent(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	id := uuid.MustParse(info.EntID)
	h.EntID = &id
	appID := uuid.MustParse(info.AppID)
	h.AppID = &appID

	handler := &deleteHandler{
		Handler: h,
	}

	now := uint32(time.Now().Unix())
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCoupons(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.deleteCoins(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if _, err := eventcrud.UpdateSet(
			tx.Event.UpdateOneID(*h.ID),
			&eventcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
