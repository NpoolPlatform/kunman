package types

import (
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
)

type PersistentWithdraw struct {
	*withdrawmwpb.Withdraw
	NewWithdrawState ledgertypes.WithdrawState
	NewReviewState   reviewtypes.ReviewState
	NeedUpdateReview bool
	Error            error
}
