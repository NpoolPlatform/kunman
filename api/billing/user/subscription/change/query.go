package change

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/user/subscription/change"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription/change"
)

func (s *Server) GetSubscriptionChange(ctx context.Context, in *npool.GetSubscriptionChangeRequest) (*npool.GetSubscriptionChangeResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSubscriptionChange(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChange",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionChangeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriptionChanges(ctx context.Context, in *npool.GetSubscriptionChangesRequest) (*npool.GetSubscriptionChangesResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetSubscriptionChanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriptionChangesResponse{
		Infos: infos,
	}, nil
}
