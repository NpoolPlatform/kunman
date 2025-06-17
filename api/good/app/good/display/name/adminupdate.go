//nolint:dupl
package displayname

import (
	"context"

	displayname1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"
)

func (s *Server) AdminUpdateDisplayName(ctx context.Context, in *npool.AdminUpdateDisplayNameRequest) (*npool.AdminUpdateDisplayNameResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithID(&in.ID, true),
		displayname1.WithEntID(&in.EntID, true),
		displayname1.WithAppID(&in.TargetAppID, true),
		displayname1.WithName(in.Name, false),
		displayname1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDisplayNameResponse{
		Info: info,
	}, nil
}
