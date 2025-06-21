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

func (s *Server) CreateFiat(ctx context.Context, in *npool.CreateFiatRequest) (*npool.CreateFiatResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithName(&in.Name, true),
		fiat1.WithUnit(&in.Unit, true),
		fiat1.WithLogo(&in.Logo, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatResponse{
		Info: info,
	}, nil
}
