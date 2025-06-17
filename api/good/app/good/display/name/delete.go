package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) DeleteDisplayName(ctx context.Context, in *npool.DeleteDisplayNameRequest) (*npool.DeleteDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithID(&in.ID, true),
		displayname1.WithEntID(&in.EntID, true),
		displayname1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDisplayNameResponse{
		Info: info,
	}, nil
}
