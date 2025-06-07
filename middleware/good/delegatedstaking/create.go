//nolint:dupl
package delegatedstaking

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	rewardcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/reward"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sqlDelegatedStaking string
	sqlGoodBase         string
}

func (h *createHandler) constructGoodBaseSQL(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), false),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, true),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, true),
		goodbase1.WithName(h.GoodBaseReq.Name, true),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, true),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, true),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, true),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
		goodbase1.WithState(h.GoodBaseReq.State, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlGoodBase = handler.ConstructCreateSQL()
	return nil
}

//nolint:goconst
func (h *createHandler) constructDelegatedStakingSQL() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into delegated_stakings "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "contract_code_url"
	_sql += comma + "contract_code_branch"
	_sql += comma + "contract_state"
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
	_sql += fmt.Sprintf("%v'%v' as contract_code_url", comma, *h.ContractCodeURL)
	_sql += fmt.Sprintf("%v'%v' as contract_code_branch", comma, *h.ContractCodeBranch)
	_sql += fmt.Sprintf("%v'%v' as contract_state", comma, h.ContractState.String())

	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	h.sqlDelegatedStaking = _sql
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create delegatedstaking %v: %v", sql, err)
	}
	return nil
}

func (h *createHandler) createDelegatedStaking(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlDelegatedStaking)
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlGoodBase)
}

func (h *createHandler) createReward(ctx context.Context, tx *ent.Tx) error {
	if _, err := rewardcrud.CreateSet(
		tx.GoodReward.Create(),
		&rewardcrud.Req{
			GoodID: h.GoodID,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) formalizeEntIDs() {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	if h.GoodBaseReq.EntID == nil {
		h.GoodBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.GoodID = h.GoodBaseReq.EntID
	}
}

func (h *Handler) CreateDelegatedStaking(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}

	h.GoodBaseReq.BenefitType = func() *types.BenefitType { e := types.BenefitType_BenefitTypeContract; return &e }()
	h.GoodBaseReq.State = func() *types.GoodState { e := types.GoodState_GoodStateReady; return &e }()
	h.ContractState = func() *types.ContractState { e := types.ContractState_ContractWaitDeployment; return &e }()
	handler.formalizeEntIDs()
	handler.constructDelegatedStakingSQL()
	if err := handler.constructGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.createReward(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createDelegatedStaking(_ctx, tx)
	})
}
