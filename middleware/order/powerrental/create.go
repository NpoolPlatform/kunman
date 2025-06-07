package powerrental

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
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
	paymenttransfer1 "github.com/NpoolPlatform/kunman/middleware/order/payment/transfer"
	powerrentalstate1 "github.com/NpoolPlatform/kunman/middleware/order/powerrental/state"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	paymentChecker        *paymentcommon.PaymentCheckHandler
	payWithBalance        bool
	sql                   string
	sqlOrderBase          string
	sqlOrderStateBase     string
	sqlPowerRentalState   string
	sqlOrderLocks         []string
	sqlPaymentBalanceLock string
	sqlOrderCoupons       []string
	sqlPaymentBase        string
	sqlPaymentBalances    []string
	sqlPaymentTransfers   []string
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
	h.sqlOrderStateBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPowerRentalStateSQL(ctx context.Context) {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *h.PowerRentalStateReq
	h.sqlPowerRentalState = handler.ConstructCreateSQL()
}

func (h *createHandler) constructOrderLockSQLs(ctx context.Context) {
	if h.OrderBaseReq.Simulate != nil && *h.OrderBaseReq.Simulate {
		return
	}
	for _, req := range h.OrderLockReqs {
		handler, _ := orderlock1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderLocks = append(h.sqlOrderLocks, handler.ConstructCreateSQL())
		if !h.payWithBalance {
			h.payWithBalance = *req.LockType == types.OrderLockType_LockBalance
		}
	}
}

func (h *createHandler) constructPaymentBalanceLockSQL(ctx context.Context) {
	if !h.payWithBalance {
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
	if !h.paymentChecker.Payable() {
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

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create powerrental: %v", err)
	}
	return nil
}

func (h *createHandler) createOrderBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderBase)
}

func (h *createHandler) createOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderStateBase)
}

func (h *createHandler) createPowerRentalState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPowerRentalState)
}

