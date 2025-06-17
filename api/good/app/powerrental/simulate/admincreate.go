//nolint:dupl
package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/kunman/gateway/good/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"
)

func (s *Server) AdminCreateSimulate(ctx context.Context, in *npool.AdminCreateSimulateRequest) (*npool.AdminCreateSimulateResponse, error) {
	handler, err := simulate1.NewHandler(
		ctx,
		simulate1.WithAppID(&in.TargetAppID, true),
		simulate1.WithAppGoodID(&in.AppGoodID, true),
		simulate1.WithOrderUnits(&in.OrderUnits, true),
		simulate1.WithOrderDurationSeconds(&in.OrderDurationSeconds, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateSimulateResponse{
		Info: info,
	}, nil
}
