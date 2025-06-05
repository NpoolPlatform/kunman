package record

import (
	"context"

	record1 "github.com/NpoolPlatform/kunman/gateway/billing/user/credit/record"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/credit/record"
)

func (s *Server) GetUserCreditRecord(ctx context.Context, in *npool.GetUserCreditRecordRequest) (*npool.GetUserCreditRecordResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecord",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetUserCreditRecord(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecord",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCreditRecordResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUserCreditRecords(ctx context.Context, in *npool.GetUserCreditRecordsRequest) (*npool.GetUserCreditRecordsResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
		record1.WithOffset(in.Offset),
		record1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecords",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetUserCreditRecords(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecords",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCreditRecordsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetUserCreditRecordsCount(ctx context.Context, in *npool.GetUserCreditRecordsCountRequest) (*npool.GetUserCreditRecordsCountResponse, error) {
	handler, err := record1.NewHandler(
		ctx,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecordsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	total, err := handler.GetUserCreditRecordsCount(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUserCreditRecordsCount",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserCreditRecordsCountResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUserCreditRecordsCountResponse{
		Total: total,
	}, nil
}
