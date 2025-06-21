package chain

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	chain1 "github.com/NpoolPlatform/kunman/gateway/chain/chain"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AdminUpdateChain(ctx context.Context, in *npool.AdminUpdateChainRequest) (*npool.AdminUpdateChainResponse, error) {
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithID(&in.ID, true),
		chain1.WithEntID(&in.EntID, true),
		chain1.WithChainType(in.ChainType, false),
		chain1.WithLogo(in.Logo, false),
		chain1.WithChainID(in.ChainID, false),
		chain1.WithNativeUnit(in.NativeUnit, false),
		chain1.WithAtomicUnit(in.AtomicUnit, false),
		chain1.WithUnitExp(in.UnitExp, false),
		chain1.WithNickname(in.NickName, false),
		chain1.WithGasType(in.GasType, false),
		chain1.WithENV(in.ENV, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateChain",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateChainResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateChain(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateChain",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateChainResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminUpdateChainResponse{
		Info: info,
	}, nil
}
