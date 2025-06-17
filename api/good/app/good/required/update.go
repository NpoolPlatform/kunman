//nolint:dupl
package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"
)

func (s *Server) UpdateRequired(ctx context.Context, in *npool.UpdateRequiredRequest) (*npool.UpdateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(&in.ID, true),
		required1.WithEntID(&in.EntID, true),
		required1.WithAppID(&in.AppID, true),
		required1.WithMust(in.Must, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRequiredResponse{
		Info: info,
	}, nil
}
