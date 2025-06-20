package feeorder

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	feeorderstate1 "github.com/NpoolPlatform/kunman/middleware/order/fee/state"
	orderlock1 "github.com/NpoolPlatform/kunman/middleware/order/order/lock"
	orderstatebase1 "github.com/NpoolPlatform/kunman/middleware/order/order/statebase"
	paymentbase1 "github.com/NpoolPlatform/kunman/middleware/order/payment"
	paymentbalance1 "github.com/NpoolPlatform/kunman/middleware/order/payment/balance"
	paymentbalancelock1 "github.com/NpoolPlatform/kunman/middleware/order/payment/balance/lock"
	paymentcommon "github.com/NpoolPlatform/kunman/middleware/order/payment/common"
	paymenttransfer1 "github.com/NpoolPlatform/kunman/middleware/order/payment/transfer"
	orderstm1 "github.com/NpoolPlatform/kunman/middleware/order/stm"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*feeOrderQueryHandler
	paymentChecker *paymentcommon.PaymentCheckHandler

	newPayment             bool
	newPaymentBalance      bool
	obseletePaymentBaseReq *paymentbasecrud.Req
	sqlObseletePaymentBase string

	sqlOrderStateBase     string
	sqlFeeOrderState      string
	sqlPaymentBase        string
	sqlLedgerLock         string
	sqlPaymentBalanceLock string
	sqlPaymentBalances    []string
	sqlPaymentTransfers   []string

	sqlPayWithMeOrderStateBases []string
	sqlPayWithMeFeeOrderStates  []string

	updateNothing bool
}

func (h *updateHandler) constructOrderStateBaseSQL(ctx context.Context) (err error) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	handler.StartMode = func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }()
	if h.sqlOrderStateBase, err = handler.ConstructUpdateSQL(); err != nil && wlog.Equal(err, cruder.ErrUpdateNothing) {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) constructPayWithMeOrderStateBaseSQLs(ctx context.Context) error {
	for _, orderID := range h._ent.PayWithMeOrderIDs() {
		_orderID := orderID
		handler, _ := orderstatebase1.NewHandler(ctx)
		handler.Req = *h.OrderStateBaseReq
		handler.OrderID = &_orderID
		handler.PaymentType = types.PaymentType_PayWithOtherOrder.Enum()
		sql, err := handler.ConstructUpdateSQL()
		if err != nil {
			if wlog.Equal(err, cruder.ErrUpdateNothing) {
				continue
			}
			return wlog.WrapError(err)
		}
		h.sqlPayWithMeOrderStateBases = append(h.sqlPayWithMeOrderStateBases, sql)
	}
	return nil
}

func (h *updateHandler) constructPayWithMeFeeOrderStateSQLs(ctx context.Context) error {
	for _, orderID := range h._ent.PayWithMeOrderIDs() {
		_orderID := orderID
		handler, _ := feeorderstate1.NewHandler(ctx)
		handler.Req = *h.FeeOrderStateReq
		handler.OrderID = &_orderID
		sql, err := handler.ConstructUpdateSQL()
		if err != nil {
			if wlog.Equal(err, cruder.ErrUpdateNothing) {
				continue
			}
			return wlog.WrapError(err)
		}
		h.sqlPayWithMeFeeOrderStates = append(h.sqlPayWithMeFeeOrderStates, sql)
	}
	return nil
}

