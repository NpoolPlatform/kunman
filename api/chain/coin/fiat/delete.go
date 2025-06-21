//nolint:nolintlint,dupl
package coinfiat

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coinfiat1 "github.com/NpoolPlatform/kunman/gateway/chain/coin/fiat"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoinFiat(ctx context.Context, in *npool.DeleteCoinFiatRequest) (*npool.DeleteCoinFiatResponse, error) {
	handler, err := coinfiat1.NewHandler(
		ctx,
		coinfiat1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoinFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinFiatResponse{
		Info: info,
	}, nil
}
