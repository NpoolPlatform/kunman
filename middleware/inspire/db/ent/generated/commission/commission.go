// Code generated by ent, DO NOT EDIT.

package commission

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the commission type in the database.
	Label = "commission"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldAmountOrPercent holds the string denoting the amount_or_percent field in the database.
	FieldAmountOrPercent = "amount_or_percent"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldSettleType holds the string denoting the settle_type field in the database.
	FieldSettleType = "settle_type"
	// FieldSettleMode holds the string denoting the settle_mode field in the database.
	FieldSettleMode = "settle_mode"
	// FieldSettleInterval holds the string denoting the settle_interval field in the database.
	FieldSettleInterval = "settle_interval"
	// FieldSettleAmountType holds the string denoting the settle_amount_type field in the database.
	FieldSettleAmountType = "settle_amount_type"
	// FieldThreshold holds the string denoting the threshold field in the database.
	FieldThreshold = "threshold"
	// FieldOrderLimit holds the string denoting the order_limit field in the database.
	FieldOrderLimit = "order_limit"
	// Table holds the table name of the commission in the database.
	Table = "commissions"
)

// Columns holds all SQL columns for commission fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldAppGoodID,
	FieldAmountOrPercent,
	FieldStartAt,
	FieldEndAt,
	FieldSettleType,
	FieldSettleMode,
	FieldSettleInterval,
	FieldSettleAmountType,
	FieldThreshold,
	FieldOrderLimit,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultAmountOrPercent holds the default value on creation for the "amount_or_percent" field.
	DefaultAmountOrPercent decimal.Decimal
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt uint32
	// DefaultSettleType holds the default value on creation for the "settle_type" field.
	DefaultSettleType string
	// DefaultSettleMode holds the default value on creation for the "settle_mode" field.
	DefaultSettleMode string
	// DefaultSettleInterval holds the default value on creation for the "settle_interval" field.
	DefaultSettleInterval string
	// DefaultSettleAmountType holds the default value on creation for the "settle_amount_type" field.
	DefaultSettleAmountType string
	// DefaultThreshold holds the default value on creation for the "threshold" field.
	DefaultThreshold decimal.Decimal
	// DefaultOrderLimit holds the default value on creation for the "order_limit" field.
	DefaultOrderLimit uint32
)

// OrderOption defines the ordering options for the Commission queries.
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

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByGoodID orders the results by the good_id field.
func ByGoodID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoodID, opts...).ToFunc()
}

// ByAppGoodID orders the results by the app_good_id field.
func ByAppGoodID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppGoodID, opts...).ToFunc()
}

// ByAmountOrPercent orders the results by the amount_or_percent field.
func ByAmountOrPercent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountOrPercent, opts...).ToFunc()
}

// ByStartAt orders the results by the start_at field.
func ByStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartAt, opts...).ToFunc()
}

// ByEndAt orders the results by the end_at field.
func ByEndAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndAt, opts...).ToFunc()
}

// BySettleType orders the results by the settle_type field.
func BySettleType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettleType, opts...).ToFunc()
}

// BySettleMode orders the results by the settle_mode field.
func BySettleMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettleMode, opts...).ToFunc()
}

// BySettleInterval orders the results by the settle_interval field.
func BySettleInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettleInterval, opts...).ToFunc()
}

// BySettleAmountType orders the results by the settle_amount_type field.
func BySettleAmountType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSettleAmountType, opts...).ToFunc()
}

// ByThreshold orders the results by the threshold field.
func ByThreshold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreshold, opts...).ToFunc()
}

// ByOrderLimit orders the results by the order_limit field.
func ByOrderLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderLimit, opts...).ToFunc()
}
