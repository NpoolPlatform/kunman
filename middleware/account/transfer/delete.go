package transfer

import (
	"context"
	"fmt"
	"time"

	transfercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/transfer"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	enttransfer "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/transfer"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/transfer"
)

func (h *Handler) DeleteTransfer(ctx context.Context) (*npool.Transfer, error) {
	info, err := h.GetTransfer(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		user, err := tx.Transfer.
			Query().
			Where(
				enttransfer.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("invalid user")
		}

		if _, err := transfercrud.UpdateSet(
			user.Update(),
			&transfercrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
