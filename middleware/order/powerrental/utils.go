package powerrental

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"

	"github.com/google/uuid"
)

func (h *Handler) ledgerLockID() *uuid.UUID {
	for _, req := range h.OrderLockReqs {
		if *req.LockType == types.OrderLockType_LockBalance {
			return req.EntID
		}
	}
	return nil
}
