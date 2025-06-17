package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
)

func (s *Server) UpdateSimulate(ctx context.Context, in *npool.UpdateSimulateRequest) (*npool.UpdateSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithID(&in.ID, true),
		simulate1.WithEntID(&in.EntID, true),
		simulate1.WithAppID(&in.AppID, true),
		simulate1.WithAppGoodID(&in.AppGoodID, true),
		simulate1.WithOrderUnits(in.OrderUnits, false),
		simulate1.WithOrderDurationSeconds(in.OrderDurationSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSimulateResponse{
		Info: info,
	}, nil
}
