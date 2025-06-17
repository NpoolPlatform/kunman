//nolint:dupl
package displaycolor

import (
	"context"

	displaycolor1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"
)

func (s *Server) AdminUpdateDisplayColor(ctx context.Context, in *npool.AdminUpdateDisplayColorRequest) (*npool.AdminUpdateDisplayColorResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithID(&in.ID, true),
		displaycolor1.WithEntID(&in.EntID, true),
		displaycolor1.WithAppID(&in.TargetAppID, true),
		displaycolor1.WithColor(in.Color, false),
		displaycolor1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDisplayColorResponse{
		Info: info,
	}, nil
}
