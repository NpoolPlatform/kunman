package appcoin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
)

func (h *Handler) UpdateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	// TODO: check appid / cointypeid / id

	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithID(h.ID, true),
		appcoinmw.WithName(h.Name, true),
		appcoinmw.WithDisplayNames(h.DisplayNames, true),
		appcoinmw.WithLogo(h.Logo, true),
		appcoinmw.WithForPay(h.ForPay, true),
		appcoinmw.WithWithdrawAutoReviewAmount(h.WithdrawAutoReviewAmount, true),
		appcoinmw.WithMarketValue(h.MarketValue, true),
		appcoinmw.WithSettlePercent(h.SettlePercent, true),
		appcoinmw.WithSettleTips(h.SettleTips, true),
		appcoinmw.WithDailyRewardAmount(h.DailyRewardAmount, true),
		appcoinmw.WithProductPage(h.ProductPage, true),
		appcoinmw.WithDisabled(h.Disabled, true),
		appcoinmw.WithDisplay(h.Display, true),
		appcoinmw.WithDisplayIndex(h.DisplayIndex, true),
		appcoinmw.WithMaxAmountPerWithdraw(h.MaxAmountPerWithdraw, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.UpdateCoin(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid coin")
	}

	h.EntID = &info.EntID

	return h.GetCoin(ctx)
}
