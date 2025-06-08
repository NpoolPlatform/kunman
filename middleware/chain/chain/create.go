package chain

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst,funlen
func (h *createHandler) constructSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into chain_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "name"
	comma = ", "
	_sql += comma + "native_unit"
	_sql += comma + "atomic_unit"
	_sql += comma + "unit_exp"
	_sql += comma + "env"
	_sql += comma + "gas_type"
	if h.Logo != nil {
		_sql += comma + "logo"
	}
	if h.ChainID != nil {
		_sql += comma + "chain_id"
	}
	if h.Nickname != nil {
		_sql += comma + "nickname"
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
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.ChainType)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as native_unit", comma, *h.NativeUnit)
	_sql += fmt.Sprintf("%v'%v' as atomic_unit", comma, *h.AtomicUnit)
	_sql += fmt.Sprintf("%v'%v' as unit_exp", comma, *h.UnitExp)
	_sql += fmt.Sprintf("%v'%v' as env", comma, *h.ENV)
	_sql += fmt.Sprintf("%v'%v' as gas_type", comma, *h.GasType)
	if h.Logo != nil {
		_sql += fmt.Sprintf("%v'%v' as logo", comma, *h.Logo)
	}
	if h.ChainID != nil {
		_sql += fmt.Sprintf("%v'%v' as chain_id", comma, *h.ChainID)
	}
	if h.Nickname != nil {
		_sql += fmt.Sprintf("%v'%v' as nickname", comma, *h.Nickname)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from chain_bases "
	_sql += fmt.Sprintf(
		"where name = '%v' and env = '%v' and deleted_at = 0",
		*h.ChainType,
		*h.ENV,
	)
	if h.ChainID != nil {
		_sql += fmt.Sprintf(
			" and chain_id = '%v'",
			*h.ChainID,
		)
	}
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createChainBase(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create devicetype: %v", err)
	}
	return nil
}

func (h *Handler) CreateChain(ctx context.Context) error {
	if h.ChainType == nil {
		return fmt.Errorf("invalid chaintype")
	}
	if h.ENV == nil {
		return fmt.Errorf("invalid env")
	}

	handler := &createHandler{
		Handler: h,
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createChainBase(_ctx, tx)
	})
}
