package sentinel

import (
	"context"

	couponwithdrawmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	couponwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/approved/types"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	for {
		couponwithdraws, _, err := couponwithdrawmwcli.GetCouponWithdraws(ctx, &couponwithdrawmwpb.Conds{
			State: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.WithdrawState_Approved)},
		}, offset, limit)
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
