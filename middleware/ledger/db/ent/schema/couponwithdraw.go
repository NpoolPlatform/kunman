package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
)

// CouponWithdraw holds the schema definition for the CouponWithdraw entity.
type CouponWithdraw struct {
	ent.Schema
}

func (CouponWithdraw) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the CouponWithdraw.
func (CouponWithdraw) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("allocated_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("state").
			Optional().
			Default(basetypes.WithdrawState_Reviewing.String()),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			UUID("review_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the CouponWithdraw.
func (CouponWithdraw) Edges() []ent.Edge {
	return nil
}
