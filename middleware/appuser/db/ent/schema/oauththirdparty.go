package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
)

// OAuthThirdParty holds the schema definition for the OAuthThirdParty entity.
type OAuthThirdParty struct {
	ent.Schema
}

func (OAuthThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OAuthThirdParty.
func (OAuthThirdParty) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("client_name").
			Optional().
			Default(basetypes.SignMethod_DefaultSignMethod.String()),
		field.
			String("client_tag").
			Optional().
			Default(""),
		field.
			String("client_logo_url").
			Optional().
			Default(""),
		field.
			String("client_oauth_url").
			Optional().
			Default(""),
		field.
			String("response_type").
			Optional().
			Default(""),
		field.
			String("scope").
			Optional().
			Default(""),
	}
}

// Edges of the OAuthThirdParty.
func (OAuthThirdParty) Edges() []ent.Edge {
	return nil
}
