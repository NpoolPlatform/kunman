package goodbase

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

//nolint:goconst,funlen
func (h *Handler) ConstructCreateSQL() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into good_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_type"
	comma = ", "
	_sql += comma + "benefit_type"
	_sql += comma + "name"
	_sql += comma + "service_start_at"
	_sql += comma + "start_mode"
	if h.TestOnly != nil {
		_sql += comma + "test_only"
	}
	if h.BenefitIntervalHours != nil {
		_sql += comma + "benefit_interval_hours"
	}
	if h.Purchasable != nil {
		_sql += comma + "purchasable"
	}
	if h.Online != nil {
		_sql += comma + "online"
	}
	if h.State != nil {
		_sql += comma + "state"
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
	_sql += fmt.Sprintf("%v'%v' as good_type", comma, h.GoodType.String())
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as benefit_type", comma, h.BenefitType.String())
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.Name)
	_sql += fmt.Sprintf("%v%v as service_start_at", comma, *h.ServiceStartAt)
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.TestOnly != nil {
		_sql += fmt.Sprintf("%v%v as test_only", comma, *h.TestOnly)
	}
	if h.BenefitIntervalHours != nil {
		_sql += fmt.Sprintf("%v'%v' as benefit_interval_hours", comma, *h.BenefitIntervalHours)
	}
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%v%v as purchasable", comma, *h.Purchasable)
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%v%v as online", comma, *h.Online)
	}
	if h.State != nil {
		_sql += fmt.Sprintf("%v'%v' as state", comma, h.State.String())
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_bases "
	_sql += fmt.Sprintf("where name = '%v' and deleted_at = 0", *h.Name)
	_sql += " limit 1)"
	return _sql
}

//nolint:gocyclo,funlen
func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.ID == nil && h.EntID == nil {
		return "", wlog.Errorf("invalid goodid")
	}
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update good_bases "
	if h.GoodType != nil {
		_sql += fmt.Sprintf("%vgood_type = '%v', ", set, h.GoodType.String())
		set = ""
	}
	if h.BenefitType != nil {
		_sql += fmt.Sprintf("%vbenefit_type = '%v', ", set, h.BenefitType.String())
		set = ""
	}
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.ServiceStartAt != nil {
		_sql += fmt.Sprintf("%vservice_start_at = %v, ", set, *h.ServiceStartAt)
		set = ""
	}
	if h.StartMode != nil {
		_sql += fmt.Sprintf("%vstart_mode = '%v', ", set, h.StartMode.String())
		set = ""
	}
	if h.TestOnly != nil {
		_sql += fmt.Sprintf("%vtest_only = %v, ", set, *h.TestOnly)
		set = ""
	}
	if h.BenefitIntervalHours != nil {
		_sql += fmt.Sprintf("%vbenefit_interval_hours = %v, ", set, *h.BenefitIntervalHours)
		set = ""
	}
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%vpurchasable = %v, ", set, *h.Purchasable)
		set = ""
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%vonline = %v, ", set, *h.Online)
		set = ""
	}
	if h.State != nil {
		_sql += fmt.Sprintf("%vstate = '%v', ", set, h.State.String())
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where"
	if h.ID != nil {
		_sql += fmt.Sprintf("%v id = %v ", whereAnd, *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
	}
	if h.Name != nil {
		_sql += " and not exists ("
		_sql += "select 1 from ("
		_sql += "select * from good_bases as gb "
		_sql += fmt.Sprintf("where gb.name = '%v' ", *h.Name)
		if h.ID != nil {
			_sql += fmt.Sprintf("and gb.id != %v", *h.ID)
		} else if h.EntID != nil {
			_sql += fmt.Sprintf("and gb.ent_id != '%v'", *h.EntID)
		}
		_sql += " limit 1) as tmp)"
	}

	return _sql, nil
}
