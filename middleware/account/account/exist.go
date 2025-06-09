package account

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
)

type existHandler struct {
	*Handler
	stm *ent.AccountQuery
}

func (h *existHandler) queryAccount(cli *ent.Client) {
	h.stm = cli.Account.
		Query().
		Where(
			entaccount.EntID(*h.EntID),
			entaccount.DeletedAt(0),
		)
}

func (h *existHandler) queryAccounts(cli *ent.Client) error {
	stm, err := accountcrud.SetQueryConds(cli.Account.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistAccount(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAccount(cli)
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistAccountConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAccounts(cli); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
