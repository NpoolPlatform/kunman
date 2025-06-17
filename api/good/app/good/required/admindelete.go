package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"
)

func (s *Server) AdminDeleteRequired(ctx context.Context, in *npool.AdminDeleteRequiredRequest) (*npool.AdminDeleteRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(&in.ID, true),
		required1.WithEntID(&in.EntID, true),
		required1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteRequiredResponse{
		Info: info,
	}, nil
}
