//nolint:dupl
package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
)

func (s *Server) AdminCreateTopMostGood(ctx context.Context, in *npool.AdminCreateTopMostGoodRequest) (*npool.AdminCreateTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithAppID(&in.TargetAppID, true),
		topmostgood1.WithTopMostID(&in.TopMostID, true),
		topmostgood1.WithAppGoodID(&in.AppGoodID, true),
		topmostgood1.WithUnitPrice(in.UnitPrice, true),
		topmostgood1.WithDisplayIndex(in.DisplayIndex, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateTopMostGoodResponse{
		Info: info,
	}, nil
}
