package orderstatement

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	orderpaymentstatementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/statement/order/payment"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appcommissionconfig"
	entappconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appconfig"
	entappgoodcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appgoodcommissionconfig"
	entcommission "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/commission"
	entorderstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderstatement"

	goodachievement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/good"
	goodcoinachievement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/good/coin"
	achievementuser1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*achievementQueryHandler
	sql                          string
	sqlCreateGoodAchievement     string
	sqlUpdateGoodAchievement     string
	sqlCreateGoodCoinAchievement string
	sqlUpdateGoodCoinAchievement string
	sqlCreateAchievementUser     string
	sqlUpdateAchievementUser     string
	selfOrder                    bool
	selfUnits                    decimal.Decimal
	selfAmountUSD                decimal.Decimal
	selfCommissionAmountUSD      decimal.Decimal
	inviteeConsumeAmount         decimal.Decimal
	updated                      bool
	updatable                    bool
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := fmt.Sprintf("insert into %v ", entorderstatement.Table)
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "user_id"
	_sql += comma + "good_id"
	_sql += comma + "app_good_id"
	_sql += comma + "order_id"
	_sql += comma + "order_user_id"
	_sql += comma + "direct_contributor_id"
	_sql += comma + "good_coin_type_id"
	_sql += comma + "units"
	_sql += comma + "good_value_usd"
	_sql += comma + "payment_amount_usd"
	_sql += comma + "commission_amount_usd"
	_sql += comma + "app_config_id"
	_sql += comma + "commission_config_id"
	_sql += comma + "commission_config_type"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	_sql += fmt.Sprintf("%v'%v' as order_id", comma, *h.OrderID)
	_sql += fmt.Sprintf("%v'%v' as order_user_id", comma, *h.OrderUserID)
	_sql += fmt.Sprintf("%v'%v' as direct_contributor_id", comma, *h.DirectContributorID)
	_sql += fmt.Sprintf("%v'%v' as good_coin_type_id", comma, *h.GoodCoinTypeID)
	_sql += fmt.Sprintf("%v'%v' as units", comma, *h.Units)
	_sql += fmt.Sprintf("%v'%v' as good_value_usd", comma, *h.GoodValueUSD)
	_sql += fmt.Sprintf("%v'%v' as payment_amount_usd", comma, *h.PaymentAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as commission_amount_usd", comma, *h.CommissionAmountUSD)
	_sql += fmt.Sprintf("%v'%v' as app_config_id", comma, *h.AppConfigID)
	commissionConfigID := uuid.Nil
	if h.CommissionConfigID != nil {
		commissionConfigID = *h.CommissionConfigID
	}
	_sql += fmt.Sprintf("%v'%v' as commission_config_id", comma, commissionConfigID)
	_sql += fmt.Sprintf("%v'%v' as commission_config_type", comma, h.CommissionConfigType.String())
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += fmt.Sprintf("select 1 from %v ", entorderstatement.Table)
	_sql += fmt.Sprintf(
		"where user_id = '%v' and order_id = '%v' and deleted_at = 0 ",
		*h.UserID,
		*h.OrderID,
	)
	_sql += "limit 1)"

	if h.AppConfigID != nil && *h.AppConfigID != uuid.Nil {
		_sql += " and exists ("
		_sql += fmt.Sprintf("select 1 from %v ", entappconfig.Table)
		_sql += fmt.Sprintf(
			"where app_id = '%v' and ent_id = '%v' and end_at = 0 and deleted_at = 0 ",
			*h.AppID,
			*h.AppConfigID,
		)
		_sql += "limit 1)"
	}

	if h.CommissionConfigID != nil && *h.CommissionConfigID != uuid.Nil {
		switch *h.CommissionConfigType {
		case types.CommissionConfigType_LegacyCommissionConfig:
			_sql += " and exists ("
			_sql += fmt.Sprintf("select 1 from %v ", entcommission.Table)
			_sql += fmt.Sprintf(
				"where app_id = '%v' and ent_id = '%v' and user_id = '%v' and good_id = '%v' and app_good_id = '%v' and deleted_at = 0 ",
				*h.AppID,
				*h.CommissionConfigID,
				*h.UserID,
				*h.GoodID,
				*h.AppGoodID,
			)
		case types.CommissionConfigType_AppCommissionConfig:
			_sql += " and exists ("
			_sql += fmt.Sprintf("select 1 from %v ", entappcommissionconfig.Table)
			_sql += fmt.Sprintf(
				"where app_id = '%v' and ent_id = '%v' and deleted_at = 0 ",
				*h.AppID,
				*h.CommissionConfigID,
			)
		case types.CommissionConfigType_AppGoodCommissionConfig:
			_sql += " and exists ("
			_sql += fmt.Sprintf("select 1 from %v ", entappgoodcommissionconfig.Table)
			_sql += fmt.Sprintf(
				"where app_id = '%v' and ent_id = '%v' and good_id = '%v' and app_good_id = '%v' and deleted_at = 0 ",
				*h.AppID,
				*h.CommissionConfigID,
				*h.GoodID,
				*h.AppGoodID,
			)
		}
		_sql += "limit 1)"
	}

	h.sql = _sql
}

