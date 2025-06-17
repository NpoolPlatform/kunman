package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"
)

func (s *Server) DeleteRequired(ctx context.Context, in *npool.DeleteRequiredRequest) (*npool.DeleteRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(&in.ID, true),
		required1.WithEntID(&in.EntID, true),
		required1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRequiredResponse{
		Info: info,
	}, nil
}
