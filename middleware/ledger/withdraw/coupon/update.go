package coupon

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	ledgercrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	statementcrud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger/statement"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	entcouponwithdraw "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/couponwithdraw"
	entledger "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledger"
)

type updateHandler struct {
	*Handler
	couponwithdraw *ent.CouponWithdraw
}

func (h *updateHandler) checkCouponWithdraw(ctx context.Context) error {
	return db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			CouponWithdraw.
			Query().
			Where(
				entcouponwithdraw.ID(*h.ID),
				entcouponwithdraw.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.couponwithdraw = info
		return nil
	})
}

func (h *updateHandler) createOrUpdateLedger(ctx context.Context, tx *ent.Tx) error {
	if *h.State != types.WithdrawState_Approved {
		return nil
	}

	info, err := tx.
		Ledger.
		Query().
		Where(
			entledger.AppID(h.couponwithdraw.AppID),
			entledger.UserID(h.couponwithdraw.UserID),
			entledger.CurrencyID(h.couponwithdraw.CoinTypeID),
			entledger.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		zero := decimal.NewFromInt(0)
		if _, err = ledgercrud.CreateSet(tx.Ledger.Create(), &ledgercrud.Req{
			AppID:      &h.couponwithdraw.AppID,
			UserID:     &h.couponwithdraw.UserID,
			CurrencyID: &h.couponwithdraw.CoinTypeID,
			Incoming:   &h.couponwithdraw.Amount,
			Spendable:  &h.couponwithdraw.Amount,
			Outcoming:  &zero,
			Locked:     &zero,
		}).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	incoming := h.couponwithdraw.Amount
	stm, err := ledgercrud.UpdateSetWithValidate(
		info,
		&ledgercrud.Req{
			Incoming:  &incoming,
			Spendable: &incoming,
		},
	)
	if err != nil {
		return err
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateCouponWithdraw(ctx context.Context, tx *ent.Tx) error {
	if _, err := crud.UpdateSet(
		tx.CouponWithdraw.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) createStatement(ctx context.Context, tx *ent.Tx) error {
	if *h.State != types.WithdrawState_Approved {
		return nil
	}
	ioExtra := fmt.Sprintf(
		`{"CouponWithdrawID":"%v","AllocatedID":"%v"}`,
		h.couponwithdraw.EntID,
		h.couponwithdraw.AllocatedID.String(),
	)

	ioType := types.IOType_Incoming
	ioSubType := types.IOSubType_RandomCouponCash
	if _, err := statementcrud.CreateSet(
		tx.Statement.Create(),
		&statementcrud.Req{
			AppID:        &h.couponwithdraw.AppID,
			UserID:       &h.couponwithdraw.UserID,
			CurrencyID:   &h.couponwithdraw.CoinTypeID,
			CurrencyType: types.CurrencyType_CurrencyCrypto.Enum(),
			Amount:       &h.couponwithdraw.Amount,
			IOType:       &ioType,
			IOSubType:    &ioSubType,
			IOExtra:      &ioExtra,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkCouponWithdraw(ctx); err != nil {
		return nil, err
	}
	switch {
	case h.State == nil:
		fallthrough //nolint
	case h.State.String() == handler.couponwithdraw.State:
		fallthrough //nolint
	case handler.couponwithdraw.State == types.WithdrawState_Approved.String():
		fallthrough //nolint
	case handler.couponwithdraw.State == types.WithdrawState_Rejected.String():
		return h.GetCouponWithdraw(ctx)
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCouponWithdraw(ctx, tx); err != nil {
			return err
		}
		if err := handler.createStatement(ctx, tx); err != nil {
			return err
		}
		if err := handler.createOrUpdateLedger(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCouponWithdraw(ctx)
}
