// Code generated by ent, DO NOT EDIT.

package topmostgoodconstraint

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the topmostgoodconstraint type in the database.
	Label = "top_most_good_constraint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTopMostGoodID holds the string denoting the top_most_good_id field in the database.
	FieldTopMostGoodID = "top_most_good_id"
	// FieldConstraint holds the string denoting the constraint field in the database.
	FieldConstraint = "constraint"
	// FieldTargetValue holds the string denoting the target_value field in the database.
	FieldTargetValue = "target_value"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// Table holds the table name of the topmostgoodconstraint in the database.
	Table = "top_most_good_constraints"
)

// Columns holds all SQL columns for topmostgoodconstraint fields.
var Columns = []string{
	FieldID,
	FieldEntID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTopMostGoodID,
	FieldConstraint,
	FieldTargetValue,
	FieldIndex,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultTopMostGoodID holds the default value on creation for the "top_most_good_id" field.
	DefaultTopMostGoodID func() uuid.UUID
	// DefaultConstraint holds the default value on creation for the "constraint" field.
	DefaultConstraint string
	// DefaultTargetValue holds the default value on creation for the "target_value" field.
	DefaultTargetValue decimal.Decimal
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint8
)

// OrderOption defines the ordering options for the TopMostGoodConstraint queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEntID orders the results by the ent_id field.
func ByEntID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByTopMostGoodID orders the results by the top_most_good_id field.
func ByTopMostGoodID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTopMostGoodID, opts...).ToFunc()
}

// ByConstraint orders the results by the constraint field.
func ByConstraint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConstraint, opts...).ToFunc()
}

// ByTargetValue orders the results by the target_value field.
func ByTargetValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetValue, opts...).ToFunc()
}

// ByIndex orders the results by the index field.
func ByIndex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndex, opts...).ToFunc()
}
