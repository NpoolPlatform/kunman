package location

import (
	"context"

	location1 "github.com/NpoolPlatform/kunman/gateway/good/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/location"
)

func (s *Server) GetLocations(ctx context.Context, in *npool.GetLocationsRequest) (*npool.GetLocationsResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithBrandID(in.BrandID, false),
		location1.WithOffset(in.Offset),
		location1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocations",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLocations(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocations",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLocationsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
