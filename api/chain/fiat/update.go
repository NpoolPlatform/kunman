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

func (s *Server) UpdateFiat(ctx context.Context, in *npool.UpdateFiatRequest) (*npool.UpdateFiatResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithID(&in.ID, true),
		fiat1.WithName(in.Name, false),
		fiat1.WithLogo(in.Logo, false),
		fiat1.WithUnit(in.Unit, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFiatResponse{
		Info: info,
	}, nil
}
