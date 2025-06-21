//nolint:nolintlint,dupl
package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coin1 "github.com/NpoolPlatform/kunman/gateway/chain/coin"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(&in.ID, true),
		coin1.WithLogo(in.Logo, false),
		coin1.WithPresale(in.Presale, false),
		coin1.WithReservedAmount(in.ReservedAmount, false),
		coin1.WithForPay(in.ForPay, false),
		coin1.WithHomePage(in.HomePage, false),
		coin1.WithSpecs(in.Specs, false),
		coin1.WithFeeCoinTypeID(in.FeeCoinTypeID, false),
		coin1.WithWithdrawFeeByStableUSD(in.WithdrawFeeByStableUSD, false),
		coin1.WithWithdrawFeeAmount(in.WithdrawFeeAmount, false),
		coin1.WithCollectFeeAmount(in.CollectFeeAmount, false),
		coin1.WithHotWalletFeeAmount(in.HotWalletFeeAmount, false),
		coin1.WithHotWalletAccountAmount(in.HotWalletAccountAmount, false),
		coin1.WithLowFeeAmount(in.LowFeeAmount, false),
		coin1.WithHotLowFeeAmount(in.HotLowFeeAmount, false),
		coin1.WithPaymentAccountCollectAmount(in.PaymentAccountCollectAmount, false),
		coin1.WithDisabled(in.Disabled, false),
		coin1.WithStableUSD(in.StableUSD, false),
		coin1.WithLeastTransferAmount(in.LeastTransferAmount, false),
		coin1.WithNeedMemo(in.NeedMemo, false),
		coin1.WithRefreshCurrency(in.RefreshCurrency, false),
		coin1.WithCheckNewAddressBalance(in.CheckNewAddressBalance, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info,
	}, nil
}
