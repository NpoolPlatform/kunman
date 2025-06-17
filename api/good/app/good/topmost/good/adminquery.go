package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
)

func (s *Server) AdminGetTopMostGoods(ctx context.Context, in *npool.AdminGetTopMostGoodsRequest) (*npool.AdminGetTopMostGoodsResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithAppID(&in.TargetAppID, true),
		topmostgood1.WithOffset(in.Offset),
		topmostgood1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMostGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetTopMostGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
