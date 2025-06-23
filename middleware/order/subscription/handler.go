//nolint:dupl
package subscriptionorder

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	ordercouponcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/coupon"
	orderlockcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/lock"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/statebase"
	paymentbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment"
	paymentbalancecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/balance"
	paymentbalancelockcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/balance/lock"
	paymentfiatcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/fiat"
	paymenttransfercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/transfer"
	subscriptionordercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription"
	subscriptionorderstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription/state"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	subscriptionordercrud.Req
	OrderBaseReq                *orderbasecrud.Req
	OrderStateBaseReq           *orderstatebasecrud.Req
	SubscriptionOrderStateReq   *subscriptionorderstatecrud.Req
	PaymentBaseReq              *paymentbasecrud.Req
	PaymentBalanceReqs          []*paymentbalancecrud.Req
	PaymentTransferReqs         []*paymenttransfercrud.Req
	PaymentFiatReqs             []*paymentfiatcrud.Req
	LedgerLockReq               *orderlockcrud.Req
	PaymentBalanceLockReq       *paymentbalancelockcrud.Req
	OrderCouponReqs             []*ordercouponcrud.Req
	SubscriptionOrderConds      *subscriptionordercrud.Conds
	OrderBaseConds              *orderbasecrud.Conds
	OrderStateBaseConds         *orderstatebasecrud.Conds
	SubscriptionOrderStateConds *subscriptionorderstatecrud.Conds
	OrderCouponConds            *ordercouponcrud.Conds
	Offset                      int32
	Limit                       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderBaseReq:              &orderbasecrud.Req{},
		OrderStateBaseReq:         &orderstatebasecrud.Req{},
		SubscriptionOrderStateReq: &subscriptionorderstatecrud.Req{},
		LedgerLockReq: &orderlockcrud.Req{
			LockType: func() *types.OrderLockType { e := types.OrderLockType_LockBalance; return &e }(),
		},
		PaymentBaseReq:              &paymentbasecrud.Req{},
		PaymentBalanceLockReq:       &paymentbalancelockcrud.Req{},
		SubscriptionOrderConds:      &subscriptionordercrud.Conds{},
		OrderBaseConds:              &orderbasecrud.Conds{},
		OrderStateBaseConds:         &orderstatebasecrud.Conds{},
		SubscriptionOrderStateConds: &subscriptionorderstatecrud.Conds{},
		OrderCouponConds:            &ordercouponcrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseReq.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseReq.UserID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseReq.GoodID = &_id
		return nil
	}
}

func WithGoodType(e *goodtypes.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case goodtypes.GoodType_Subscription:
		case goodtypes.GoodType_OneShot:
		default:
			return wlog.Errorf("invalid goodtype")
		}
		h.OrderBaseReq.GoodType = e
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseReq.AppGoodID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
		h.OrderBaseReq.EntID = &_id
		h.OrderStateBaseReq.OrderID = &_id
		h.SubscriptionOrderStateReq.OrderID = &_id
		h.LedgerLockReq.OrderID = &_id
		h.PaymentBaseReq.OrderID = &_id
		return nil
	}
}

func WithOrderType(orderType *types.OrderType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderType == nil {
			if must {
				return wlog.Errorf("invalid ordertype")
			}
			return nil
		}
		switch *orderType {
		case types.OrderType_Airdrop:
		case types.OrderType_Offline:
		case types.OrderType_Normal:
		default:
			return wlog.Errorf("invalid ordertype")
		}
		h.OrderBaseReq.OrderType = orderType
		return nil
	}
}

func WithPaymentType(paymentType *types.PaymentType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if paymentType == nil {
			if must {
				return wlog.Errorf("invalid paymenttype")
			}
			return nil
		}
		switch *paymentType {
		case types.PaymentType_PayWithBalanceOnly:
		case types.PaymentType_PayWithTransferOnly:
		case types.PaymentType_PayWithTransferAndBalance:
		case types.PaymentType_PayWithOffline:
		case types.PaymentType_PayWithNoPayment:
		case types.PaymentType_PayWithFiatOnly:
		case types.PaymentType_PayWithFiatAndBalance:
		default:
			return wlog.Errorf("invalid paymentType")
		}
		h.OrderStateBaseReq.PaymentType = paymentType
		return nil
	}
}

func WithGoodValueUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid goodvalueusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid goodvalueusd")
		}
		h.GoodValueUSD = &amount
		return nil
	}
}

func WithPaymentAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid paymentamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid paymentamountusd")
		}
		h.PaymentAmountUSD = &amount
		return nil
	}
}

func WithDiscountAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid discountamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid discountamountusd")
		}
		h.DiscountAmountUSD = &amount
		return nil
	}
}

func WithPromotionID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid promotionid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PromotionID = &_id
		return nil
	}
}

func WithLifeSeconds(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return wlog.Errorf("invalid lifeseconds")
			}
			return nil
		}
		if *duration <= 0 {
			return wlog.Errorf("invalid lifeseconds")
		}
		h.LifeSeconds = duration
		return nil
	}
}

//nolint:gocyclo
func WithOrderState(state *types.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid orderstate")
			}
			return nil
		}
		switch *state {
		case types.OrderState_OrderStateCreated:
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaymentTransferReceived:
		case types.OrderState_OrderStatePaymentTransferBookKeeping:
		case types.OrderState_OrderStatePaymentSpendBalance:
		case types.OrderState_OrderStateTransferGoodStockLocked:
		case types.OrderState_OrderStateAddCommission:
		case types.OrderState_OrderStateAchievementBookKeeping:
		case types.OrderState_OrderStatePaymentUnlockAccount:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateTransferGoodStockWaitStart:
		case types.OrderState_OrderStateInService:
		case types.OrderState_OrderStatePaymentTimeout:
		case types.OrderState_OrderStatePreCancel:
		case types.OrderState_OrderStatePreExpired:
		case types.OrderState_OrderStateRestoreExpiredStock:
		case types.OrderState_OrderStateRestoreCanceledStock:
		case types.OrderState_OrderStateCancelAchievement:
		case types.OrderState_OrderStateDeductLockedCommission:
		case types.OrderState_OrderStateReturnCanceledBalance:
		case types.OrderState_OrderStateCanceledTransferBookKeeping:
		case types.OrderState_OrderStateCancelUnlockPaymentAccount:
		case types.OrderState_OrderStateCanceled:
		case types.OrderState_OrderStateExpired:
		default:
			return wlog.Errorf("invalid orderstate")
		}
		h.OrderStateBaseReq.OrderState = state
		return nil
	}
}

func WithCreateMethod(e *types.OrderCreateMethod, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid createmethod")
			}
			return nil
		}
		switch *e {
		case types.OrderCreateMethod_OrderCreatedByPurchase:
		case types.OrderCreateMethod_OrderCreatedByAdmin:
		case types.OrderCreateMethod_OrderCreatedByRenew:
		default:
			return wlog.Errorf("invalid createmethod")
		}
		h.OrderBaseReq.CreateMethod = e
		return nil
	}
}

func WithUserSetPaid(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid usersetpaid")
			}
			return nil
		}
		h.SubscriptionOrderStateReq.UserSetPaid = value
		return nil
	}
}

func WithUserSetCanceled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid usersetcanceled")
			}
			return nil
		}
		h.SubscriptionOrderStateReq.UserSetCanceled = value
		return nil
	}
}

func WithAdminSetCanceled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid adminsetcanceled")
			}
			return nil
		}
		h.SubscriptionOrderStateReq.AdminSetCanceled = value
		return nil
	}
}

func WithPaymentState(state *types.PaymentState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid paymentstate")
			}
			return nil
		}
		switch *state {
		case types.PaymentState_PaymentStateWait:
		case types.PaymentState_PaymentStateCanceled:
		case types.PaymentState_PaymentStateTimeout:
		case types.PaymentState_PaymentStateDone:
		case types.PaymentState_PaymentStateNoPayment:
		default:
			return wlog.Errorf("invalid paymentstate")
		}
		h.SubscriptionOrderStateReq.PaymentState = state
		return nil
	}
}

func WithLedgerLockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid ledgerlockid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.LedgerLockReq.EntID = &_id
		h.PaymentBalanceLockReq.LedgerLockID = &_id
		return nil
	}
}

func WithPaymentID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid paymentid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PaymentBaseReq.EntID = &_id
		h.SubscriptionOrderStateReq.PaymentID = &_id
		h.PaymentBalanceLockReq.PaymentID = &_id
		return nil
	}
}

func WithDealEventID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid dealeventid")
			}
			return nil
		}
		h.SubscriptionOrderStateReq.DealEventID = id
		return nil
	}
}

func WithCouponIDs(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			id, err := uuid.Parse(s)
			if err != nil {
				return wlog.WrapError(err)
			}
			// Fill order id later
			h.OrderCouponReqs = append(h.OrderCouponReqs, &ordercouponcrud.Req{
				CouponID: &id,
			})
		}
		return nil
	}
}

