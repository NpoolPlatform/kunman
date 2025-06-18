//nolint:nolintlint,dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
)

// ReadAnnouncement holds the schema definition for the ReadAnnouncement entity.
type ReadAnnouncement struct {
	ent.Schema
}

func (ReadAnnouncement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the ReadAnnouncement.
func (ReadAnnouncement) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("announcement_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
	}
}

// Edges of the ReadAnnouncement.
func (ReadAnnouncement) Edges() []ent.Edge {
	return nil
}
