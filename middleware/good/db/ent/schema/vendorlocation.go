package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// VendorLocation holds the schema definition for the VendorLocation entity.
type VendorLocation struct {
	ent.Schema
}

func (VendorLocation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the VendorLocation.
func (VendorLocation) Fields() []ent.Field {
	maxLen := 128
	addressMaxLen := 256
	return []ent.Field{
		field.String("country").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("province").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("city").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("address").
			Optional().
			Default("").
			MaxLen(addressMaxLen),
		field.UUID("brand_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

// Edges of the VendorLocation.
func (VendorLocation) Edges() []ent.Edge {
	return nil
}
