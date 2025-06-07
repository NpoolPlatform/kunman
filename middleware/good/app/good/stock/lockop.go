package appstock

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appstocklockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/stock/lock"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappstocklock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstocklock"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LockStock struct {
	EntID         *uuid.UUID
	AppGoodID     *uuid.UUID
	Locked        *decimal.Decimal
	AppSpotLocked *decimal.Decimal
	LockID        *uuid.UUID
}

type lockopHandler struct {
	*Handler
	locks []*ent.AppStockLock
	state *types.AppStockLockState
}

func (h *lockopHandler) getLocks(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		locks, err := cli.
			AppStockLock.
			Query().
			Where(
				entappstocklock.ExLockID(*h.LockID),
				entappstocklock.DeletedAt(0),
			).
			All(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.locks = locks
		if len(locks) > 0 {
			h.EntID = &locks[0].AppStockID
		}
		return nil
	})
}

func (h *lockopHandler) checkState() error {
	stateMap := map[types.AppStockLockState][]types.AppStockLockState{
		types.AppStockLockState_AppStockLocked: {
			types.AppStockLockState_AppStockWaitStart,
			types.AppStockLockState_AppStockRollback,
			types.AppStockLockState_AppStockCanceled,
		},
		types.AppStockLockState_AppStockWaitStart: {types.AppStockLockState_AppStockInService, types.AppStockLockState_AppStockChargeBack},
		types.AppStockLockState_AppStockInService: {types.AppStockLockState_AppStockExpired, types.AppStockLockState_AppStockChargeBack},
	}
	for _, lock := range h.locks {
		validState := false
		for _, state := range stateMap[types.AppStockLockState(types.AppStockLockState_value[lock.LockState])] {
			if state == *h.state {
				validState = true
				break
			}
		}
		if !validState {
			return wlog.Errorf("invalid state")
		}
	}
	return nil
}

func (h *lockopHandler) updateLocks(ctx context.Context, tx *ent.Tx) error {
	if len(h.locks) == 0 {
		return wlog.Errorf("invalid locks")
	}
	if err := h.checkState(); err != nil {
		return wlog.WrapError(err)
	}
	lockIDs := []uint32{}
	for _, lock := range h.locks {
		lockIDs = append(lockIDs, lock.ID)
	}
	stm := tx.
		AppStockLock.
		Update().
		Where(
			entappstocklock.IDIn(lockIDs...),
			entappstocklock.DeletedAt(0),
		).
		SetLockState(h.state.String())
	switch *h.state {
	case types.AppStockLockState_AppStockRollback:
		fallthrough //nolint
	case types.AppStockLockState_AppStockChargeBack:
		stm.SetChargeBackState(h.locks[0].LockState)
	}
	if _, err := stm.Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

//nolint:goconst
func (h *lockopHandler) constructCreateSQL(req *appstocklockcrud.Req, checkExist bool) string {
	now := uint32(time.Now().Unix())

	_sql := "insert into app_stock_locks "
	_sql += "("
	_sql += "ent_id, "
	_sql += "app_good_id, "
	_sql += "app_stock_id, "
	_sql += "units, "
	_sql += "app_spot_units, "
	_sql += "ex_lock_id, "
	_sql += "created_at, "
	_sql += "updated_at, "
	_sql += "deleted_at"
	_sql += ")"
	_sql += " select * from (select "
	_sql += fmt.Sprintf("'%v' as ent_id, ", *req.EntID)
	_sql += fmt.Sprintf("'%v' as app_good_id, ", *req.AppGoodID)
	_sql += fmt.Sprintf("'%v' as app_stock_id, ", *req.AppStockID)
	_sql += fmt.Sprintf("'%v' as units, ", *req.Units)
	if req.AppSpotUnits != nil {
		_sql += fmt.Sprintf("'%v' as app_spot_units, ", *req.AppSpotUnits)
	} else {
		_sql += "'0' as app_spot_units, "
	}
	_sql += fmt.Sprintf("'%v' as ex_lock_id, ", *req.ExLockID)
	_sql += fmt.Sprintf("%v as created_at, ", now)
	_sql += fmt.Sprintf("%v as updated_at, ", now)
	_sql += "0 as deleted_at"
	_sql += ") as tmp "
	_sql += "where "
	if checkExist {
		_sql += "not exists ("
		_sql += "select 1 from app_stock_locks "
		_sql += fmt.Sprintf("where ent_id = '%v' ", *req.EntID)
		_sql += fmt.Sprintf("or ex_lock_id = '%v' ", *req.ExLockID)
		_sql += " limit 1) and "
	}
	_sql += "exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *req.AppGoodID)
	_sql += "limit 1) and ("
	_sql += "exists ("
	_sql += "select 1 from app_stocks "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *req.AppStockID)
	_sql += "limit 1) or "
	_sql += "exists ("
	_sql += "select 1 from app_mining_good_stocks "
	_sql += fmt.Sprintf("where ent_id = '%v' ", *req.AppStockID)
	_sql += "limit 1)"
	_sql += ")"

	return _sql
}

func (h *lockopHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail lock stock: %v", err)
	}
	return nil
}

func (h *lockopHandler) createLocks(ctx context.Context, tx *ent.Tx) error {
	if len(h.Stocks) == 0 {
		return h.execSQL(ctx, tx, h.constructCreateSQL(&appstocklockcrud.Req{
			EntID:        h.LockID,
			AppStockID:   h.EntID,
			AppGoodID:    h.AppGoodID,
			Units:        h.Locked,
			AppSpotUnits: h.AppSpotLocked,
			ExLockID:     h.LockID,
		}, true))
	}
	for i, stock := range h.Stocks {
		if err := h.execSQL(ctx, tx, h.constructCreateSQL(&appstocklockcrud.Req{
			EntID:        func() *uuid.UUID { uid := uuid.New(); return &uid }(),
			AppStockID:   stock.EntID,
			AppGoodID:    stock.AppGoodID,
			Units:        stock.Locked,
			AppSpotUnits: stock.AppSpotLocked,
			ExLockID:     h.LockID,
		}, i == 0)); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *lockopHandler) lock2Stocks() (reqs []*LockStock) {
	for _, lock := range h.locks {
		reqs = append(reqs, &LockStock{
			EntID:         &lock.AppStockID,
			AppGoodID:     &lock.AppGoodID,
			Locked:        &lock.Units,
			AppSpotLocked: &lock.AppSpotUnits,
			LockID:        &lock.ExLockID,
		})
	}
	return
}
