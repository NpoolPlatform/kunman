package description

import (
	"context"

	description1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

func (s *Server) AdminDeleteDescription(ctx context.Context, in *npool.AdminDeleteDescriptionRequest) (*npool.AdminDeleteDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID, true),
		description1.WithEntID(&in.EntID, true),
		description1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteDescriptionResponse{
		Info: info,
	}, nil
}
