// Code generated by ent, DO NOT EDIT.

package tran

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the tran type in the database.
	Label = "tran"
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
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldFromAccountID holds the string denoting the from_account_id field in the database.
	FieldFromAccountID = "from_account_id"
	// FieldToAccountID holds the string denoting the to_account_id field in the database.
	FieldToAccountID = "to_account_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldFeeAmount holds the string denoting the fee_amount field in the database.
	FieldFeeAmount = "fee_amount"
	// FieldChainTxID holds the string denoting the chain_tx_id field in the database.
	FieldChainTxID = "chain_tx_id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldExtra holds the string denoting the extra field in the database.
	FieldExtra = "extra"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the tran in the database.
	Table = "trans"
)

// Columns holds all SQL columns for tran fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldCoinTypeID,
	FieldFromAccountID,
	FieldToAccountID,
	FieldAmount,
	FieldFeeAmount,
	FieldChainTxID,
	FieldState,
	FieldExtra,
	FieldType,
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
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultFromAccountID holds the default value on creation for the "from_account_id" field.
	DefaultFromAccountID func() uuid.UUID
	// DefaultToAccountID holds the default value on creation for the "to_account_id" field.
	DefaultToAccountID func() uuid.UUID
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount decimal.Decimal
	// DefaultFeeAmount holds the default value on creation for the "fee_amount" field.
	DefaultFeeAmount decimal.Decimal
	// DefaultChainTxID holds the default value on creation for the "chain_tx_id" field.
	DefaultChainTxID string
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultExtra holds the default value on creation for the "extra" field.
	DefaultExtra string
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
)

// OrderOption defines the ordering options for the Tran queries.
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

// ByCoinTypeID orders the results by the coin_type_id field.
func ByCoinTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinTypeID, opts...).ToFunc()
}

// ByFromAccountID orders the results by the from_account_id field.
func ByFromAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromAccountID, opts...).ToFunc()
}

// ByToAccountID orders the results by the to_account_id field.
func ByToAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToAccountID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByFeeAmount orders the results by the fee_amount field.
func ByFeeAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeeAmount, opts...).ToFunc()
}

// ByChainTxID orders the results by the chain_tx_id field.
func ByChainTxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChainTxID, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByExtra orders the results by the extra field.
func ByExtra(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtra, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}
