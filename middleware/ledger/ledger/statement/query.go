package statement

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entstatement "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/statement"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.StatementSelect
	infos     []*npool.Statement
	total     uint32
}

func (h *queryHandler) selectStatement(stm *ent.StatementQuery) {
	h.stmSelect = stm.Select(
		entstatement.FieldID,
		entstatement.FieldEntID,
		entstatement.FieldAppID,
		entstatement.FieldUserID,
		entstatement.FieldCoinTypeID,
		entstatement.FieldIoType,
		entstatement.FieldIoSubType,
		entstatement.FieldAmount,
		entstatement.FieldIoExtraV1,
		entstatement.FieldCreatedAt,
		entstatement.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryStatement(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Statement.Query().Where(entstatement.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entstatement.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entstatement.EntID(*h.EntID))
	}
	h.selectStatement(stm)
	return nil
}

func (h *queryHandler) queryStatements(ctx context.Context, cli *ent.Client) error {
	stm, err := crud.SetQueryConds(cli.Statement.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectStatement(stm)
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

		info.IOType = basetypes.IOType(basetypes.IOType_value[info.IOTypeStr])
		info.IOSubType = basetypes.IOSubType(basetypes.IOSubType_value[info.IOSubTypeStr])
	}
}

func (h *Handler) GetStatement(ctx context.Context) (*npool.Statement, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatement(cli); err != nil {
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

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	handler := &queryHandler{
		Handler: h,
		infos:   []*npool.Statement{},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(ctx, cli); err != nil {
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

func (h *Handler) GetStatementOnly(ctx context.Context) (*npool.Statement, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStatements(_ctx, cli); err != nil {
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
