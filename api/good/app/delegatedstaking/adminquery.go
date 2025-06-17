package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) AdminGetAppDelegatedStakings(ctx context.Context, in *npool.AdminGetAppDelegatedStakingsRequest) (*npool.AdminGetAppDelegatedStakingsResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithAppID(&in.TargetAppID, true),
		delegatedstaking1.WithOffset(in.Offset),
		delegatedstaking1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDelegatedStakings(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppDelegatedStakingsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
