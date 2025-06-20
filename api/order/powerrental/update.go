//nolint:dupl
package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) UpdatePowerRentalOrder(ctx context.Context, in *npool.UpdatePowerRentalOrderRequest) (*npool.UpdatePowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(&in.UserID, true),
		powerrental1.WithOrderID(&in.OrderID, true),
		powerrental1.WithPaymentBalances(in.Balances, true),
		powerrental1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		powerrental1.WithUserSetPaid(in.Paid, true),
		powerrental1.WithUserSetCanceled(in.Canceled, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdatePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdatePowerRentalOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateUserPowerRentalOrder(ctx context.Context, in *npool.UpdateUserPowerRentalOrderRequest) (*npool.UpdateUserPowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(&in.TargetUserID, true),
		powerrental1.WithOrderID(&in.OrderID, true),
		powerrental1.WithAdminSetCanceled(in.Canceled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserPowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdatePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserPowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateUserPowerRentalOrderResponse{
		Info: info,
	}, nil
}
