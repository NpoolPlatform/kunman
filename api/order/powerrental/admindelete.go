package powerrental

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminDeletePowerRentalOrder(ctx context.Context, in *npool.AdminDeletePowerRentalOrderRequest) (*npool.AdminDeletePowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.TargetAppID, true),
		powerrental1.WithUserID(&in.TargetUserID, true),
		powerrental1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeletePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeletePowerRentalOrderResponse{
		Info: info,
	}, nil
}
