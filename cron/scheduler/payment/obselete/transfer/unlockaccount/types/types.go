package types

import (
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
)

type PersistentPayment struct {
	*paymentmwpb.Payment
	UnlockAccountID *uint32
}
