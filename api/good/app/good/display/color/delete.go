package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) DeleteDisplayColor(ctx context.Context, in *npool.DeleteDisplayColorRequest) (*npool.DeleteDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithID(&in.ID, true),
		displaycolor1.WithEntID(&in.EntID, true),
		displaycolor1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDisplayColorResponse{
		Info: info,
	}, nil
}
