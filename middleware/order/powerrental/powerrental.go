package powerrental

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PowerRental interface {
	OrderID() uuid.UUID
	PaymentID() uuid.UUID
	LedgerLockID() uuid.UUID
}

type powerRental struct {
	entOrderBase        *ent.OrderBase
	entOrderStateBase   *ent.OrderStateBase
	entPowerRental      *ent.PowerRental
	entPowerRentalState *ent.PowerRentalState
	entPaymentBase      *ent.PaymentBase
	entPaymentBalances  []*ent.PaymentBalance
	entPaymentTransfers []*ent.PaymentTransfer
	entLedgerLock       *ent.OrderLock
	entStockLock        *ent.OrderLock
	entOrderCoupons     []*ent.OrderCoupon
	payWithMeOrderIDs   []uuid.UUID
	childOrderIDs       []uuid.UUID
	poolOrderUser       *ent.PoolOrderUser
}

func (f *powerRental) OrderID() uuid.UUID {
	return f.entPowerRental.OrderID
}

func (f *powerRental) UserID() uuid.UUID {
	return f.entOrderBase.UserID
}

func (f *powerRental) OrderType() types.OrderType {
	return types.OrderType(types.OrderType_value[f.entOrderBase.OrderType])
}

func (f *powerRental) PaymentID() uuid.UUID {
	return f.entPaymentBase.EntID
}

func (f *powerRental) Exist() bool {
	return f.entOrderBase != nil
}

func (f *powerRental) LedgerLockID() uuid.UUID {
	if f.entLedgerLock == nil {
		return uuid.Nil
	}
	return f.entLedgerLock.EntID
}

func (f *powerRental) OrderBaseID() uint32 {
	return f.entOrderBase.ID
}

func (f *powerRental) OrderStateBaseID() uint32 {
	return f.entOrderStateBase.ID
}

func (f *powerRental) PowerRentalID() uint32 {
	return f.entPowerRental.ID
}

func (f *powerRental) PowerRentalStateID() uint32 {
	return f.entPowerRentalState.ID
}

func (f *powerRental) PoolOrderUserRecordID() *uint32 {
	if f.poolOrderUser == nil {
		return nil
	}
	return &f.poolOrderUser.ID
}

func (f *powerRental) PaymentState() types.PaymentState {
	return types.PaymentState(types.PaymentState_value[f.entPowerRentalState.PaymentState])
}

func (f *powerRental) PaymentType() types.PaymentType {
	return types.PaymentType(types.PaymentType_value[f.entOrderStateBase.PaymentType])
}

func (f *powerRental) OrderState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entOrderStateBase.OrderState])
}

func (f *powerRental) BenefitState() types.BenefitState {
	return types.BenefitState(types.BenefitState_value[f.entOrderStateBase.BenefitState])
}

func (f *powerRental) CancelState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entPowerRentalState.CancelState])
}

func (f *powerRental) RenewState() types.OrderRenewState {
	return types.OrderRenewState(types.OrderRenewState_value[f.entPowerRentalState.RenewState])
}

func (f *powerRental) UserSetCanceled() bool {
	return f.entPowerRentalState.UserSetCanceled
}

func (f *powerRental) AdminSetCanceled() bool {
	return f.entPowerRentalState.AdminSetCanceled
}

func (f *powerRental) PaymentAmountUSD() decimal.Decimal {
	return f.entPowerRental.PaymentAmountUsd
}

func (f *powerRental) DiscountAmountUSD() decimal.Decimal {
	return f.entPowerRental.DiscountAmountUsd
}

func (f *powerRental) PayWithMeOrderIDs() []uuid.UUID {
	return f.payWithMeOrderIDs
}

func (f *powerRental) ChildOrderIDs() []uuid.UUID {
	return f.childOrderIDs
}
