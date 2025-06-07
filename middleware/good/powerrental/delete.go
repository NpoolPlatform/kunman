package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	rewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"
	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock/mining"
	powerrentalcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/powerrental"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*powerRentalGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.goodBase == nil {
		return wlog.Errorf("invalid goodbase")
	}
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deletePowerRental(ctx context.Context, tx *ent.Tx) error {
	if h.powerRental == nil {
		return wlog.Errorf("invalid powerrental")
	}
	if _, err := powerrentalcrud.UpdateSet(
		tx.PowerRental.UpdateOneID(h.powerRental.ID),
		&powerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteStock(ctx context.Context, tx *ent.Tx) error {
	if h.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		&stockcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteReward(ctx context.Context, tx *ent.Tx) error {
	if h.goodReward == nil {
		return nil
	}
	if _, err := rewardcrud.UpdateSet(
		tx.GoodReward.UpdateOneID(h.goodReward.ID),
		&rewardcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, poolStock := range h.miningGoodStocks {
		if _, err := mininggoodstockcrud.UpdateSet(
			tx.MiningGoodStock.UpdateOneID(poolStock.ID),
			&mininggoodstockcrud.Req{
				DeletedAt: &h.now,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) DeletePowerRental(ctx context.Context) error {
	handler := &deleteHandler{
		powerRentalGoodQueryHandler: &powerRentalGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getPowerRentalGood(ctx); err != nil {
		return err
	}
	if handler.powerRental == nil {
		return nil
	}
	h.ID = &handler.powerRental.ID

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteMiningGoodStocks(_ctx, tx); err != nil {
			return err
		}
		return handler.deletePowerRental(_ctx, tx)
	})
}
