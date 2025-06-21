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

func (s *Server) DeleteFiat(ctx context.Context, in *npool.DeleteFiatRequest) (*npool.DeleteFiatResponse, error) {
	handler, err := appfiat1.NewHandler(
		ctx,
		appfiat1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFiatResponse{
		Info: info,
	}, nil
}
