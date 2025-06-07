package delegatedstaking

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/delegatedstaking"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	rewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodcoin "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoin"
)

type deleteHandler struct {
	*delegatedstakingGoodQueryHandler
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

func (h *deleteHandler) deleteDelegatedStaking(ctx context.Context, tx *ent.Tx) error {
	if h.delegatedstaking == nil {
		return wlog.Errorf("invalid delegatedstaking")
	}
	if _, err := delegatedstakingcrud.UpdateSet(
		tx.DelegatedStaking.UpdateOneID(h.delegatedstaking.ID),
		&delegatedstakingcrud.Req{
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

func (h *deleteHandler) deleteGoodCoin(ctx context.Context, tx *ent.Tx) error {
	fmt.Println("--h.goodCoins: ", h.goodCoins)
	if h.goodCoins == nil {
		return nil
	}
	ids := []uint32{}
	for _, coin := range h.goodCoins {
		ids = append(ids, coin.ID)
	}
	if len(ids) == 0 {
		return nil
	}
	if _, err := tx.GoodCoin.
		Update().
		Where(
			entgoodcoin.IDIn(ids...),
			entgoodcoin.DeletedAt(0),
		).
		SetDeletedAt(h.now).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteDelegatedStaking(ctx context.Context) error {
	handler := &deleteHandler{
		delegatedstakingGoodQueryHandler: &delegatedstakingGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getDelegatedStakingGood(ctx); err != nil {
		return err
	}
	if handler.delegatedstaking == nil {
		return nil
	}
	h.ID = &handler.delegatedstaking.ID

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteGoodCoin(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteDelegatedStaking(_ctx, tx)
	})
}
