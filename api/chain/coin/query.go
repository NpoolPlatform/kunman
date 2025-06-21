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

func (s *Server) GetCoins(ctx context.Context, in *npool.GetCoinsRequest) (*npool.GetCoinsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithOffset(in.GetOffset()),
		coin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
