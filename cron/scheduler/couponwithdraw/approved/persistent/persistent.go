package persistent

import (
	"context"
	"fmt"

	allocatedmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"

	allocatedmwcli "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/approved/types"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, couponwithdraw interface{}, reward, notif, done chan interface{}) error {
	_couponwithdraw, ok := couponwithdraw.(*types.PersistentCouponWithdraw)
	if !ok {
		return fmt.Errorf("invalid coupon withdraw")
	}

	defer asyncfeed.AsyncFeed(ctx, _couponwithdraw, done)

	coupon, err := allocatedmwcli.GetCoupon(ctx, _couponwithdraw.AllocatedID)
	if err != nil {
		return err
	}
	if coupon == nil {
		return fmt.Errorf("coupon not found")
	}
	if coupon.Used {
		return nil
	}
	used := true
	if _, err := allocatedmwcli.UpdateCoupon(ctx, &allocatedmwpb.CouponReq{
		ID:            &coupon.ID,
		Used:          &used,
		UsedByOrderID: &_couponwithdraw.EntID,
	}); err != nil {
		return err
	}
	return nil
}
