package orderstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entorderstatement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderstatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderStatementSelect
	infos    []*npool.Statement
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CommissionConfigType = types.CommissionConfigType(types.CommissionConfigType_value[info.CommissionConfigTypeStr])
		info.GoodValueUSD = func() string {
			amount := decimal.RequireFromString(info.GoodValueUSD)
			return amount.String()
		}()
		info.PaymentAmountUSD = func() string {
			amount := decimal.RequireFromString(info.PaymentAmountUSD)
			return amount.String()
		}()
		info.CommissionAmountUSD = func() string {
			amount := decimal.RequireFromString(info.CommissionAmountUSD)
			return amount.String()
		}()
		info.Units = func() string {
			units := decimal.RequireFromString(info.Units)
			return units.String()
		}()
		info.DirectContributorID = func() string {
			id, err := uuid.Parse(info.DirectContributorID)
			if err != nil {
				return uuid.Nil.String()
			}
			return id.String()
		}()
	}
}

func (h *Handler) GetStatement(ctx context.Context) (*npool.Statement, error) {
	info := &npool.Statement{}
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		_info, err := h._getStatement(ctx, cli)
		if err != nil {
			return err
		}
		info = _info
		return nil
	})
	return info, err
}

func (h *Handler) GetStatementWithTx(ctx context.Context, tx *ent.Tx) (*npool.Statement, error) {
	return h._getStatement(ctx, tx.Client())
}

func (h *Handler) _getStatement(ctx context.Context, cli *ent.Client) (*npool.Statement, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := handler.queryOrderStatement(cli); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.queryJoin()
	if err := handler.scan(ctx); err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderStatements(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryOrderStatements(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entorderstatement.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
