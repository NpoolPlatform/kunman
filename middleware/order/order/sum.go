package order

import (
	"context"

	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entfeeorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type sumHandler struct {
	*baseQueryHandler
}

func (h *sumHandler) sumOrdersPaymentUSD(ctx context.Context, cli *ent.Client) (amount string, err error) {
	if h.stmSelect, err = h.queryOrderBases(cli); err != nil {
		return decimal.NewFromInt(0).String(), wlog.WrapError(err)
	}
	h.queryJoin()
	amounts := []struct {
		AppID  uuid.UUID       `json:"app_id"`
		Amount decimal.Decimal `json:"payment_amount_usd"`
	}{}
	if err := h.stmSelect.GroupBy(
		entorderbase.FieldAppID,
	).Aggregate(func(s *sql.Selector) string {
		t1 := sql.Table(entpowerrental.Table)
		s.LeftJoin(t1).On(
			s.C(entorderbase.FieldEntID),
			t1.C(entpowerrental.FieldOrderID),
		)

		t2 := sql.Table(entfeeorder.Table)
		s.LeftJoin(t2).On(
			s.C(entorderbase.FieldEntID),
			t2.C(entpowerrental.FieldOrderID),
		)

		return sql.As(sql.Sum(
			"ifnull("+t1.C(entpowerrental.FieldPaymentAmountUsd)+", 0) + ifnull("+t2.C(entfeeorder.FieldPaymentAmountUsd)+", 0)",
		), "payment_amount_usd")
	}).Scan(ctx, &amounts); err != nil {
		return decimal.NewFromInt(0).String(), wlog.WrapError(err)
	}
	if len(amounts) > 1 {
		return decimal.NewFromInt(0).String(), wlog.Errorf("invalid paymentamounts")
	}
	if len(amounts) == 0 {
		return decimal.NewFromInt(0).String(), nil
	}
	return amounts[0].Amount.String(), nil
}

func (h *Handler) SumOrdersPaymentUSD(ctx context.Context) (amount string, err error) {
	handler := &sumHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		amount, err = handler.sumOrdersPaymentUSD(_ctx, cli)
		return wlog.WrapError(err)
	}); err != nil {
		return decimal.NewFromInt(0).String(), wlog.WrapError(err)
	}
	return amount, wlog.WrapError(err)
}

func (h *Handler) SubOrdersValueUSD(ctx context.Context) (string, error) {
	return "", nil
}
