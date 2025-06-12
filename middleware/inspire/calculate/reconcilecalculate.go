package calculate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entorderpaymentstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderstatement"
	appConfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/commission"
	registration1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/registration"
	statementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	orderpaymentmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	commisisonmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"

	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

type reconcileCalculateHandler struct {
	*Handler
	inviters           []*registrationmwpb.Registration
	inviterIDs         []string
	appConfig          *appconfigmwpb.AppConfig
	statements         []*ent.OrderStatement
	payments           map[uuid.UUID][]*ent.OrderPaymentStatement
	orderUserStatement *ent.OrderStatement
	commissions        map[string]*commisisonmwpb.Commission
	ratios             map[string]decimal.Decimal
	infos              []*statementmwpb.StatementReq
}

func (h *reconcileCalculateHandler) CalculateUserCommissionRatio(ctx context.Context) error {
	for _, registration := range h.inviters {
		inviterID, err := uuid.Parse(registration.InviterID)
		if err != nil {
			return wlog.WrapError(err)
		}
		if inviterID == uuid.Nil {
			break
		}
		if _, err := uuid.Parse(registration.InviteeID); err != nil {
			return wlog.WrapError(err)
		}

		inviterRatio := decimal.NewFromInt(0)
		inviteeRatio := decimal.NewFromInt(0)

		inviterComm, inviterFound := h.commissions[registration.InviterID]
		if inviterFound {
			inviterRatio, err = decimal.NewFromString(inviterComm.AmountOrPercent)
			if err != nil {
				return wlog.WrapError(err)
			}
		}
		inviteeComm, inviteeFound := h.commissions[registration.InviteeID]
		if inviteeFound {
			inviteeRatio, err = decimal.NewFromString(inviteeComm.AmountOrPercent)
			if err != nil {
				return wlog.WrapError(err)
			}
		}

		if inviterRatio.Cmp(inviteeRatio) < 0 {
			return wlog.Errorf("inviter percent(%v) less than invitee percent(%v)", inviterRatio.String(), inviteeRatio.String())
		}
		if inviterRatio.Cmp(inviteeRatio) == 0 {
			h.ratios[registration.InviterID] = decimal.NewFromInt(0)
			continue
		}

		h.ratios[registration.InviterID] = inviterRatio.Sub(inviteeRatio)
	}

	// order user ratio
	comm, ok := h.commissions[h.UserID.String()]
	if ok {
		ratio, err := decimal.NewFromString(comm.AmountOrPercent)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ratios[h.UserID.String()] = ratio
	}
	return nil
}

