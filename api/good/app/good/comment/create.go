package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
)

func (s *Server) CreateComment(ctx context.Context, in *npool.CreateCommentRequest) (*npool.CreateCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithAppID(&in.AppID, true),
		comment1.WithCommentUserID(&in.UserID, true),
		comment1.WithAppGoodID(&in.AppGoodID, true),
		comment1.WithOrderID(in.OrderID, false),
		comment1.WithContent(&in.Content, true),
		comment1.WithReplyToID(in.ReplyToID, false),
		comment1.WithAnonymous(in.Anonymous, false),
		comment1.WithScore(in.Score, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCommentResponse{
		Info: info,
	}, nil
}
