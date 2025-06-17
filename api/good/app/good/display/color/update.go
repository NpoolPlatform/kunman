//nolint:dupl
package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) UpdateDisplayColor(ctx context.Context, in *npool.UpdateDisplayColorRequest) (*npool.UpdateDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithID(&in.ID, true),
		displaycolor1.WithEntID(&in.EntID, true),
		displaycolor1.WithAppID(&in.AppID, true),
		displaycolor1.WithColor(in.Color, false),
		displaycolor1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDisplayColorResponse{
		Info: info,
	}, nil
}
