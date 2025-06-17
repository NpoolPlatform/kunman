package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
)

func (s *Server) GetAppSubscription(ctx context.Context, in *npool.GetAppSubscriptionRequest) (*npool.GetAppSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppSubscriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppSubscriptions(ctx context.Context, in *npool.GetAppSubscriptionsRequest) (*npool.GetAppSubscriptionsResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSubscriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppSubscriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
