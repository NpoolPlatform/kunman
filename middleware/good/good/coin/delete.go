package coin

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoincrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	goodID     uuid.UUID
	coinTypeID uuid.UUID
	now        uint32
}

func (h *deleteHandler) deleteGoodCoin(ctx context.Context, tx *ent.Tx) error {
	_, err := goodcoincrud.UpdateSet(
		tx.GoodCoin.UpdateOneID(*h.ID),
		&goodcoincrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteGoodCoinReward(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.GoodCoinReward.
		Update().
		Where(
			entgoodcoinreward.GoodID(h.goodID),
			entgoodcoinreward.CoinTypeID(h.coinTypeID),
			entgoodcoinreward.DeletedAt(0),
		).
		SetDeletedAt(h.now).
		Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteGoodCoin(ctx context.Context) error {
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID
	handler := &deleteHandler{
		Handler:    h,
		goodID:     uuid.MustParse(info.GoodID),
		coinTypeID: uuid.MustParse(info.CoinTypeID),
		now:        uint32(time.Now().Unix()),
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodCoinReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.deleteGoodCoin(_ctx, tx)
	})
}
