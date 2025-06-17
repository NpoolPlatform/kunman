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

func (s *Server) CreateTopMostGoodConstraint(ctx context.Context, in *npool.CreateTopMostGoodConstraintRequest) (*npool.CreateTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithTopMostGoodID(&in.TopMostGoodID, true),
		constraint1.WithConstraint(&in.Constraint, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
