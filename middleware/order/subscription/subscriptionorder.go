package subscriptionorder

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SubscriptionOrder interface {
	OrderID() uuid.UUID
	PaymentID() uuid.UUID
	LedgerLockID() uuid.UUID
}

type subscriptionOrder struct {
	entOrderBase              *ent.OrderBase
	entOrderStateBase         *ent.OrderStateBase
	entSubscriptionOrder      *ent.SubscriptionOrder
	entSubscriptionOrderState *ent.SubscriptionOrderState
	entPaymentBase            *ent.PaymentBase
	entPaymentBalances        []*ent.PaymentBalance
	entPaymentTransfers       []*ent.PaymentTransfer
	entLedgerLock             *ent.OrderLock
	entOrderCoupons           []*ent.OrderCoupon
	payWithMeOrderIDs         []uuid.UUID
}

func (f *subscriptionOrder) OrderID() uuid.UUID {
	return f.entSubscriptionOrder.OrderID
}

func (f *subscriptionOrder) UserID() uuid.UUID {
	return f.entOrderBase.UserID
}

func (f *subscriptionOrder) PaymentID() uuid.UUID {
	if f.entPaymentBase == nil {
		return uuid.Nil
	}
	return f.entPaymentBase.EntID
}

func (f *subscriptionOrder) LedgerLockID() uuid.UUID {
	if f.entLedgerLock == nil {
		return uuid.Nil
	}
	return f.entLedgerLock.EntID
}

func (f *subscriptionOrder) Exist() bool {
	return f.entOrderBase != nil
}

func (f *subscriptionOrder) OrderBaseID() uint32 {
	return f.entOrderBase.ID
}

func (f *subscriptionOrder) OrderType() types.OrderType {
	return types.OrderType(types.OrderType_value[f.entOrderBase.OrderType])
}

func (f *subscriptionOrder) OrderStateBaseID() uint32 {
	return f.entOrderStateBase.ID
}

func (f *subscriptionOrder) SubscriptionOrderID() uint32 {
	return f.entSubscriptionOrder.ID
}

func (f *subscriptionOrder) SubscriptionOrderStateID() uint32 {
	return f.entSubscriptionOrderState.ID
}

func (f *subscriptionOrder) PaymentState() types.PaymentState {
	return types.PaymentState(types.PaymentState_value[f.entSubscriptionOrderState.PaymentState])
}

func (f *subscriptionOrder) PaymentType() types.PaymentType {
	return types.PaymentType(types.PaymentType_value[f.entOrderStateBase.PaymentType])
}

func (f *subscriptionOrder) OrderState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entOrderStateBase.OrderState])
}

func (f *subscriptionOrder) CancelState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entSubscriptionOrderState.CancelState])
}

func (f *subscriptionOrder) UserSetCanceled() bool {
	return f.entSubscriptionOrderState.UserSetCanceled
}

func (f *subscriptionOrder) AdminSetCanceled() bool {
	return f.entSubscriptionOrderState.AdminSetCanceled
}

func (f *subscriptionOrder) PaymentAmountUSD() decimal.Decimal {
	return f.entSubscriptionOrder.PaymentAmountUsd
}

func (f *subscriptionOrder) DiscountAmountUSD() decimal.Decimal {
	return f.entSubscriptionOrder.DiscountAmountUsd
}

func (f *subscriptionOrder) PayWithMeOrderIDs() []uuid.UUID {
	return f.payWithMeOrderIDs
}
