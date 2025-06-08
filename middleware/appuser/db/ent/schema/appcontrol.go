package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	appusertypes "github.com/NpoolPlatform/kunman/message/basetypes/appuser/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/google/uuid"
)

// AppControl holds the schema definition for the AppControl entity.
type AppControl struct {
	ent.Schema
}

func (AppControl) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppControl.
func (AppControl) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			JSON("signup_methods", []string{}).
			Optional().
			Default(func() []string {
				return []string{
					basetypes.SignMethod_Mobile.String(),
					basetypes.SignMethod_Email.String(),
				}
			}),
		field.
			JSON("extern_signin_methods", []string{}).
			Optional().
			Default(func() []string {
				return []string{}
			}),
		field.
			String("recaptcha_method").
			Optional().
			Default(basetypes.RecaptchaMethod_GoogleRecaptchaV3.String()),
		field.
			Bool("kyc_enable").
			Optional().
			Default(false),
		field.
			Bool("signin_verify_enable").
			Optional().
			Default(false),
		field.
			Bool("invitation_code_must").
			Optional().
			Default(false),
		field.
			String("create_invitation_code_when").
			Optional().
			Default(basetypes.CreateInvitationCodeWhen_Registration.String()),
		field.
			Uint32("max_typed_coupons_per_order").
			Optional().
			Default(1),
		field.
			Bool("maintaining").
			Optional().
			Default(false),
		field.
			Bool("coupon_withdraw_enable").
			Optional().
			Default(false),
		field.
			JSON("commit_button_targets", []string{}).
			Optional().
			Default(func() []string {
				return []string{}
			}),
		field.
			String("reset_user_method").
			Optional().
			Default(appusertypes.ResetUserMethod_Normal.String()),
	}
}

// Edges of the AppControl.
func (AppControl) Edges() []ent.Edge {
	return nil
}
