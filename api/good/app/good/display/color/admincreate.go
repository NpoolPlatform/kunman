package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) AdminCreateDisplayColor(ctx context.Context, in *npool.AdminCreateDisplayColorRequest) (*npool.AdminCreateDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithAppID(&in.TargetAppID, true),
		displaycolor1.WithAppGoodID(&in.AppGoodID, true),
		displaycolor1.WithColor(&in.Color, true),
		displaycolor1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDisplayColorResponse{
		Info: info,
	}, nil
}
