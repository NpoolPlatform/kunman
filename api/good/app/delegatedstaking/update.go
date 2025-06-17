package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) UpdateAppDelegatedStaking(ctx context.Context, in *npool.UpdateAppDelegatedStakingRequest) (*npool.UpdateAppDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(&in.ID, true),
		delegatedstaking1.WithEntID(&in.EntID, true),
		delegatedstaking1.WithAppID(&in.AppID, true),
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
		delegatedstaking1.WithEnableSetCommission(in.EnableSetCommission, false),
		delegatedstaking1.WithStartMode(in.StartMode, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppDelegatedStakingResponse{
		Info: info,
	}, nil
}
