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

func (s *Server) CreateFeed(ctx context.Context, in *npool.CreateFeedRequest) (*npool.CreateFeedResponse, error) {
	handler, err := feed1.NewHandler(
		ctx,
		feed1.WithFiatID(&in.FiatID, true),
		feed1.WithFeedType(&in.FeedType, true),
		feed1.WithFeedFiatName(&in.FeedFiatName, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFeed(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeedResponse{
		Info: info,
	}, nil
}
