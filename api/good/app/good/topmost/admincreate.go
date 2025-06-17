package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
)

func (s *Server) AdminCreateTopMost(ctx context.Context, in *npool.AdminCreateTopMostRequest) (*npool.AdminCreateTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithAppID(&in.TargetAppID, true),
		topmost1.WithTopMostType(&in.TopMostType, true),
		topmost1.WithTitle(&in.Title, true),
		topmost1.WithMessage(&in.Message, true),
		topmost1.WithTargetURL(in.TargetUrl, true),
		topmost1.WithStartAt(&in.StartAt, true),
		topmost1.WithEndAt(&in.EndAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateTopMostResponse{
		Info: info,
	}, nil
}
