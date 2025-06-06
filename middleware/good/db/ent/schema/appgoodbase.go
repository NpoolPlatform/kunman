package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// AppGoodBase holds the schema definition for the AppGoodBase entity.
type AppGoodBase struct {
	ent.Schema
}

func (AppGoodBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

func (AppGoodBase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Bool("purchasable").
			Optional().
			Default(false),
		field.
			Bool("enable_product_page").
			Optional().
			Default(false),
		field.
			String("product_page").
			Optional().
			Default(""),
		field.
			Bool("online").
			Optional().
			Default(false),
		field.
			Bool("visible").
			Optional().
			Default(false),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			Int32("display_index").
			Optional().
			Default(0),
		field.
			String("banner").
			Optional().
			Default(""),
	}
}

// Edges of the AppGoodBase.
func (AppGoodBase) Edges() []ent.Edge {
	return nil
}

func (AppGoodBase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id"),
	}
}
