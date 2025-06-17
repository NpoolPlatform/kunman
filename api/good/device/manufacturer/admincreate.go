//nolint:dupl
package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/kunman/gateway/good/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device/manufacturer"
)

func (s *Server) AdminCreateManufacturer(ctx context.Context, in *npool.AdminCreateManufacturerRequest) (*npool.AdminCreateManufacturerResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithName(&in.Name, true),
		manufacturer1.WithLogo(&in.Logo, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateManufacturer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateManufacturerResponse{
		Info: info,
	}, nil
}
