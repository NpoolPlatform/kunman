package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) AdminDeleteDisplayName(ctx context.Context, in *npool.AdminDeleteDisplayNameRequest) (*npool.AdminDeleteDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithID(&in.ID, true),
		displayname1.WithEntID(&in.EntID, true),
		displayname1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteDisplayNameResponse{
		Info: info,
	}, nil
}
