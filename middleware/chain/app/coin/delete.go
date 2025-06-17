package appcoin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entdescription "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coindescription"
	entappexrate "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/exchangerate"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	appcoincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin"
	descriptioncrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/description"
	appexratecrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/exrate"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppCoin(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	info, err := appcoincrud.UpdateSet(
		tx.AppCoin.UpdateOneID(*h.ID),
		&appcoincrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.AppID = &info.AppID
	h.CoinTypeID = &info.CoinTypeID

	return nil
}

func (h *deleteHandler) deleteExrate(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		ExchangeRate.
		Query().
		Where(
			entappexrate.AppID(*h.AppID),
			entappexrate.CoinTypeID(*h.CoinTypeID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		return nil
	}

	now := uint32(time.Now().Unix())
	if _, err := appexratecrud.UpdateSet(
		info.Update(),
		&appexratecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *deleteHandler) deleteCoinDescription(ctx context.Context, tx *ent.Tx) error {
	infos, err := tx.
		CoinDescription.
		Query().
		Where(
			entdescription.AppID(*h.AppID),
			entdescription.CoinTypeID(*h.CoinTypeID),
		).
		ForUpdate().
		All(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if len(infos) == 0 {
		return nil
	}

	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if _, err := descriptioncrud.UpdateSet(
			info.Update(),
			&descriptioncrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) DeleteCoin(ctx context.Context) (*npool.Coin, error) {
	info, err := h.GetCoin(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	h.ID = &info.ID

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppCoin(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteCoinDescription(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteExrate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
