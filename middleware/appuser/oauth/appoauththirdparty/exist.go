package appoauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent"

	appoauththirdpartycrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/oauth/appoauththirdparty"
	entappoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/appoauththirdparty"
)

type existHandler struct {
	*Handler
	stm *ent.AppOAuthThirdPartyQuery
}

func (h *existHandler) queryOAuthThirdParty(cli *ent.Client) {
	h.stm = cli.AppOAuthThirdParty.
		Query().
		Where(
			entappoauththirdparty.EntID(*h.EntID),
			entappoauththirdparty.DeletedAt(0),
		)
}

func (h *existHandler) queryOAuthThirdParties(cli *ent.Client) error {
	stm, err := appoauththirdpartycrud.SetQueryConds(cli.AppOAuthThirdParty.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistOAuthThirdParty(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOAuthThirdParty(cli)
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

func (h *Handler) ExistOAuthThirdPartyConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOAuthThirdParties(cli); err != nil {
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
