package orderstm

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

var forwards = map[types.OrderState][]types.OrderState{
	types.OrderState_OrderStateCreated: {types.OrderState_OrderStateWaitPayment},
	types.OrderState_OrderStateWaitPayment: {
		types.OrderState_OrderStatePaymentTransferReceived,
		types.OrderState_OrderStatePreCancel,
		types.OrderState_OrderStatePaymentTimeout,
	},
	types.OrderState_OrderStatePaymentTransferReceived:     {types.OrderState_OrderStatePaymentTransferBookKeeping},
	types.OrderState_OrderStatePaymentTransferBookKeeping:  {types.OrderState_OrderStatePaymentSpendBalance},
	types.OrderState_OrderStatePaymentSpendBalance:         {types.OrderState_OrderStateTransferGoodStockLocked},
	types.OrderState_OrderStateTransferGoodStockLocked:     {types.OrderState_OrderStateAchievementBookKeeping},
	types.OrderState_OrderStateAchievementBookKeeping:      {types.OrderState_OrderStateAddCommission},
	types.OrderState_OrderStateAddCommission:               {types.OrderState_OrderStatePaymentUnlockAccount},
	types.OrderState_OrderStatePaymentUnlockAccount:        {types.OrderState_OrderStatePaid},
	types.OrderState_OrderStatePaid:                        {types.OrderState_OrderStateTransferGoodStockWaitStart, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStateTransferGoodStockWaitStart:  {types.OrderState_OrderStateCreateOrderUser},
	types.OrderState_OrderStateCreateOrderUser:             {types.OrderState_OrderStateSetProportion},
	types.OrderState_OrderStateSetProportion:               {types.OrderState_OrderStateSetRevenueAddress},
	types.OrderState_OrderStateSetRevenueAddress:           {types.OrderState_OrderStateInService},
	types.OrderState_OrderStateInService:                   {types.OrderState_OrderStatePreExpired, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStatePreExpired:                  {types.OrderState_OrderStateDeleteProportion},
	types.OrderState_OrderStateDeleteProportion:            {types.OrderState_OrderStateCheckProportion},
	types.OrderState_OrderStateCheckProportion:             {types.OrderState_OrderStateRestoreExpiredStock},
	types.OrderState_OrderStateRestoreExpiredStock:         {types.OrderState_OrderStateCheckPoolBalance},
	types.OrderState_OrderStateCheckPoolBalance:            {types.OrderState_OrderStateExpired},
	types.OrderState_OrderStateExpired:                     {},
	types.OrderState_OrderStatePreCancel:                   {types.OrderState_OrderStateRestoreCanceledStock},
	types.OrderState_OrderStateRestoreCanceledStock:        {types.OrderState_OrderStateDeductLockedCommission},
	types.OrderState_OrderStateDeductLockedCommission:      {types.OrderState_OrderStateCancelAchievement},
	types.OrderState_OrderStateCancelAchievement:           {types.OrderState_OrderStateReturnCanceledBalance},
	types.OrderState_OrderStateReturnCanceledBalance:       {types.OrderState_OrderStateCanceledTransferBookKeeping},
	types.OrderState_OrderStateCanceledTransferBookKeeping: {types.OrderState_OrderStateCancelUnlockPaymentAccount},
	types.OrderState_OrderStateCancelUnlockPaymentAccount:  {types.OrderState_OrderStateCanceled},
	types.OrderState_OrderStateCanceled:                    {},
	types.OrderState_OrderStatePaymentTimeout:              {types.OrderState_OrderStatePreCancel},
}

type forwardHandler struct {
	*orderQueryHandler
}

func (h *forwardHandler) forward() (*types.OrderState, error) {
	states, ok := forwards[h._ent.OrderState()]
	if !ok {
		return nil, wlog.Errorf("invalid orderstate")
	}
	for _, state := range states {
		if state == *h.OrderState {
			return &state, nil
		}
	}
	return nil, wlog.Errorf("invalid orderstate")
}
