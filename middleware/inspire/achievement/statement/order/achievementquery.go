package orderstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entgoodachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodachievement"
	entgoodcoinachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodcoinachievement"
)

type achievementQueryHandler struct {
	*Handler
	entGoodAchievement     *ent.GoodAchievement
	entGoodCoinAchievement *ent.GoodCoinAchievement
}

func (h *achievementQueryHandler) getGoodAchievement(ctx context.Context, cli *ent.Client, must bool) (err error) {
	h.entGoodAchievement, err = cli.
		GoodAchievement.
		Query().
		Where(
			entgoodachievement.AppID(*h.AppID),
			entgoodachievement.UserID(*h.UserID),
			entgoodachievement.AppGoodID(*h.AppGoodID),
		).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) && !must {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *achievementQueryHandler) getGoodCoinAchievement(ctx context.Context, cli *ent.Client, must bool) (err error) {
	h.entGoodCoinAchievement, err = cli.
		GoodCoinAchievement.
		Query().
		Where(
			entgoodcoinachievement.AppID(*h.AppID),
			entgoodcoinachievement.UserID(*h.UserID),
			entgoodcoinachievement.GoodCoinTypeID(*h.GoodCoinTypeID),
		).
		Only(ctx)
	if err != nil && ent.IsNotFound(err) && !must {
		return nil
	}
	return err
}

func (h *achievementQueryHandler) _processAchievement(ctx context.Context, cli *ent.Client, must bool) error {
	if h.AppGoodID == nil || h.UserID == nil || h.GoodCoinTypeID == nil {
		return wlog.Errorf("invalid goodid")
	}
	if err := h.getGoodAchievement(ctx, cli, must); err != nil {
		return wlog.WrapError(err)
	}
	return h.getGoodCoinAchievement(ctx, cli, must)
}

func (h *achievementQueryHandler) _getAchievement(ctx context.Context, tx *ent.Tx, must bool) error {
	if tx != nil {
		return h._processAchievement(ctx, tx.Client(), must)
	}
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return h._processAchievement(ctx, cli, must)
	})
}

func (h *achievementQueryHandler) getAchievementWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getAchievement(ctx, tx, false)
}

func (h *achievementQueryHandler) requireAchievementWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getAchievement(ctx, tx, true)
}
