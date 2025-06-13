package coupon

import (
	"context"
	"fmt"

	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

func (h *Handler) CreateCouponWithdraw(ctx context.Context) (*npool.CouponWithdraw, error) {
	h.Conds = &crud.Conds{
		AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:      &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AllocatedID: &cruder.Cond{Op: cruder.EQ, Val: *h.AllocatedID},
		States: &cruder.Cond{Op: cruder.IN, Val: []ledgertypes.WithdrawState{
			ledgertypes.WithdrawState_Approved,
			ledgertypes.WithdrawState_Reviewing,
		}},
	}
	exist, err := h.ExistCouponWithdrawConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coupon withdraw already exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.CouponWithdraw.Create(),
			&h.Req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCouponWithdraw(ctx)
}