func (h *createHandler) createOrderLocks(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderLocks {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
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

func (h *createHandler) createPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *createHandler) formalizeOrderID() {
	if h.OrderID != nil {
		return
	}
	h.OrderID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.OrderBaseReq.EntID = h.OrderID
	h.OrderStateBaseReq.OrderID = h.OrderID
	h.PowerRentalStateReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
	h.PoolOrderUserReq.OrderID = h.OrderID
}

func (h *createHandler) formalizeOrderLocks() {
	for _, req := range h.OrderLockReqs {
		req.OrderID = h.OrderID
		req.UserID = h.OrderBaseReq.UserID
	}
}

func (h *createHandler) formalizeEntIDs() {
	if h.OrderStateBaseReq.EntID == nil {
		h.OrderStateBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PowerRentalStateReq.EntID == nil {
		h.PowerRentalStateReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
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

func (h *createHandler) formalizePaymentType() error {
	if h.OrderBaseReq.Simulate != nil && *h.OrderBaseReq.Simulate {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithNoPayment; return &e }()
		return nil
	}
	if *h.OrderBaseReq.OrderType == types.OrderType_Offline {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithOffline; return &e }()
		return nil
	}
	if *h.OrderBaseReq.OrderType == types.OrderType_Airdrop {
		h.OrderStateBaseReq.PaymentType = func() *types.PaymentType { e := types.PaymentType_PayWithNoPayment; return &e }()
		return nil
	}
	// For shop cart
	if h.OrderStateBaseReq.PaymentType != nil && *h.OrderStateBaseReq.PaymentType == types.PaymentType_PayWithOtherOrder {
		return nil
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
	if h.OrderStateBaseReq.PaymentType == nil {
		return wlog.Errorf("invalid paymenttype")
	}
	return nil
}

func (h *createHandler) formalizePaymentState() {
	if *h.OrderStateBaseReq.PaymentType != types.PaymentType_PayWithNoPayment {
		return
	}
	h.PowerRentalStateReq.PaymentState = types.PaymentState_PaymentStateNoPayment.Enum()
}

func (h *createHandler) formalizeStartAt() {
	switch *h.OrderStateBaseReq.StartMode {
	case types.OrderStartMode_OrderStartInstantly:
		h.OrderStateBaseReq.StartAt = func() *uint32 { u := uint32(time.Now().Unix() + 60); return &u }() //nolint
	case types.OrderStartMode_OrderStartNextDay:
		h.OrderStateBaseReq.StartAt = func() *uint32 { u := uint32(timedef.TomorrowStart().Unix()); return &u }()
	}
}

func (h *createHandler) formalizeFeeOrders() {
	for _, handler := range h.FeeMultiHandler.GetHandlers() {
		handler.OrderBaseReq.AppID = h.OrderBaseReq.AppID
		handler.OrderBaseReq.UserID = h.OrderBaseReq.UserID
		handler.OrderBaseReq.ParentOrderID = h.OrderBaseReq.EntID
		handler.OrderBaseReq.OrderType = h.OrderBaseReq.OrderType
		handler.OrderBaseReq.CreateMethod = h.OrderBaseReq.CreateMethod
		handler.PaymentBaseReq.EntID = h.PowerRentalStateReq.PaymentID
		handler.FeeOrderStateReq.PaymentID = h.PowerRentalStateReq.PaymentID
	}
}

func (h *createHandler) formalizePaymentID() {
	if !h.paymentChecker.Payable() {
		return
	}
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	h.PowerRentalStateReq.PaymentID = h.PaymentBaseReq.EntID
	h.PaymentBalanceLockReq.PaymentID = h.PaymentBaseReq.EntID
}

func (h *createHandler) createFeeOrders(ctx context.Context, tx *ent.Tx) error {
	return h.FeeMultiHandler.CreateFeeOrdersWithTx(ctx, tx)
}

//nolint:gocyclo
func (h *createHandler) validatePaymentType() error {
	if h.OrderStateBaseReq.PaymentType == nil {
		if h.ledgerLockID() != nil || h.PaymentBaseReq.EntID != nil {
			return wlog.Errorf("invalid paymenttype")
		}
		return nil
	}
	switch *h.OrderStateBaseReq.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
		fallthrough //nolint
	case types.PaymentType_PayWithTransferAndBalance:
		if h.PaymentBaseReq.EntID == nil || h.ledgerLockID() == nil {
			return wlog.Errorf("invalid paymentid or ledgerlockid")
		}
	case types.PaymentType_PayWithTransferOnly:
		if h.PaymentBaseReq.EntID == nil || h.ledgerLockID() != nil {
			return wlog.Errorf("invalid paymentid or ledgerlockid")
		}
	case types.PaymentType_PayWithOtherOrder:
		fallthrough //nolint
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

func (h *createHandler) validatePayment() error {
	if !h.paymentChecker.Payable() {
		if len(h.PaymentBalanceReqs) > 0 || len(h.PaymentTransferReqs) > 0 {
			return wlog.Errorf("invalid payment")
		}
		return nil
	}
	return h.paymentChecker.ValidatePayment()
}

func (h *createHandler) validateGoodStockMode() error {
	if h.OrderBaseReq.Simulate == nil || !*h.OrderBaseReq.Simulate {
		return nil
	}
	if *h.GoodStockMode == goodtypes.GoodStockMode_GoodStockByMiningPool {
		return wlog.Errorf("disable simulate order of good is goodstockbyminingpool")
	}
	return nil
}

func (h *createHandler) validateAppGoodStock() error {
	if h.OrderBaseReq.Simulate != nil && *h.OrderBaseReq.Simulate {
		return nil
	}
	if h.AppGoodStockID == nil {
		return wlog.Errorf("invalid appgoodstockid")
	}
	for _, req := range h.OrderLockReqs {
		if *req.LockType == types.OrderLockType_LockStock {
			return nil
		}
	}
	return wlog.Errorf("invalid appgoodstocklock")
}

func (h *createHandler) validateOrderType() error {
	if h.OrderBaseReq.Simulate != nil && *h.OrderBaseReq.Simulate {
		if *h.OrderBaseReq.CreateMethod != types.OrderCreateMethod_OrderCreatedByPurchase {
			return wlog.Errorf("invalid createmethod")
		}
		if len(h.OrderCouponReqs) > 0 {
			return wlog.Errorf("invalid ordercoupons")
		}
	}
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

//nolint:funlen,gocyclo
func (h *Handler) CreatePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &createHandler{
		Handler: h,
		paymentChecker: &paymentcommon.PaymentCheckHandler{
			PaymentType:         h.OrderStateBaseReq.PaymentType,
			PaymentBalanceReqs:  h.PaymentBalanceReqs,
			PaymentTransferReqs: h.PaymentTransferReqs,
			PaymentAmountUSD:    h.PaymentAmountUSD,
			DiscountAmountUSD:   h.DiscountAmountUSD,
			Simulate:            h.OrderBaseReq.Simulate,
		},
	}

	if err := handler.validateAppGoodStock(); err != nil {
		return err
	}

	if err := handler.validateGoodStockMode(); err != nil {
		return err
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	handler.formalizeOrderID()
	handler.formalizeOrderLocks()
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
	if err := handler.validateOrderType(); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizePaymentBalances()
	handler.formalizePaymentTransfers()
	handler.formalizeStartAt()
	handler.formalizeFeeOrders()

	handler.constructOrderBaseSQL(ctx)
	handler.constructOrderStateBaseSQL(ctx)
	handler.constructPowerRentalStateSQL(ctx)
	handler.constructOrderLockSQLs(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructOrderCouponSQLs(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	handler.constructPaymentTransferSQLs(ctx)
	handler.constructSQL()

	if err := handler.createOrderBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPowerRentalState(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrderLocks(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if handler.paymentChecker.Payable() {
		if err := handler.createPaymentBase(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
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
	if err := handler.createFeeOrders(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.createPowerRental(ctx, tx)
}

func (h *Handler) CreatePowerRental(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreatePowerRentalWithTx(_ctx, tx)
	})
}
