package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppUserExtra holds the schema definition for the AppUserExtra entity.
type AppUserExtra struct {
	ent.Schema
}

func (AppUserExtra) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppUserExtra.
func (AppUserExtra) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("username").
			Default(""),
		field.String("first_name").
			Default(""),
		field.String("last_name").
			Default(""),
		field.JSON("address_fields", []string{}).
			Default([]string{}),
		field.String("gender").
			Default(""),
		field.String("postal_code").
			Default(""),
		field.Uint32("age").
			Default(0),
		field.Uint32("birthday").
			Default(0),
		field.String("avatar").
			Default(""),
		field.String("organization").
			Default(""),
		field.String("id_number").
			Default(""),
	}
}

// Edges of the AppUserExtra.
func (AppUserExtra) Edges() []ent.Edge {
	return nil
}

func (AppUserExtra) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id").Unique(),
	}
}
