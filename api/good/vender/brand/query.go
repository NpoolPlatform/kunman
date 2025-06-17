package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/kunman/gateway/good/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/brand"
)

func (s *Server) GetBrands(ctx context.Context, in *npool.GetBrandsRequest) (*npool.GetBrandsResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithOffset(in.Offset),
		brand1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrands",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetBrands(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrands",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetBrandsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
