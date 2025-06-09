package contract

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entcontract "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	accountSQL         string
	contractAccountSQL string
}

func (h *createHandler) checkBackupAccount(ctx context.Context, tx *ent.Tx) (err error) {
	stm := tx.Contract.Query()

	backup := true
	stm.Where(entcontract.Backup(backup))
	stm.Where(entcontract.GoodID(*h.GoodID))

	if _, err = stm.Exist(ctx); err != nil {
		if ent.IsNotFound(err) {
			h.Backup = &backup
			return nil
		}
		return err
	}
	backup = false
	h.Backup = &backup
	return nil
}

//nolint:goconst
func (h *createHandler) constructCreateContractAccountSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into contracts "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "delegated_staking_id"
	_sql += comma + "account_id"
	if h.Backup != nil {
		_sql += comma + "backup"
	}
	_sql += comma + "contract_operator_type"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as delegated_staking_id", comma, *h.DelegatedStakingID)
	_sql += fmt.Sprintf("%v'%v' as account_id", comma, *h.AccountID)
	if h.Backup != nil {
		_sql += fmt.Sprintf("%v%v as backup", comma, *h.Backup)
	}
	_sql += fmt.Sprintf("%v'%v' as contract_operator_type", comma, *h.ContractOperatorType)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp"
	h.contractAccountSQL = _sql
}

func (h *createHandler) constructCreateaccountSQL() {
	usedFor := basetypes.AccountUsedFor_GoodBenefit
	privateKey := true
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into accounts "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "coin_type_id"
	comma = ", "
	_sql += comma + "address"
	_sql += comma + "used_for"
	_sql += comma + "platform_hold_private_key"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.AccountID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as address", comma, *h.Address)
	_sql += fmt.Sprintf("%v'%v' as used_for", comma, usedFor)
	_sql += fmt.Sprintf("%v%v as platform_hold_private_key", comma, privateKey)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from accounts "
	_sql += fmt.Sprintf(
		"where coin_type_id = '%v' and address = '%v' and deleted_at = 0",
		*h.CoinTypeID,
		*h.Address,
	)
	_sql += " limit 1)"
	h.accountSQL = _sql
}

func (h *createHandler) createAccount(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.accountSQL)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create account: %v", err)
	}
	return nil
}

func (h *createHandler) createContractAccount(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.contractAccountSQL)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create delegatedtaking: %v", err)
	}
	return nil
}

func (h *Handler) CreateAccount(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	if h.AccountID == nil {
		h.AccountID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkBackupAccount(_ctx, tx); err != nil {
			return err
		}
		handler.constructCreateaccountSQL()
		handler.constructCreateContractAccountSQL()
		if err := handler.createAccount(_ctx, tx); err != nil {
			return err
		}
		return handler.createContractAccount(_ctx, tx)
	})
}
