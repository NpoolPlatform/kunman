// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApisColumns holds the columns for the "apis" table.
	ApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "protocol", Type: field.TypeString, Nullable: true, Default: "DefaultProtocol"},
		{Name: "service_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: "DefaultMethod"},
		{Name: "method_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "exported", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "path_prefix", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "domains", Type: field.TypeJSON, Nullable: true},
		{Name: "deprecated", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// ApisTable holds the schema information for the "apis" table.
	ApisTable = &schema.Table{
		Name:       "apis",
		Columns:    ApisColumns,
		PrimaryKey: []*schema.Column{ApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ApisColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApisTable,
	}
)

func init() {
	ApisTable.Annotation = &entsql.Annotation{
		Table: "apis",
	}
}
