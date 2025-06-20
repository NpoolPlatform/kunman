package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/comment"
)

// nolint
func (s *Server) AdminUpdateComment(ctx context.Context, in *npool.AdminUpdateCommentRequest) (*npool.AdminUpdateCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
		comment1.WithEntID(&in.EntID, true),
		comment1.WithAppID(&in.TargetAppID, true),
		comment1.WithCommentUserID(&in.TargetUserID, true),
		comment1.WithHide(in.Hide, false),
		comment1.WithHideReason(in.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateCommentResponse{
		Info: info,
	}, nil
}
