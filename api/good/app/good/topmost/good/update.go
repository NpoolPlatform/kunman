//nolint:dupl
package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
)

func (s *Server) UpdateTopMostGood(ctx context.Context, in *npool.UpdateTopMostGoodRequest) (*npool.UpdateTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithEntID(&in.EntID, true),
		topmostgood1.WithAppID(&in.AppID, true),
		topmostgood1.WithUnitPrice(in.UnitPrice, false),
		topmostgood1.WithDisplayIndex(in.DisplayIndex, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostGoodResponse{
		Info: info,
	}, nil
}
