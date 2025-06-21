//nolint:nolintlint,dupl
package coinfiat

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	fiat1 "github.com/NpoolPlatform/kunman/gateway/chain/coin/fiat"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinFiat(ctx context.Context, in *npool.CreateCoinFiatRequest) (*npool.CreateCoinFiatResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithCoinTypeID(&in.CoinTypeID, true),
		fiat1.WithFiatID(&in.FiatID, true),
		fiat1.WithFeedType(&in.FeedType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinFiatResponse{
		Info: info,
	}, nil
}
