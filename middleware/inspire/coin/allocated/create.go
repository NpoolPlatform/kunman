package allocated

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	coinconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coin/config"
	usercoinrewardcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/user/coin/reward"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoinconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coinconfig"
	entusercoinreward "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/usercoinreward"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into coin_allocateds "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "coin_config_id"
	_sql += comma + "coin_type_id"
	_sql += comma + "user_id"
	_sql += comma + "value"
	if h.Extra != nil {
		_sql += comma + "extra"
	}
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
	_sql += fmt.Sprintf("%v'%v' as coin_config_id", comma, *h.CoinConfigID)
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	_sql += fmt.Sprintf("%v'%v' as user_id", comma, *h.UserID)
	_sql += fmt.Sprintf("%v'%v' as value", comma, *h.Value)
	if h.Extra != nil {
		_sql += fmt.Sprintf("%v'%v' as extra", comma, *h.Extra)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from (select * from coin_configs) as di "
	_sql += fmt.Sprintf("where di.ent_id = '%v' and di.app_id = '%v' and di.coin_type_id = '%v' and di.deleted_at=0", *h.CoinConfigID, *h.AppID, *h.CoinTypeID)
	_sql += " limit 1)"
	if h.Extra != nil {
		_sql += " and not exists ("
		_sql += "select 1 from (select * from coin_allocateds) as di "
		_sql += fmt.Sprintf("where di.coin_config_id = '%v' and di.app_id = '%v' and di.coin_type_id = '%v' and di.user_id = '%v' and di.extra = '%v' and di.deleted_at=0",
			*h.CoinConfigID, *h.AppID, *h.CoinTypeID, *h.UserID, *h.Extra)
		_sql += " limit 1)"
	}

	h.sql = _sql
}

func (h *createHandler) createCoinAllocated(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create coinallocated: %v", err)
	}
	return nil
}

func (h *createHandler) createOrUpdateUserCoinReward(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		UserCoinReward.
		Query().
		Where(
			entusercoinreward.AppID(*h.AppID),
			entusercoinreward.UserID(*h.UserID),
			entusercoinreward.CoinTypeID(*h.CoinTypeID),
			entusercoinreward.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return wlog.WrapError(err)
		}
	}
	coinRewards := decimal.NewFromInt(0)

	if info == nil {
		id := uuid.New()
		if _, err := usercoinrewardcrud.CreateSet(
			tx.UserCoinReward.Create(),
			&usercoinrewardcrud.Req{
				EntID:       &id,
				AppID:       h.AppID,
				UserID:      h.UserID,
				CoinTypeID:  h.CoinTypeID,
				CoinRewards: h.Value,
			},
		).Save(ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	}

	coinRewards = info.CoinRewards.Add(*h.Value)
	if _, err := usercoinrewardcrud.UpdateSet(
		info.Update(),
		&usercoinrewardcrud.Req{
			CoinRewards: &coinRewards,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) updateCoinAllocated(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		CoinConfig.
		Query().
		Where(
			entcoinconfig.EntID(*h.CoinConfigID),
			entcoinconfig.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	maxValue, err := decimal.NewFromString(info.MaxValue.String())
	if err != nil {
		return wlog.WrapError(err)
	}
	allocated, err := decimal.NewFromString(info.Allocated.String())
	if err != nil {
		return wlog.WrapError(err)
	}

	allocated = allocated.Add(*h.Value)
	if allocated.Cmp(maxValue) > 0 {
		return wlog.Errorf("invalid value")
	}
	if _, err := coinconfigcrud.UpdateSet(
		info.Update(),
		&coinconfigcrud.Req{
			Allocated: &allocated,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateCoinAllocated(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOrUpdateUserCoinReward(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateCoinAllocated(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createCoinAllocated(_ctx, tx)
	})
}
