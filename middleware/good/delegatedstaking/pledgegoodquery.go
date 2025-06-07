package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	entgoodreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodreward"
)

type delegatedstakingGoodQueryHandler struct {
	*Handler
	delegatedstaking *ent.DelegatedStaking
	goodBase         *ent.GoodBase
	goodReward       *ent.GoodReward
	goodCoins        []*ent.GoodCoin
	coinRewards      []*ent.GoodCoinReward
}

func (h *delegatedstakingGoodQueryHandler) getDelegatedStaking(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.DelegatedStaking.Query()
	if h.ID != nil {
		stm.Where(entdelegatedstaking.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdelegatedstaking.EntID(*h.EntID))
	}
	if h.GoodID != nil {
		stm.Where(entdelegatedstaking.GoodID(*h.GoodID))
	}
	if h.delegatedstaking, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h.delegatedstaking.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingGoodQueryHandler) getGoodReward(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodReward, err = cli.
		GoodReward.
		Query().
		Where(
			entgoodreward.GoodID(h.delegatedstaking.GoodID),
			entgoodreward.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *delegatedstakingGoodQueryHandler) getGoodCoins(ctx context.Context, cli *ent.Client) (err error) {
	h.goodCoins, err = cli.
		GoodCoin.
		Query().
		Where(
			entgoodcoin.GoodID(h.delegatedstaking.GoodID),
			entgoodcoin.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *delegatedstakingGoodQueryHandler) getGoodCoinRewards(ctx context.Context, cli *ent.Client) (err error) {
	h.coinRewards, err = cli.
		GoodCoinReward.
		Query().
		Where(
			entgoodcoinreward.GoodID(h.delegatedstaking.GoodID),
			entgoodcoinreward.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *delegatedstakingGoodQueryHandler) _getDelegatedStakingGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getDelegatedStaking(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h.delegatedstaking == nil {
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
		if err := h.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *delegatedstakingGoodQueryHandler) getDelegatedStakingGood(ctx context.Context) error {
	return h._getDelegatedStakingGood(ctx, false)
}

func (h *delegatedstakingGoodQueryHandler) requireDelegatedStakingGood(ctx context.Context) error {
	return h._getDelegatedStakingGood(ctx, true)
}
