//nolint:dupl
package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/kunman/gateway/good/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/brand"
)

func (s *Server) AdminCreateBrand(ctx context.Context, in *npool.AdminCreateBrandRequest) (*npool.AdminCreateBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithName(&in.Name, true),
		brand1.WithLogo(&in.Logo, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateBrandResponse{
		Info: info,
	}, nil
}
