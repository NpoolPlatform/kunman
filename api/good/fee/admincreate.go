package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/kunman/gateway/good/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/fee"
)

func (s *Server) AdminCreateFee(ctx context.Context, in *npool.AdminCreateFeeRequest) (*npool.AdminCreateFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithGoodType(&in.GoodType, true),
		fee1.WithName(&in.Name, true),
		fee1.WithSettlementType(&in.SettlementType, true),
		fee1.WithUnitValue(&in.UnitValue, true),
		fee1.WithDurationDisplayType(&in.DurationDisplayType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateFeeResponse{
		Info: info,
	}, nil
}
