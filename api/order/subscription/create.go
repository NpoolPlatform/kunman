package subscription

import (
	"context"
	"os"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) CreateSubscriptionOrder(ctx context.Context, in *npool.CreateSubscriptionOrderRequest) (*npool.CreateSubscriptionOrderResponse, error) {
	domain := os.Getenv("PAYMENT_CALLBACK_DOMAIN")
	metadata, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if authorities := metadata[":authority"]; len(authorities) > 0 {
			domain = authorities[0]
		}
	}

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
		subscription1.WithFiatPaymentChannel(in.FiatPaymentChannel, false),
		subscription1.WithCouponIDs(in.CouponIDs, true),
		subscription1.WithLifeSeconds(in.LifeSeconds, false),
		subscription1.WithDomain(&domain, true),
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
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),
		subscription1.WithOrderType(types.OrderType_Airdrop.Enum(), true),
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
