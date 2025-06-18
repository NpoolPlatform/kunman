package types

import (
	depositaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/deposit"
)

type PersistentAccount struct {
	*depositaccmwpb.Account
	CollectAmount          string
	FeeAmount              string
	DepositAccountID       string
	DepositAddress         string
	CollectAccountID       string
	CollectAddress         string
	CollectingTIDCandidate *string
	Error                  error
}
