package like

import (
	"context"

	like1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/like"
)

func (s *Server) GetMyLikes(ctx context.Context, in *npool.GetMyLikesRequest) (*npool.GetMyLikesResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(&in.UserID, true),
		like1.WithGoodID(in.GoodID, false),
		like1.WithAppGoodID(in.AppGoodID, false),
		like1.WithOffset(in.Offset),
		like1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLikes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMyLikesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetLikes(ctx context.Context, in *npool.GetLikesRequest) (*npool.GetLikesResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithAppID(&in.AppID, true),
		like1.WithUserID(in.TargetUserID, false),
		like1.WithGoodID(in.GoodID, false),
		like1.WithAppGoodID(in.AppGoodID, false),
		like1.WithOffset(in.Offset),
		like1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLikes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLikesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
