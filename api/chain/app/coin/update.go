//nolint:nolintlint,dupl
package appcoin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appcoin1 "github.com/NpoolPlatform/kunman/gateway/chain/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(&in.ID, true),
		appcoin1.WithAppID(&in.AppID, true),
		appcoin1.WithName(in.Name, false),
		appcoin1.WithDisplayNames(in.DisplayNames, false),
		appcoin1.WithLogo(in.Logo, false),
		appcoin1.WithForPay(in.ForPay, false),
		appcoin1.WithWithdrawAutoReviewAmount(in.WithdrawAutoReviewAmount, false),
		appcoin1.WithMarketValue(in.MarketValue, false),
		appcoin1.WithSettlePercent(in.SettlePercent, false),
		appcoin1.WithSettleTips(in.SettleTips, false),
		appcoin1.WithProductPage(in.ProductPage, false),
		appcoin1.WithDisabled(in.Disabled, false),
		appcoin1.WithDisplay(in.Display, false),
		appcoin1.WithDisplayIndex(in.DisplayIndex, false),
		appcoin1.WithMaxAmountPerWithdraw(in.MaxAmountPerWithdraw, false),
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
