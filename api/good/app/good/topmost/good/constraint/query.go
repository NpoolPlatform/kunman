package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/constraint"
)

func (s *Server) GetTopMostGoodConstraints(ctx context.Context, in *npool.GetTopMostGoodConstraintsRequest) (*npool.GetTopMostGoodConstraintsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithAppID(&in.AppID, true),
		constraint1.WithOffset(in.Offset),
		constraint1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetConstraints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodConstraintsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
