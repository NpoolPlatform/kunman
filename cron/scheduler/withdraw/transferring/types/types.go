package types

import (
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
)

type PersistentWithdraw struct {
	*withdrawmwpb.Withdraw
	NewWithdrawState ledgertypes.WithdrawState
	ChainTxID        string
	Error            error
}
