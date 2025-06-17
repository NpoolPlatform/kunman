//nolint:dupl
package appgood

import (
	"context"

	appgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good"
)

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithAppID(&in.AppID, true),
		appgood1.WithOffset(in.Offset),
		appgood1.WithLimit(in.Limit),
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
