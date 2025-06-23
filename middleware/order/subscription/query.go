package subscriptionorder

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ordercouponmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	ordercouponcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/coupon"
	paymentbalancecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/balance"
	paymentfiatcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/fiat"
	paymenttransfercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/transfer"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entpaymentbalance "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalance"
	entpaymentfiat "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentfiat"
	entpaymenttransfer "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymenttransfer"
	entsubscriptionorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/subscriptionorder"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount         *ent.OrderBaseSelect
	infos            []*npool.SubscriptionOrder
	orderCoupons     []*ordercouponmiddlewarepb.OrderCouponInfo
	paymentBalances  []*paymentmwpb.PaymentBalanceInfo
	paymentTransfers []*paymentmwpb.PaymentTransferInfo
	paymentFiats     []*paymentmwpb.PaymentFiatInfo
	total            uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinSubscriptionOrder(s); err != nil {
			logger.Sugar().Errorw("queryJoinSubscriptionOrder", "Error", err)
		}
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinSubscriptionOrderState(s); err != nil {
			logger.Sugar().Errorw("queryJoinSubscriptionOrderState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) queryOrderCoupons(ctx context.Context, cli *ent.Client) error {
	orderIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.OrderID))
		}
		return
	}()

	stm, err := ordercouponcrud.SetQueryConds(
		cli.OrderCoupon.Query(),
		&ordercouponcrud.Conds{
			OrderIDs: &cruder.Cond{Op: cruder.IN, Val: orderIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entordercoupon.FieldOrderID,
		entordercoupon.FieldCouponID,
		entordercoupon.FieldCreatedAt,
	).Scan(ctx, &h.orderCoupons)
}

func (h *queryHandler) queryPaymentBalances(ctx context.Context, cli *ent.Client) error {
	paymentIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.PaymentID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.PaymentID))
		}
		return
	}()

	stm, err := paymentbalancecrud.SetQueryConds(
		cli.PaymentBalance.Query(),
		&paymentbalancecrud.Conds{
			PaymentIDs: &cruder.Cond{Op: cruder.IN, Val: paymentIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entpaymentbalance.FieldEntID,
		entpaymentbalance.FieldPaymentID,
		entpaymentbalance.FieldCoinTypeID,
		entpaymentbalance.FieldAmount,
		entpaymentbalance.FieldCoinUsdCurrency,
		entpaymentbalance.FieldLocalCoinUsdCurrency,
		entpaymentbalance.FieldLiveCoinUsdCurrency,
		entpaymentbalance.FieldCreatedAt,
	).Scan(ctx, &h.paymentBalances)
}

func (h *queryHandler) queryPaymentTransfers(ctx context.Context, cli *ent.Client) error {
	paymentIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.PaymentID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.PaymentID))
		}
		return
	}()

	stm, err := paymenttransfercrud.SetQueryConds(
		cli.PaymentTransfer.Query(),
		&paymenttransfercrud.Conds{
			PaymentIDs: &cruder.Cond{Op: cruder.IN, Val: paymentIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entpaymenttransfer.FieldEntID,
		entpaymenttransfer.FieldPaymentID,
		entpaymenttransfer.FieldCoinTypeID,
		entpaymenttransfer.FieldAmount,
		entpaymenttransfer.FieldAccountID,
		entpaymenttransfer.FieldStartAmount,
		entpaymenttransfer.FieldCoinUsdCurrency,
		entpaymenttransfer.FieldLocalCoinUsdCurrency,
		entpaymenttransfer.FieldLiveCoinUsdCurrency,
		entpaymenttransfer.FieldFinishAmount,
		entpaymenttransfer.FieldCreatedAt,
	).Scan(ctx, &h.paymentTransfers)
}

func (h *queryHandler) queryPaymentFiats(ctx context.Context, cli *ent.Client) error {
	paymentIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if _, err := uuid.Parse(info.PaymentID); err != nil {
				continue
			}
			uids = append(uids, uuid.MustParse(info.PaymentID))
		}
		return
	}()

	stm, err := paymentfiatcrud.SetQueryConds(
		cli.PaymentFiat.Query(),
		&paymentfiatcrud.Conds{
			PaymentIDs: &cruder.Cond{Op: cruder.IN, Val: paymentIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entpaymentfiat.FieldEntID,
		entpaymentfiat.FieldPaymentID,
		entpaymentfiat.FieldFiatID,
		entpaymentfiat.FieldPaymentChannel,
		entpaymentfiat.FieldAmount,
		entpaymentfiat.FieldChannelPaymentID,
		entpaymentfiat.FieldUsdCurrency,
		entpaymentfiat.FieldCreatedAt,
	).Scan(ctx, &h.paymentFiats)
}

func (h *queryHandler) formalize() {
	orderCoupons := map[string][]*ordercouponmiddlewarepb.OrderCouponInfo{}
	paymentBalances := map[string][]*paymentmwpb.PaymentBalanceInfo{}
	paymentTransfers := map[string][]*paymentmwpb.PaymentTransferInfo{}
	paymentFiats := map[string][]*paymentmwpb.PaymentFiatInfo{}

	for _, orderCoupon := range h.orderCoupons {
		orderCoupons[orderCoupon.OrderID] = append(orderCoupons[orderCoupon.OrderID], orderCoupon)
	}
	for _, paymentBalance := range h.paymentBalances {
		paymentBalances[paymentBalance.PaymentID] = append(paymentBalances[paymentBalance.PaymentID], paymentBalance)
		paymentBalance.Amount = func() string { amount, _ := decimal.NewFromString(paymentBalance.Amount); return amount.String() }()
		paymentBalance.CoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.CoinUSDCurrency)
			return amount.String()
		}()
		paymentBalance.LocalCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.LocalCoinUSDCurrency)
			return amount.String()
		}()
		paymentBalance.LiveCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentBalance.LiveCoinUSDCurrency)
			return amount.String()
		}()
	}
	for _, paymentTransfer := range h.paymentTransfers {
		paymentTransfers[paymentTransfer.PaymentID] = append(paymentTransfers[paymentTransfer.PaymentID], paymentTransfer)
		paymentTransfer.Amount = func() string { amount, _ := decimal.NewFromString(paymentTransfer.Amount); return amount.String() }()
		paymentTransfer.StartAmount = func() string { amount, _ := decimal.NewFromString(paymentTransfer.StartAmount); return amount.String() }()
		paymentTransfer.FinishAmount = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.FinishAmount)
			return amount.String()
		}()
		paymentTransfer.CoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.CoinUSDCurrency)
			return amount.String()
		}()
		paymentTransfer.LocalCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.LocalCoinUSDCurrency)
			return amount.String()
		}()
		paymentTransfer.LiveCoinUSDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentTransfer.LiveCoinUSDCurrency)
			return amount.String()
		}()
	}
	for _, paymentFiat := range h.paymentFiats {
		paymentFiats[paymentFiat.PaymentID] = append(paymentFiats[paymentFiat.PaymentID], paymentFiat)
		paymentFiat.Amount = func() string { amount, _ := decimal.NewFromString(paymentFiat.Amount); return amount.String() }()
		paymentFiat.PaymentChannel = types.FiatPaymentChannel(types.FiatPaymentChannel_value[paymentFiat.PaymentChannelStr])
		paymentFiat.USDCurrency = func() string {
			amount, _ := decimal.NewFromString(paymentFiat.USDCurrency)
			return amount.String()
		}()
	}

	for _, info := range h.infos {
		info.GoodValueUSD = func() string { amount, _ := decimal.NewFromString(info.GoodValueUSD); return amount.String() }()
		info.PaymentGoodValueUSD = func() string { amount, _ := decimal.NewFromString(info.GoodValueUSD); return amount.String() }()
		info.PaymentAmountUSD = func() string { amount, _ := decimal.NewFromString(info.PaymentAmountUSD); return amount.String() }()
		info.DiscountAmountUSD = func() string { amount, _ := decimal.NewFromString(info.DiscountAmountUSD); return amount.String() }()
		info.GoodType = goodtypes.GoodType(goodtypes.GoodType_value[info.GoodTypeStr])
		info.OrderType = types.OrderType(types.OrderType_value[info.OrderTypeStr])
		info.PaymentType = types.PaymentType(types.PaymentType_value[info.PaymentTypeStr])
		info.PaymentState = types.PaymentState(types.PaymentState_value[info.PaymentStateStr])
		info.OrderState = types.OrderState(types.OrderState_value[info.OrderStateStr])
		info.CancelState = types.OrderState(types.OrderState_value[info.CancelStateStr])
		info.CreateMethod = types.OrderCreateMethod(types.OrderCreateMethod_value[info.CreateMethodStr])
		info.Coupons = orderCoupons[info.OrderID]
		info.PaymentBalances = paymentBalances[info.PaymentID]
		info.PaymentTransfers = paymentTransfers[info.PaymentID]
		info.PaymentFiats = paymentFiats[info.PaymentID]
	}
}

func (h *Handler) GetSubscriptionOrder(ctx context.Context) (*npool.SubscriptionOrder, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentBalances(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentTransfers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentFiats(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.queryOrderCoupons(_ctx, cli)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetSubscriptionOrders(ctx context.Context) ([]*npool.SubscriptionOrder, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entsubscriptionorder.FieldCreatedAt))

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentBalances(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentTransfers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentFiats(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.queryOrderCoupons(_ctx, cli)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}

func (h *Handler) GetSubscriptionOrderOnly(ctx context.Context) (*npool.SubscriptionOrder, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(0).
			Limit(2).
			Order(ent.Desc(entsubscriptionorder.FieldCreatedAt))

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentBalances(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentTransfers(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.queryPaymentFiats(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.queryOrderCoupons(_ctx, cli)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("invalid order")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()

	return handler.infos[0], nil
}
