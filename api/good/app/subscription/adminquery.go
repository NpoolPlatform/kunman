package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
)

func (s *Server) AdminGetAppSubscriptions(ctx context.Context, in *npool.AdminGetAppSubscriptionsRequest) (*npool.AdminGetAppSubscriptionsResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSubscriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppSubscriptions",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppSubscriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppSubscriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
