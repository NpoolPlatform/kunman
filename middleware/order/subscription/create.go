package subscriptionorder

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	ordercoupon1 "github.com/NpoolPlatform/kunman/middleware/order/order/coupon"
	orderlock1 "github.com/NpoolPlatform/kunman/middleware/order/order/lock"
	orderbase1 "github.com/NpoolPlatform/kunman/middleware/order/order/orderbase"
	orderstatebase1 "github.com/NpoolPlatform/kunman/middleware/order/order/statebase"
	paymentbase1 "github.com/NpoolPlatform/kunman/middleware/order/payment"
	paymentbalance1 "github.com/NpoolPlatform/kunman/middleware/order/payment/balance"
	paymentbalancelock1 "github.com/NpoolPlatform/kunman/middleware/order/payment/balance/lock"
	paymentcommon "github.com/NpoolPlatform/kunman/middleware/order/payment/common"
	paymentfiat1 "github.com/NpoolPlatform/kunman/middleware/order/payment/fiat"
	paymenttransfer1 "github.com/NpoolPlatform/kunman/middleware/order/payment/transfer"
	subscriptionorderstate1 "github.com/NpoolPlatform/kunman/middleware/order/subscription/state"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	paymentChecker            *paymentcommon.PaymentCheckHandler
	payWithBalance            bool
	sql                       string
	sqlOrderBase              string
	sqlOrderStateBase         string
	sqlSubscriptionOrderState string
	sqlLedgerLock             string
	sqlPaymentBalanceLock     string
	sqlOrderCoupons           []string
	sqlPaymentBase            string
	sqlPaymentBalances        []string
	sqlPaymentTransfers       []string
	sqlPaymentFiats           []string
}

func (h *createHandler) constructSQL() {
	h.sql = h.ConstructCreateSQL()
}

func (h *createHandler) constructOrderBaseSQL(ctx context.Context) {
	handler, _ := orderbase1.NewHandler(ctx)
	handler.Req = *h.OrderBaseReq
	h.sqlOrderBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructOrderStateBaseSQL(ctx context.Context) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	handler.StartMode = func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }()
	h.sqlOrderStateBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructSubscriptionOrderStateSQL(ctx context.Context) {
	handler, _ := subscriptionorderstate1.NewHandler(ctx)
	handler.Req = *h.SubscriptionOrderStateReq
	h.sqlSubscriptionOrderState = handler.ConstructCreateSQL()
}

