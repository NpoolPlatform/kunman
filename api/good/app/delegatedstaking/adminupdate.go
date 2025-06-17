package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) AdminUpdateAppDelegatedStaking(ctx context.Context, in *npool.AdminUpdateAppDelegatedStakingRequest) (*npool.AdminUpdateAppDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(&in.ID, true),
		delegatedstaking1.WithEntID(&in.EntID, true),
		delegatedstaking1.WithAppID(&in.TargetAppID, true),
		delegatedstaking1.WithAppGoodID(&in.AppGoodID, true),

		delegatedstaking1.WithPurchasable(in.Purchasable, false),
		delegatedstaking1.WithEnableProductPage(in.EnableProductPage, false),
		delegatedstaking1.WithProductPage(in.ProductPage, false),
		delegatedstaking1.WithOnline(in.Online, false),
		delegatedstaking1.WithVisible(in.Visible, false),
		delegatedstaking1.WithName(in.Name, false),
		delegatedstaking1.WithDisplayIndex(in.DisplayIndex, false),
		delegatedstaking1.WithBanner(in.Banner, false),

		delegatedstaking1.WithServiceStartAt(in.ServiceStartAt, false),
		delegatedstaking1.WithStartMode(in.StartMode, false),
		delegatedstaking1.WithEnableSetCommission(in.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateAppDelegatedStakingResponse{
		Info: info,
	}, nil
}
