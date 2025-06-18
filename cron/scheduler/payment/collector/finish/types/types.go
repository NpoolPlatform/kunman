package types

import (
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
)

type PersistentAccount struct {
	*paymentaccountmwpb.Account
	Error error
}
