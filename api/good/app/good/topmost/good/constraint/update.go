//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
)

func (s *Server) UpdateTopMostGoodConstraint(ctx context.Context, in *npool.UpdateTopMostGoodConstraintRequest) (*npool.UpdateTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(&in.ID, true),
		constraint1.WithEntID(&in.EntID, true),
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
