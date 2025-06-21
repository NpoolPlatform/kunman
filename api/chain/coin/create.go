//nolint:nolintlint,dupl
package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	coin1 "github.com/NpoolPlatform/kunman/gateway/chain/coin"
	npool "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *npool.CreateCoinRequest) (*npool.CreateCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithName(&in.Name, true),
		coin1.WithUnit(&in.Unit, true),
		coin1.WithENV(&in.ENV, true),
		coin1.WithChainType(&in.ChainType, true),
		coin1.WithChainNativeUnit(&in.ChainNativeUnit, true),
		coin1.WithChainAtomicUnit(&in.ChainAtomicUnit, true),
		coin1.WithChainUnitExp(&in.ChainUnitExp, true),
		coin1.WithGasType(&in.GasType, true),
		coin1.WithChainID(&in.ChainID, true),
		coin1.WithChainNickname(&in.ChainNickname, true),
		coin1.WithChainNativeCoinName(&in.ChainNativeCoinName, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info,
	}, nil
}
