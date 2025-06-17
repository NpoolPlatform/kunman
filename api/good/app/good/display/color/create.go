package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) CreateDisplayColor(ctx context.Context, in *npool.CreateDisplayColorRequest) (*npool.CreateDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithAppID(&in.AppID, true),
		displaycolor1.WithAppGoodID(&in.AppGoodID, true),
		displaycolor1.WithColor(&in.Color, true),
		displaycolor1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDisplayColorResponse{
		Info: info,
	}, nil
}
