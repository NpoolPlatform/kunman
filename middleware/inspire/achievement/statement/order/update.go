package orderstatement

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodachievement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/good"
	goodcoinachievement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/good/coin"
	achievementuser1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommission "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/commission"
	entorderpaymentstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderpaymentstatement"
	entorderstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderstatement"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*achievementQueryHandler
	selfOrder                    bool
	selfCommissionAmountUSD      decimal.Decimal
	statement                    *ent.OrderStatement
	payments                     map[uuid.UUID]*ent.OrderPaymentStatement
	sql                          string
	sqlUpdateGoodAchievement     string
	sqlUpdateGoodCoinAchievement string
	sqlUpdateAchievementUser     string
}

func (h *updateHandler) constructUpdateSQL() {
	now := time.Now().Unix()
	sql := fmt.Sprintf(
		`update %v set updated_at = %v`,
		entorderstatement.Table,
		now,
	)
	if h.CommissionAmountUSD != nil {
		sql += fmt.Sprintf(
			`, commission_amount_usd = commission_amount_usd + %v`,
			*h.CommissionAmountUSD,
		)
	}
	if h.AppConfigID != nil {
		sql += fmt.Sprintf(
			`, app_config_id = '%v'`,
			*h.AppConfigID,
		)
	}
	if h.CommissionConfigID != nil {
		sql += fmt.Sprintf(
			`, commission_config_id = '%v'`,
			*h.CommissionConfigID,
		)
	}
	sql += " where deleted_at = 0"

	if h.ID != nil {
		sql += fmt.Sprintf(
			" and id = %v ",
			*h.ID,
		)
	}
	if h.EntID != nil {
		sql += fmt.Sprintf(
			" and ent_id = '%v' ",
			*h.EntID,
		)
	}
	h.sql = sql
}

func (h *updateHandler) updateOrderStatement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n > 1 {
		return wlog.Errorf("fail update: %v", err)
	}
	if n == 0 {
		return wlog.WrapError(cruder.ErrCreateNothing)
	}
	return nil
}

func (h *updateHandler) constructUpdateGoodAchievementSQL(ctx context.Context) error {
	handler, err := goodachievement1.NewHandler(
		ctx,
		goodachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodachievement1.WithGoodID(func() *string { s := h.GoodID.String(); return &s }(), true),
		goodachievement1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
		goodachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlUpdateGoodAchievement = handler.ConstructUpdateSQL()
	return nil
}

func (h *updateHandler) updateGoodAchievement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateGoodAchievement)
}

