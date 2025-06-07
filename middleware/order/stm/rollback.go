package orderstm

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
)

var rollbacks = map[types.OrderState]types.OrderState{
	types.OrderState_OrderStatePaymentSpendBalance:         types.OrderState_OrderStatePaymentTransferBookKeeping,
	types.OrderState_OrderStateTransferGoodStockLocked:     types.OrderState_OrderStatePaymentSpendBalance,
	types.OrderState_OrderStateAchievementBookKeeping:      types.OrderState_OrderStateTransferGoodStockLocked,
	types.OrderState_OrderStateAddCommission:               types.OrderState_OrderStateAchievementBookKeeping,
	types.OrderState_OrderStatePaymentUnlockAccount:        types.OrderState_OrderStateAddCommission,
	types.OrderState_OrderStatePaid:                        types.OrderState_OrderStatePaymentUnlockAccount,
	types.OrderState_OrderStateCreateOrderUser:             types.OrderState_OrderStateTransferGoodStockWaitStart,
	types.OrderState_OrderStateSetProportion:               types.OrderState_OrderStateCreateOrderUser,
	types.OrderState_OrderStateSetRevenueAddress:           types.OrderState_OrderStateSetProportion,
	types.OrderState_OrderStateInService:                   types.OrderState_OrderStateSetRevenueAddress,
	types.OrderState_OrderStateDeleteProportion:            types.OrderState_OrderStatePreExpired,
	types.OrderState_OrderStateCheckProportion:             types.OrderState_OrderStateDeleteProportion,
	types.OrderState_OrderStateRestoreExpiredStock:         types.OrderState_OrderStateCheckProportion,
	types.OrderState_OrderStateCheckPoolBalance:            types.OrderState_OrderStateRestoreExpiredStock,
	types.OrderState_OrderStateExpired:                     types.OrderState_OrderStateCheckPoolBalance,
	types.OrderState_OrderStateDeductLockedCommission:      types.OrderState_OrderStateRestoreCanceledStock,
	types.OrderState_OrderStateCancelAchievement:           types.OrderState_OrderStateDeductLockedCommission,
	types.OrderState_OrderStateReturnCanceledBalance:       types.OrderState_OrderStateCancelAchievement,
	types.OrderState_OrderStateCanceledTransferBookKeeping: types.OrderState_OrderStateReturnCanceledBalance,
	types.OrderState_OrderStateCancelUnlockPaymentAccount:  types.OrderState_OrderStateCanceledTransferBookKeeping,
	types.OrderState_OrderStateCanceled:                    types.OrderState_OrderStateCancelUnlockPaymentAccount,
}

type rollbackHandler struct {
	*orderQueryHandler
}

func (h *rollbackHandler) rollback() (*types.OrderState, error) {
	if h.OrderState == nil || *h.OrderState != h._ent.OrderState() {
		return nil, wlog.Errorf("invalid orderstate")
	}
	state, ok := rollbacks[h._ent.OrderState()]
	if !ok {
		return nil, wlog.Errorf("invalid orderstate")
	}
	return &state, nil
}
