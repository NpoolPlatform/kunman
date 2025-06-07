package outofgas

import (
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
)

type powerRentalState struct {
	entPowerRentalState *ent.PowerRentalState
}

func (f *powerRentalState) PowerRentalStateID() uint32 {
	return f.entPowerRentalState.ID
}

func (f *powerRentalState) OrderID() uuid.UUID {
	return f.entPowerRentalState.OrderID
}
