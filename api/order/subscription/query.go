package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) GetSubscriptionOrder(ctx context.Context, in *npool.GetSubscriptionOrderRequest) (*npool.GetSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.UserID, true),
		subscription1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriptionOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriptionOrders(ctx context.Context, in *npool.GetSubscriptionOrdersRequest) (*npool.GetSubscriptionOrdersResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(in.TargetUserID, false),
		subscription1.WithAppGoodID(in.AppGoodID, false),
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSubscriptionOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriptionOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMySubscriptionOrders(ctx context.Context, in *npool.GetMySubscriptionOrdersRequest) (*npool.GetMySubscriptionOrdersResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithUserID(&in.UserID, true),
		subscription1.WithAppGoodID(in.AppGoodID, false),
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMySubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMySubscriptionOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSubscriptionOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMySubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMySubscriptionOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMySubscriptionOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
