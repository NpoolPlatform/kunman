package persistent

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/approved/types"
	allocatedmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
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

	handler, err := allocatedmw.NewHandler(
		ctx,
		allocatedmw.WithEntID(&_couponwithdraw.AllocatedID, true),
	)
	if err != nil {
		return err
	}

	coupon, err := handler.GetCoupon(ctx)
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

	handler, err = allocatedmw.NewHandler(
		ctx,
		allocatedmw.WithID(&coupon.ID, true),
		allocatedmw.WithUsed(&used, true),
		allocatedmw.WithUsedByOrderID(&_couponwithdraw.EntID, true),
	)
	if err != nil {
		return err
	}

	if err := handler.UpdateCoupon(ctx); err != nil {
		return err
	}
	return nil
}
