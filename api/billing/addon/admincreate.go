package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/kunman/gateway/billing/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/addon"
)

func (s *Server) AdminCreateAddon(ctx context.Context, in *npool.AdminCreateAddonRequest) (*npool.AdminCreateAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithAppID(&in.TargetAppID, true),
		addon1.WithUsdPrice(&in.UsdPrice, true),
		addon1.WithCredit(&in.Credit, true),
		addon1.WithSortOrder(&in.SortOrder, false),
		addon1.WithEnabled(&in.Enabled, false),
		addon1.WithDescription(&in.Description, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAddon",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAddonResponse{
		Info: info,
	}, nil
}
