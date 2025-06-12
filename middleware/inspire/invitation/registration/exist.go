package registration

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	registrationcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/invitation/registration"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entregistration "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/registration"
)

func (h *Handler) ExistRegistration(ctx context.Context) (bool, error) {
	exist := false
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Registration.
			Query().
			Where(
				entregistration.ID(*h.ID),
				entregistration.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}

func (h *Handler) ExistRegistrationConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := registrationcrud.SetQueryConds(cli.Registration.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
