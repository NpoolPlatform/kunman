package allocated

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/credit/allocated"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/credit/allocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcreditallocated "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/creditallocated"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CreditAllocatedSelect
	stmCount  *ent.CreditAllocatedSelect
	infos     []*npool.CreditAllocated
	total     uint32
}

func (h *queryHandler) selectCreditAllocated(stm *ent.CreditAllocatedQuery) {
	h.stmSelect = stm.Select(entcreditallocated.FieldID)
}

func (h *queryHandler) queryCreditAllocated(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.CreditAllocated.Query().Where(entcreditallocated.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcreditallocated.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcreditallocated.EntID(*h.EntID))
	}
	h.selectCreditAllocated(stm)
	return nil
}

func (h *queryHandler) queryCreditAllocateds(ctx context.Context, cli *ent.Client) error {
	stm, err := allocatedcrud.SetQueryConds(cli.CreditAllocated.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectCreditAllocated(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entcreditallocated.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcreditallocated.FieldEntID),
			t1.C(entcreditallocated.FieldEntID),
		).
		AppendSelect(
			t1.C(entcreditallocated.FieldEntID),
			t1.C(entcreditallocated.FieldAppID),
			t1.C(entcreditallocated.FieldUserID),
			t1.C(entcreditallocated.FieldValue),
			t1.C(entcreditallocated.FieldExtra),
			t1.C(entcreditallocated.FieldCreatedAt),
			t1.C(entcreditallocated.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Value)
		if err != nil {
			info.Value = decimal.NewFromInt(0).String()
		} else {
			info.Value = amount.String()
		}
	}
}

func (h *Handler) GetCreditAllocated(ctx context.Context) (*npool.CreditAllocated, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCreditAllocated(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
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

func (h *Handler) GetCreditAllocateds(ctx context.Context) ([]*npool.CreditAllocated, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCreditAllocateds(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
