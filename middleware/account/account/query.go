package account

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
)

type queryHandler struct {
	*Handler
	stm   *ent.AccountSelect
	infos []*npool.Account
	total uint32
}

func (h *queryHandler) selectAccount(stm *ent.AccountQuery) {
	h.stm = stm.Select(
		entaccount.FieldID,
		entaccount.FieldEntID,
		entaccount.FieldCoinTypeID,
		entaccount.FieldAddress,
		entaccount.FieldUsedFor,
		entaccount.FieldPlatformHoldPrivateKey,
		entaccount.FieldActive,
		entaccount.FieldLocked,
		entaccount.FieldLockedBy,
		entaccount.FieldBlocked,
		entaccount.FieldCreatedAt,
		entaccount.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryAccount(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Account.Query().Where(entaccount.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entaccount.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entaccount.EntID(*h.EntID))
	}
	h.selectAccount(stm)
	return nil
}

func (h *queryHandler) queryAccounts(ctx context.Context, cli *ent.Client) error {
	stm, err := accountcrud.SetQueryConds(cli.Account.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectAccount(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.UsedFor = basetypes.AccountUsedFor(basetypes.AccountUsedFor_value[info.UsedForStr])
		info.LockedBy = basetypes.AccountLockedBy(basetypes.AccountLockedBy_value[info.LockedByStr])
	}
}

func (h *Handler) GetAccount(ctx context.Context) (*npool.Account, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccount(cli); err != nil {
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

func (h *Handler) GetAccounts(ctx context.Context) ([]*npool.Account, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(_ctx, cli); err != nil {
			return err
		}

		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entaccount.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
