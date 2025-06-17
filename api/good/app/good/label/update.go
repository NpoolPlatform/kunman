//nolint:dupl
package label

import (
	"context"

	label1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"
)

func (s *Server) UpdateLabel(ctx context.Context, in *npool.UpdateLabelRequest) (*npool.UpdateLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithID(&in.ID, true),
		label1.WithEntID(&in.EntID, true),
		label1.WithAppID(&in.AppID, true),
		label1.WithIcon(in.Icon, false),
		label1.WithIconBgColor(in.IconBgColor, false),
		label1.WithLabelBgColor(in.LabelBgColor, false),
		label1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateLabelResponse{
		Info: info,
	}, nil
}
