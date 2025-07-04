package types

import (
	payaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
)

type PersistentAccount struct {
	*payaccmwpb.Account
	CollectAmount          string
	FeeAmount              string
	PaymentAccountID       string
	PaymentAddress         string
	CollectAccountID       string
	CollectAddress         string
	CollectingTIDCandidate *string
	Error                  error
}
