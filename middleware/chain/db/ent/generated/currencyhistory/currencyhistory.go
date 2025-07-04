// Code generated by ent, DO NOT EDIT.

package currencyhistory

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the currencyhistory type in the database.
	Label = "currency_history"
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
	// FieldFeedType holds the string denoting the feed_type field in the database.
	FieldFeedType = "feed_type"
	// FieldMarketValueHigh holds the string denoting the market_value_high field in the database.
	FieldMarketValueHigh = "market_value_high"
	// FieldMarketValueLow holds the string denoting the market_value_low field in the database.
	FieldMarketValueLow = "market_value_low"
	// Table holds the table name of the currencyhistory in the database.
	Table = "currency_histories"
)

// Columns holds all SQL columns for currencyhistory fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldCoinTypeID,
	FieldFeedType,
	FieldMarketValueHigh,
	FieldMarketValueLow,
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
	// DefaultFeedType holds the default value on creation for the "feed_type" field.
	DefaultFeedType string
	// DefaultMarketValueHigh holds the default value on creation for the "market_value_high" field.
	DefaultMarketValueHigh decimal.Decimal
	// DefaultMarketValueLow holds the default value on creation for the "market_value_low" field.
	DefaultMarketValueLow decimal.Decimal
)

// OrderOption defines the ordering options for the CurrencyHistory queries.
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

// ByFeedType orders the results by the feed_type field.
func ByFeedType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeedType, opts...).ToFunc()
}

// ByMarketValueHigh orders the results by the market_value_high field.
func ByMarketValueHigh(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMarketValueHigh, opts...).ToFunc()
}

// ByMarketValueLow orders the results by the market_value_low field.
func ByMarketValueLow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMarketValueLow, opts...).ToFunc()
}
