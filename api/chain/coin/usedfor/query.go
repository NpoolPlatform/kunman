//nolint:nolintlint,dupl
package coinusedfor

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coinusedfor1 "github.com/NpoolPlatform/kunman/gateway/chain/coin/usedfor"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/usedfor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoinUsedFors(ctx context.Context, in *npool.GetCoinUsedForsRequest) (*npool.GetCoinUsedForsResponse, error) {
	handler, err := coinusedfor1.NewHandler(
		ctx,
		coinusedfor1.WithCoinTypeIDs(in.GetCoinTypeIDs(), false),
		coinusedfor1.WithUsedFors(in.GetUsedFors(), false),
		coinusedfor1.WithOffset(in.GetOffset()),
		coinusedfor1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFors",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoinUsedFors(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFors",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinUsedForsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