func (h *updateHandler) constructFeeOrderStateSQL(ctx context.Context) (err error) {
	handler, _ := feeorderstate1.NewHandler(ctx)
	handler.Req = *h.FeeOrderStateReq
	if h.sqlFeeOrderState, err = handler.ConstructUpdateSQL(); err != nil && wlog.Equal(err, cruder.ErrUpdateNothing) {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) constructLedgerLockSQL(ctx context.Context) {
	if !h.newPaymentBalance {
		return
	}
	handler, _ := orderlock1.NewHandler(ctx)
	handler.Req = *h.LedgerLockReq
	h.sqlLedgerLock = handler.ConstructCreateSQL()
}

func (h *updateHandler) constructPaymentBalanceLockSQL(ctx context.Context) {
	if !h.newPaymentBalance {
		return
	}
	handler, _ := paymentbalancelock1.NewHandler(ctx)
	handler.Req = *h.PaymentBalanceLockReq
	h.sqlPaymentBalanceLock = handler.ConstructCreateSQL()
}

func (h *updateHandler) constructPaymentBaseSQL(ctx context.Context) {
	if !h.newPayment {
		return
	}
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.PaymentBaseReq
	h.sqlPaymentBase = handler.ConstructCreateSQL()
}

func (h *updateHandler) constructObseletePaymentBaseSQL(ctx context.Context) (err error) {
	if !h.newPayment {
		return
	}
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.obseletePaymentBaseReq
	if h.sqlObseletePaymentBase, err = handler.ConstructUpdateSQL(); err != nil && wlog.Equal(err, cruder.ErrUpdateNothing) {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) constructPaymentBalanceSQLs(ctx context.Context) {
	if !h.newPaymentBalance {
		return
	}
	for _, req := range h.PaymentBalanceReqs {
		handler, _ := paymentbalance1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentBalances = append(h.sqlPaymentBalances, handler.ConstructCreateSQL())
	}
}

func (h *updateHandler) constructPaymentTransferSQLs(ctx context.Context) error {
	for _, req := range h.PaymentTransferReqs {
		handler, _ := paymenttransfer1.NewHandler(ctx)
		handler.Req = *req
		if h.newPayment {
			h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, handler.ConstructCreateSQL())
		} else {
			sql, err := handler.ConstructUpdateSQL()
			if err == cruder.ErrUpdateNothing {
				continue
			}
			if err != nil {
				return wlog.WrapError(err)
			}
			h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, sql)
		}
	}
	return nil
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) (int64, error) {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return n, nil
}

func (h *updateHandler) updateOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlOrderStateBase == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlOrderStateBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n == 1 {
		h.updateNothing = false
	}
	return nil
}

func (h *updateHandler) updatePayWithMeOrderStateBases(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPayWithMeOrderStateBases {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n == 1 {
			h.updateNothing = false
		}
	}
	return nil
}

func (h *updateHandler) updatePayWithMeFeeOrderStates(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPayWithMeFeeOrderStates {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n == 1 {
			h.updateNothing = false
		}
	}
	return nil
}

func (h *updateHandler) updateFeeOrderState(ctx context.Context, tx *ent.Tx) error {
	if h.sqlFeeOrderState == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlFeeOrderState)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n == 1 {
		h.updateNothing = false
	}
	return nil
}

func (h *updateHandler) createLedgerLock(ctx context.Context, tx *ent.Tx) error {
	if !h.newPaymentBalance {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlLedgerLock)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n != 1 {
		return wlog.Errorf("fail create ledgerlock")
	}
	h.updateNothing = false
	return nil
}

func (h *updateHandler) createPaymentBalanceLock(ctx context.Context, tx *ent.Tx) error {
	if !h.newPaymentBalance {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlPaymentBalanceLock)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n != 1 {
		return wlog.Errorf("fail create paymentbalancelock")
	}
	h.updateNothing = false
	return nil
}

func (h *updateHandler) createPaymentBase(ctx context.Context, tx *ent.Tx) error {
	if !h.newPayment {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlPaymentBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n != 1 {
		return wlog.Errorf("fail create paymentbase")
	}
	h.updateNothing = false
	return nil
}

func (h *updateHandler) updateObseletePaymentBase(ctx context.Context, tx *ent.Tx) error {
	if !h.newPayment {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlObseletePaymentBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n == 1 {
		h.updateNothing = false
	}
	return nil
}

func (h *updateHandler) createPaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n != 1 {
			return wlog.Errorf("fail create paymentbalance")
		}
		h.updateNothing = false
	}
	return nil
}

func (h *updateHandler) createOrUpdatePaymentTransfers(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentTransfers {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n == 1 {
			h.updateNothing = false
			continue
		}
		if h.newPayment {
			return wlog.Errorf("fail create paymenttransfer")
		}
	}
	return nil
}

