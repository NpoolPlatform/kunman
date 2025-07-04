// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldCoinType holds the string denoting the coin_type field in the database.
	FieldCoinType = "coin_type"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldUtxo holds the string denoting the utxo field in the database.
	FieldUtxo = "utxo"
	// FieldPre holds the string denoting the pre field in the database.
	FieldPre = "pre"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldRecentBhash holds the string denoting the recent_bhash field in the database.
	FieldRecentBhash = "recent_bhash"
	// FieldTxData holds the string denoting the tx_data field in the database.
	FieldTxData = "tx_data"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldCid holds the string denoting the cid field in the database.
	FieldCid = "cid"
	// FieldExitCode holds the string denoting the exit_code field in the database.
	FieldExitCode = "exit_code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldEntID,
	FieldCoinType,
	FieldNonce,
	FieldUtxo,
	FieldPre,
	FieldTransactionType,
	FieldRecentBhash,
	FieldTxData,
	FieldTransactionID,
	FieldCid,
	FieldExitCode,
	FieldName,
	FieldFrom,
	FieldTo,
	FieldMemo,
	FieldAmount,
	FieldPayload,
	FieldState,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
	// DefaultCoinType holds the default value on creation for the "coin_type" field.
	DefaultCoinType int32
	// DefaultNonce holds the default value on creation for the "nonce" field.
	DefaultNonce uint64
	// DefaultUtxo holds the default value on creation for the "utxo" field.
	DefaultUtxo []*sphinxplugin.Unspent
	// DefaultPre holds the default value on creation for the "pre" field.
	DefaultPre *sphinxplugin.Unspent
	// DefaultTransactionType holds the default value on creation for the "transaction_type" field.
	DefaultTransactionType int8
	// DefaultRecentBhash holds the default value on creation for the "recent_bhash" field.
	DefaultRecentBhash string
	// DefaultTxData holds the default value on creation for the "tx_data" field.
	DefaultTxData []byte
	// DefaultCid holds the default value on creation for the "cid" field.
	DefaultCid string
	// DefaultExitCode holds the default value on creation for the "exit_code" field.
	DefaultExitCode int64
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultFrom holds the default value on creation for the "from" field.
	DefaultFrom string
	// DefaultTo holds the default value on creation for the "to" field.
	DefaultTo string
	// DefaultMemo holds the default value on creation for the "memo" field.
	DefaultMemo string
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount uint64
	// DefaultPayload holds the default value on creation for the "payload" field.
	DefaultPayload []byte
	// PayloadValidator is a validator for the "payload" field. It is called by the builders before save.
	PayloadValidator func([]byte) error
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState uint8
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
)

// OrderOption defines the ordering options for the Transaction queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEntID orders the results by the ent_id field.
func ByEntID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntID, opts...).ToFunc()
}

// ByCoinType orders the results by the coin_type field.
func ByCoinType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinType, opts...).ToFunc()
}

// ByNonce orders the results by the nonce field.
func ByNonce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNonce, opts...).ToFunc()
}

// ByTransactionType orders the results by the transaction_type field.
func ByTransactionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionType, opts...).ToFunc()
}

// ByRecentBhash orders the results by the recent_bhash field.
func ByRecentBhash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecentBhash, opts...).ToFunc()
}

// ByTransactionID orders the results by the transaction_id field.
func ByTransactionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionID, opts...).ToFunc()
}

// ByCid orders the results by the cid field.
func ByCid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCid, opts...).ToFunc()
}

// ByExitCode orders the results by the exit_code field.
func ByExitCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExitCode, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFrom orders the results by the from field.
func ByFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrom, opts...).ToFunc()
}

// ByTo orders the results by the to field.
func ByTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTo, opts...).ToFunc()
}

// ByMemo orders the results by the memo field.
func ByMemo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
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
