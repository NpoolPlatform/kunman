package allocated

import (
	"context"
	"math/rand"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"

	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon/allocated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	coupon *ent.Coupon
}

func (h *createHandler) getCoupon(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		coupon, err := cli.
			Coupon.
			Query().
			Where(
				entcoupon.EntID(*h.CouponID),
				entcoupon.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		now := time.Now().Unix()
		if now < int64(coupon.StartAt) || now > int64(coupon.EndAt) {
			return wlog.Errorf("coupon can not be issued in current time")
		}
		h.coupon = coupon
		return nil
	})
}

func (h *createHandler) cashable() bool {
	if h.Cashable != nil {
		return *h.Cashable
	}
	probability := h.coupon.CashableProbability
	if probability.Cmp(decimal.NewFromInt(0)) <= 0 {
		return false
	}
	if probability.Cmp(decimal.NewFromInt(1)) >= 0 {
		return true
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	value := r.Float64()
	return decimal.NewFromFloat(value).Cmp(h.coupon.CashableProbability) <= 0
}

func (h *createHandler) createAllocatedCoupon(ctx context.Context, tx *ent.Tx) error {
	startAt := uint32(time.Now().Unix())
	couponScope := inspiretypes.CouponScope(inspiretypes.CouponScope_value[h.coupon.CouponScope])
	_cashable := h.cashable()
	if _, err := allocatedcrud.CreateSet(
		tx.CouponAllocated.Create(),
		&allocatedcrud.Req{
			EntID:        h.EntID,
			AppID:        h.AppID,
			CouponID:     h.CouponID,
			UserID:       h.UserID,
			StartAt:      &startAt,
			Denomination: &h.coupon.Denomination,
			CouponScope:  &couponScope,
			Cashable:     &_cashable,
			Extra:        h.Extra,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) validateAllocated() error {
	allocated := h.coupon.Allocated
	switch h.coupon.CouponType {
	case inspiretypes.CouponType_FixAmount.String():
		allocated = allocated.Add(h.coupon.Denomination)
	case inspiretypes.CouponType_Discount.String():
		allocated = allocated.Add(decimal.NewFromInt(1))
	default:
		return wlog.Errorf("invalid coupontype")
	}
	if allocated.Cmp(h.coupon.Circulation) > 0 {
		return wlog.Errorf("insufficient circulation")
	}
	return nil
}

func (h *createHandler) updateCoupon(ctx context.Context, tx *ent.Tx) error {
	allocated := h.coupon.Allocated
	switch h.coupon.CouponType {
	case inspiretypes.CouponType_FixAmount.String():
		allocated = allocated.Add(h.coupon.Denomination)
	case inspiretypes.CouponType_Discount.String():
		allocated = allocated.Add(decimal.NewFromInt(1))
	default:
		return wlog.Errorf("invalid coupontype")
	}
	if allocated.Cmp(h.coupon.Circulation) > 0 {
		return wlog.Errorf("insufficient circulation")
	}
	if _, err := tx.
		Coupon.
		UpdateOne(h.coupon).
		SetAllocated(allocated).
		Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateCoupon(ctx context.Context) error {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoupon(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAllocatedCoupon(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateCoupon(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *Handler) CalcluateAllocatedCoupon(ctx context.Context) (*npool.Coupon, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.getCoupon(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.validateAllocated(); err != nil {
		return nil, wlog.WrapError(err)
	}

	_cashable := handler.cashable()
	coupon := &npool.Coupon{
		AppID:        h.AppID.String(),
		CouponID:     h.CouponID.String(),
		UserID:       h.UserID.String(),
		Denomination: handler.coupon.Denomination.String(),
		Cashable:     _cashable,
	}

	return coupon, nil
}