func (h *updateHandler) formalizeOrderID() {
	if h.OrderID != nil {
		return
	}
	h.OrderID = func() *uuid.UUID { uid := h._ent.OrderID(); return &uid }()
	h.OrderBaseReq.EntID = h.OrderID
	h.OrderStateBaseReq.OrderID = h.OrderID
	h.FeeOrderStateReq.OrderID = h.OrderID
	h.LedgerLockReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
}

func (h *updateHandler) formalizeUserID() {
	h.LedgerLockReq.UserID = func() *uuid.UUID { uid := h._ent.UserID(); return &uid }()
}

func (h *updateHandler) formalizePaymentBalances() {
	if !h.newPaymentBalance {
		return
	}
	for _, req := range h.PaymentBalanceReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentTransfers() {
	if !h.newPayment {
		return
	}
	for _, req := range h.PaymentTransferReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentType() error {
	switch h._ent.OrderType() {
	case types.OrderType_Offline:
		fallthrough //nolint
	case types.OrderType_Airdrop:
		return wlog.Errorf("permission denied")
	}
	switch h._ent.PaymentType() {
	case types.PaymentType_PayWithBalanceOnly:
	case types.PaymentType_PayWithTransferOnly:
	case types.PaymentType_PayWithTransferAndBalance:
	default:
		return wlog.Errorf("permission denied")
	}
	if len(h.PaymentBalanceReqs) > 0 && len(h.PaymentTransferReqs) > 0 {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithTransferAndBalance; return &e }()
		return nil
	}
	if len(h.PaymentBalanceReqs) > 0 {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithBalanceOnly; return &e }()
		return nil
	}
	if len(h.PaymentTransferReqs) > 0 {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithTransferOnly; return &e }()
		return nil
	}
	return nil
}

func (h *updateHandler) formalizePaymentID() error {
	if (h.PaymentBaseReq.EntID == nil || h._ent.PaymentID() == *h.PaymentBaseReq.EntID) &&
		(h.OrderStateBaseReq.PaymentType == nil || h._ent.PaymentType() == *h.OrderStateBaseReq.PaymentType) {
		return nil
	}

	h.newPayment = true
	h.newPaymentBalance = h.LedgerLockReq.EntID != nil

	if h.newPaymentBalance && *h.LedgerLockReq.EntID == h._ent.LedgerLockID() {
		return wlog.Errorf("invalid ledgerlock")
	}

	h.obseletePaymentBaseReq.EntID = func() *uuid.UUID { uid := h._ent.PaymentID(); return &uid }()
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	h.FeeOrderStateReq.PaymentID = h.PaymentBaseReq.EntID
	h.PaymentBalanceLockReq.PaymentID = h.PaymentBaseReq.EntID
	return nil
}

func (h *updateHandler) formalizeEntIDs() {
	if h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *updateHandler) validateCancelState() error {
	if h.FeeOrderStateReq.CancelState == nil {
		return nil
	}
	if h._ent.CancelState() != types.OrderState_DefaultOrderState {
		return wlog.Errorf("invalid cancelstate")
	}
	h.FeeOrderStateReq.CanceledAt = func() *uint32 { u := uint32(time.Now().Unix()); return &u }()
	return nil
}

