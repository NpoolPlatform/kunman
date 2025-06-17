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

func (s *Server) AdminUpdateRequired(ctx context.Context, in *npool.AdminUpdateRequiredRequest) (*npool.AdminUpdateRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(&in.ID, true),
		required1.WithEntID(&in.EntID, true),
		required1.WithAppID(&in.TargetAppID, true),
		required1.WithMust(in.Must, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateRequiredResponse{
		Info: info,
	}, nil
}
