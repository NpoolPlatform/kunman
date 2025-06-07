package required

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql  string
	must bool
}

//nolint:goconst
func (h *createHandler) constructSQL() { //nolint:funlen
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into required_app_goods ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "main_app_good_id"
	comma = ", "
	_sql += comma + "required_app_good_id"
	if h.Must != nil {
		_sql += comma + "must"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from ( select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as main_app_good_id", comma, *h.MainAppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as required_app_good_id", comma, *h.RequiredAppGoodID)
	if h.Must != nil {
		_sql += fmt.Sprintf("%v%v as must", comma, *h.Must)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from required_app_goods as rg "
	_sql += fmt.Sprintf(
		"where rg.main_app_good_id = '%v' and rg.required_app_good_id='%v' and deleted_at = 0 ",
		*h.MainAppGoodID,
		*h.RequiredAppGoodID,
	)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from required_app_goods as rg "
	_sql += fmt.Sprintf(
		"where rg.required_app_good_id='%v' and deleted_at = 0",
		*h.MainAppGoodID,
	)
	_sql += " limit 1) and not exists ("
	_sql += "select 1 from required_app_goods as rg "
	_sql += fmt.Sprintf(
		"where rg.main_app_good_id='%v' and deleted_at = 0",
		*h.RequiredAppGoodID,
	)
	_sql += " limit 1) and exists ("
	/**
	 * select good_id_1, good_id_2 from (
		 select agb1.good_id as good_id_1, agb2.good_id as good_id_2 from app_good_bases as agb1, app_good_bases as agb2 where agb1.id=349 and agb2.id=238
	   ) as tmp
	   join good_bases as gb on gb.ent_id = tmp.good_id_1
	   join good_bases as gb2 on gb2.ent_id = tmp.good_id_2
	   join required_goods on main_good_id=good_id_1 and required_good_id=good_id_2;
	*/
	_sql += "select tmp.main_good_id, tmp.required_good_id from ("
	_sql += "select agb1.good_id as main_good_id, agb2.good_id as required_good_id "
	_sql += "from app_good_bases as agb1, app_good_bases as agb2 "
	_sql += fmt.Sprintf("where agb1.ent_id = '%v' and agb2.ent_id = '%v'", *h.MainAppGoodID, *h.RequiredAppGoodID)
	_sql += ") as tmp "
	_sql += "join good_bases as gb1 on gb1.ent_id = tmp.main_good_id "
	_sql += "join good_bases as gb2 on gb2.ent_id = tmp.required_good_id "
	_sql += "join required_goods as rg on rg.main_good_id = tmp.main_good_id and rg.required_good_id = tmp.required_good_id"
	if !h.must {
		_sql += fmt.Sprintf(" where rg.must=%v and rg.deleted_at=0", h.must)
	}
	_sql += ")"
	h.sql = _sql
}

func (h *createHandler) createRequired(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create requireappdgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateRequired(ctx context.Context) error {
	if *h.MainAppGoodID == *h.RequiredAppGoodID {
		return wlog.Errorf("invalid appgoodid")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	handler := &createHandler{
		Handler: h,
		must:    false,
	}
	if h.Must != nil {
		handler.must = *h.Must
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createRequired(_ctx, tx)
	})
}
