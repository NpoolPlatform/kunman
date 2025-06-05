package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/kunman/gateway/billing/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/addon"
)

func (s *Server) AdminDeleteAddon(ctx context.Context, in *npool.AdminDeleteAddonRequest) (*npool.AdminDeleteAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithID(&in.ID, true),
		addon1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAddonResponse{
		Info: info,
	}, nil
}
