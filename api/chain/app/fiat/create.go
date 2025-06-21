//nolint:nolintlint,dupl
package appfiat

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appfiat1 "github.com/NpoolPlatform/kunman/gateway/chain/app/fiat"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFiat(ctx context.Context, in *npool.CreateFiatRequest) (*npool.CreateFiatResponse, error) {
	handler, err := appfiat1.NewHandler(
		ctx,
		appfiat1.WithAppID(&in.TargetAppID, true),
		appfiat1.WithFiatID(&in.FiatID, true),
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
