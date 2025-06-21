package chain

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	chain1 "github.com/NpoolPlatform/kunman/gateway/chain/chain"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetChains(ctx context.Context, in *npool.GetChainsRequest) (*npool.GetChainsResponse, error) {
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithOffset(in.GetOffset()),
		chain1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChains",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetChains(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChains",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChainsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
