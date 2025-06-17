package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) GetDisplayColors(ctx context.Context, in *npool.GetDisplayColorsRequest) (*npool.GetDisplayColorsResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithAppID(&in.AppID, true),
		displaycolor1.WithAppGoodID(in.AppGoodID, false),
		displaycolor1.WithOffset(in.Offset),
		displaycolor1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayColors(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayColorsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
