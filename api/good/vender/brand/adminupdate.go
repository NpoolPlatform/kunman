package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/kunman/gateway/good/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/brand"
)

func (s *Server) AdminUpdateBrand(ctx context.Context, in *npool.AdminUpdateBrandRequest) (*npool.AdminUpdateBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(&in.ID, true),
		brand1.WithEntID(&in.EntID, true),
		brand1.WithName(in.Name, false),
		brand1.WithLogo(in.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateBrandResponse{
		Info: info,
	}, nil
}
