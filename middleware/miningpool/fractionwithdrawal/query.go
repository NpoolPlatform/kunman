package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	fractionwithdrawalent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawal"

	fractionwithdrawalcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawal"
)

type queryHandler struct {
	*Handler
	stm   *ent.FractionWithdrawalSelect
	infos []*npool.FractionWithdrawal
	total uint32
}

func (h *queryHandler) selectFractionWithdrawal(stm *ent.FractionWithdrawalQuery) {
	h.stm = stm.Select(
		fractionwithdrawalent.FieldID,
		fractionwithdrawalent.FieldEntID,
		fractionwithdrawalent.FieldAppID,
		fractionwithdrawalent.FieldUserID,
		fractionwithdrawalent.FieldOrderUserID,
		fractionwithdrawalent.FieldCoinTypeID,
		fractionwithdrawalent.FieldFractionWithdrawState,
		fractionwithdrawalent.FieldWithdrawAt,
		fractionwithdrawalent.FieldPromisePayAt,
		fractionwithdrawalent.FieldMsg,
		fractionwithdrawalent.FieldCreatedAt,
		fractionwithdrawalent.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryFractionWithdrawal(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.FractionWithdrawal.Query().Where(fractionwithdrawalent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(fractionwithdrawalent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(fractionwithdrawalent.EntID(*h.EntID))
	}
	h.selectFractionWithdrawal(stm)
	return nil
}

func (h *queryHandler) queryFractionWithdrawals(ctx context.Context, cli *ent.Client) error {
	stm, err := fractionwithdrawalcrud.SetQueryConds(cli.FractionWithdrawal.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}

	stmCount, err := fractionwithdrawalcrud.SetQueryConds(cli.FractionWithdrawal.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stmCount.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)

	h.selectFractionWithdrawal(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.FractionWithdrawState = basetypes.FractionWithdrawState(basetypes.FractionWithdrawState_value[info.FractionWithdrawStateStr])
	}
}

func (h *Handler) GetFractionWithdrawal(ctx context.Context) (*npool.FractionWithdrawal, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionWithdrawal(cli); err != nil {
			return wlog.WrapError(err)
		}
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

func (h *Handler) GetFractionWithdrawals(ctx context.Context) ([]*npool.FractionWithdrawal, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFractionWithdrawals(ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(fractionwithdrawalent.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
