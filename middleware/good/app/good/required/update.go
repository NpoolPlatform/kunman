package required

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/required"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructCheckMustSQL() {
	must := true
	_sql := "select 1 from required_goods where main_good_id=( "
	_sql += "select tmp.main_good_id from ( "
	_sql += "select agb1.good_id as main_good_id, agb2.good_id as required_good_id "
	_sql += "from app_good_bases as agb1, app_good_bases as agb2 "
	_sql += "where agb1.ent_id = ( "
	_sql += fmt.Sprintf("select rag1.main_app_good_id as ent_id from required_app_goods as rag1 where id='%v' ", *h.ID)
	_sql += ") and agb2.ent_id = ( "
	_sql += fmt.Sprintf("select rag2.required_app_good_id as ent_id from required_app_goods as rag2 where id='%v' ", *h.ID)
	_sql += ")) as tmp "
	_sql += "join good_bases as gb1 on gb1.ent_id = tmp.main_good_id "
	_sql += "join good_bases as gb2 on gb2.ent_id = tmp.required_good_id "
	_sql += ") and required_good_id=( "
	_sql += "select tmp.required_good_id from ( "
	_sql += "select agb1.good_id as main_good_id, agb2.good_id as required_good_id "
	_sql += "from app_good_bases as agb1, app_good_bases as agb2 "
	_sql += "where agb1.ent_id = ( "
	_sql += fmt.Sprintf("select rag1.main_app_good_id as ent_id from required_app_goods as rag1 where id='%v' ", *h.ID)
	_sql += ") and agb2.ent_id = ( "
	_sql += fmt.Sprintf("select rag2.required_app_good_id as ent_id from required_app_goods as rag2 where id='%v' ", *h.ID)
	_sql += ")) as tmp "
	_sql += "join good_bases as gb1 on gb1.ent_id = tmp.main_good_id "
	_sql += "join good_bases as gb2 on gb2.ent_id = tmp.required_good_id "
	_sql += fmt.Sprintf(") and must=%v and deleted_at=0", must)

	h.sql = _sql
}

func (h *updateHandler) checkMust(ctx context.Context, tx *ent.Tx) error {
	if h.Must == nil || *h.Must {
		return nil
	}

	rows, err := tx.QueryContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
	}
	if err != nil || count != 0 {
		return wlog.Errorf("fail update required must: %v", err)
	}
	return nil
}

func (h *Handler) UpdateRequired(ctx context.Context) error {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid required")
	}
	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	handler.constructCheckMustSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkMust(ctx, tx); err != nil {
			return err
		}
		if _, err := requiredcrud.UpdateSet(
			tx.RequiredAppGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				Must: h.Must,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
