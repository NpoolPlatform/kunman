package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// PoolOrderUser holds the schema definition for the PoolOrderUser entity.
type PoolOrderUser struct {
	ent.Schema
}

func (PoolOrderUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the PoolOrderUser.
func (PoolOrderUser) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("order_id", uuid.UUID{}),
		field.
			UUID("pool_order_user_id", uuid.UUID{}),
	}
}

// Edges of the PoolOrderUser.
func (PoolOrderUser) Edges() []ent.Edge {
	return nil
}

func (PoolOrderUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
