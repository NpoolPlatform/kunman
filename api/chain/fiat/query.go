//nolint:nolintlint,dupl
package fiat

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	fiat1 "github.com/NpoolPlatform/kunman/gateway/chain/fiat"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFiats(ctx context.Context, in *npool.GetFiatsRequest) (*npool.GetFiatsResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithOffset(in.GetOffset()),
		fiat1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFiats(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
