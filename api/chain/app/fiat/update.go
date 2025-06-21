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

func (s *Server) UpdateFiat(ctx context.Context, in *npool.UpdateFiatRequest) (*npool.UpdateFiatResponse, error) {
	handler, err := appfiat1.NewHandler(
		ctx,
		appfiat1.WithID(&in.ID, true),
		appfiat1.WithAppID(&in.AppID, true),
		appfiat1.WithName(in.Name, false),
		appfiat1.WithDisplayNames(in.DisplayNames, false),
		appfiat1.WithLogo(in.Logo, false),
		appfiat1.WithDisabled(in.Disabled, false),
		appfiat1.WithDisplay(in.Display, false),
		appfiat1.WithDisplayIndex(in.DisplayIndex, false),
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
