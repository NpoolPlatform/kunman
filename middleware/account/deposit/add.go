//nolint:dupl
package deposit

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
	depositcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/deposit"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entdeposit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/deposit"

	"github.com/shopspring/decimal"
)

func (h *Handler) AddBalance(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		deposit, err := tx.Deposit.
			Query().
			Where(
				entdeposit.ID(*h.ID),
				entdeposit.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if deposit == nil {
			return fmt.Errorf("invalid deposit")
		}

		incoming := deposit.Incoming
		if h.Incoming != nil {
			incoming = incoming.Add(*h.Incoming)
		}
		outcoming := deposit.Outcoming
		if h.Outcoming != nil {
			outcoming = outcoming.Add(*h.Outcoming)
		}

		if incoming.Cmp(outcoming) < 0 {
			return fmt.Errorf("incoming (%v) < outcoming (%v)", incoming, outcoming)
		}
		if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid incoming")
		}
		if outcoming.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid outcoming")
		}

		if _, err := depositcrud.UpdateSet(
			deposit.Update(),
			&depositcrud.Req{
				Incoming:  &incoming,
				Outcoming: &outcoming,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAccount(ctx)
}
