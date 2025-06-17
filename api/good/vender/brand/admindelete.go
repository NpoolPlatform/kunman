//nolint:dupl
package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/kunman/gateway/good/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/brand"
)

func (s *Server) AdminDeleteBrand(ctx context.Context, in *npool.AdminDeleteBrandRequest) (*npool.AdminDeleteBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(&in.ID, true),
		brand1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteBrandResponse{
		Info: info,
	}, nil
}
