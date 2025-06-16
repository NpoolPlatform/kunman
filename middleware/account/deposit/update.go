package deposit

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	depositcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/deposit"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entdeposit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/deposit"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
)

func (h *Handler) UpdateAccount(ctx context.Context) (*npool.Account, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		deposit, err := tx.Deposit.
			Query().
			Where(
				entdeposit.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if deposit == nil {
			return fmt.Errorf("invalid deposit")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(deposit.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		var _scannableAt *uint32
		if account.Locked && h.Locked != nil && !*h.Locked {
			scannableAt := uint32(time.Now().Unix()) + timedef.SecondsPerHour
			_scannableAt = &scannableAt
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

		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Active:   h.Active,
				Locked:   h.Locked,
				LockedBy: h.LockedBy,
				Blocked:  h.Blocked,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := depositcrud.UpdateSet(
			deposit.Update(),
			&depositcrud.Req{
				CollectingTID: h.CollectingTID,
				Incoming:      &incoming,
				Outcoming:     &outcoming,
				ScannableAt:   _scannableAt,
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
