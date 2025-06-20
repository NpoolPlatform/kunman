package event

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	couponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	allocated1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	eventcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event"
	registration1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/registration"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
}

//nolint:dupl
func (h *rewardHandler) condGood() error {
	switch *h.EventType {
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		if h.GoodID == nil {
			return wlog.Errorf("need goodid")
		}
		if h.AppGoodID == nil {
			return wlog.Errorf("need appgoodid")
		}
		h.Conds.GoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID}
		h.Conds.AppGoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID}
	}
	return nil
}

func (h *rewardHandler) getEvent(ctx context.Context) (*npool.Event, error) {
	h.Conds = &eventcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		EventType: &cruder.Cond{Op: cruder.EQ, Val: *h.EventType},
	}
	if err := h.condGood(); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetEventOnly(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}

func (h *rewardHandler) calculateCredits(ev *npool.Event) (decimal.Decimal, error) {
	credits, err := decimal.NewFromString(ev.Credits)
	if err != nil {
		return decimal.NewFromInt(0), wlog.WrapError(err)
	}

	_credits, err := decimal.NewFromString(ev.CreditsPerUSD)
	if err != nil {
		return decimal.NewFromInt(0), wlog.WrapError(err)
	}

	credits = credits.Add(_credits.Mul(*h.Amount))
	return credits, nil
}

func (h *rewardHandler) allocateCoupons(ctx context.Context, ev *npool.Event) error {
	coups := []*couponmwpb.Coupon{}
	for _, id := range ev.CouponIDs {
		_id := id
		handler, err := coupon1.NewHandler(
			ctx,
			coupon1.WithEntID(&_id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		_coupon, err := handler.GetCoupon(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _coupon == nil {
			return wlog.Errorf("invalid coupon")
		}

		now := time.Now().Unix()
		if now < int64(_coupon.StartAt) || now > int64(_coupon.EndAt) {
			logger.Sugar().Errorw("coupon can not be issued in current time")
			continue
		}
		coups = append(coups, _coupon)
	}

	for _, coup := range coups {
		userID := h.UserID.String()

		handler, err := allocated1.NewHandler(
			ctx,
			allocated1.WithAppID(&coup.AppID, true),
			allocated1.WithUserID(&userID, true),
			allocated1.WithCouponID(&coup.EntID, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		if err := handler.CreateCoupon(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}

	return nil
}

func (h *rewardHandler) rewardSelf(ctx context.Context) ([]*npool.Credit, error) {
	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if ev == nil {
		return nil, nil
	}

	if *h.Consecutive > ev.MaxConsecutive {
		return nil, nil
	}

	credits, err := h.calculateCredits(ev)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	// We don't care about result of allocate coupon
	if err := h.allocateCoupons(ctx, ev); err != nil {
		logger.Sugar().Warnw(
			"rewardSelf",
			"Event", ev,
			"Error", err,
		)
	}

	_credits := []*npool.Credit{}
	if credits.Cmp(decimal.NewFromInt(0)) > 0 {
		_credits = append(_credits, &npool.Credit{
			AppID:   h.AppID.String(),
			UserID:  h.UserID.String(),
			Credits: credits.String(),
		})
	}

	return _credits, nil
}

func (h *rewardHandler) rewardAffiliate(ctx context.Context) ([]*npool.Credit, error) {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.AppID = h.AppID
	handler.InviteeID = h.UserID

	_, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(inviterIDs) == 0 {
		return nil, nil
	}

	ev, err := h.getEvent(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if ev == nil {
		return nil, nil
	}

	if ev.InviterLayers == 0 {
		return nil, nil
	}

	credits := []*npool.Credit{}
	i := uint32(0)
	const inviterIgnore = 2
	j := len(inviterIDs) - inviterIgnore

	appID := h.AppID.String()
	goodID := h.GoodID.String()
	appGoodID := h.AppGoodID.String()
	amount := h.Amount.String()

	for ; i < ev.InviterLayers && j >= 0; i++ {
		handler, err := NewHandler(
			ctx,
			WithAppID(&appID, true),
			WithUserID(&inviterIDs[j], true),
			WithEventType(h.EventType, true),
			WithGoodID(&goodID, true),
			WithAppGoodID(&appGoodID, true),
			WithConsecutive(h.Consecutive, true),
			WithAmount(&amount, true),
		)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		_handler := &rewardHandler{
			Handler: handler,
		}

		credit, err := _handler.rewardSelf(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		j--
		if len(credit) == 0 {
			continue
		}

		credits = append(credits, credit...)
	}

	return credits, nil
}

func (h *Handler) RewardEvent(ctx context.Context) ([]*npool.Credit, error) {
	handler := &rewardHandler{
		Handler: h,
	}

	switch *h.EventType {
	case basetypes.UsedFor_Signup:
		fallthrough //nolint
	case basetypes.UsedFor_Purchase:
		fallthrough //nolint
	case basetypes.UsedFor_SimulateOrderProfit:
		return handler.rewardSelf(ctx)
	case basetypes.UsedFor_AffiliateSignup:
		fallthrough //nolint
	case basetypes.UsedFor_AffiliatePurchase:
		return handler.rewardAffiliate(ctx)
	default:
		return nil, wlog.Errorf("not implemented")
	}
}
