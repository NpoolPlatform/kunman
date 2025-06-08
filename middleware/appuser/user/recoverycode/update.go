package recoverycode

import (
	"context"
	"fmt"

	recoverycodecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/user/recoverycode"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	npool "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user/recoverycode"
)

func (h *Handler) UpdateRecoveryCode(ctx context.Context) (*npool.RecoveryCode, error) {
	code, err := h.GetRecoveryCode(ctx)
	if err != nil {
		return nil, err
	}
	if code == nil {
		return nil, fmt.Errorf("recovery code not found")
	}
	if code.Used {
		return code, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := recoverycodecrud.UpdateSet(
			cli.RecoveryCode.UpdateOneID(*h.ID),
			&recoverycodecrud.Req{
				Used: h.Used,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetRecoveryCode(ctx)
}
