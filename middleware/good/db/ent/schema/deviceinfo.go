package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// DeviceInfo holds the schema definition for the DeviceInfo entity.
type DeviceInfo struct {
	ent.Schema
}

func (DeviceInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

// Fields of the DeviceInfo.
func (DeviceInfo) Fields() []ent.Field {
	const maxLen = 64
	return []ent.Field{
		field.
			String("type").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.
			UUID("manufacturer_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("power_consumption").
			Optional().
			Default(0),
		field.
			Uint32("shipment_at").
			Optional().
			Default(0),
	}
}

// Edges of the DeviceInfo.
func (DeviceInfo) Edges() []ent.Edge {
	return nil
}
