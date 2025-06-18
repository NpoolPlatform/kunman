package sentinel

import (
	"context"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/approved/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	couponwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	couponwithdrawmw "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw/coupon"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &couponwithdrawmwpb.Conds{
		State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.WithdrawState_Approved)},
	}

	for {
		handler, err := couponwithdrawmw.NewHandler(
			ctx,
			couponwithdrawmw.WithConds(conds),
			couponwithdrawmw.WithOffset(offset),
			couponwithdrawmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		couponwithdraws, _, err := handler.GetCouponWithdraws(ctx)
		if err != nil {
			return err
		}
		if len(couponwithdraws) == 0 {
			return nil
		}
		for _, cw := range couponwithdraws {
			cancelablefeed.CancelableFeed(ctx, cw, exec)
		}
		offset += limit
	}
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if couponwithdraw, ok := ent.(*types.PersistentCouponWithdraw); ok {
		return couponwithdraw.EntID
	}
	return ent.(*couponwithdrawmwpb.CouponWithdraw).EntID
}
