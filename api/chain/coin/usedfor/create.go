//nolint:nolintlint,dupl
package coinusedfor

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	usedfor1 "github.com/NpoolPlatform/kunman/gateway/chain/coin/usedfor"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/usedfor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinUsedFor(ctx context.Context, in *npool.CreateCoinUsedForRequest) (*npool.CreateCoinUsedForResponse, error) {
	handler, err := usedfor1.NewHandler(
		ctx,
		usedfor1.WithCoinTypeID(&in.CoinTypeID, true),
		usedfor1.WithUsedFor(&in.UsedFor, true),
		usedfor1.WithPriority(in.Priority, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinUsedForResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinUsedFor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinUsedForResponse{
		Info: info,
	}, nil
}
