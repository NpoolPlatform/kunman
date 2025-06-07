package feeorder

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FeeOrder interface {
	OrderID() uuid.UUID
	PaymentID() uuid.UUID
	LedgerLockID() uuid.UUID
}

type feeOrder struct {
	entOrderBase        *ent.OrderBase
	entOrderStateBase   *ent.OrderStateBase
	entFeeOrder         *ent.FeeOrder
	entFeeOrderState    *ent.FeeOrderState
	entPaymentBase      *ent.PaymentBase
	entPaymentBalances  []*ent.PaymentBalance
	entPaymentTransfers []*ent.PaymentTransfer
	entLedgerLock       *ent.OrderLock
	entOrderCoupons     []*ent.OrderCoupon
	payWithMeOrderIDs   []uuid.UUID
}

func (f *feeOrder) OrderID() uuid.UUID {
	return f.entFeeOrder.OrderID
}

func (f *feeOrder) UserID() uuid.UUID {
	return f.entOrderBase.UserID
}

func (f *feeOrder) PaymentID() uuid.UUID {
	if f.entPaymentBase == nil {
		return uuid.Nil
	}
	return f.entPaymentBase.EntID
}

func (f *feeOrder) LedgerLockID() uuid.UUID {
	if f.entLedgerLock == nil {
		return uuid.Nil
	}
	return f.entLedgerLock.EntID
}

func (f *feeOrder) Exist() bool {
	return f.entOrderBase != nil
}

func (f *feeOrder) OrderBaseID() uint32 {
	return f.entOrderBase.ID
}

func (f *feeOrder) OrderType() types.OrderType {
	return types.OrderType(types.OrderType_value[f.entOrderBase.OrderType])
}

func (f *feeOrder) OrderStateBaseID() uint32 {
	return f.entOrderStateBase.ID
}

func (f *feeOrder) FeeOrderID() uint32 {
	return f.entFeeOrder.ID
}

func (f *feeOrder) FeeOrderStateID() uint32 {
	return f.entFeeOrderState.ID
}

func (f *feeOrder) PaymentState() types.PaymentState {
	return types.PaymentState(types.PaymentState_value[f.entFeeOrderState.PaymentState])
}

func (f *feeOrder) PaymentType() types.PaymentType {
	return types.PaymentType(types.PaymentType_value[f.entOrderStateBase.PaymentType])
}

func (f *feeOrder) OrderState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entOrderStateBase.OrderState])
}

func (f *feeOrder) CancelState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entFeeOrderState.CancelState])
}

func (f *feeOrder) UserSetCanceled() bool {
	return f.entFeeOrderState.UserSetCanceled
}

func (f *feeOrder) AdminSetCanceled() bool {
	return f.entFeeOrderState.AdminSetCanceled
}

func (f *feeOrder) PaymentAmountUSD() decimal.Decimal {
	return f.entFeeOrder.PaymentAmountUsd
}

func (f *feeOrder) DiscountAmountUSD() decimal.Decimal {
	return f.entFeeOrder.DiscountAmountUsd
}

func (f *feeOrder) PayWithMeOrderIDs() []uuid.UUID {
	return f.payWithMeOrderIDs
}
