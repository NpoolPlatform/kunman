package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	poolorderusermiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental/poolorderuser"
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
	"github.com/NpoolPlatform/kunman/middleware/order/powerrental/poolorderuser"
	powerrentalstate1 "github.com/NpoolPlatform/kunman/middleware/order/powerrental/state"
	orderstm1 "github.com/NpoolPlatform/kunman/middleware/order/stm"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*powerRentalQueryHandler
	paymentChecker *paymentcommon.PaymentCheckHandler

	newPayment             bool
	newPaymentBalance      bool
	oldPoolOrderUser       *poolorderusermiddlewarepb.PoolOrderUser
	obseletePaymentBaseReq *paymentbasecrud.Req
	sqlObseletePaymentBase string

	sqlOrderStateBase     string
	sqlPowerRentalState   string
	sqlPaymentBase        string
	sqlPoolOrderUser      string
	sqlOrderLocks         []string
	sqlPaymentBalanceLock string
	sqlPaymentBalances    []string
	sqlPaymentTransfers   []string

	sqlStatedOrderStateBases   []string
	sqlPayWithMeFeeOrderStates []string

	updateNothing bool
}

func (h *updateHandler) constructOrderStateBaseSQL(ctx context.Context) (err error) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	if h.sqlOrderStateBase, err = handler.ConstructUpdateSQL(); wlog.Equal(err, cruder.ErrUpdateNothing) {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) constructStatedOrderStateBaseSQLs(ctx context.Context, orderIDs []uuid.UUID) error {
	for _, orderID := range orderIDs {
		_orderID := orderID
		handler, _ := orderstatebase1.NewHandler(ctx)
		handler.Req = *h.OrderStateBaseReq
		handler.OrderID = &_orderID
		handler.PaymentType = types.PaymentType_PayWithParentOrder.Enum()
		handler.BenefitState = nil
		handler.LastBenefitAt = nil
		handler.StartAt = nil
		sql, err := handler.ConstructUpdateSQL()
		if err != nil {
			if wlog.Equal(err, cruder.ErrUpdateNothing) {
				continue
			}
			return wlog.WrapError(err)
		}
		h.sqlStatedOrderStateBases = append(h.sqlStatedOrderStateBases, sql)
	}
	return nil
}

