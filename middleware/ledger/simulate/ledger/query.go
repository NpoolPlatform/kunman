package ledger

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/simulate/ledger"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulateledger"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.SimulateLedgerSelect
	infos     []*npool.Ledger
	total     uint32
}

func (h *queryHandler) selectLedger(stm *ent.SimulateLedgerQuery) {
	h.stmSelect = stm.Select(
		entledger.FieldID,
		entledger.FieldEntID,
		entledger.FieldAppID,
		entledger.FieldUserID,
		entledger.FieldCoinTypeID,
		entledger.FieldIncoming,
		entledger.FieldOutcoming,
		entledger.FieldCreatedAt,
		entledger.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryLedger(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.SimulateLedger.Query().Where(entledger.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entledger.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entledger.EntID(*h.EntID))
	}
	h.selectLedger(stm)
	return nil
}

func (h *queryHandler) queryLedgers(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.SimulateLedger.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectLedger(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		incoming, err := decimal.NewFromString(info.Incoming)
		if err != nil {
			info.Incoming = decimal.NewFromInt(0).String()
		} else {
			info.Incoming = incoming.String()
		}

		outcoming, err := decimal.NewFromString(info.Outcoming)
		if err != nil {
			info.Outcoming = decimal.NewFromInt(0).String()
		} else {
			info.Outcoming = outcoming.String()
		}
	}
}

func (h *Handler) GetLedger(ctx context.Context) (*npool.Ledger, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Ledger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLedger(cli); err != nil {
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

func (h *Handler) GetLedgers(ctx context.Context) ([]*npool.Ledger, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Ledger{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLedgers(ctx, cli); err != nil {
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

func (h *Handler) GetLedgerOnly(ctx context.Context) (*npool.Ledger, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLedgers(_ctx, cli); err != nil {
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
