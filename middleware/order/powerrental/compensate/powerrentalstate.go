package compensate

import (
	"github.com/google/uuid"
)

type powerRentalState struct {
	ID      uint32    `sql:"id"`
	OrderID uuid.UUID `sql:"ent_id"`
}

type powerRentalStates struct {
	powerRentalStates []*powerRentalState
}

func (f *powerRentalStates) Exhausted() bool {
	return len(f.powerRentalStates) == 0
}

func (f *powerRentalStates) Drain() {
	f.powerRentalStates = []*powerRentalState{}
}
