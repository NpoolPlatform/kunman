package goodcoin

import (
	"context"

	goodcoin1 "github.com/NpoolPlatform/kunman/gateway/good/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"
)

func (s *Server) AdminDeleteGoodCoin(ctx context.Context, in *npool.AdminDeleteGoodCoinRequest) (*npool.AdminDeleteGoodCoinResponse, error) {
	handler, err := goodcoin1.NewHandler(
		ctx,
		goodcoin1.WithID(&in.ID, true),
		goodcoin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteGoodCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteGoodCoinResponse{
		Info: info,
	}, nil
}
