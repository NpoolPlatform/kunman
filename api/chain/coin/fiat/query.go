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

func (s *Server) GetCoinFiats(ctx context.Context, in *npool.GetCoinFiatsRequest) (*npool.GetCoinFiatsResponse, error) {
	handler, err := coinfiat1.NewHandler(
		ctx,
		coinfiat1.WithCoinTypeIDs(in.GetCoinTypeIDs(), false),
		coinfiat1.WithOffset(in.GetOffset()),
		coinfiat1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinFiatsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoinFiats(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinFiatsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinFiatsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
