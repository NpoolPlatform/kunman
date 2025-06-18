package executor

import (
	"context"
	"fmt"

	appcoinmwcli "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/reviewing/types"

	couponwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	reviewmwcli "github.com/NpoolPlatform/review-middleware/pkg/client/review"
	"github.com/google/uuid"
)

type couponwithdrawHandler struct {
	*couponwithdrawmwpb.CouponWithdraw
	persistent chan interface{}
	notif      chan interface{}
	done       chan interface{}
}

func (h *couponwithdrawHandler) checkCouponWithdrawReview(ctx context.Context) error {
	if _, err := uuid.Parse(h.ReviewID); err != nil {
		return err
	}
	review, err := reviewmwcli.GetReview(ctx, h.ReviewID)
	if err != nil {
		return err
	}
	if review == nil {
		return fmt.Errorf("invalid review")
	}
	if review.ObjectID != h.CouponWithdraw.EntID {
		return fmt.Errorf("objectid mismatch")
	}
	if review.AppID != h.CouponWithdraw.AppID {
		return fmt.Errorf("appid mismatch")
	}
	return nil
}

func (h *couponwithdrawHandler) checkAppCoin(ctx context.Context) error {
	coin, err := appcoinmwcli.GetCoinOnly(ctx, &appcoinmwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
	})
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}
	return nil
}

func (h *couponwithdrawHandler) final(ctx context.Context, err *error) { //nolint
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"CouponWithdraw", h.CouponWithdraw,
			"Error", *err,
		)
	}
	persistentCouponWithdraw := &types.PersistentCouponWithdraw{
		CouponWithdraw: h.CouponWithdraw,
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentCouponWithdraw, h.notif)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentCouponWithdraw, h.persistent)
}

func (h *couponwithdrawHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)

	if err := h.checkCouponWithdrawReview(ctx); err != nil {
		return err
	}
	if err := h.checkAppCoin(ctx); err != nil {
		return err
	}
	return nil
}
