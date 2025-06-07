package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	crudermixin "github.com/NpoolPlatform/kunman/pkg/cruder/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppGood holds the schema definition for the AppGood entity.
type AppGood struct {
	ent.Schema
}

func (AppGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		crudermixin.TimeMixin{},
	}
}

//nolint:funlen
func (AppGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			Bool("online").
			Optional().
			Default(false),
		field.
			Bool("visible").
			Optional().
			Default(false),
		field.
			String("good_name").
			Optional().
			Default(""),
		field.
			Other("unit_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("package_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Int32("display_index").
			Optional().
			Default(0),
		field.
			Uint32("sale_start_at").
			Optional().
			Default(0),
		field.
			Uint32("sale_end_at").
			Optional().
			Default(0),
		field.
			Uint32("service_start_at").
			Optional().
			Default(0),
		field.
			Other("technical_fee_ratio", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("electricity_fee_ratio", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			JSON("descriptions", []string{}).
			Optional().
			Default([]string{}),
		field.
			String("good_banner").
			Optional().
			Default(""),
		field.
			JSON("display_names", []string{}).
			Optional().
			Default([]string{}),
		field.
			Bool("enable_purchase").
			Optional().
			Default(true),
		field.
			Bool("enable_product_page").
			Optional().
			Default(true),
		field.
			String("cancel_mode").
			Optional().
			Default(types.CancelMode_Uncancellable.String()),
		field.
			JSON("display_colors", []string{}).
			Optional().
			Default([]string{}),
		field.
			Uint32("cancellable_before_start").
			Optional().
			Default(0),
		field.
			String("product_page").
			Optional().
			Default(""),
		field.
			Bool("enable_set_commission").
			Optional().
			Default(true),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
		field.
			Other("min_order_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Other("max_order_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Other("max_user_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.NewFromInt(0)),
		field.
			Uint32("min_order_duration").
			Optional().
			Default(0),
		field.
			Uint32("max_order_duration").
			Optional().
			Default(0),
		field.
			Bool("package_with_requireds").
			Optional().
			Default(true),
	}
}

// Edges of the AppGood.
func (AppGood) Edges() []ent.Edge {
	return nil
}

func (AppGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id", "online"),
	}
}
