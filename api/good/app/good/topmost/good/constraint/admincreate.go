//nolint:dupl
package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
)

func (s *Server) AdminCreateTopMostGoodConstraint(ctx context.Context, in *npool.AdminCreateTopMostGoodConstraintRequest) (*npool.AdminCreateTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.TargetAppID, true),
		constraint1.WithTopMostGoodID(&in.TopMostGoodID, true),
		constraint1.WithConstraint(&in.Constraint, true),
		constraint1.WithTargetValue(in.TargetValue, false),
		constraint1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}
