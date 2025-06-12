package coupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (h *Handler) CreateCoupon(ctx context.Context) (*npool.Coupon, error) {
	if *h.EndAt <= *h.StartAt {
		return nil, wlog.Errorf("endat less than startat")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	switch *h.CouponType {
	case types.CouponType_FixAmount:
		if h.Denomination.Cmp(*h.Circulation) > 0 {
			return nil, wlog.Errorf("denomination > circulation")
		}
	case types.CouponType_Discount:
		if h.Denomination.Cmp(decimal.NewFromInt(100)) > 0 { //nolint
			return nil, wlog.Errorf("100 discounat not allowed")
		}
		if h.CashableProbability != nil {
			return nil, wlog.Errorf("probability must set with fix amount")
		}
	}

	if h.CouponConstraint != nil {
		switch *h.CouponConstraint {
		case types.CouponConstraint_Normal:
			threshold := decimal.RequireFromString("0")
			h.Threshold = &threshold
		case types.CouponConstraint_PaymentThreshold:
			if h.Threshold == nil {
				return nil, wlog.Errorf("threshold is must")
			}
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := couponcrud.CreateSet(
			cli.Coupon.Create(),
			&couponcrud.Req{
				EntID:               h.EntID,
				CouponType:          h.CouponType,
				AppID:               h.AppID,
				Denomination:        h.Denomination,
				Circulation:         h.Circulation,
				IssuedBy:            h.IssuedBy,
				StartAt:             h.StartAt,
				EndAt:               h.EndAt,
				DurationDays:        h.DurationDays,
				Message:             h.Message,
				Name:                h.Name,
				CouponConstraint:    h.CouponConstraint,
				CouponScope:         h.CouponScope,
				Threshold:           h.Threshold,
				Random:              h.Random,
				CashableProbability: h.CashableProbability,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetCoupon(ctx)
}
