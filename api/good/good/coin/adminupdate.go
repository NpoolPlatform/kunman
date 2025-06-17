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

func (s *Server) AdminUpdateGoodCoin(ctx context.Context, in *npool.AdminUpdateGoodCoinRequest) (*npool.AdminUpdateGoodCoinResponse, error) {
	handler, err := goodcoin1.NewHandler(
		ctx,
		goodcoin1.WithID(&in.ID, true),
		goodcoin1.WithEntID(&in.EntID, true),
		goodcoin1.WithMain(in.Main, false),
		goodcoin1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGoodCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateGoodCoinResponse{
		Info: info,
	}, nil
}