func (h *reconcileCalculateHandler) GetCommissions(ctx context.Context) error {
	if _, ok := h.payments[h.orderUserStatement.EntID]; !ok {
		return wlog.Errorf("invalid payment")
	}

	h2, err := commission1.NewHandler(
		ctx,
		commission1.WithConds(&commisisonmwpb.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.SettleType_GoodOrderPayment)},
			UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: h.inviterIDs},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
			AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppGoodID.String()},
			StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
			EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
		}),
		commission1.WithOffset(0),
		commission1.WithLimit(int32(len(h.inviterIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	commissions, _, err := h2.GetCommissions(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, comm := range commissions {
		h.commissions[comm.UserID] = comm
	}
	return nil
}

func (h *reconcileCalculateHandler) getLayeredInviters(ctx context.Context) error {
	handler, err := registration1.NewHandler(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	handler.AppID = &h.AppID
	handler.InviteeID = &h.UserID

	inviters, inviterIDs, err := handler.GetSortedInviters(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.inviters = inviters
	h.inviterIDs = inviterIDs
	return nil
}

func (h *reconcileCalculateHandler) getAppConfig(ctx context.Context) error {
	h1, err := appConfig1.NewHandler(
		ctx,
		appConfig1.WithConds(&appconfigmwpb.Conds{
			AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			EndAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
		}),
		appConfig1.WithOffset(0),
		appConfig1.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	appConfigs, _, err := h1.GetAppConfigs(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(appConfigs) == 0 {
		return wlog.Errorf("invalid app config")
	}
	if appConfigs[0].CommissionType != types.CommissionType_LegacyCommission {
		return wlog.Errorf("invalid commission type: %v", appConfigs[0].CommissionType.String())
	}
	h.appConfig = appConfigs[0]
	return nil
}

func (h *reconcileCalculateHandler) getOrderPaymentStatements(ctx context.Context, cli *ent.Client) error {
	ids := []uuid.UUID{}
	for _, statement := range h.statements {
		ids = append(ids, statement.EntID)
	}
	payments, err := cli.
		OrderPaymentStatement.
		Query().
		Where(
			entorderpaymentstatement.StatementIDIn(ids...),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, payment := range payments {
		h.payments[payment.StatementID] = append(h.payments[payment.StatementID], payment)
	}
	return nil
}

func (h *reconcileCalculateHandler) requireOrderStatement() error {
	found := false
	// find order user statement
	for _, statement := range h.statements {
		if statement.UserID != statement.OrderUserID || statement.CommissionConfigType != types.CommissionConfigType_LegacyCommissionConfig.String() {
			continue
		}
		found = true
		h.AppID = statement.AppID
		h.UserID = statement.UserID
		h.GoodID = statement.GoodID
		h.AppGoodID = statement.AppGoodID
		h.GoodCoinTypeID = statement.GoodCoinTypeID
		h.Units = statement.Units
		h.GoodValueUSD = statement.GoodValueUsd
		h.PaymentAmountUSD = statement.PaymentAmountUsd
		h.OrderCreatedAt = statement.CreatedAt
		h.orderUserStatement = statement
	}
	if !found {
		return wlog.Errorf("order user statement not found")
	}
	return nil
}

func (h *reconcileCalculateHandler) getOrderStatements(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		statements, err := cli.
			OrderStatement.
			Query().
			Where(
				entorderstatement.OrderID(h.OrderID),
			).
			All(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(statements) == 0 {
			return wlog.Errorf("invalid statement")
		}
		h.statements = statements

		if err := h.requireOrderStatement(); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getOrderPaymentStatements(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *Handler) ReconcileCalculate(ctx context.Context) ([]*statementmwpb.StatementReq, error) {
	handler := &reconcileCalculateHandler{
		Handler:     h,
		inviters:    []*registrationmwpb.Registration{},
		inviterIDs:  []string{},
		statements:  []*ent.OrderStatement{},
		payments:    map[uuid.UUID][]*ent.OrderPaymentStatement{},
		commissions: map[string]*commisisonmwpb.Commission{},
		ratios:      map[string]decimal.Decimal{},
		infos:       []*statementmwpb.StatementReq{},
	}
	if err := handler.getOrderStatements(ctx); err != nil {
		return nil, err
	}
	if err := handler.getAppConfig(ctx); err != nil {
		return nil, err
	}
	if err := handler.getLayeredInviters(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCommissions(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.CalculateUserCommissionRatio(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, nil
}

func (h *reconcileCalculateHandler) formalize() {
	for _, statement := range h.statements {
		dbPayments, ok := h.payments[statement.EntID]
		if !ok {
			continue
		}
		req := statementmwpb.StatementReq{
			ID:                  &statement.ID,
			EntID:               func() *string { s := statement.EntID.String(); return &s }(),
			AppID:               func() *string { s := statement.AppID.String(); return &s }(),
			UserID:              func() *string { s := statement.UserID.String(); return &s }(),
			GoodID:              func() *string { s := statement.GoodID.String(); return &s }(),
			AppGoodID:           func() *string { s := statement.AppGoodID.String(); return &s }(),
			OrderID:             func() *string { s := statement.OrderID.String(); return &s }(),
			OrderUserID:         func() *string { s := statement.OrderUserID.String(); return &s }(),
			DirectContributorID: func() *string { s := statement.DirectContributorID.String(); return &s }(),
			GoodCoinTypeID:      func() *string { s := statement.GoodCoinTypeID.String(); return &s }(),
			Units:               func() *string { s := statement.Units.String(); return &s }(),
			GoodValueUSD:        func() *string { s := statement.GoodValueUsd.String(); return &s }(),
			PaymentAmountUSD:    func() *string { s := h.PaymentAmountUSD.String(); return &s }(),
			CommissionConfigID:  func() *string { s := statement.CommissionConfigID.String(); return &s }(),
			CommissionAmountUSD: func() *string { s := decimal.NewFromInt(0).String(); return &s }(),
			AppConfigID:         &h.appConfig.EntID,
			CommissionConfigType: func() *types.CommissionConfigType {
				s := types.CommissionConfigType(types.CommissionConfigType_value[statement.CommissionConfigType])
				return &s
			}(),
		}

		ratio, ratioFound := h.ratios[statement.UserID.String()]
		if ratioFound {
			req.CommissionAmountUSD = func() *string {
				amount := statement.PaymentAmountUsd.Mul(ratio).Div(decimal.NewFromInt(100)).String() //nolint
				return &amount
			}()
		}
		for _, dbPayment := range dbPayments {
			payment := &orderpaymentmwpb.StatementReq{
				EntID:             func() *string { id := dbPayment.EntID.String(); return &id }(),
				Amount:            func() *string { amount := dbPayment.Amount.String(); return &amount }(),
				CommissionAmount:  func() *string { amount := "0"; return &amount }(),
				PaymentCoinTypeID: func() *string { id := dbPayment.PaymentCoinTypeID.String(); return &id }(),
			}

			comm, ok := h.commissions[statement.UserID.String()]
			if !ok {
				req.PaymentStatements = append(req.PaymentStatements, payment)
				continue
			}
			if !ratioFound {
				req.PaymentStatements = append(req.PaymentStatements, payment)
				req.CommissionConfigID = &comm.EntID
				continue
			}
			commissionAmount := dbPayment.Amount.Mul(ratio).Div(decimal.NewFromInt(100)).String() //nolint
			payment.CommissionAmount = &commissionAmount

			req.PaymentStatements = append(req.PaymentStatements, payment)
			req.CommissionConfigID = &comm.EntID
		}

		h.infos = append(h.infos, &req)
	}
}
