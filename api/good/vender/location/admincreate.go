package location

import (
	"context"

	location1 "github.com/NpoolPlatform/kunman/gateway/good/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/location"
)

func (s *Server) AdminCreateLocation(ctx context.Context, in *npool.AdminCreateLocationRequest) (*npool.AdminCreateLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithCountry(&in.Country, true),
		location1.WithProvince(&in.Province, true),
		location1.WithCity(&in.City, true),
		location1.WithAddress(&in.Address, true),
		location1.WithBrandID(&in.BrandID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateLocationResponse{
		Info: info,
	}, nil
}
