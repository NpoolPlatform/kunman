package cruder

import (
	"github.com/google/uuid"
)

type CrudBase interface {
	ID() uint32
	EntID() uuid.UUID
	CreatedAt() uint32
	UpdatedAt() uint32
	DeletedAt() uint32
}
