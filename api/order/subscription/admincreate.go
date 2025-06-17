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

func (s *Server) AdminCreateSubscriptionOrder(ctx context.Context, in *npool.AdminCreateSubscriptionOrderRequest) (*npool.AdminCreateSubscriptionOrderResponse, error) {
	if err := ordercommon.ValidateAdminCreateOrderType(in.GetOrderType()); err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSubscriptionOrder",
			"In", in,
		)
		return &npool.AdminCreateSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),
		subscription1.WithOrderType(types.OrderType_Airdrop.Enum(), true),
		subscription1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminCreateSubscriptionOrderResponse{
		Info: info,
	}, nil
}
