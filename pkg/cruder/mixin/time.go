package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

type TimePolicy interface {
	Policy() ent.Policy
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.
			Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.
			Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}
