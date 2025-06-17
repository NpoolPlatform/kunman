package good

import (
	"context"

	good1 "github.com/NpoolPlatform/kunman/gateway/good/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good"
)

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithOffset(in.Offset),
		good1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
