package label

import (
	"context"

	label1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
)

func (s *Server) DeleteLabel(ctx context.Context, in *npool.DeleteLabelRequest) (*npool.DeleteLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithID(&in.ID, true),
		label1.WithEntID(&in.EntID, true),
		label1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLabelResponse{
		Info: info,
	}, nil
}
