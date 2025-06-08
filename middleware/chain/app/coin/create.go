package appcoin

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	appcoincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin"
	appexratecrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin/exrate"
	coincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppCoin(ctx context.Context, tx *ent.Tx) error {
	if _, err := appcoincrud.CreateSet(
		tx.AppCoin.Create(),
		&appcoincrud.Req{
			ID:                       h.ID,
			EntID:                    h.EntID,
			AppID:                    h.AppID,
			CoinTypeID:               h.CoinTypeID,
			Name:                     h.Name,
			DisplayNames:             h.DisplayNames,
			Logo:                     h.Logo,
			ForPay:                   h.ForPay,
			ProductPage:              h.ProductPage,
			WithdrawAutoReviewAmount: h.WithdrawAutoReviewAmount,
			DailyRewardAmount:        h.DailyRewardAmount,
			Display:                  h.Display,
			DisplayIndex:             h.DisplayIndex,
			MaxAmountPerWithdraw:     h.MaxAmountPerWithdraw,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createExrate(ctx context.Context, tx *ent.Tx) error {
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

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	// TODO: deduplicate

	h.Conds = &appcoincrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
	}
	exist, err := h.ExistCoinConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("appcoin exist")
	}

	coinHandler, err := coinmw.NewHandler(
		ctx,
	)
	if err != nil {
		return nil, err
	}
	coinHandler.Conds = &coincrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
	}
	coin, err := coinHandler.GetCoinOnly(ctx)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, fmt.Errorf("coin not exist")
	}
	if h.Name == nil {
		h.Name = &coin.Name
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppCoin(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createExrate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
