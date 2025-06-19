package executor

import (
	"context"
	"encoding/json"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/commission/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	achievementorderpaymentstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"
	orderlockmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/lock"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	achievementorderpaymentstatementmw "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order/payment"
	orderlockmw "github.com/NpoolPlatform/kunman/middleware/order/order/lock"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent        chan interface{}
	notif             chan interface{}
	done              chan interface{}
	statements        []*achievementorderpaymentstatementmwpb.Statement
	commissionLocks   map[string]*orderlockmwpb.OrderLock
	commissionRevokes map[string]*types.CommissionRevoke
}

func (h *orderHandler) getOrderCommissionLock(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.commissionLocks = map[string]*orderlockmwpb.OrderLock{}

	conds := &orderlockmwpb.Conds{
		OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID},
		LockType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ordertypes.OrderLockType_LockCommission)},
	}

	for {
		handler, err := orderlockmw.NewHandler(
			ctx,
			orderlockmw.WithConds(conds),
			orderlockmw.WithOffset(offset),
			orderlockmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		locks, _, err := handler.GetOrderLocks(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(locks) == 0 {
			break
		}
		for _, lock := range locks {
			h.commissionLocks[lock.UserID] = lock
		}
		offset += limit
	}
	return nil
}

func (h *orderHandler) getOrderAchievement(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &achievementorderpaymentstatementmwpb.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID},
	}

	for {
		handler, err := achievementorderpaymentstatementmw.NewHandler(
			ctx,
			achievementorderpaymentstatementmw.WithConds(conds),
			achievementorderpaymentstatementmw.WithOffset(offset),
			achievementorderpaymentstatementmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		statements, _, err := handler.GetStatements(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(statements) == 0 {
			return nil
		}
		h.statements = append(h.statements, statements...)
		offset += limit
	}
}

func (h *orderHandler) constructCommissionRevoke() error {
	h.commissionRevokes = map[string]*types.CommissionRevoke{}

	for _, statement := range h.statements {
		amount, err := decimal.NewFromString(statement.CommissionAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			continue
		}
		extra := struct {
			AppID                   string   `json:"AppID"`
			UserID                  string   `json:"UserID"`
			AchievementStatementIDs []string `json:"AchievementStatementIDs"`
			CancelOrder             bool     `json:"CancelOrder"`
		}{
			CancelOrder: true,
		}
		revoke, ok := h.commissionRevokes[statement.UserID]
		if !ok {
			lock, ok := h.commissionLocks[statement.UserID]
			if !ok {
				return wlog.Errorf("invalid commission lock")
			}
			revoke = &types.CommissionRevoke{
				LockID: lock.EntID,
			}
			extra.AppID = statement.AppID
			extra.UserID = statement.UserID
			extra.AchievementStatementIDs = []string{statement.EntID}
		} else {
			if err := json.Unmarshal([]byte(revoke.IOExtra), &extra); err != nil {
				return wlog.WrapError(err)
			}
			extra.AchievementStatementIDs = append(extra.AchievementStatementIDs, statement.EntID)
		}
		_extra, err := json.Marshal(&extra)
		if err != nil {
			return wlog.WrapError(err)
		}
		revoke.IOExtra = string(_extra)
		revoke.StatementIDs = append(revoke.StatementIDs, uuid.NewString())
		h.commissionRevokes[statement.UserID] = revoke
	}
	return nil
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Order", h.PowerRentalOrder,
			"CommissionStatements", h.statements,
			"CommissionLocks", h.commissionLocks,
			"CommissionRevokes", h.commissionRevokes,
			"Error", *err,
		)
	}

	persistentOrder := &types.PersistentPowerRentalOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		CommissionRevokes: func() (revokes []*types.CommissionRevoke) {
			for _, revoke := range h.commissionRevokes {
				revokes = append(revokes, revoke)
			}
			return
		}(),
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getOrderCommissionLock(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getOrderAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.constructCommissionRevoke(); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
