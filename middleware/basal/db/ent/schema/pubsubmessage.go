package schema

import (
	"entgo.io/ent"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	pubsub "github.com/NpoolPlatform/libent-cruder/pkg/pubsub"
)

// PubsubMessage holds the schema definition for the PubsubMessage entity.
type PubsubMessage struct {
	ent.Schema
}

func (PubsubMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		pubsub.PubsubMessage{},
		crudermixin.AutoIDMixin{},
	}
}
