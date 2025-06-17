package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"
)

func (s *Server) AdminGetRequireds(ctx context.Context, in *npool.AdminGetRequiredsRequest) (*npool.AdminGetRequiredsResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithAppID(&in.TargetAppID, true),
		required1.WithAppGoodID(in.AppGoodID, false),
		required1.WithOffset(in.Offset),
		required1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRequireds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetRequiredsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
