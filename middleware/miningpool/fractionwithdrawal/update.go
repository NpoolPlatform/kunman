package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	fractionwithdrawalcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawal"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	fractionwithdrawalent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawal"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) validateState(info *ent.FractionWithdrawal) error {
	if info.FractionWithdrawState == basetypes.FractionWithdrawState_DefaultFractionWithdrawState.String() {
		return wlog.Errorf("invalid withdrawstate")
	}
	return nil
}

func (h *Handler) UpdateFractionWithdrawal(ctx context.Context) error {
	info, err := h.GetFractionWithdrawal(ctx)
	if err != nil {
		return wlog.Errorf("invalid id or ent_id")
	}
	if info == nil {
		return wlog.Errorf("invalid id or ent_id")
	}

	h.ID = &info.ID

	handler := &updateHandler{
		Handler: h,
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			FractionWithdrawal.
			Query().
			Where(
				fractionwithdrawalent.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		if err := handler.validateState(info); err != nil {
			return wlog.WrapError(err)
		}

		stm, err := fractionwithdrawalcrud.UpdateSet(
			info.Update(),
			&fractionwithdrawalcrud.Req{
				AppID:                 h.AppID,
				UserID:                h.UserID,
				OrderUserID:           h.OrderUserID,
				CoinTypeID:            h.CoinTypeID,
				FractionWithdrawState: h.FractionWithdrawState,
				WithdrawAt:            h.WithdrawAt,
				PromisePayAt:          h.PromisePayAt,
				Msg:                   h.Msg,
			},
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _, err := stm.Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
