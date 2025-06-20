//nolint:funlen,dupl
package calculate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	statementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"
	achievementusermwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/user"
	appcommissionconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	appConfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	appgoodcommissionconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/good/commission/config"
	commmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"
	achievementuser1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user"
	common1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
	appcommissionconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/commission/config"
	appConfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	appgoodcommissionconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/good/commission/config"
	commission2 "github.com/NpoolPlatform/kunman/middleware/inspire/calculate/commission"
	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/commission"
	registration1 "github.com/NpoolPlatform/kunman/middleware/inspire/invitation/registration"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type calculateHandler struct {
	*Handler
	inviters   []*registrationmwpb.Registration
	inviterIDs []string
}

func (h *calculateHandler) getLayeredInviters(ctx context.Context) error {
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

func (h *calculateHandler) getDirectInviters(ctx context.Context) error {
	handler, err := registration1.NewHandler(
		ctx,
		registration1.WithConds(&registrationmwpb.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			InviteeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID.String()},
		}),
		registration1.WithOffset(0),
		registration1.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	inviters, _, err := handler.GetRegistrations(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if len(inviters) == 0 {
		return nil
	}
	h.inviters = inviters
	h.inviterIDs = append(h.inviterIDs, inviters[0].InviterID)
	return nil
}

func (h *calculateHandler) getAchievementUsers(ctx context.Context) (map[string]*achievementusermwpb.AchievementUser, error) {
	achievementUserMap := map[string]*achievementusermwpb.AchievementUser{}
	handler, err := achievementuser1.NewHandler(
		ctx,
		common1.WithConds(&achievementusermwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.inviterIDs},
		}),
		common1.WithOffset(0),
		common1.WithLimit(int32(len(h.inviterIDs))),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	achievmentUsers, _, err := handler.GetAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(achievmentUsers) == 0 {
		return achievementUserMap, nil
	}
	for _, id := range h.inviterIDs {
		for _, achievementUser := range achievmentUsers {
			if achievementUser.UserID == id {
				achievementUserMap[id] = achievementUser
			}
		}
	}
	return achievementUserMap, nil
}

