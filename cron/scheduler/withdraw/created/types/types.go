package types

import (
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
)

type PersistentWithdraw struct {
	*withdrawmwpb.Withdraw
	ReviewTrigger reviewtypes.ReviewTriggerType
	Error         error
}
