package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) AdminCreateDisplayName(ctx context.Context, in *npool.AdminCreateDisplayNameRequest) (*npool.AdminCreateDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithAppID(&in.TargetAppID, true),
		displayname1.WithAppGoodID(&in.AppGoodID, true),
		displayname1.WithName(&in.Name, true),
		displayname1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDisplayNameResponse{
		Info: info,
	}, nil
}
