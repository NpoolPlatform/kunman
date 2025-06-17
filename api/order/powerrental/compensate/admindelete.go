package compensate

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental/compensate"
	compensate1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental/compensate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminDeleteCompensate(ctx context.Context, in *npool.AdminDeleteCompensateRequest) (*npool.AdminDeleteCompensateResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithAppID(&in.TargetAppID, true),
		compensate1.WithUserID(&in.TargetUserID, true),
		compensate1.WithID(&in.ID, true),
		compensate1.WithEntID(&in.EntID, true),
		compensate1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCompensate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeleteCompensateResponse{
		Info: info,
	}, nil
}
