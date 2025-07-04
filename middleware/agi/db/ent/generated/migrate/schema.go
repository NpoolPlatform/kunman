// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CapacitiesColumns holds the columns for the "capacities" table.
	CapacitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "capacity_key", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "capacity_value", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CapacitiesTable holds the schema information for the "capacities" table.
	CapacitiesTable = &schema.Table{
		Name:       "capacities",
		Columns:    CapacitiesColumns,
		PrimaryKey: []*schema.Column{CapacitiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "capacity_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CapacitiesColumns[1]},
			},
		},
	}
	// QuotasColumns holds the columns for the "quotas" table.
	QuotasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "quota", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "consumed_quota", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "expired_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// QuotasTable holds the schema information for the "quotas" table.
	QuotasTable = &schema.Table{
		Name:       "quotas",
		Columns:    QuotasColumns,
		PrimaryKey: []*schema.Column{QuotasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "quota_ent_id",
				Unique:  true,
				Columns: []*schema.Column{QuotasColumns[1]},
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "next_extend_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "permanent_quota", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "consumed_quota", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "pay_with_coin_balance", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "subscription_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "fiat_payment_channel", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "last_payment_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "last_updated_event_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "activated_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "activated_event_id", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subscription_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SubscriptionsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CapacitiesTable,
		QuotasTable,
		SubscriptionsTable,
	}
)

func init() {
	QuotasTable.Annotation = &entsql.Annotation{
		Table: "quotas",
	}
}
