//nolint:nolintlint,dupl
package feed

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	feed1 "github.com/NpoolPlatform/kunman/gateway/chain/fiat/currency/feed"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/fiat/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFeed(ctx context.Context, in *npool.UpdateFeedRequest) (*npool.UpdateFeedResponse, error) {
	handler, err := feed1.NewHandler(
		ctx,
		feed1.WithID(&in.ID, true),
		feed1.WithFeedFiatName(in.FeedFiatName, false),
		feed1.WithDisabled(in.Disabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFeed(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFeedResponse{
		Info: info,
	}, nil
}
