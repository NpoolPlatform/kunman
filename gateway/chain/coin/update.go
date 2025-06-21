package coin

import (
	"context"

	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
)

func (h *Handler) UpdateCoin(ctx context.Context) (*coinmwpb.Coin, error) {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithID(h.ID, true),
		coinmw.WithLogo(h.Logo, false),
		coinmw.WithPresale(h.Presale, false),
		coinmw.WithReservedAmount(h.ReservedAmount, false),
		coinmw.WithForPay(h.ForPay, false),
		coinmw.WithHomePage(h.HomePage, false),
		coinmw.WithSpecs(h.Specs, false),
		coinmw.WithFeeCoinTypeID(h.FeeCoinTypeID, false),
		coinmw.WithWithdrawFeeByStableUSD(h.WithdrawFeeByStableUSD, false),
		coinmw.WithWithdrawFeeAmount(h.WithdrawFeeAmount, false),
		coinmw.WithCollectFeeAmount(h.CollectFeeAmount, false),
		coinmw.WithHotWalletFeeAmount(h.HotWalletFeeAmount, false),
		coinmw.WithHotWalletAccountAmount(h.HotWalletAccountAmount, false),
		coinmw.WithLowFeeAmount(h.LowFeeAmount, false),
		coinmw.WithHotLowFeeAmount(h.HotLowFeeAmount, false),
		coinmw.WithPaymentAccountCollectAmount(h.PaymentAccountCollectAmount, false),
		coinmw.WithDisabled(h.Disabled, false),
		coinmw.WithStableUSD(h.StableUSD, false),
		coinmw.WithLeastTransferAmount(h.LeastTransferAmount, false),
		coinmw.WithNeedMemo(h.NeedMemo, false),
		coinmw.WithRefreshCurrency(h.RefreshCurrency, false),
		coinmw.WithCheckNewAddressBalance(h.CheckNewAddressBalance, false),
	)
	if err != nil {
		return nil, err
	}

	return handler.UpdateCoin(ctx)
}
