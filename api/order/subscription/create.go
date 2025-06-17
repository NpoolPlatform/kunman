//nolint:dupl
package subscription

import (
	"context"

	ordercommon "github.com/NpoolPlatform/kunman/api/order/order/common"
	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) CreateSubscriptionOrder(ctx context.Context, in *npool.CreateSubscriptionOrderRequest) (*npool.CreateSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.UserID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),
		subscription1.WithOrderType(func() *types.OrderType { e := types.OrderType_Normal; return &e }(), true),
		subscription1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByPurchase; return &e }(), true),
		subscription1.WithPaymentBalances(in.Balances, true),
		subscription1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		subscription1.WithPaymentFiatID(in.PaymentFiatID, false),
		subscription1.WithCouponIDs(in.CouponIDs, true),
		subscription1.WithLifeSeconds(in.LifeSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubscriptionOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateUserSubscriptionOrder(ctx context.Context, in *npool.CreateUserSubscriptionOrderRequest) (*npool.CreateUserSubscriptionOrderResponse, error) {
	if err := ordercommon.ValidateAdminCreateOrderType(in.GetOrderType()); err != nil {
		logger.Sugar().Errorw(
			"CreateUserSubscriptionOrder",
			"In", in,
		)
		return &npool.CreateUserSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),
		subscription1.WithOrderType(&in.OrderType, true),
		subscription1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserSubscriptionOrderResponse{
		Info: info,
	}, nil
}