func (h *updateHandler) constructPowerRentalStateSQL(ctx context.Context) (err error) {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *h.PowerRentalStateReq
	if h.sqlPowerRentalState, err = handler.ConstructUpdateSQL(); wlog.Equal(err, cruder.ErrUpdateNothing) {
		return nil
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) constructPayWithMeFeeOrderStateSQLs(ctx context.Context) error {
	for _, orderID := range h._ent.PayWithMeOrderIDs() {
		_orderID := orderID
		handler, err := feeorderstate1.NewHandler(
			ctx,
			feeorderstate1.WithOrderID(func() *string { s := orderID.String(); return &s }(), true),
			feeorderstate1.WithCancelState(h.PowerRentalStateReq.CancelState, false),
			feeorderstate1.WithPaidAt(h.PowerRentalStateReq.PaidAt, false),
			feeorderstate1.WithUserSetPaid(h.PowerRentalStateReq.UserSetPaid, false),
			feeorderstate1.WithUserSetCanceled(h.PowerRentalStateReq.UserSetCanceled, false),
			feeorderstate1.WithAdminSetCanceled(h.PowerRentalStateReq.AdminSetCanceled, false),
			feeorderstate1.WithPaymentState(h.PowerRentalStateReq.PaymentState, false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.PaymentID = h.PaymentBaseReq.EntID
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

func (h *updateHandler) constructOrderLockSQLs(ctx context.Context) {
	for _, req := range h.OrderLockReqs {
		handler, _ := orderlock1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderLocks = append(h.sqlOrderLocks, handler.ConstructCreateSQL())
	}
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
	if h.sqlObseletePaymentBase, err = handler.ConstructUpdateSQL(); wlog.Equal(err, cruder.ErrUpdateNothing) {
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
			if wlog.Equal(err, cruder.ErrUpdateNothing) {
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

func (h *updateHandler) constructPoolOrderUserSQL(ctx context.Context) error {
	if h.PoolOrderUserReq.PoolOrderUserID == nil {
		return nil
	}
	handler, err := poolorderuser.NewHandler(ctx, poolorderuser.WithReq(h.PoolOrderUserReq, true))
	if err != nil {
		return wlog.WrapError(err)
	}

	if h.oldPoolOrderUser != nil {
		handler.ID = &h.oldPoolOrderUser.ID
		h.sqlPoolOrderUser, err = handler.ConstructUpdateSQL()
	} else {
		handler.EntID = func() *uuid.UUID { id := uuid.New(); return &id }()
		h.sqlPoolOrderUser = handler.ConstructCreateSQL()
	}
	return wlog.WrapError(err)
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) (int64, error) {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return n, wlog.WrapError(err)
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

func (h *updateHandler) updateStatedOrderStateBases(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlStatedOrderStateBases {
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

func (h *updateHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	if h.sqlPowerRentalState == "" {
		return nil
	}
	n, err := h.execSQL(ctx, tx, h.sqlPowerRentalState)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n == 1 {
		h.updateNothing = false
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

func (h *updateHandler) createOrderLocks(ctx context.Context, tx *ent.Tx) error {
	if !h.newPayment {
		return nil
	}
	for _, sql := range h.sqlOrderLocks {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n != 1 {
			return wlog.Errorf("fail create orderlock")
		}
		h.updateNothing = false
	}
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

func (h *updateHandler) upsertPoolOrderUser(ctx context.Context, tx *ent.Tx) (int64, error) {
	if h.PoolOrderUserReq.PoolOrderUserID == nil {
		return 0, nil
	}
	return h.execSQL(ctx, tx, h.sqlPoolOrderUser)
}

func (h *updateHandler) createPaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		n, err := h.execSQL(ctx, tx, sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		if n != 1 {
			return wlog.Errorf("fail create orderlock")
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
		if n != 1 {
			return wlog.Errorf("fail create orderlock")
		}
		h.updateNothing = false
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
	h.PowerRentalStateReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
	h.PoolOrderUserReq.OrderID = h.OrderID
}

func (h *updateHandler) formalizeOrderLocks() {
	for _, req := range h.OrderLockReqs {
		req.OrderID = h.OrderID
		req.UserID = func() *uuid.UUID { uid := h._ent.UserID(); return &uid }()
	}
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
	ledgerLockID := h.ledgerLockID()
	h.newPaymentBalance = ledgerLockID != nil
	if h.newPaymentBalance && *ledgerLockID == h._ent.LedgerLockID() {
		return wlog.Errorf("invalid ledgerlock")
	}

	h.obseletePaymentBaseReq.EntID = func() *uuid.UUID { uid := h._ent.PaymentID(); return &uid }()
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	h.PowerRentalStateReq.PaymentID = h.PaymentBaseReq.EntID
	h.PaymentBalanceLockReq.PaymentID = h.PaymentBaseReq.EntID
	return nil
}

func (h *updateHandler) validatePoolOrderUser(ctx context.Context) error {
	if h.PoolOrderUserReq.PoolOrderUserID == nil {
		return nil
	}
	handler, err := poolorderuser.NewHandler(
		ctx,
		poolorderuser.WithConds(&poolorderusermiddlewarepb.Conds{
			OrderID: &v1.StringVal{
				Op:    cruder.EQ,
				Value: h.PoolOrderUserReq.OrderID.String(),
			},
		}),
		poolorderuser.WithLimit(1),
		poolorderuser.WithOffset(0),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	infos, _, err := handler.GetPoolOrderUsers(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(infos) > 0 {
		h.oldPoolOrderUser = &infos[0]
	}

	return nil
}

func (h *updateHandler) validatePaymentState() error {
	if h.PowerRentalStateReq.PaymentState == nil {
		return nil
	}
	if h._ent.PaymentState() != types.PaymentState_PaymentStateWait {
		return wlog.Errorf("permission denied")
	}
	// If you need to update payment state, you should put it in the last state which cannot be rollback
	if h.OrderStateBaseReq.OrderState == nil || (h.Rollback != nil && *h.Rollback) {
		return wlog.Errorf("permission denied")
	}
	switch *h.OrderStateBaseReq.OrderState {
	case types.OrderState_OrderStatePaid:
		if *h.PowerRentalStateReq.PaymentState != types.PaymentState_PaymentStateDone {
			return wlog.Errorf("permission denied")
		}
	case types.OrderState_OrderStatePaymentTimeout:
		if *h.PowerRentalStateReq.PaymentState != types.PaymentState_PaymentStateTimeout {
			return wlog.Errorf("permission denied")
		}
	case types.OrderState_OrderStatePreCancel:
		if *h.PowerRentalStateReq.PaymentState != types.PaymentState_PaymentStateCanceled {
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
		if h.ledgerLockID() == nil {
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
		if h.PaymentBaseReq.EntID != nil || h.ledgerLockID() != nil {
			return wlog.Errorf("invalid paymenttype")
		}
	}
	return nil
}

func (h *updateHandler) validatePayment() error {
	if !h.paymentChecker.Payable() {
		if len(h.PaymentBalanceReqs) > 0 || len(h.PaymentTransferReqs) > 0 {
			return wlog.Errorf("invalid payment")
		}
		return nil
	}
	return h.paymentChecker.ValidatePayment()
}

func (h *updateHandler) formalizeEntIDs() {
	if h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *updateHandler) validateCancelState() error {
	if h.PowerRentalStateReq.CancelState == nil {
		return nil
	}
	if h._ent.CancelState() != types.OrderState_DefaultOrderState {
		return wlog.Errorf("invalid cancelstate")
	}
	h.PowerRentalStateReq.CanceledAt = func() *uint32 { u := uint32(time.Now().Unix()); return &u }()
	return nil
}

func (h *updateHandler) validateBenefitState() error {
	if h.OrderStateBaseReq.BenefitState == nil {
		return nil
	}
	if h._ent.OrderState() != types.OrderState_OrderStateInService {
		return wlog.Errorf("permission denied")
	}
	switch h._ent.BenefitState() {
	case types.BenefitState_BenefitWait:
		if *h.OrderStateBaseReq.BenefitState != types.BenefitState_BenefitCalculated {
			return wlog.Errorf("permission denied")
		}
	case types.BenefitState_BenefitCalculated:
		switch *h.OrderStateBaseReq.BenefitState {
		case types.BenefitState_BenefitBookKept:
		case types.BenefitState_BenefitWait:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.BenefitState_BenefitBookKept:
		if *h.OrderStateBaseReq.BenefitState != types.BenefitState_BenefitWait {
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

func (h *updateHandler) validateRenewState() error { //nolint:gocyclo
	if h.PowerRentalStateReq.RenewState == nil {
		return nil
	}
	if h._ent.OrderState() != types.OrderState_OrderStateInService {
		return wlog.Errorf("permission denied")
	}
	switch h._ent.RenewState() {
	case types.OrderRenewState_OrderRenewWait:
		switch *h.PowerRentalStateReq.RenewState {
		case types.OrderRenewState_OrderRenewCheck:
		case types.OrderRenewState_OrderRenewFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.OrderRenewState_OrderRenewCheck:
		switch *h.PowerRentalStateReq.RenewState {
		case types.OrderRenewState_OrderRenewWait:
		case types.OrderRenewState_OrderRenewNotify:
		case types.OrderRenewState_OrderRenewFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.OrderRenewState_OrderRenewNotify:
		switch *h.PowerRentalStateReq.RenewState {
		case types.OrderRenewState_OrderRenewWait:
		case types.OrderRenewState_OrderRenewExecute:
		case types.OrderRenewState_OrderRenewFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.OrderRenewState_OrderRenewExecute:
		fallthrough //nolint
	case types.OrderRenewState_OrderRenewFail:
		if *h.PowerRentalStateReq.RenewState != types.OrderRenewState_OrderRenewWait {
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

func (h *updateHandler) formalizeCancelState() {
	if h.OrderStateBaseReq.OrderState != nil && *h.OrderStateBaseReq.OrderState == types.OrderState_OrderStatePreCancel {
		h.PowerRentalStateReq.CancelState = func() *types.OrderState { e := h._ent.OrderState(); return &e }()
	}
}

func (h *updateHandler) formalizePaidAt() {
	if h.PowerRentalStateReq.PaymentState != nil && *h.PowerRentalStateReq.PaymentState == types.PaymentState_PaymentStateDone {
		h.PowerRentalStateReq.PaidAt = func() *uint32 { u := uint32(time.Now().Unix()); return &u }()
	}
}

func (h *updateHandler) validateUpdate(ctx context.Context, tx *ent.Tx) error {
	handler, err := orderstm1.NewHandler(
		ctx,
		orderstm1.WithOrderID(h.OrderID, true),
		orderstm1.WithOrderState(h.OrderStateBaseReq.OrderState, false),
		orderstm1.WithCurrentPaymentState(func() *types.PaymentState { e := h._ent.PaymentState(); return &e }(), true),
		orderstm1.WithNewPaymentState(h.PowerRentalStateReq.PaymentState, false),
		orderstm1.WithUserSetPaid(h.PowerRentalStateReq.UserSetPaid, false),
		orderstm1.WithUserSetCanceled(h.PowerRentalStateReq.UserSetCanceled, false),
		orderstm1.WithUserCanceled(func() *bool { b := h._ent.UserSetCanceled(); return &b }(), false),
		orderstm1.WithAdminSetCanceled(h.PowerRentalStateReq.AdminSetCanceled, false),
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

//nolint:gocyclo,funlen
func (h *Handler) UpdatePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &updateHandler{
		powerRentalQueryHandler: &powerRentalQueryHandler{
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

	if err := handler.requirePowerRentalWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	handler.paymentChecker.PaymentAmountUSD = func() *decimal.Decimal { d := handler._ent.PaymentAmountUSD(); return &d }()
	handler.paymentChecker.DiscountAmountUSD = func() *decimal.Decimal { d := handler._ent.DiscountAmountUSD(); return &d }()
	if handler.paymentChecker.PaymentType == nil {
		handler.paymentChecker.PaymentType = func() *types.PaymentType { e := handler._ent.PaymentType(); return &e }()
	}
	handler.formalizeOrderID()
	if err := handler.validateUpdate(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateBenefitState(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateRenewState(); err != nil {
		return wlog.WrapError(err)
	}

	handler.formalizeOrderLocks()
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
		if err := handler.validatePayment(); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.validatePaymentType(); err != nil {
			return wlog.WrapError(err)
		}
	}
	if err := handler.validatePaymentState(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizeCancelState()
	if err := handler.validateCancelState(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaidAt()

	if err := handler.validatePoolOrderUser(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.constructOrderStateBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	statedOrderIDs := handler._ent.PayWithMeOrderIDs()
	if h.OrderStateBaseReq.OrderState != nil && *h.OrderStateBaseReq.OrderState == types.OrderState_OrderStateExpired {
		statedOrderIDs = handler._ent.ChildOrderIDs()
	}
	if err := handler.constructStatedOrderStateBaseSQLs(ctx, statedOrderIDs); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPowerRentalStateSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPayWithMeFeeOrderStateSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	handler.constructOrderLockSQLs(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	if err := handler.constructPoolOrderUserSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPaymentTransferSQLs(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructObseletePaymentBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if _, err := handler.upsertPoolOrderUser(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateStatedOrderStateBases(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePowerRentalState(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePayWithMeFeeOrderStates(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updateObseletePaymentBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrderLocks(ctx, tx); err != nil {
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
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	return wlog.WrapError(handler.FeeMultiHandler.UpdateFeeOrdersWithTx(ctx, tx))
}

func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	return wlog.WrapError(db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return wlog.WrapError(h.UpdatePowerRentalWithTx(_ctx, tx))
	}))
}
