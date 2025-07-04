package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	npool "github.com/NpoolPlatform/kunman/message/basal/middleware/v1/api"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
)

// API holds the schema definition for the API entity.
type API struct {
	ent.Schema
}

func (API) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

func (API) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "apis"},
	}
}

// Fields of the API.
func (API) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("protocol").
			Optional().
			Default(npool.Protocol_DefaultProtocol.String()),
		field.
			String("service_name").
			Optional().
			Default(""),
		field.
			String("method").
			Optional().
			Default(npool.Method_DefaultMethod.String()),
		field.
			String("method_name").
			Optional().
			Default(""),
		field.
			String("path").
			Optional().
			Default(""),
		field.
			Bool("exported").
			Optional().
			Default(false),
		field.
			String("path_prefix").
			Optional().
			Default(""),
		field.
			JSON("domains", []string{}).
			Optional().
			Default([]string{}),
		field.
			Bool("deprecated").
			Optional().
			Default(false),
	}
}

// Edges of the API.
func (API) Edges() []ent.Edge {
	return nil
}
