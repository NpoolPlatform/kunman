package addon

import (
	"context"

	addon1 "github.com/NpoolPlatform/kunman/gateway/billing/addon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/addon"
)

func (s *Server) GetAddon(ctx context.Context, in *npool.GetAddonRequest) (*npool.GetAddonResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddon",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetAddon(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddon",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAddonResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAddons(ctx context.Context, in *npool.GetAddonsRequest) (*npool.GetAddonsResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
		addon1.WithOffset(in.Offset),
		addon1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddons",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetAddons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddons",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAddonsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetAddonsCount(ctx context.Context, in *npool.GetAddonsCountRequest) (*npool.GetAddonsCountResponse, error) {
	handler, err := addon1.NewHandler(
		ctx,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddonsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.GetAddonsCount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAddonsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetAddonsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAddonsCountResponse{
		Total: total,
	}, nil
}
