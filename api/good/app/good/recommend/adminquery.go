//nolint:dupl
package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/recommend"
)

func (s *Server) AdminGetRecommends(ctx context.Context, in *npool.AdminGetRecommendsRequest) (*npool.AdminGetRecommendsResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithAppID(&in.TargetAppID, true),
		recommend1.WithRecommenderID(in.TargetUserID, false),
		recommend1.WithAppGoodID(in.AppGoodID, false),
		recommend1.WithOffset(in.Offset),
		recommend1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRecommends(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetRecommendsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
