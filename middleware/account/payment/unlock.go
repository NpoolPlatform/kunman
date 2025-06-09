package payment

import (
	"context"
	"fmt"
	"time"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"
	paymentcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/payment"
	"github.com/NpoolPlatform/kunman/middleware/account/db"
	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entaccount "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/account"
	entpayment "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/payment"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
)

func (h *Handler) UnlockAccount(ctx context.Context) (*npool.Account, error) {
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		payment, err := tx.Payment.
			Query().
			Where(
				entpayment.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if payment == nil {
			return fmt.Errorf("invalid payment")
		}

		account, err := tx.Account.
			Query().
			Where(
				entaccount.EntID(payment.AccountID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if !account.Locked {
			return fmt.Errorf("not locked")
		}

		const coolDown = uint32(60 * 60)
		availableAt := uint32(time.Now().Unix()) + coolDown
		h.AvailableAt = &availableAt

		locked := false
		if _, err := accountcrud.UpdateSet(
			account.Update(),
			&accountcrud.Req{
				Locked: &locked,
			},
		).Save(_ctx); err != nil {
			return err
		}
		if _, err := paymentcrud.UpdateSet(
			payment.Update(),
			&paymentcrud.Req{
				AvailableAt: h.AvailableAt,
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
