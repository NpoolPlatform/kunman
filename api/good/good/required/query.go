package required

import (
	"context"

	required1 "github.com/NpoolPlatform/kunman/gateway/good/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/required"
)

func (s *Server) GetRequireds(ctx context.Context, in *npool.GetRequiredsRequest) (*npool.GetRequiredsResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithGoodID(in.GoodID, false),
		required1.WithOffset(in.Offset),
		required1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRequireds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRequiredsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
