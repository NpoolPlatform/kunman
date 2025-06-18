package types

import (
	couponwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw/coupon"
)

type PersistentCouponWithdraw struct {
	*couponwithdrawmwpb.CouponWithdraw
}
