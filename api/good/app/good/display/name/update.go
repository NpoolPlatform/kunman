//nolint:dupl
package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) UpdateDisplayName(ctx context.Context, in *npool.UpdateDisplayNameRequest) (*npool.UpdateDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithID(&in.ID, true),
		displayname1.WithEntID(&in.EntID, true),
		displayname1.WithAppID(&in.AppID, true),
		displayname1.WithName(in.Name, false),
		displayname1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDisplayNameResponse{
		Info: info,
	}, nil
}
