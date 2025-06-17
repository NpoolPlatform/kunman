//nolint:dupl
package description

import (
	"context"

	description1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

func (s *Server) UpdateDescription(ctx context.Context, in *npool.UpdateDescriptionRequest) (*npool.UpdateDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID, true),
		description1.WithEntID(&in.EntID, true),
		description1.WithAppID(&in.AppID, true),
		description1.WithDescription(in.Description, false),
		description1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDescriptionResponse{
		Info: info,
	}, nil
}
