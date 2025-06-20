package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// CurrencyFeed holds the schema definition for the CurrencyFeed entity.
type CurrencyFeed struct {
	ent.Schema
}

func (CurrencyFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CurrencyFeed.
func (CurrencyFeed) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(basetypes.CurrencyFeedType_DefaultFeedType.String()),
		field.
			String("feed_coin_name").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the CurrencyFeed.
func (CurrencyFeed) Edges() []ent.Edge {
	return nil
}

func (CurrencyFeed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "id"),
		index.Fields("coin_type_id"),
	}
}
