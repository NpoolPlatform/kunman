package chain

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/chain"
	chaincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/chain"
	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	entchain "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/chainbase"
)

type queryHandler struct {
	*Handler
	stm   *ent.ChainBaseSelect
	infos []*npool.Chain
	total uint32
}

func (h *queryHandler) selectChainBase(stm *ent.ChainBaseQuery) {
	h.stm = stm.Select(
		entchain.FieldID,
		entchain.FieldEntID,
		entchain.FieldLogo,
		entchain.FieldNativeUnit,
		entchain.FieldAtomicUnit,
		entchain.FieldUnitExp,
		entchain.FieldEnv,
		entchain.FieldChainID,
		entchain.FieldNickname,
		entchain.FieldGasType,
		entchain.FieldCreatedAt,
		entchain.FieldUpdatedAt,
	).Modify(func(s *sql.Selector) {
		t := sql.Table(entchain.Table)
		s.AppendSelect(
			sql.As(t.C(entchain.FieldName), "chain_type"),
		)
	})
}

func (h *queryHandler) queryChainBase(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.ChainBase.Query().Where(entchain.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entchain.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entchain.EntID(*h.EntID))
	}
	h.selectChainBase(stm)
	return nil
}

func (h *queryHandler) queryChainBases(ctx context.Context, cli *ent.Client) error {
	stm, err := chaincrud.SetQueryConds(cli.ChainBase.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectChainBase(stm)
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.GasType = basetypes.GasType(basetypes.GasType_value[info.GasTypeStr])
	}
}

func (h *Handler) GetChain(ctx context.Context) (*npool.Chain, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChainBase(cli); err != nil {
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

func (h *Handler) GetChains(ctx context.Context) ([]*npool.Chain, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryChainBases(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
