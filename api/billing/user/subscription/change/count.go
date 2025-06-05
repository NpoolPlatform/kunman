package change

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/user/subscription/change"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription/change"
)

func (s *Server) CountSubscriptionChanges(ctx context.Context, in *npool.CountSubscriptionChangesRequest) (*npool.CountSubscriptionChangesResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.CountSubscriptionChanges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountSubscriptionChanges",
			"In", in,
			"Error", err,
		)
		return &npool.CountSubscriptionChangesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountSubscriptionChangesResponse{
		Total: total,
	}, nil
}
