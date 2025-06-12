package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	allocatedcouponmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type AllocatedCouponCheckHandler struct {
	UserCheckHandler
	AllocatedCouponID *string
}

func (h *AllocatedCouponCheckHandler) CheckAllocatedCouponWithAllocatedCouponID(ctx context.Context, allocatedCouponID string) error {
	conds := &allocatedcouponmwpb.Conds{
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: allocatedCouponID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
	}
	handler, err := allocatedcouponmw.NewHandler(
		ctx,
		allocatedcouponmw.WithConds(conds),
		allocatedcouponmw.WithOffset(0),
		allocatedcouponmw.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	// TODO: should be replaced with exist api
	info, err := handler.GetCouponOnly(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid allocatedcoupon")
	}
	return nil
}
