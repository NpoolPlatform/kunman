//nolint:dupl
package goodcoin

import (
	"context"

	goodcoin1 "github.com/NpoolPlatform/kunman/gateway/good/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
)

func (s *Server) AdminCreateGoodCoin(ctx context.Context, in *npool.AdminCreateGoodCoinRequest) (*npool.AdminCreateGoodCoinResponse, error) {
	handler, err := goodcoin1.NewHandler(
		ctx,
		goodcoin1.WithGoodID(&in.GoodID, true),
		goodcoin1.WithCoinTypeID(&in.CoinTypeID, true),
		goodcoin1.WithMain(in.Main, false),
		goodcoin1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGoodCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateGoodCoinResponse{
		Info: info,
	}, nil
}