func (h *createHandler) constructLedgerLockSQL(ctx context.Context) {
	if h.LedgerLockReq.EntID == nil {
		return
	}
	handler, _ := orderlock1.NewHandler(ctx)
	handler.Req = *h.LedgerLockReq
	h.sqlLedgerLock = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPaymentBalanceLockSQL(ctx context.Context) {
	if h.PaymentBalanceLockReq.EntID == nil {
		return
	}
	handler, _ := paymentbalancelock1.NewHandler(ctx)
	handler.Req = *h.PaymentBalanceLockReq
	h.sqlPaymentBalanceLock = handler.ConstructCreateSQL()
}

func (h *createHandler) constructOrderCouponSQLs(ctx context.Context) {
	for _, req := range h.OrderCouponReqs {
		handler, _ := ordercoupon1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderCoupons = append(h.sqlOrderCoupons, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) constructPaymentBaseSQL(ctx context.Context) {
	switch *h.OrderStateBaseReq.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
	case types.PaymentType_PayWithTransferOnly:
	case types.PaymentType_PayWithTransferAndBalance:
	case types.PaymentType_PayWithFiatOnly:
	case types.PaymentType_PayWithFiatAndBalance:
	default:
		return
	}
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.PaymentBaseReq
	h.sqlPaymentBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPaymentBalanceSQLs(ctx context.Context) {
	for _, req := range h.PaymentBalanceReqs {
		handler, _ := paymentbalance1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentBalances = append(h.sqlPaymentBalances, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) constructPaymentTransferSQLs(ctx context.Context) {
	for _, req := range h.PaymentTransferReqs {
		handler, _ := paymenttransfer1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) constructPaymentFiatSQLs(ctx context.Context) {
	for _, req := range h.PaymentFiatReqs {
		handler, _ := paymentfiat1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentFiats = append(h.sqlPaymentFiats, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create subscription: %v", err)
	}
	return nil
}

func (h *createHandler) createOrderBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderBase)
}

func (h *createHandler) createOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderStateBase)
}

func (h *createHandler) createSubscriptionOrderState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlSubscriptionOrderState)
}

func (h *createHandler) createLedgerLock(ctx context.Context, tx *ent.Tx) error {
	if !h.payWithBalance {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlLedgerLock)
}

func (h *createHandler) createPaymentBalanceLock(ctx context.Context, tx *ent.Tx) error {
	if !h.payWithBalance {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPaymentBalanceLock)
}

func (h *createHandler) createOrderCoupons(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderCoupons {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createPaymentBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlPaymentBase == "" {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPaymentBase)
}

func (h *createHandler) createPaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createPaymentTransfers(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentTransfers {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createPaymentFiats(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentFiats {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createSubscriptionOrder(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *createHandler) formalizeOrderID() {
	if h.OrderID != nil {
		return
	}
	h.OrderID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.OrderBaseReq.EntID = h.OrderID
	h.OrderStateBaseReq.OrderID = h.OrderID
	h.SubscriptionOrderStateReq.OrderID = h.OrderID
	h.LedgerLockReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
}

func (h *createHandler) formalizeUserID() {
	h.LedgerLockReq.UserID = h.OrderBaseReq.UserID
}

func (h *createHandler) formalizeEntIDs() {
	if h.OrderStateBaseReq.EntID == nil {
		h.OrderStateBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.SubscriptionOrderStateReq.EntID == nil {
		h.SubscriptionOrderStateReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *createHandler) formalizeOrderCoupons() {
	for _, req := range h.OrderCouponReqs {
		req.OrderID = h.OrderID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentBalances() {
	for _, req := range h.PaymentBalanceReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentTransfers() {
	for _, req := range h.PaymentTransferReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentFiats() {
	for _, req := range h.PaymentFiatReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentType() error {
	if *h.OrderBaseReq.OrderType == types.OrderType_Offline {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithOffline; return &e }()
		return nil
	}
	if *h.OrderBaseReq.OrderType == types.OrderType_Airdrop {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithNoPayment; return &e }()
		return nil
	}

	hasTransfer := len(h.PaymentTransferReqs) > 0
	hasBalance := len(h.PaymentBalanceReqs) > 0
	hasFiat := len(h.PaymentFiatReqs) > 0

	newPaymentType := func(ptypes types.PaymentType) *types.PaymentType {
		return &ptypes
	}

	if hasFiat && hasTransfer {
		return wlog.Errorf("invalid payment type: fiat and transfer cannot coexist")
	}

	switch {
	case hasTransfer && hasBalance:
		h.OrderStateBaseReq.PaymentType = newPaymentType(types.PaymentType_PayWithTransferAndBalance)
	case hasFiat && hasBalance:
		h.OrderStateBaseReq.PaymentType = newPaymentType(types.PaymentType_PayWithFiatAndBalance)
	case hasTransfer:
		h.OrderStateBaseReq.PaymentType = newPaymentType(types.PaymentType_PayWithTransferOnly)
	case hasBalance:
		h.OrderStateBaseReq.PaymentType = newPaymentType(types.PaymentType_PayWithBalanceOnly)
	case hasFiat:
		h.OrderStateBaseReq.PaymentType = newPaymentType(types.PaymentType_PayWithFiatOnly)
	default:
		return wlog.Errorf("invalid payment type: no payment method provided")
	}

	return nil
}

func (h *createHandler) formalizePaymentState() {
	if *h.OrderStateBaseReq.PaymentType != types.PaymentType_PayWithNoPayment {
		return
	}
	h.SubscriptionOrderStateReq.PaymentState = types.PaymentState_PaymentStateNoPayment.Enum()
}

func (h *createHandler) formalizePaymentID() {
	if !h.paymentChecker.Payable() {
		return
	}
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	h.SubscriptionOrderStateReq.PaymentID = h.PaymentBaseReq.EntID
	h.PaymentBalanceLockReq.PaymentID = h.PaymentBaseReq.EntID
	if h.payWithBalance && h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

//nolint:gocyclo
func (h *createHandler) validatePaymentType() error {
	if h.OrderStateBaseReq.PaymentType == nil {
		if h.LedgerLockReq.EntID != nil || h.PaymentBaseReq.EntID != nil {
			return wlog.Errorf("invalid paymenttype")
		}
		return nil
	}
	switch *h.OrderStateBaseReq.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
		fallthrough //nolint
	case types.PaymentType_PayWithTransferAndBalance:
		if h.PaymentBaseReq.EntID == nil || h.LedgerLockReq.EntID == nil {
			return wlog.Errorf("invalid paymentid or ledgerlockid")
		}
	case types.PaymentType_PayWithTransferOnly:
		if h.PaymentBaseReq.EntID == nil || h.LedgerLockReq.EntID != nil {
			return wlog.Errorf("invalid paymentid or ledgerlockid")
		}
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

func (h *createHandler) validatePayment() error {
	if !h.paymentChecker.Payable() {
		if len(h.PaymentBalanceReqs) > 0 || len(h.PaymentTransferReqs) > 0 || len(h.PaymentFiatReqs) > 0 {
			return wlog.Errorf("invalid payment")
		}
		return nil
	}
	return h.paymentChecker.ValidatePayment()
}

func (h *createHandler) validateOrderType() error {
	switch *h.OrderBaseReq.CreateMethod {
	case types.OrderCreateMethod_OrderCreatedByPurchase:
		fallthrough //nolint
	case types.OrderCreateMethod_OrderCreatedByRenew:
		switch *h.OrderBaseReq.OrderType {
		case types.OrderType_Offline:
			fallthrough //nolint
		case types.OrderType_Airdrop:
			return wlog.Errorf("invalid ordertype")
		}
	case types.OrderCreateMethod_OrderCreatedByAdmin:
		switch *h.OrderBaseReq.OrderType {
		case types.OrderType_Offline:
			fallthrough //nolint
		case types.OrderType_Airdrop:
			if len(h.OrderCouponReqs) > 0 {
				return wlog.Errorf("invalid ordercoupons")
			}
		default:
			return wlog.Errorf("invalid ordertype")
		}
	}
	return nil
}

func (h *Handler) CreateSubscriptionOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &createHandler{
		Handler: h,
		paymentChecker: &paymentcommon.PaymentCheckHandler{
			PaymentType:         h.OrderStateBaseReq.PaymentType,
			PaymentBalanceReqs:  h.PaymentBalanceReqs,
			PaymentTransferReqs: h.PaymentTransferReqs,
			PaymentFiatReqs:     h.PaymentFiatReqs,
			PaymentAmountUSD:    h.PaymentAmountUSD,
			DiscountAmountUSD:   h.DiscountAmountUSD,
		},
		payWithBalance: len(h.PaymentBalanceReqs) > 0,
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if err := handler.validateOrderType(); err != nil {
		return wlog.WrapError(err)
	}

	handler.formalizeOrderID()
	handler.formalizeUserID()
	handler.formalizeEntIDs()
	handler.formalizeOrderCoupons()
	if err := handler.formalizePaymentType(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaymentState()
	handler.paymentChecker.PaymentType = h.OrderStateBaseReq.PaymentType
	handler.formalizePaymentID()
	if err := handler.validatePaymentType(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validatePayment(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaymentBalances()
	handler.formalizePaymentTransfers()
	handler.formalizePaymentFiats()

	handler.constructOrderBaseSQL(ctx)
	handler.constructOrderStateBaseSQL(ctx)
	handler.constructSubscriptionOrderStateSQL(ctx)
	handler.constructLedgerLockSQL(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructOrderCouponSQLs(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	handler.constructPaymentTransferSQLs(ctx)
	handler.constructPaymentFiatSQLs(ctx)
	handler.constructSQL()

	if err := handler.createOrderBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createSubscriptionOrderState(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createLedgerLock(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBalanceLock(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrderCoupons(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentBalances(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentTransfers(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentFiats(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.createSubscriptionOrder(ctx, tx)
}

func (h *Handler) CreateSubscriptionOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateSubscriptionOrderWithTx(_ctx, tx)
	})
}
