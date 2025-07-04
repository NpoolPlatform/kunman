// Code generated by ent, DO NOT EDIT.

package orderpaymentstatement

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the orderpaymentstatement type in the database.
	Label = "order_payment_statement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldStatementID holds the string denoting the statement_id field in the database.
	FieldStatementID = "statement_id"
	// FieldPaymentCoinTypeID holds the string denoting the payment_coin_type_id field in the database.
	FieldPaymentCoinTypeID = "payment_coin_type_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCommissionAmount holds the string denoting the commission_amount field in the database.
	FieldCommissionAmount = "commission_amount"
	// Table holds the table name of the orderpaymentstatement in the database.
	Table = "order_payment_statements"
)

// Columns holds all SQL columns for orderpaymentstatement fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldStatementID,
	FieldPaymentCoinTypeID,
	FieldAmount,
	FieldCommissionAmount,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultStatementID holds the default value on creation for the "statement_id" field.
	DefaultStatementID func() uuid.UUID
	// DefaultPaymentCoinTypeID holds the default value on creation for the "payment_coin_type_id" field.
	DefaultPaymentCoinTypeID func() uuid.UUID
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount decimal.Decimal
	// DefaultCommissionAmount holds the default value on creation for the "commission_amount" field.
	DefaultCommissionAmount decimal.Decimal
)

// OrderOption defines the ordering options for the OrderPaymentStatement queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
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

// ByEntID orders the results by the ent_id field.
func ByEntID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntID, opts...).ToFunc()
}

// ByStatementID orders the results by the statement_id field.
func ByStatementID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatementID, opts...).ToFunc()
}

// ByPaymentCoinTypeID orders the results by the payment_coin_type_id field.
func ByPaymentCoinTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentCoinTypeID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByCommissionAmount orders the results by the commission_amount field.
func ByCommissionAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommissionAmount, opts...).ToFunc()
}