//nolint:funlen
func (h *Handler) Calculate(ctx context.Context) ([]*statementmwpb.StatementReq, error) {
	h1, err := appConfig1.NewHandler(
		ctx,
		appConfig1.WithConds(&appConfigmwpb.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			StartAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
			EndAt:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
		}),
		appConfig1.WithOffset(0),
		appConfig1.WithLimit(1),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	appConfigs, _, err := h1.GetAppConfigs(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler := &calculateHandler{
		Handler:    h,
		inviters:   []*registrationmwpb.Registration{},
		inviterIDs: []string{},
	}

	commissionConfigType := types.CommissionConfigType_WithoutCommissionConfig

	if len(appConfigs) == 0 {
		return handler.generateStatements(map[string]map[uuid.UUID][]*commission2.Commission{}, uuid.Nil.String(), commissionConfigType)
	}
	appConfig := appConfigs[0]

	switch appConfig.CommissionType {
	case types.CommissionType_LegacyCommission:
		commissionConfigType = types.CommissionConfigType_LegacyCommissionConfig
		fallthrough
	case types.CommissionType_LayeredCommission:
		err := handler.getLayeredInviters(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
	case types.CommissionType_DirectCommission:
		err := handler.getDirectInviters(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if len(handler.inviters) == 0 {
			return handler.generateStatements(map[string]map[uuid.UUID][]*commission2.Commission{}, appConfig.EntID, commissionConfigType)
		}
	case types.CommissionType_WithoutCommission:
	default:
		return nil, wlog.Errorf("invalid commissiontype")
	}

	achievementUsers, err := handler.getAchievementUsers(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	commMap := map[string]map[uuid.UUID][]*commission2.Commission{} // userid->cointypeid->commissions
	if h.HasCommission {
		for _, payment := range h.Payments {
			_comms := []*commission2.Commission{}

			switch appConfig.CommissionType {
			case types.CommissionType_LegacyCommission:
				h2, err := commission1.NewHandler(
					ctx,
					commission1.WithConds(&commmwpb.Conds{
						AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
						UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: handler.inviterIDs},
						GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
						AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppGoodID.String()},
						SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
						StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
						EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
					}),
					commission1.WithOffset(0),
					commission1.WithLimit(int32(len(handler.inviterIDs))),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}

				comms, _, err := h2.GetCommissions(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				handler, err := commission2.NewHandler(
					ctx,
					commission2.WithSettleType(types.SettleType_GoodOrderPayment),
					commission2.WithSettleAmountType(h.SettleAmountType),
					commission2.WithInviters(handler.inviters),
					commission2.WithAppConfig(appConfig),
					commission2.WithCommissions(comms),
					commission2.WithPaymentAmount(payment.Amount.String()),
					commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
					commission2.WithAchievementUsers(achievementUsers),
					commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				_comms, err = handler.Calculate(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
			case types.CommissionType_LayeredCommission:
				fallthrough //nolint
			case types.CommissionType_DirectCommission:
				h2, err := appgoodcommissionconfig1.NewHandler(
					ctx,
					appgoodcommissionconfig1.WithConds(&appgoodcommissionconfigmwpb.Conds{
						AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
						GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID.String()},
						AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AppGoodID.String()},
						SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
						EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
						StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
						Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
					}),
					appgoodcommissionconfig1.WithOffset(0),
					appgoodcommissionconfig1.WithLimit(0),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}

				goodcomms, _, err := h2.GetCommissionConfigs(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				if len(goodcomms) > 0 {
					handler, err := commission2.NewHandler(
						ctx,
						commission2.WithSettleType(types.SettleType_GoodOrderPayment),
						commission2.WithSettleAmountType(h.SettleAmountType),
						commission2.WithInviters(handler.inviters),
						commission2.WithAppConfig(appConfig),
						commission2.WithAppGoodCommissionConfigs(goodcomms),
						commission2.WithPaymentAmount(payment.Amount.String()),
						commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
						commission2.WithAchievementUsers(achievementUsers),
						commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
					)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					_comms, err = handler.CalculateByAppGoodCommConfig(ctx)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					break
				}

				h3, err := appcommissionconfig1.NewHandler(
					ctx,
					appcommissionconfig1.WithConds(&appcommissionconfigmwpb.Conds{
						AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
						SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(h.SettleType)},
						EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(0)},
						StartAt:    &basetypes.Uint32Val{Op: cruder.LTE, Value: h.OrderCreatedAt},
						Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
					}),
					appcommissionconfig1.WithOffset(0),
					appcommissionconfig1.WithLimit(0),
				)
				if err != nil {
					return nil, wlog.WrapError(err)
				}

				appcomms, _, err := h3.GetCommissionConfigs(ctx)
				if err != nil {
					return nil, wlog.WrapError(err)
				}
				if len(appcomms) > 0 {
					handler, err := commission2.NewHandler(
						ctx,
						commission2.WithSettleType(types.SettleType_GoodOrderPayment),
						commission2.WithSettleAmountType(h.SettleAmountType),
						commission2.WithInviters(handler.inviters),
						commission2.WithAppConfig(appConfig),
						commission2.WithAppCommissionConfigs(appcomms),
						commission2.WithPaymentAmount(payment.Amount.String()),
						commission2.WithPaymentAmountUSD(h.PaymentAmountUSD.String()),
						commission2.WithAchievementUsers(achievementUsers),
						commission2.WithGoodValueUSD(h.GoodValueUSD.String()),
					)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
					_comms, err = handler.CalculateByAppCommConfig(ctx)
					if err != nil {
						return nil, wlog.WrapError(err)
					}
				}
			case types.CommissionType_WithoutCommission:
			default:
				return nil, wlog.Errorf("invalid commissiontype")
			}

			for _, _com := range _comms {
				coinCommMap, ok := commMap[_com.UserID]
				if !ok {
					coinCommMap = map[uuid.UUID][]*commission2.Commission{}
				}
				commissions, ok := coinCommMap[*payment.PaymentCoinTypeID]
				if !ok {
					commissions = []*commission2.Commission{}
				}
				commissions = append(commissions, _com)
				coinCommMap[*payment.PaymentCoinTypeID] = commissions
				commMap[_com.UserID] = coinCommMap
			}
			if len(_comms) == 0 {
				_com := &commission2.Commission{
					AppID:                h.AppID.String(),
					UserID:               h.UserID.String(),
					PaymentAmount:        payment.Amount.String(),
					Amount:               "0",
					CommissionAmountUSD:  "0",
					AppConfigID:          appConfig.EntID,
					CommissionConfigID:   uuid.Nil.String(),
					CommissionConfigType: types.CommissionConfigType(appConfig.CommissionType),
				}
				coinCommMap, ok := commMap[_com.UserID]
				if !ok {
					coinCommMap = map[uuid.UUID][]*commission2.Commission{}
				}
				commissions, ok := coinCommMap[*payment.PaymentCoinTypeID]
				if !ok {
					commissions = []*commission2.Commission{}
				}
				commissions = append(commissions, _com)
				coinCommMap[*payment.PaymentCoinTypeID] = commissions
				commMap[_com.UserID] = coinCommMap
			}
		}
	}

	return handler.generateStatements(commMap, appConfigs[0].EntID, commissionConfigType)
}

//nolint:funlen,dupl
func (h *calculateHandler) generateStatements(
	userCoinCommMap map[string]map[uuid.UUID][]*commission2.Commission,
	appConfigID string,
	commissionConfigType types.CommissionConfigType,
) ([]*statementmwpb.StatementReq, error) {
	if len(userCoinCommMap) == 0 {
		for _, payment := range h.Payments {
			_com := &commission2.Commission{
				AppID:                h.AppID.String(),
				UserID:               h.UserID.String(),
				PaymentAmount:        payment.Amount.String(),
				Amount:               "0",
				CommissionAmountUSD:  "0",
				AppConfigID:          appConfigID,
				CommissionConfigID:   uuid.Nil.String(),
				CommissionConfigType: commissionConfigType,
			}
			coinCommMap, ok := userCoinCommMap[_com.UserID]
			if !ok {
				coinCommMap = map[uuid.UUID][]*commission2.Commission{}
			}
			commissions, ok := coinCommMap[*payment.PaymentCoinTypeID]
			if !ok {
				commissions = []*commission2.Commission{}
			}
			commissions = append(commissions, _com)
			coinCommMap[*payment.PaymentCoinTypeID] = commissions
			userCoinCommMap[h.UserID.String()] = coinCommMap
		}
	}

	statements := []*statementmwpb.StatementReq{}
	for _, inviter := range h.inviters {
		if inviter.InviterID == uuid.Nil.String() {
			continue
		}

		commissionAmountUSD := decimal.NewFromInt(0).String()
		inviterCommissionConfigID := uuid.Nil.String()
		inviterCommissionConfigType := commissionConfigType

		inviterPayments := []*paymentmwpb.StatementReq{}

		userCommMap, ok := userCoinCommMap[inviter.InviterID]
		if ok {
			for key, commissions := range userCommMap {
				_key := key.String()
				for _, commission := range commissions {
					inviterPayments = append(inviterPayments, &paymentmwpb.StatementReq{
						PaymentCoinTypeID: &_key,
						Amount:            &commission.PaymentAmount,
						CommissionAmount:  &commission.Amount,
					})
					inviterCommissionConfigID = commission.CommissionConfigID
					inviterCommissionConfigType = commission.CommissionConfigType
					if h.HasCommission {
						commissionAmountUSD = commission.CommissionAmountUSD
					}
				}
			}
		} else {
			for _, payment := range h.Payments {
				coinTypeID := payment.PaymentCoinTypeID.String()
				amount := payment.Amount.String()
				commissionAmount := "0"
				inviterPayments = append(inviterPayments, &paymentmwpb.StatementReq{
					PaymentCoinTypeID: &coinTypeID,
					Amount:            &amount,
					CommissionAmount:  &commissionAmount,
				})
			}
		}

		statements = append(statements, &statementmwpb.StatementReq{
			AppID:               func() *string { id := h.AppID.String(); return &id }(),
			UserID:              &inviter.InviterID,
			OrderUserID:         func() *string { s := h.UserID.String(); return &s }(),
			DirectContributorID: &inviter.InviteeID,
			GoodID: func() *string {
				id := h.GoodID.String()
				return &id
			}(),
			AppGoodID: func() *string {
				id := h.AppGoodID.String()
				return &id
			}(),
			OrderID: func() *string {
				id := h.OrderID.String()
				return &id
			}(),

			GoodCoinTypeID: func() *string {
				id := h.GoodCoinTypeID.String()
				return &id
			}(),
			Units: func() *string {
				units := h.Units.String()
				return &units
			}(),
			GoodValueUSD: func() *string {
				goodValueUSD := h.GoodValueUSD.String()
				return &goodValueUSD
			}(),
			PaymentAmountUSD: func() *string {
				paymentAmountUSD := h.PaymentAmountUSD.String()
				return &paymentAmountUSD
			}(),
			CommissionAmountUSD:  &commissionAmountUSD,
			AppConfigID:          &appConfigID,
			CommissionConfigID:   &inviterCommissionConfigID,
			CommissionConfigType: &inviterCommissionConfigType,
			PaymentStatements:    inviterPayments,
		})
	}

	commissionAmountUSD := decimal.NewFromInt(0).String()
	commissionConfigID := uuid.Nil.String()
	payments := []*paymentmwpb.StatementReq{}

	userCommMap, ok := userCoinCommMap[h.UserID.String()]
	if ok {
		for key, commissions := range userCommMap {
			_key := key.String()
			for _, commission := range commissions {
				payments = append(payments, &paymentmwpb.StatementReq{
					PaymentCoinTypeID: &_key,
					Amount:            &commission.PaymentAmount,
					CommissionAmount:  &commission.Amount,
				})
				commissionConfigID = commission.CommissionConfigID
				commissionConfigType = commission.CommissionConfigType
				if h.HasCommission {
					commissionAmountUSD = commission.CommissionAmountUSD
				}
			}
		}
	} else {
		for _, payment := range h.Payments {
			coinTypeID := payment.PaymentCoinTypeID.String()
			amount := payment.Amount.String()
			commissionAmount := "0"
			payments = append(payments, &paymentmwpb.StatementReq{
				PaymentCoinTypeID: &coinTypeID,
				Amount:            &amount,
				CommissionAmount:  &commissionAmount,
			})
		}
	}

	statements = append(statements, &statementmwpb.StatementReq{
		AppID: func() *string {
			id := h.AppID.String()
			return &id
		}(),
		UserID: func() *string {
			id := h.UserID.String()
			return &id
		}(),
		OrderUserID: func() *string {
			id := h.UserID.String()
			return &id
		}(),
		DirectContributorID: func() *string {
			id := h.UserID.String()
			return &id
		}(),
		GoodID: func() *string {
			id := h.GoodID.String()
			return &id
		}(),
		AppGoodID: func() *string {
			id := h.AppGoodID.String()
			return &id
		}(),
		OrderID: func() *string {
			id := h.OrderID.String()
			return &id
		}(),
		GoodCoinTypeID: func() *string {
			id := h.GoodCoinTypeID.String()
			return &id
		}(),
		Units: func() *string {
			units := h.Units.String()
			return &units
		}(),
		GoodValueUSD: func() *string {
			goodValueUSD := h.GoodValueUSD.String()
			return &goodValueUSD
		}(),
		PaymentAmountUSD: func() *string {
			paymentAmountUSD := h.PaymentAmountUSD.String()
			return &paymentAmountUSD
		}(),
		CommissionAmountUSD:  &commissionAmountUSD,
		AppConfigID:          &appConfigID,
		CommissionConfigID:   &commissionConfigID,
		CommissionConfigType: &commissionConfigType,
		PaymentStatements:    payments,
	})

	return statements, nil
}
