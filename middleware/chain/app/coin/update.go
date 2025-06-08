package appcoin

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entappexrate "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/exchangerate"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	appcoincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin"
	appexratecrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/exrate"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateAppCoin(ctx context.Context, tx *ent.Tx) error {
	info, err := appcoincrud.UpdateSet(
		tx.AppCoin.UpdateOneID(*h.ID),
		&appcoincrud.Req{
			Name:                     h.Name,
			DisplayNames:             h.DisplayNames,
			Logo:                     h.Logo,
			ForPay:                   h.ForPay,
			ProductPage:              h.ProductPage,
			WithdrawAutoReviewAmount: h.WithdrawAutoReviewAmount,
			DailyRewardAmount:        h.DailyRewardAmount,
			Disabled:                 h.Disabled,
			Display:                  h.Display,
			DisplayIndex:             h.DisplayIndex,
			MaxAmountPerWithdraw:     h.MaxAmountPerWithdraw,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.AppID = &info.AppID
	h.CoinTypeID = &info.CoinTypeID

	return nil
}

func (h *updateHandler) updateExrate(ctx context.Context, tx *ent.Tx) error {
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

	if info != nil {
		if _, err := appexratecrud.UpdateSet(
			info.Update(),
			&appexratecrud.Req{
				MarketValue:   h.MarketValue,
				SettlePercent: h.SettlePercent,
				SettleTips:    h.SettleTips,
				Setter:        h.Setter,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := appexratecrud.CreateSet(
		tx.ExchangeRate.Create(),
		&appexratecrud.Req{
			AppID:         h.AppID,
			CoinTypeID:    h.CoinTypeID,
			MarketValue:   h.MarketValue,
			SettlePercent: h.SettlePercent,
			SettleTips:    h.SettleTips,
			Setter:        h.Setter,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppCoin(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateExrate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
