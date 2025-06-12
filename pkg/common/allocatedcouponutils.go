//nolint:dupl
package common

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	allocatedcouponmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/coupon/allocated"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	allocatedcouponmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"

	"github.com/google/uuid"
)

func GetAllocatedCoupons(ctx context.Context, allocatedCouponIDs []string) (map[string]*allocatedcouponmwpb.Coupon, error) {
	for _, allocatedCouponID := range allocatedCouponIDs {
		if _, err := uuid.Parse(allocatedCouponID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	allocatedCoupons, _, err := allocatedcouponmwcli.GetCoupons(ctx, &allocatedcouponmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: allocatedCouponIDs},
	}, int32(0), int32(len(allocatedCouponIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	allocatedCouponMap := map[string]*allocatedcouponmwpb.Coupon{}
	for _, allocatedCoupon := range allocatedCoupons {
		allocatedCouponMap[allocatedCoupon.EntID] = allocatedCoupon
	}
	return allocatedCouponMap, nil
}
