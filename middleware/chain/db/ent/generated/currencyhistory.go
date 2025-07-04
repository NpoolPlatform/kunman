// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currencyhistory"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CurrencyHistory is the model entity for the CurrencyHistory schema.
type CurrencyHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// FeedType holds the value of the "feed_type" field.
	FeedType string `json:"feed_type,omitempty"`
	// MarketValueHigh holds the value of the "market_value_high" field.
	MarketValueHigh decimal.Decimal `json:"market_value_high,omitempty"`
	// MarketValueLow holds the value of the "market_value_low" field.
	MarketValueLow decimal.Decimal `json:"market_value_low,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CurrencyHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case currencyhistory.FieldMarketValueHigh, currencyhistory.FieldMarketValueLow:
			values[i] = new(decimal.Decimal)
		case currencyhistory.FieldID, currencyhistory.FieldCreatedAt, currencyhistory.FieldUpdatedAt, currencyhistory.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case currencyhistory.FieldFeedType:
			values[i] = new(sql.NullString)
		case currencyhistory.FieldEntID, currencyhistory.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CurrencyHistory fields.
func (ch *CurrencyHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case currencyhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ch.ID = uint32(value.Int64)
		case currencyhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ch.CreatedAt = uint32(value.Int64)
			}
		case currencyhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ch.UpdatedAt = uint32(value.Int64)
			}
		case currencyhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ch.DeletedAt = uint32(value.Int64)
			}
		case currencyhistory.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ch.EntID = *value
			}
		case currencyhistory.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ch.CoinTypeID = *value
			}
		case currencyhistory.FieldFeedType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_type", values[i])
			} else if value.Valid {
				ch.FeedType = value.String
			}
		case currencyhistory.FieldMarketValueHigh:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_high", values[i])
			} else if value != nil {
				ch.MarketValueHigh = *value
			}
		case currencyhistory.FieldMarketValueLow:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_low", values[i])
			} else if value != nil {
				ch.MarketValueLow = *value
			}
		default:
			ch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CurrencyHistory.
// This includes values selected through modifiers, order, etc.
func (ch *CurrencyHistory) Value(name string) (ent.Value, error) {
	return ch.selectValues.Get(name)
}

// Update returns a builder for updating this CurrencyHistory.
// Note that you need to call CurrencyHistory.Unwrap() before calling this method if this CurrencyHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *CurrencyHistory) Update() *CurrencyHistoryUpdateOne {
	return NewCurrencyHistoryClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the CurrencyHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *CurrencyHistory) Unwrap() *CurrencyHistory {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("generated: CurrencyHistory is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *CurrencyHistory) String() string {
	var builder strings.Builder
	builder.WriteString("CurrencyHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ch.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.EntID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ch.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("feed_type=")
	builder.WriteString(ch.FeedType)
	builder.WriteString(", ")
	builder.WriteString("market_value_high=")
	builder.WriteString(fmt.Sprintf("%v", ch.MarketValueHigh))
	builder.WriteString(", ")
	builder.WriteString("market_value_low=")
	builder.WriteString(fmt.Sprintf("%v", ch.MarketValueLow))
	builder.WriteByte(')')
	return builder.String()
}

// CurrencyHistories is a parsable slice of CurrencyHistory.
type CurrencyHistories []*CurrencyHistory
