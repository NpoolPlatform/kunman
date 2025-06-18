package persistent

import (
	"context"
	"fmt"

	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	reviewmw "github.com/NpoolPlatform/kunman/middleware/review/review"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/reviewing/types"
	allocatedmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	couponwithdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw/coupon"
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

	reviewHandler, err := reviewmw.NewHandler(
		ctx,
		reviewmw.WithEntID(&_couponwithdraw.ReviewID, true),
	)
	if err != nil {
		return err
	}

	review, err := reviewHandler.GetReview(ctx)
	if err != nil {
		return err
	}
	if review == nil {
		return fmt.Errorf("review not found")
	}

	state := ledgertypes.WithdrawState_DefaultWithdrawState
	switch review.State {
	case reviewtypes.ReviewState_Rejected:
		state = ledgertypes.WithdrawState_Rejected
	case reviewtypes.ReviewState_Approved:
		state = ledgertypes.WithdrawState_Approved
	default:
		return nil
	}

	couponHandler, err := couponwithdrawmw.NewHandler(
		ctx,
		couponwithdrawmw.WithID(&_couponwithdraw.ID, true),
		couponwithdrawmw.WithState(&state, true),
	)
	if err != nil {
		return err
	}

	if _, err := couponHandler.UpdateCouponWithdraw(ctx); err != nil {
		return err
	}

	if state != ledgertypes.WithdrawState_Approved {
		return nil
	}

	allocatedHandler, err := allocatedmw.NewHandler(
		ctx,
		allocatedmw.WithEntID(&_couponwithdraw.AllocatedID, true),
	)
	if err != nil {
		return err
	}

	coupon, err := allocatedHandler.GetCoupon(ctx)
	if err != nil {
		return err
	}
	if coupon == nil {
		return fmt.Errorf("coupon not found")
	}
	if coupon.Used {
		return fmt.Errorf("coupon already used")
	}
	used := true

	allocatedHandler, err = allocatedmw.NewHandler(
		ctx,
		allocatedmw.WithID(&coupon.ID, true),
		allocatedmw.WithUsed(&used, true),
		allocatedmw.WithUsedByOrderID(&_couponwithdraw.EntID, true),
	)
	if err != nil {
		return err
	}

	if err := allocatedHandler.UpdateCoupon(ctx); err != nil {
		return err
	}
	return nil
}
