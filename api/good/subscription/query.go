package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription"
)

func (s *Server) GetSubscription(ctx context.Context, in *npool.GetSubscriptionRequest) (*npool.GetSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriptions(ctx context.Context, in *npool.GetSubscriptionsRequest) (*npool.GetSubscriptionsResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetSubscriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionsResponse{
		Infos: infos,
	}, nil
}
