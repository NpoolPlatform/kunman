package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"
	entmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/mininggoodstock"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"
	entstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/stock"
)

type powerRentalGoodQueryHandler struct {
	*Handler
	powerRental      *ent.PowerRental
	goodBase         *ent.GoodBase
	goodReward       *ent.GoodReward
	coinRewards      []*ent.GoodCoinReward
	stock            *ent.Stock
	miningGoodStocks []*ent.MiningGoodStock
}

func (h *powerRentalGoodQueryHandler) getPowerRental(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.PowerRental.Query()
	if h.ID != nil {
		stm.Where(entpowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entpowerrental.EntID(*h.EntID))
	}
	if h.GoodID != nil {
		stm.Where(entpowerrental.GoodID(*h.GoodID))
	}
	if h.powerRental, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h.powerRental.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalGoodQueryHandler) getGoodReward(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodReward, err = cli.
		GoodReward.
		Query().
		Where(
			entgoodreward.GoodID(h.powerRental.GoodID),
			entgoodreward.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalGoodQueryHandler) getGoodCoinRewards(ctx context.Context, cli *ent.Client) (err error) {
	h.coinRewards, err = cli.
		GoodCoinReward.
		Query().
		Where(
			entgoodcoinreward.GoodID(h.powerRental.GoodID),
			entgoodcoinreward.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalGoodQueryHandler) getGoodStock(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.stock, err = cli.
		Stock.
		Query().
		Where(
			entstock.GoodID(h.powerRental.GoodID),
			entstock.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalGoodQueryHandler) getMiningGoodStock(ctx context.Context, cli *ent.Client) (err error) {
	if h.miningGoodStocks, err = cli.
		MiningGoodStock.
		Query().
		Where(
			entmininggoodstock.GoodStockID(h.stock.EntID),
			entmininggoodstock.DeletedAt(0),
		).All(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalGoodQueryHandler) _getPowerRentalGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getPowerRental(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h.powerRental == nil {
			return nil
		}
		if err := h.getGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodReward(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodCoinRewards(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodStock(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h.stock == nil {
			return nil
		}
		return h.getMiningGoodStock(_ctx, cli)
	})
}

func (h *powerRentalGoodQueryHandler) getPowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, false)
}

func (h *powerRentalGoodQueryHandler) requirePowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, true)
}
