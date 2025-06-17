package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
)

func (s *Server) GetSimulates(ctx context.Context, in *npool.GetSimulatesRequest) (*npool.GetSimulatesResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithOffset(in.Offset),
		simulate1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSimulates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSimulatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