func WithPaymentBalances(bs []*paymentmiddlewarepb.PaymentBalanceReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymentbalancecrud.Req{}

			id, err := uuid.Parse(b.GetCoinTypeID())
			if err != nil {
				return wlog.WrapError(err)
			}
			req.CoinTypeID = &id

			if b.LocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(b.GetLocalCoinUSDCurrency())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.LocalCoinUSDCurrency = &amount
			}

			amount1, err := decimal.NewFromString(b.GetAmount())
			if err != nil {
				return wlog.WrapError(err)
			}
			req.Amount = &amount1

			amount2, err := decimal.NewFromString(b.GetLiveCoinUSDCurrency())
			if err != nil {
				return wlog.WrapError(err)
			}
			req.LiveCoinUSDCurrency = &amount2

			h.PaymentBalanceReqs = append(h.PaymentBalanceReqs, req)
		}
		return nil
	}
}

func WithPaymentTransfers(bs []*paymentmiddlewarepb.PaymentTransferReq, must bool) func(context.Context, *Handler) error { //nolint:gocyclo
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymenttransfercrud.Req{}

			if b.EntID != nil {
				id0, err := uuid.Parse(b.GetEntID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.EntID = &id0
			}

			if b.CoinTypeID != nil {
				id1, err := uuid.Parse(b.GetCoinTypeID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.CoinTypeID = &id1
			}

			if b.LocalCoinUSDCurrency != nil {
				amount, err := decimal.NewFromString(*b.LocalCoinUSDCurrency)
				if err != nil {
					return wlog.WrapError(err)
				}
				req.LocalCoinUSDCurrency = &amount
			}

			if b.AccountID != nil {
				id2, err := uuid.Parse(b.GetAccountID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.AccountID = &id2
			}

			if b.Amount != nil {
				amount1, err := decimal.NewFromString(b.GetAmount())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.Amount = &amount1
			}

			if b.StartAmount != nil {
				amount2, err := decimal.NewFromString(b.GetStartAmount())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.StartAmount = &amount2
			}

			if b.LiveCoinUSDCurrency != nil {
				amount3, err := decimal.NewFromString(*b.LiveCoinUSDCurrency)
				if err != nil {
					return wlog.WrapError(err)
				}
				req.LiveCoinUSDCurrency = &amount3
			}

			if b.FinishAmount != nil {
				amount4, err := decimal.NewFromString(b.GetFinishAmount())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.FinishAmount = &amount4
			}

			h.PaymentTransferReqs = append(h.PaymentTransferReqs, req)
		}
		return nil
	}
}

func WithPaymentFiats(bs []*paymentmiddlewarepb.PaymentFiatReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymentfiatcrud.Req{}

			if b.EntID != nil {
				id0, err := uuid.Parse(b.GetEntID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.EntID = &id0
			}

			if b.FiatID != nil {
				id1, err := uuid.Parse(b.GetFiatID())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.FiatID = &id1
			}

			if b.Amount != nil {
				amount1, err := decimal.NewFromString(b.GetAmount())
				if err != nil {
					return wlog.WrapError(err)
				}
				req.Amount = &amount1
			}

			if b.USDCurrency != nil {
				amount3, err := decimal.NewFromString(*b.USDCurrency)
				if err != nil {
					return wlog.WrapError(err)
				}
				req.USDCurrency = &amount3
			}

			req.PaymentChannel = b.PaymentChannel
			req.ChannelPaymentID = b.ChannelPaymentID

			h.PaymentFiatReqs = append(h.PaymentFiatReqs, req)
		}
		return nil
	}
}

