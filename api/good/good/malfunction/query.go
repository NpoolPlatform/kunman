package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/kunman/gateway/good/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"
)

func (s *Server) GetMalfunctions(ctx context.Context, in *npool.GetMalfunctionsRequest) (*npool.GetMalfunctionsResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithGoodID(in.GoodID, false),
		malfunction1.WithOffset(in.Offset),
		malfunction1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunctions",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetMalfunctions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunctions",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMalfunctionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
