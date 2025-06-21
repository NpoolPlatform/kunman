package tx

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	enttx "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/tran"

	txcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/tx"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stm   *ent.TranSelect
	infos []*npool.Tx
	total uint32
}

func (h *queryHandler) selectTx(stm *ent.TranQuery) {
	h.stm = stm.Select(
		enttx.FieldID,
		enttx.FieldEntID,
		enttx.FieldCoinTypeID,
		enttx.FieldFromAccountID,
		enttx.FieldToAccountID,
		enttx.FieldAmount,
		enttx.FieldFeeAmount,
		enttx.FieldState,
		enttx.FieldChainTxID,
		enttx.FieldType,
		enttx.FieldExtra,
		enttx.FieldCreatedAt,
		enttx.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTx(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Tran.Query().Where(enttx.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttx.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttx.EntID(*h.EntID))
	}
	h.selectTx(stm)
	return nil
}

func (h *queryHandler) queryTxs(ctx context.Context, cli *ent.Client) error {
	stm, err := txcrud.SetQueryConds(cli.Tran.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectTx(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t).
		On(
			s.C(enttx.FieldCoinTypeID),
			t.C(entcoinbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoin(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Type = basetypes.TxType(basetypes.TxType_value[info.TypeStr])
		info.State = basetypes.TxState(basetypes.TxState_value[info.StateStr])
		amount, err := decimal.NewFromString(info.Amount)
		if err != nil {
			info.Amount = decimal.NewFromInt(0).String()
		} else {
			info.Amount = amount.String()
		}
		amount, err = decimal.NewFromString(info.FeeAmount)
		if err != nil {
			info.FeeAmount = decimal.NewFromInt(0).String()
		} else {
			info.FeeAmount = amount.String()
		}
	}
}

func (h *Handler) GetTx(ctx context.Context) (*npool.Tx, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTx(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetTxs(ctx context.Context) ([]*npool.Tx, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTxs(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(enttx.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetTxOnly(ctx context.Context) (*npool.Tx, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTxs(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(0).
			Limit(2).
			Order(ent.Desc(enttx.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("invalid tx")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}
