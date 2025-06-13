package coupon

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entcouponwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/couponwithdraw"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CouponWithdrawSelect
	infos     []*npool.CouponWithdraw
	total     uint32
}

func (h *queryHandler) selectCouponWithdraw(stm *ent.CouponWithdrawQuery) {
	h.stmSelect = stm.Select(
		entcouponwithdraw.FieldID,
		entcouponwithdraw.FieldEntID,
		entcouponwithdraw.FieldAppID,
		entcouponwithdraw.FieldUserID,
		entcouponwithdraw.FieldCoinTypeID,
		entcouponwithdraw.FieldAllocatedID,
		entcouponwithdraw.FieldReviewID,
		entcouponwithdraw.FieldState,
		entcouponwithdraw.FieldAmount,
		entcouponwithdraw.FieldCreatedAt,
		entcouponwithdraw.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCouponWithdraw(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.CouponWithdraw.Query().Where(entcouponwithdraw.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcouponwithdraw.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcouponwithdraw.EntID(*h.EntID))
	}
	h.selectCouponWithdraw(stm)
	return nil
}

func (h *queryHandler) queryCouponWithdraws(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.CouponWithdraw.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectCouponWithdraw(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount := decimal.NewFromInt(0).String()
		if _amount, err := decimal.NewFromString(info.Amount); err == nil {
			amount = _amount.String()
		}
		info.Amount = amount
		info.State = basetypes.WithdrawState(basetypes.WithdrawState_value[info.StateStr])
	}
}

func (h *Handler) GetCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CouponWithdraw{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCouponWithdraw(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCouponWithdraws(ctx context.Context) ([]*npool.CouponWithdraw, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.CouponWithdraw{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCouponWithdraws(ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