func (h *updateHandler) validateUserSetPaid() error {
	if h.FeeOrderStateReq.UserSetPaid != nil && *h.FeeOrderStateReq.UserSetPaid {
		switch h._ent.PaymentType() {
		case types.PaymentType_PayWithBalanceOnly:
			fallthrough //nolint
		case types.PaymentType_PayWithTransferAndBalance:
			fallthrough //nolint
		case types.PaymentType_PayWithTransferOnly:
		default:
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

func (h *updateHandler) formalizeCancelState() error {
	if (h.FeeOrderStateReq.UserSetCanceled != nil && *h.FeeOrderStateReq.UserSetCanceled) ||
		(h.FeeOrderStateReq.AdminSetCanceled != nil && *h.FeeOrderStateReq.AdminSetCanceled) {
		switch h._ent.PaymentType() {
		case types.PaymentType_PayWithBalanceOnly:
			fallthrough //nolint
		case types.PaymentType_PayWithTransferAndBalance:
			fallthrough //nolint
		case types.PaymentType_PayWithTransferOnly:
			fallthrough //nolint
		case types.PaymentType_PayWithOffline:
			fallthrough //nolint
		case types.PaymentType_PayWithNoPayment:
		default:
			return wlog.Errorf("permission denied")
		}
	}
	if h.OrderStateBaseReq.OrderState != nil && *h.OrderStateBaseReq.OrderState == types.OrderState_OrderStatePreCancel {
		h.FeeOrderStateReq.CancelState = func() *types.OrderState { e := h._ent.OrderState(); return &e }()
	}
	return nil
}

func (h *updateHandler) formalizePaidAt() {
	if h.FeeOrderStateReq.PaymentState != nil && *h.FeeOrderStateReq.PaymentState == types.PaymentState_PaymentStateDone {
		h.FeeOrderStateReq.PaidAt = func() *uint32 { u := uint32(time.Now().Unix()); return &u }()
	}
}

func (h *updateHandler) validateUpdate(ctx context.Context, tx *ent.Tx) error {
	handler, err := orderstm1.NewHandler(
		ctx,
		orderstm1.WithOrderID(h.OrderID, true),
		orderstm1.WithOrderState(h.OrderStateBaseReq.OrderState, false),
		orderstm1.WithCurrentPaymentState(func() *types.PaymentState { e := h._ent.PaymentState(); return &e }(), true),
		orderstm1.WithNewPaymentState(h.FeeOrderStateReq.PaymentState, false),
		orderstm1.WithUserSetPaid(h.FeeOrderStateReq.UserSetPaid, false),
		orderstm1.WithUserSetCanceled(h.FeeOrderStateReq.UserSetCanceled, false),
		orderstm1.WithUserCanceled(func() *bool { b := h._ent.UserSetCanceled(); return &b }(), false),
		orderstm1.WithAdminSetCanceled(h.FeeOrderStateReq.AdminSetCanceled, false),
		orderstm1.WithAdminCanceled(func() *bool { b := h._ent.AdminSetCanceled(); return &b }(), false),
		orderstm1.WithRollback(h.Rollback, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	state, err := handler.ValidateUpdateForNewState(ctx, tx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.OrderStateBaseReq.OrderState = state
	return nil
}

func (h *updateHandler) validatePaymentState() error {
	if h.FeeOrderStateReq.PaymentState == nil {
		return nil
	}
	if h._ent.PaymentType() == types.PaymentType_PayWithOtherOrder {
		return wlog.Errorf("permission denied")
	}
	if h._ent.PaymentState() != types.PaymentState_PaymentStateWait {
		return wlog.Errorf("permission denied")
	}
	if h.OrderStateBaseReq.OrderState == nil || (h.Rollback != nil && *h.Rollback) {
		return wlog.Errorf("permission denied")
	}
	switch *h.OrderStateBaseReq.OrderState {
	case types.OrderState_OrderStatePaid:
		if *h.FeeOrderStateReq.PaymentState != types.PaymentState_PaymentStateDone {
			return wlog.Errorf("permission denied")
		}
	case types.OrderState_OrderStatePaymentTimeout:
		if *h.FeeOrderStateReq.PaymentState != types.PaymentState_PaymentStateTimeout {
			return wlog.Errorf("permission denied")
		}
	case types.OrderState_OrderStatePreCancel:
		if *h.FeeOrderStateReq.PaymentState != types.PaymentState_PaymentStateCanceled {
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

//nolint:gocyclo
func (h *updateHandler) validatePaymentType() error {
	switch h._ent.OrderState() {
	case types.OrderState_OrderStateCreated:
	case types.OrderState_OrderStateWaitPayment:
	default:
		return wlog.Errorf("permission denied")
	}
	if h._ent.PaymentState() != types.PaymentState_PaymentStateWait {
		return wlog.Errorf("permission denied")
	}
	paymentType := h._ent.PaymentType()
	if h.OrderStateBaseReq.PaymentType != nil {
		paymentType = *h.OrderStateBaseReq.PaymentType
	}
	switch paymentType {
	case types.PaymentType_PayWithBalanceOnly:
		fallthrough //nolint
	case types.PaymentType_PayWithTransferAndBalance:
		if h.LedgerLockReq.EntID == nil {
			return wlog.Errorf("invalid ledgerlockid")
		}
		fallthrough
	case types.PaymentType_PayWithTransferOnly:
		if h.PaymentBaseReq.EntID == nil {
			return wlog.Errorf("invalid paymentid")
		}
	case types.PaymentType_PayWithParentOrder:
		fallthrough //nolint
	case types.PaymentType_PayWithContract:
		fallthrough //nolint
	case types.PaymentType_PayWithOffline:
		fallthrough //nolint
	case types.PaymentType_PayWithNoPayment:
		if h.PaymentBaseReq.EntID != nil || h.LedgerLockReq.EntID != nil {
			return wlog.Errorf("invalid paymenttype")
		}
	}
	return nil
}

//nolint:funlen,gocyclo
func (h *Handler) UpdateFeeOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &updateHandler{
		feeOrderQueryHandler: &feeOrderQueryHandler{
			Handler: h,
		},
		obseletePaymentBaseReq: &paymentbasecrud.Req{
			OrderID:       h.OrderID,
			ObseleteState: func() *types.PaymentObseleteState { e := types.PaymentObseleteState_PaymentObseleteWait; return &e }(),
		},
		updateNothing: true,
		paymentChecker: &paymentcommon.PaymentCheckHandler{
			PaymentType:         h.OrderStateBaseReq.PaymentType,
			PaymentBalanceReqs:  h.PaymentBalanceReqs,
			PaymentTransferReqs: h.PaymentTransferReqs,
			PaymentAmountUSD:    h.PaymentAmountUSD,
			DiscountAmountUSD:   h.DiscountAmountUSD,
		},
	}

	if err := handler.requireFeeOrderWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if h.MainOrder != nil && *h.MainOrder {
		switch handler._ent.PaymentType() {
		case types.PaymentType_PayWithBalanceOnly:
		case types.PaymentType_PayWithTransferOnly:
		case types.PaymentType_PayWithTransferAndBalance:
		case types.PaymentType_PayWithOffline:
		case types.PaymentType_PayWithNoPayment:
		default:
			return wlog.Errorf("permission denied")
		}
	}
	handler.paymentChecker.PaymentAmountUSD = func() *decimal.Decimal { d := handler._ent.PaymentAmountUSD(); return &d }()
	handler.paymentChecker.DiscountAmountUSD = func() *decimal.Decimal { d := handler._ent.DiscountAmountUSD(); return &d }()
	if handler.paymentChecker.PaymentType == nil {
		handler.paymentChecker.PaymentType = func() *types.PaymentType { e := handler._ent.PaymentType(); return &e }()
	}

	handler.formalizeOrderID()
	handler.formalizeUserID()
	if err := handler.validateUpdate(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}

	handler.formalizeEntIDs()
	if err := handler.formalizePaymentID(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaymentBalances()
	handler.formalizePaymentTransfers()
	if handler.newPayment {
		if err := handler.formalizePaymentType(); err != nil {
			return wlog.WrapError(err)
		}
		if h.OrderStateBaseReq.PaymentType != nil {
			handler.paymentChecker.PaymentType = h.OrderStateBaseReq.PaymentType
		}
		if err := handler.paymentChecker.ValidatePayment(); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.validatePaymentType(); err != nil {
			return wlog.WrapError(err)
		}
	}
	if err := handler.validatePaymentState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateUserSetPaid(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.formalizeCancelState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateCancelState(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaidAt()

	if err := handler.constructOrderStateBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPayWithMeOrderStateBaseSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPayWithMeFeeOrderStateSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructFeeOrderStateSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	handler.constructLedgerLockSQL(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	if err := handler.constructPaymentTransferSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructObseletePaymentBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.updateOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePayWithMeOrderStateBases(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePayWithMeFeeOrderStates(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateFeeOrderState(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateObseletePaymentBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createLedgerLock(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBalanceLock(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBalances(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrUpdatePaymentTransfers(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if handler.updateNothing {
		return cruder.ErrUpdateNothing
	}
	return nil
}

func (h *Handler) UpdateFeeOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateFeeOrderWithTx(_ctx, tx)
	})
}
