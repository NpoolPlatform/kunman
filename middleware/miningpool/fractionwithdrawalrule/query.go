package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawalrule"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	fractionwithdrawalruleent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawalrule"

	fractionwithdrawalrulecrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawalrule"
)

type queryHandler struct {
	*Handler
	stm   *ent.FractionWithdrawalRuleSelect
	infos []*npool.FractionWithdrawalRule
	total uint32
}

func (h *queryHandler) selectFractionWithdrawalRule(stm *ent.FractionWithdrawalRuleQuery) {
	h.stm = stm.Select(
		fractionwithdrawalruleent.FieldID,
		fractionwithdrawalruleent.FieldCreatedAt,
		fractionwithdrawalruleent.FieldUpdatedAt,
		fractionwithdrawalruleent.FieldEntID,
		fractionwithdrawalruleent.FieldPoolCoinTypeID,
		fractionwithdrawalruleent.FieldWithdrawInterval,
		fractionwithdrawalruleent.FieldPayoutThreshold,
		fractionwithdrawalruleent.FieldLeastWithdrawalAmount,
		fractionwithdrawalruleent.FieldWithdrawFee,
	)
}

func (h *queryHandler) queryFractionWithdrawalRule(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.FractionWithdrawalRule.Query().Where(fractionwithdrawalruleent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(fractionwithdrawalruleent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(fractionwithdrawalruleent.EntID(*h.EntID))
	}
	h.selectFractionWithdrawalRule(stm)
	return nil
}

func (h *queryHandler) queryFractionWithdrawalRules(ctx context.Context, cli *ent.Client) error {
	stm, err := fractionwithdrawalrulecrud.SetQueryConds(cli.FractionWithdrawalRule.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := fractionwithdrawalrulecrud.SetQueryConds(cli.FractionWithdrawalRule.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmCount.Modify(h.queryJoinCoinAndPool)
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectFractionWithdrawalRule(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.MiningPoolType = mpbasetypes.MiningPoolType(mpbasetypes.MiningPoolType_value[info.MiningPoolTypeStr])
		info.CoinType = basetypes.CoinType(basetypes.CoinType_value[info.CoinTypeStr])
	}
}

func (h *Handler) GetFractionWithdrawalRule(ctx context.Context) (*npool.FractionWithdrawalRule, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionWithdrawalRule(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin(handler.stm.FractionWithdrawalRuleQuery)
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetFractionWithdrawalRules(ctx context.Context) ([]*npool.FractionWithdrawalRule, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionWithdrawalRules(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin(handler.stm.FractionWithdrawalRuleQuery)
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(fractionwithdrawalruleent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