func (h *updateHandler) constructUpdateGoodCoinAchievementSQL(ctx context.Context) error {
	handler, err := goodcoinachievement1.NewHandler(
		ctx,
		goodcoinachievement1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		goodcoinachievement1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		goodcoinachievement1.WithGoodCoinTypeID(func() *string { s := h.GoodCoinTypeID.String(); return &s }(), true),
		goodcoinachievement1.WithTotalCommissionUSD(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		goodcoinachievement1.WithSelfCommissionUSD(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlUpdateGoodCoinAchievement = handler.ConstructUpdateSQL()
	return nil
}

func (h *updateHandler) updateGoodCoinAchievement(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateGoodCoinAchievement)
}

func (h *updateHandler) constructUpdateAchievementUserSQL(ctx context.Context) error {
	handler, err := achievementuser1.NewHandler(
		ctx,
		achievementuser1.WithAppID(func() *string { s := h.AppID.String(); return &s }(), true),
		achievementuser1.WithUserID(func() *string { s := h.UserID.String(); return &s }(), true),
		achievementuser1.WithTotalCommission(func() *string { s := h.CommissionAmountUSD.String(); return &s }(), true),
		achievementuser1.WithSelfCommission(func() *string { s := h.selfCommissionAmountUSD.String(); return &s }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlUpdateAchievementUser = handler.ConstructUpdateSQL()
	return nil
}

func (h *updateHandler) updateAchievementUser(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlUpdateAchievementUser)
}

func (h *updateHandler) updatePaymentStatements(ctx context.Context, tx *ent.Tx) error {
	if h.statement.CommissionConfigType != types.CommissionConfigType_LegacyCommissionConfig.String() {
		return nil
	}

	for _, req := range h.PaymentStatementReqs {
		dbPayment, ok := h.payments[*req.EntID]
		if !ok {
			return wlog.Errorf("invalid payment")
		}
		if dbPayment.Amount.Cmp(*req.Amount) != 0 {
			return wlog.Errorf("invalid payment amount")
		}
		if dbPayment.PaymentCoinTypeID != *req.PaymentCoinTypeID {
			return wlog.Errorf("invalid payment cointypeid")
		}
		if dbPayment.CommissionAmount.Cmp(decimal.NewFromInt(0)) != 0 {
			return wlog.Errorf("permission denied")
		}
		if _, err := tx.
			OrderPaymentStatement.
			UpdateOneID(dbPayment.ID).
			SetCommissionAmount(*req.CommissionAmount).
			Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	// update commission config id even commission amount is 0
	if err := h.updateOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if h.CommissionAmountUSD.Cmp(decimal.NewFromInt(0)) > 0 {
		if err := h.updateGoodAchievement(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updateGoodCoinAchievement(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updateAchievementUser(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *updateHandler) requireOrderPaymentStatements(ctx context.Context, tx *ent.Tx) error {
	ids := []uuid.UUID{}
	for _, payment := range h.PaymentStatementReqs {
		if payment.EntID == nil {
			return wlog.Errorf("invalid payment ent id")
		}
		ids = append(ids, *payment.EntID)
	}

	payments, err := tx.
		OrderPaymentStatement.
		Query().
		Where(
			entorderpaymentstatement.EntIDIn(ids...),
			entorderpaymentstatement.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(payments) != len(h.PaymentStatementReqs) {
		return wlog.Errorf("payment statements mismatch")
	}
	for _, payment := range payments {
		h.payments[payment.EntID] = payment
	}
	return nil
}

func (h *updateHandler) requireOrderStatement(ctx context.Context, tx *ent.Tx) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invaild id or entid")
	}
	stm := tx.
		OrderStatement.
		Query().
		Where(
			entorderstatement.AppID(*h.AppID),
			entorderstatement.UserID(*h.UserID),
			entorderstatement.OrderID(*h.OrderID),
			entorderstatement.OrderUserID(*h.OrderUserID),
			entorderstatement.GoodCoinTypeID(*h.GoodCoinTypeID),
			entorderstatement.CommissionConfigType(types.CommissionConfigType_LegacyCommissionConfig.String()),
			entorderstatement.DeletedAt(0),
		)
	if h.ID != nil {
		stm.Where(entorderstatement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderstatement.EntID(*h.EntID))
	}

	statement, err := stm.Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.statement = statement
	h.ID = &statement.ID
	h.EntID = &statement.EntID
	return nil
}

func (h *updateHandler) requireCommission(ctx context.Context, tx *ent.Tx) error {
	if h.CommissionAmountUSD.Cmp(decimal.NewFromInt(0)) > 0 {
		if _, err := tx.
			Commission.
			Query().
			Where(
				entcommission.EntID(*h.CommissionConfigID),
				entcommission.AppID(*h.AppID),
				entcommission.UserID(*h.UserID),
				entcommission.AppGoodID(*h.AppGoodID),
				entcommission.EndAt(0),
				entcommission.DeletedAt(0),
			).
			Only(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *updateHandler) validateCommissionAmount() error {
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

func (h *Handler) UpdateStatementWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &updateHandler{
		payments: map[uuid.UUID]*ent.OrderPaymentStatement{},
		achievementQueryHandler: &achievementQueryHandler{
			Handler: h,
		},
		selfOrder: *h.OrderUserID == *h.UserID,
		selfCommissionAmountUSD: func() decimal.Decimal {
			if *h.OrderUserID == *h.UserID {
				return *h.CommissionAmountUSD
			}
			return decimal.NewFromInt(0)
		}(),
	}
	if err := handler.validateCommissionAmount(); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.requireCommission(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.requireOrderStatement(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.requireOrderPaymentStatements(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	handler.constructUpdateSQL()
	if err := handler.constructUpdateGoodAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructUpdateGoodCoinAchievementSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructUpdateAchievementUserSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.updatePaymentStatements(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateStatement(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateStatementWithTx(_ctx, tx)
	})
}
