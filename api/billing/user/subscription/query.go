package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription"
)

func (s *Server) GetSubscription(ctx context.Context, in *npool.GetSubscriptionRequest) (*npool.GetSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(&in.EntID, true),
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

func (s *Server) GetSubscriptionsCount(ctx context.Context, in *npool.GetSubscriptionsCountRequest) (*npool.GetSubscriptionsCountResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.GetSubscriptionsCount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionsCountResponse{
		Total: total,
	}, nil
}
