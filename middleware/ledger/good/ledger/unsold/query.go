package unsold

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/unsold"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger/unsold"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entunsoldstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/unsoldstatement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.UnsoldStatementSelect
	infos     []*npool.UnsoldStatement
	total     uint32
}

func (h *queryHandler) selectUnsoldStatement(stm *ent.UnsoldStatementQuery) {
	h.stmSelect = stm.Select(
		entunsoldstatement.FieldID,
		entunsoldstatement.FieldEntID,
		entunsoldstatement.FieldGoodID,
		entunsoldstatement.FieldCoinTypeID,
		entunsoldstatement.FieldAmount,
		entunsoldstatement.FieldBenefitDate,
		entunsoldstatement.FieldCreatedAt,
		entunsoldstatement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryUnsoldStatement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}

	stm := cli.UnsoldStatement.Query().Where(entunsoldstatement.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entunsoldstatement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entunsoldstatement.EntID(*h.EntID))
	}
	h.selectUnsoldStatement(stm)
	return nil
}

func (h *queryHandler) queryUnsoldStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.UnsoldStatement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectUnsoldStatement(stm)
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
	}
}

func (h *Handler) GetUnsoldStatement(ctx context.Context) (*npool.UnsoldStatement, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.UnsoldStatement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUnsoldStatement(cli); err != nil {
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

func (h *Handler) GetUnsoldStatements(ctx context.Context) ([]*npool.UnsoldStatement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.UnsoldStatement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUnsoldStatements(ctx, cli); err != nil {
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

func (h *Handler) GetUnsoldStatementOnly(ctx context.Context) (*npool.UnsoldStatement, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryUnsoldStatements(_ctx, cli); err != nil {
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
