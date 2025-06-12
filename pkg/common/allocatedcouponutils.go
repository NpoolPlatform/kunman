//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	allocatedcouponmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func GetAllocatedCoupons(ctx context.Context, allocatedCouponIDs []string) (map[string]*allocatedcouponmwpb.Coupon, error) {
	for _, allocatedCouponID := range allocatedCouponIDs {
		if _, err := uuid.Parse(allocatedCouponID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	conds := &allocatedcouponmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: allocatedCouponIDs},
	}
	handler, err := allocatedcouponmw.NewHandler(
		ctx,
		allocatedcouponmw.WithConds(conds),
		allocatedcouponmw.WithOffset(0),
		allocatedcouponmw.WithLimit(int32(len(allocatedCouponIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	allocatedCoupons, _, err := handler.GetCoupons(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	allocatedCouponMap := map[string]*allocatedcouponmwpb.Coupon{}
	for _, allocatedCoupon := range allocatedCoupons {
		allocatedCouponMap[allocatedCoupon.EntID] = allocatedCoupon
	}
	return allocatedCouponMap, nil
}
