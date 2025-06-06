//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
)

// DeviceManufacturer holds the schema definition for the DeviceManufacturer entity.
type DeviceManufacturer struct {
	ent.Schema
}

func (DeviceManufacturer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the DeviceManufacturer.
func (DeviceManufacturer) Fields() []ent.Field {
	const maxLen = 128
	return []ent.Field{
		field.String("name").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("logo").
			Optional().
			Default("").
			MaxLen(maxLen),
	}
}

// Edges of the DeviceManufacturer.
func (DeviceManufacturer) Edges() []ent.Edge {
	return nil
}