func (h *createHandler) constructCreateGoodAchievementSQL(ctx context.Context) error {
	handler, err := goodachievement1.NewHandler(
		ctx,
		goodachievement1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		goodachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodachievement1.WithGoodID(func() *string { s := h.GoodID.String(); return &s }(), true),
		goodachievement1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
		goodachievement1.WithTotalAmountUSD(func() *string { s := h.GoodValueUSD.String(); return &s }(), true),
		goodachievement1.WithSelfAmountUSD(func() *string { s := h.selfAmountUSD.String(); return &s }(), true),
		goodachievement1.WithTotalUnits(func() *string { s := h.Units.String(); return &s }(), true),
		goodachievement1.WithSelfUnits(func() *string { s := h.selfUnits.String(); return &s }(), true),
		goodachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateGoodAchievement = handler.ConstructCreateSQL()
	h.sqlUpdateGoodAchievement = handler.ConstructUpdateSQL()
	return nil
}

func (h *createHandler) constructCreateGoodCoinAchievementSQL(ctx context.Context) error {
	handler, err := goodcoinachievement1.NewHandler(
		ctx,
		goodcoinachievement1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		goodcoinachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodcoinachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodcoinachievement1.WithGoodCoinTypeID(func() *string { s := h.GoodCoinTypeID.String(); return &s }(), true),
		goodcoinachievement1.WithTotalAmountUSD(func() *string { s := h.GoodValueUSD.String(); return &s }(), true),
		goodcoinachievement1.WithSelfAmountUSD(func() *string { s := h.selfAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithTotalUnits(func() *string { s := h.Units.String(); return &s }(), true),
		goodcoinachievement1.WithSelfUnits(func() *string { s := h.selfUnits.String(); return &s }(), true),
		goodcoinachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateGoodCoinAchievement = handler.ConstructCreateSQL()
	h.sqlUpdateGoodCoinAchievement = handler.ConstructUpdateSQL()
	return nil
}

func (h *createHandler) constructCreateAchievementUserSQL(ctx context.Context) error {
	handler, err := achievementuser1.NewHandler(
		ctx,
		achievementuser1.WithEntID(func() *string { s := uuid.NewString(); return &s }(), true),
		achievementuser1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		achievementuser1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		achievementuser1.WithTotalCommission(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		achievementuser1.WithSelfCommission(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
		achievementuser1.WithDirectConsumeAmount(func() *string { s := h.selfAmountUSD.String(); return &s }(), true),
		achievementuser1.WithInviteeConsumeAmount(func() *string { s := h.inviteeConsumeAmount.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlCreateAchievementUser = handler.ConstructCreateSQL()
	h.sqlUpdateAchievementUser = handler.ConstructUpdateSQL()
	return nil
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	h.updatable = true
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n > 1 {
		return wlog.Errorf("fail create orderstatement: %v", err)
	}
	if n == 1 {
		h.updated = true
		return nil
	}
	if n == 0 {
		return wlog.WrapError(cruder.ErrCreateNothing)
	}
	return nil
}

func (h *createHandler) createOrderStatement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *createHandler) createPaymentStatements(ctx context.Context, tx *ent.Tx) error {
	for _, req := range h.PaymentStatementReqs {
		if _, err := orderpaymentstatementcrud.CreateSet(
			tx.OrderPaymentStatement.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) updateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateGoodAchievement)
}

func (h *createHandler) createOrUpdateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	err := h.execSQL(ctx, tx, h.sqlCreateGoodAchievement)
	if err == nil {
		return nil
	}
	if !wlog.Equal(err, cruder.ErrCreateNothing) {
		return wlog.WrapError(err)
	}
	return h.updateGoodAchievement(ctx, tx)
}

func (h *createHandler) updateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateGoodCoinAchievement)
}

func (h *createHandler) createOrUpdateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	err := h.execSQL(ctx, tx, h.sqlCreateGoodCoinAchievement)
	if err == nil {
		return nil
	}
	if !wlog.Equal(err, cruder.ErrCreateNothing) {
		return wlog.WrapError(err)
	}
	return h.updateGoodCoinAchievement(ctx, tx)
}

func (h *createHandler) updateAchievementUser(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateAchievementUser)
}

func (h *createHandler) createOrUpdateAchievementUser(ctx context.Context, tx *ent.Tx) error {
	err := h.execSQL(ctx, tx, h.sqlCreateAchievementUser)
	if err == nil {
		return nil
	}
	if !wlog.Equal(err, cruder.ErrCreateNothing) {
		return wlog.WrapError(err)
	}
	return h.updateAchievementUser(ctx, tx)
}

func (h *createHandler) validateCommissionAmount() error {
	if h.CommissionConfigID == nil || h.CommissionConfigID.String() == uuid.Nil.String() {
		if h.CommissionAmountUSD.Cmp(decimal.NewFromInt(0)) > 0 {
			return wlog.Errorf("commission config id mismatch commission amount usd")
		}
	}
	for _, req := range h.PaymentStatementReqs {
		if h.CommissionAmountUSD.Cmp(decimal.NewFromInt(0)) <= 0 {
			if req.CommissionAmount.Cmp(decimal.NewFromInt(0)) > 0 {
				return wlog.Errorf("invalid commission amount")
			}
		}
	}
	if *h.CommissionConfigType == types.CommissionConfigType_WithoutCommissionConfig {
		if h.CommissionAmountUSD.Cmp(decimal.NewFromInt(0)) > 0 {
			return wlog.Errorf("commission amount usd mismatch commission config type")
		}
	}
	return nil
}

func (h *createHandler) validateDirectContributorID() error {
	if *h.UserID == *h.OrderUserID && *h.DirectContributorID != *h.UserID {
		return wlog.Errorf("invalid direct contributor id")
	}
	return nil
}

func (h *Handler) CreateStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	for _, req := range h.PaymentStatementReqs {
		req.StatementID = h.EntID
	}
	handler := &createHandler{
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		selfOrder: *h.OrderUserID == *h.UserID,
		selfUnits: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.Units
			}
			return decimal.NewFromInt(0)
		}(),
		selfAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.GoodValueUSD
			}
			return decimal.NewFromInt(0)
		}(),
		selfCommissionAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.CommissionAmountUSD
			}
			return decimal.NewFromInt(0)
		}(),
		inviteeConsumeAmount: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return decimal.NewFromInt(0)
			}
			return *h.GoodValueUSD
		}(),
	}

	if err := handler.validateCommissionAmount(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.validateDirectContributorID(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.getAchievementWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}

	handler.constructSQL()
	if err := handler.constructCreateGoodAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCreateGoodCoinAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructCreateAchievementUserSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if err := handler.createOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createPaymentStatements(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrUpdateGoodAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrUpdateGoodCoinAchievement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.createOrUpdateAchievementUser(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if handler.updatable && !handler.updated {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	return nil
}

func (h *Handler) CreateStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateStatementWithTx(_ctx, tx)
	})
}
