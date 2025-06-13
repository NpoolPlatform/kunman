package profit

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger/profit"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/simulate/ledger/profit"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entprofit "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulateprofit"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.SimulateProfitSelect
	infos     []*npool.Profit
	total     uint32
}

func (h *queryHandler) selectProfit(stm *ent.SimulateProfitQuery) {
	h.stmSelect = stm.Select(
		entprofit.FieldID,
		entprofit.FieldEntID,
		entprofit.FieldAppID,
		entprofit.FieldUserID,
		entprofit.FieldCoinTypeID,
		entprofit.FieldIncoming,
		entprofit.FieldCreatedAt,
		entprofit.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryProfit(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.SimulateProfit.Query().Where(entprofit.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entprofit.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entprofit.EntID(*h.EntID))
	}
	h.selectProfit(stm)
	return nil
}

func (h *queryHandler) queryProfits(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.SimulateProfit.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectProfit(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		incoming := decimal.NewFromInt(0).String()
		if _incoming, err := decimal.NewFromString(info.Incoming); err == nil {
			incoming = _incoming.String()
		}
		info.Incoming = incoming
	}
}

func (h *Handler) GetProfit(ctx context.Context) (*npool.Profit, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Profit{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryProfit(cli); err != nil {
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

func (h *Handler) GetProfits(ctx context.Context) ([]*npool.Profit, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Profit{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryProfits(ctx, cli); err != nil {
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

func (h *Handler) GetProfitOnly(ctx context.Context) (*npool.Profit, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryProfits(_ctx, cli); err != nil {
			return err
		}
		_, err := handler.stmSelect.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		handler.formalize()
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("to many record")
	}
	return handler.infos[0], nil
}
