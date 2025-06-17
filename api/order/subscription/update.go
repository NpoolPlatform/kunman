//nolint:dupl
package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) UpdateSubscriptionOrder(ctx context.Context, in *npool.UpdateSubscriptionOrderRequest) (*npool.UpdateSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.UserID, true),
		subscription1.WithOrderID(&in.OrderID, true),
		subscription1.WithPaymentBalances(in.Balances, true),
		subscription1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		subscription1.WithUserSetCanceled(in.Canceled, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSubscriptionOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateUserSubscriptionOrder(ctx context.Context, in *npool.UpdateUserSubscriptionOrderRequest) (*npool.UpdateUserSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithOrderID(&in.OrderID, true),
		subscription1.WithAdminSetCanceled(in.Canceled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateUserSubscriptionOrderResponse{
		Info: info,
	}, nil
}
