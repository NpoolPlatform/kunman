//nolint:dupl
package delegatedstaking

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	extrainfocrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/extrainfo"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*delegatedstakingAppGoodQueryHandler
	sqlAppDelegatedStaking string
	sqlAppGoodBase         string
}

func (h *createHandler) constructAppGoodBaseSQL(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), false),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), true),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), true),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, true),
		appgoodbase1.WithPurchasable(h.AppGoodBaseReq.Purchasable, false),
		appgoodbase1.WithEnableProductPage(h.AppGoodBaseReq.EnableProductPage, false),
		appgoodbase1.WithProductPage(h.AppGoodBaseReq.ProductPage, false),
		appgoodbase1.WithOnline(h.AppGoodBaseReq.Online, false),
		appgoodbase1.WithVisible(h.AppGoodBaseReq.Visible, false),
		appgoodbase1.WithDisplayIndex(h.AppGoodBaseReq.DisplayIndex, false),
		appgoodbase1.WithBanner(h.AppGoodBaseReq.Banner, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlAppGoodBase = handler.ConstructCreateSQL()
	return nil
}

//nolint:goconst
func (h *createHandler) constructAppDelegatedStakingSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_delegated_stakings "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "service_start_at"
	_sql += comma + "start_mode"
	if h.EnableSetCommission != nil {
		_sql += comma + "enable_set_commission"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v%v as service_start_at", comma, *h.ServiceStartAt)
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%v%v as enable_set_commission", comma, *h.EnableSetCommission)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from ("
	_sql += "select * from app_delegated_stakings as apr "
	_sql += fmt.Sprintf("where app_good_id = '%v' and deleted_at=0", *h.AppGoodID)
	_sql += " limit 1) as tmp)"
	_sql += "and exists ("
	_sql += "select 1 from delegated_stakings "
	_sql += fmt.Sprintf("where good_id = '%v' and deleted_at=0", *h.AppGoodBaseReq.GoodID)
	_sql += " limit 1)"
	h.sqlAppDelegatedStaking = _sql
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create appDelegatedStaking: %v", err)
	}
	return nil
}

func (h *createHandler) createAppDelegatedStaking(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppDelegatedStaking)
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppGoodBase)
}

func (h *createHandler) createExtraInfo(ctx context.Context, tx *ent.Tx) error {
	if _, err := extrainfocrud.CreateSet(tx.ExtraInfo.Create(), h.ExtraInfoReq).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) formalizeEntIDs() {
	if h.AppGoodBaseReq.EntID == nil {
		h.AppGoodBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.ExtraInfoReq.AppGoodID = h.AppGoodBaseReq.EntID
		h.AppGoodID = h.AppGoodBaseReq.EntID
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

//nolint:gocyclo
func (h *Handler) CreateDelegatedStaking(ctx context.Context) error {
	handler := &createHandler{
		delegatedstakingAppGoodQueryHandler: &delegatedstakingAppGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireDelegatedStakingGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	handler.formalizeEntIDs()
	if h.AppGoodBaseReq.Purchasable != nil && (!handler._ent.goodBase.Purchasable && *h.AppGoodBaseReq.Purchasable) {
		return wlog.Errorf("invalid purchasable")
	}
	if h.ServiceStartAt == nil {
		h.ServiceStartAt = func() *uint32 { u := handler._ent.GoodServiceStartAt(); return &u }()
	}
	if h.StartMode == nil {
		h.StartMode = func() *types.GoodStartMode { u := handler._ent.GoodStartMode(); return &u }()
	} else {
		switch handler._ent.GoodStartMode() {
		case types.GoodStartMode_GoodStartModeTBD:
			switch *h.StartMode {
			case types.GoodStartMode_GoodStartModeTBD:
			case types.GoodStartMode_GoodStartModePreset:
			default:
				return wlog.Errorf("invalid startmode")
			}
		case types.GoodStartMode_GoodStartModeInstantly:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModeNextDay:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModePreset:
			fallthrough //nolint
		case types.GoodStartMode_GoodStartModeWithParent:
			if *h.StartMode != handler._ent.GoodStartMode() {
				return wlog.Errorf("invalid startmode")
			}
		default:
			return wlog.Errorf("invalid startmode")
		}
	}

	handler.constructAppDelegatedStakingSQL()
	if err := handler.constructAppGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			fmt.Println("--fail createAppGoodBase")
			return wlog.WrapError(err)
		}
		if err := handler.createExtraInfo(_ctx, tx); err != nil {
			fmt.Println("--fail createExtraInfo")
			return wlog.WrapError(err)
		}
		return handler.createAppDelegatedStaking(_ctx, tx)
	})
}
