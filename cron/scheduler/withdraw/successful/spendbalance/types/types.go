package types

import (
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
)

type PersistentWithdraw struct {
	*withdrawmwpb.Withdraw
	LockedBalanceAmount string
	WithdrawFeeAmount   string
}
