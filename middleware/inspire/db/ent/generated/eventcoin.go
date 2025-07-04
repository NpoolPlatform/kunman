// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// EventCoin is the model entity for the EventCoin schema.
type EventCoin struct {
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
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// EventID holds the value of the "event_id" field.
	EventID uuid.UUID `json:"event_id,omitempty"`
	// CoinConfigID holds the value of the "coin_config_id" field.
	CoinConfigID uuid.UUID `json:"coin_config_id,omitempty"`
	// CoinValue holds the value of the "coin_value" field.
	CoinValue decimal.Decimal `json:"coin_value,omitempty"`
	// CoinPerUsd holds the value of the "coin_per_usd" field.
	CoinPerUsd   decimal.Decimal `json:"coin_per_usd,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EventCoin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case eventcoin.FieldCoinValue, eventcoin.FieldCoinPerUsd:
			values[i] = new(decimal.Decimal)
		case eventcoin.FieldID, eventcoin.FieldCreatedAt, eventcoin.FieldUpdatedAt, eventcoin.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case eventcoin.FieldEntID, eventcoin.FieldAppID, eventcoin.FieldEventID, eventcoin.FieldCoinConfigID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EventCoin fields.
func (ec *EventCoin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case eventcoin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = uint32(value.Int64)
		case eventcoin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ec.CreatedAt = uint32(value.Int64)
			}
		case eventcoin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ec.UpdatedAt = uint32(value.Int64)
			}
		case eventcoin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ec.DeletedAt = uint32(value.Int64)
			}
		case eventcoin.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ec.EntID = *value
			}
		case eventcoin.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ec.AppID = *value
			}
		case eventcoin.FieldEventID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value != nil {
				ec.EventID = *value
			}
		case eventcoin.FieldCoinConfigID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_config_id", values[i])
			} else if value != nil {
				ec.CoinConfigID = *value
			}
		case eventcoin.FieldCoinValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field coin_value", values[i])
			} else if value != nil {
				ec.CoinValue = *value
			}
		case eventcoin.FieldCoinPerUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field coin_per_usd", values[i])
			} else if value != nil {
				ec.CoinPerUsd = *value
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EventCoin.
// This includes values selected through modifiers, order, etc.
func (ec *EventCoin) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// Update returns a builder for updating this EventCoin.
// Note that you need to call EventCoin.Unwrap() before calling this method if this EventCoin
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EventCoin) Update() *EventCoinUpdateOne {
	return NewEventCoinClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EventCoin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EventCoin) Unwrap() *EventCoin {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("generated: EventCoin is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EventCoin) String() string {
	var builder strings.Builder
	builder.WriteString("EventCoin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ec.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.AppID))
	builder.WriteString(", ")
	builder.WriteString("event_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.EventID))
	builder.WriteString(", ")
	builder.WriteString("coin_config_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.CoinConfigID))
	builder.WriteString(", ")
	builder.WriteString("coin_value=")
	builder.WriteString(fmt.Sprintf("%v", ec.CoinValue))
	builder.WriteString(", ")
	builder.WriteString("coin_per_usd=")
	builder.WriteString(fmt.Sprintf("%v", ec.CoinPerUsd))
	builder.WriteByte(')')
	return builder.String()
}

// EventCoins is a parsable slice of EventCoin.
type EventCoins []*EventCoin
