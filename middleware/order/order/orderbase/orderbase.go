package orderbase

import (
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type OrderBase interface {
	cruder.CrudBase
}

type orderBase struct {
	_ent *ent.OrderBase
}

func (gb *orderBase) ID() uint32 {
	return gb._ent.ID
}

func (gb *orderBase) EntID() uuid.UUID {
	return gb._ent.EntID
}

func (gb *orderBase) CreatedAt() uint32 {
	return gb._ent.CreatedAt
}

func (gb *orderBase) UpdatedAt() uint32 {
	return gb._ent.UpdatedAt
}

func (gb *orderBase) DeletedAt() uint32 {
	return gb._ent.DeletedAt
}