//nolint:funlen,gocyclo
func (h *Handler) withOrderBaseConds(conds *npool.Conds) error {
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetOrderIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.UserID = &cruder.Cond{
			Op:  conds.GetUserID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.OrderType != nil {
		h.OrderBaseConds.OrderType = &cruder.Cond{
			Op:  conds.GetOrderType().GetOp(),
			Val: types.OrderType(conds.GetOrderType().GetValue()),
		}
	}
	if conds.CreatedAt != nil {
		h.OrderBaseConds.CreatedAt = &cruder.Cond{
			Op:  conds.GetCreatedAt().GetOp(),
			Val: conds.GetCreatedAt().GetValue(),
		}
	}
	if conds.UpdatedAt != nil {
		h.OrderBaseConds.UpdatedAt = &cruder.Cond{
			Op:  conds.GetUpdatedAt().GetOp(),
			Val: conds.GetUpdatedAt().GetValue(),
		}
	}
	return nil
}

func (h *Handler) withSubscriptionOrderConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.SubscriptionOrderConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.IDs != nil {
		h.SubscriptionOrderConds.IDs = &cruder.Cond{
			Op:  conds.GetIDs().GetOp(),
			Val: conds.GetIDs().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionOrderConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.EntIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetEntIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.SubscriptionOrderConds.EntIDs = &cruder.Cond{
			Op:  conds.GetEntIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionOrderConds.OrderID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.SubscriptionOrderConds.OrderIDs = &cruder.Cond{
			Op:  conds.GetOrderIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withOrderStateBaseConds(conds *npool.Conds) {
	if conds.OrderState != nil {
		h.OrderStateBaseConds.OrderState = &cruder.Cond{
			Op:  conds.GetOrderState().GetOp(),
			Val: types.OrderState(conds.GetOrderState().GetValue()),
		}
	}
	if conds.OrderStates != nil {
		_types := []types.OrderState{}
		for _, _type := range conds.GetOrderStates().GetValue() {
			_types = append(_types, types.OrderState(_type))
		}
		h.OrderStateBaseConds.OrderStates = &cruder.Cond{
			Op:  conds.GetOrderStates().GetOp(),
			Val: _types,
		}
	}
	if conds.PaymentType != nil {
		h.OrderStateBaseConds.PaymentType = &cruder.Cond{
			Op:  conds.GetPaymentType().GetOp(),
			Val: types.PaymentType(conds.GetPaymentType().GetValue()),
		}
	}
	if conds.PaymentTypes != nil {
		_types := []types.PaymentType{}
		for _, _type := range conds.GetPaymentTypes().GetValue() {
			_types = append(_types, types.PaymentType(_type))
		}
		h.OrderStateBaseConds.PaymentTypes = &cruder.Cond{
			Op:  conds.GetPaymentTypes().GetOp(),
			Val: _types,
		}
	}
}

func (h *Handler) withSubscriptionOrderStateConds(conds *npool.Conds) error {
	if conds.PaymentState != nil {
		h.SubscriptionOrderStateConds.PaymentState = &cruder.Cond{
			Op:  conds.GetPaymentState().GetOp(),
			Val: types.PaymentState(conds.GetPaymentState().GetValue()),
		}
	}
	if conds.PaymentStates != nil {
		_types := []types.PaymentState{}
		for _, _type := range conds.GetPaymentStates().GetValue() {
			_types = append(_types, types.PaymentState(_type))
		}
		h.SubscriptionOrderStateConds.PaymentStates = &cruder.Cond{
			Op:  conds.GetPaymentStates().GetOp(),
			Val: _types,
		}
	}
	if conds.UserSetCanceled != nil {
		h.SubscriptionOrderStateConds.UserSetCanceled = &cruder.Cond{
			Op:  conds.GetUserSetCanceled().GetOp(),
			Val: conds.GetUserSetCanceled().GetValue(),
		}
	}
	if conds.AdminSetCanceled != nil {
		h.SubscriptionOrderStateConds.AdminSetCanceled = &cruder.Cond{
			Op:  conds.GetAdminSetCanceled().GetOp(),
			Val: conds.GetAdminSetCanceled().GetValue(),
		}
	}
	if conds.PaidAt != nil {
		h.SubscriptionOrderStateConds.PaidAt = &cruder.Cond{
			Op:  conds.GetPaidAt().GetOp(),
			Val: conds.GetPaidAt().GetValue(),
		}
	}
	return nil
}

func (h *Handler) withOrderCouponConds(conds *npool.Conds) error {
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderCouponConds.OrderID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderCouponConds.OrderIDs = &cruder.Cond{
			Op:  conds.GetOrderIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CouponID != nil {
		id, err := uuid.Parse(conds.GetCouponID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderCouponConds.CouponID = &cruder.Cond{
			Op:  conds.GetCouponID().GetOp(),
			Val: id,
		}
	}
	if conds.CouponIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetCouponIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderCouponConds.CouponIDs = &cruder.Cond{
			Op:  conds.GetCouponIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withOrderBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withSubscriptionOrderConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		h.withOrderStateBaseConds(conds)
		if err := h.withOrderCouponConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withSubscriptionOrderStateConds(conds)
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
