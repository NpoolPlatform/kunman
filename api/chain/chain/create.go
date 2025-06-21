package chain

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	chain1 "github.com/NpoolPlatform/kunman/gateway/chain/chain"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AdminCreateChain(ctx context.Context, in *npool.AdminCreateChainRequest) (*npool.AdminCreateChainResponse, error) {
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithChainType(&in.ChainType, true),
		chain1.WithLogo(in.Logo, false),
		chain1.WithChainID(in.ChainID, false),
		chain1.WithNativeUnit(&in.NativeUnit, true),
		chain1.WithAtomicUnit(&in.AtomicUnit, true),
		chain1.WithUnitExp(&in.UnitExp, true),
		chain1.WithNickname(in.NickName, false),
		chain1.WithGasType(&in.GasType, true),
		chain1.WithENV(&in.ENV, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateChain",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateChainResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateChain(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateChain",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateChainResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminCreateChainResponse{
		Info: info,
	}, nil
}
