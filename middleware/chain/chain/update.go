package chain

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*Handler
	sql       string
	chainType string
	env       string
	chainID   string
}

func (h *updateHandler) constructSQL() error {
	now := uint32(time.Now().Unix())
	set := "set "
	_sql := "update chain_bases "

	if h.ChainType != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.ChainType)
		set = ""
	}
	if h.NativeUnit != nil {
		_sql += fmt.Sprintf("%vnative_unit = '%v', ", set, *h.NativeUnit)
		set = ""
	}
	if h.AtomicUnit != nil {
		_sql += fmt.Sprintf("%vatomic_unit = '%v', ", set, *h.AtomicUnit)
		set = ""
	}
	if h.UnitExp != nil {
		_sql += fmt.Sprintf("%vunit_exp = '%v', ", set, *h.UnitExp)
		set = ""
	}
	if h.ENV != nil {
		_sql += fmt.Sprintf("%venv = '%v', ", set, *h.ENV)
		set = ""
	}
	if h.GasType != nil {
		_sql += fmt.Sprintf("%vgas_type = '%v', ", set, *h.GasType)
		set = ""
	}
	if h.Logo != nil {
		_sql += fmt.Sprintf("%vlogo = '%v', ", set, *h.Logo)
		set = ""
	}
	if h.ChainID != nil {
		_sql += fmt.Sprintf("%vchain_id = '%v', ", set, *h.ChainID)
		set = ""
	}
	if h.Nickname != nil {
		_sql += fmt.Sprintf("%vnickname = '%v', ", set, *h.Nickname)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from chain_bases) as cb "
	_sql += fmt.Sprintf(
		"where cb.name = '%v' and cb.env = '%v' and cb.id != %v and deleted_at = 0",
		h.chainType,
		h.env,
		*h.ID,
	)
	if h.ChainID != nil {
		_sql += fmt.Sprintf(
			" and chain_id = '%v'",
			h.chainID,
		)
	}
	_sql += " limit 1)"

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateChain(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update chain: %v", err)
	}
	return nil
}

func (h *Handler) UpdateChain(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetChain(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid chain")
	}

	if h.ChainType == nil {
		handler.chainType = info.ChainType
	} else {
		handler.chainType = *h.ChainType
	}
	if h.ENV == nil {
		handler.env = info.ENV
	} else {
		handler.env = *h.ENV
	}
	if h.ChainID == nil {
		handler.chainID = info.ChainID
	} else {
		handler.chainID = *h.ChainID
	}

	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateChain(_ctx, tx)
	})
}
