package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// CoinFiat holds the schema definition for the CoinFiat entity.
type CoinFiat struct {
	ent.Schema
}

func (CoinFiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CoinFiat.
func (CoinFiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(basetypes.CurrencyFeedType_DefaultFeedType.String()),
	}
}

// Edges of the CoinFiat.
func (CoinFiat) Edges() []ent.Edge {
	return nil
}
