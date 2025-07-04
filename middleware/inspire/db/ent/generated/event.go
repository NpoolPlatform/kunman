// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/event"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Event is the model entity for the Event schema.
type Event struct {
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
	// EventType holds the value of the "event_type" field.
	EventType string `json:"event_type,omitempty"`
	// Credits holds the value of the "credits" field.
	Credits decimal.Decimal `json:"credits,omitempty"`
	// CreditsPerUsd holds the value of the "credits_per_usd" field.
	CreditsPerUsd decimal.Decimal `json:"credits_per_usd,omitempty"`
	// MaxConsecutive holds the value of the "max_consecutive" field.
	MaxConsecutive uint32 `json:"max_consecutive,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// InviterLayers holds the value of the "inviter_layers" field.
	InviterLayers uint32 `json:"inviter_layers,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldCredits, event.FieldCreditsPerUsd:
			values[i] = new(decimal.Decimal)
		case event.FieldID, event.FieldCreatedAt, event.FieldUpdatedAt, event.FieldDeletedAt, event.FieldMaxConsecutive, event.FieldInviterLayers:
			values[i] = new(sql.NullInt64)
		case event.FieldEventType:
			values[i] = new(sql.NullString)
		case event.FieldEntID, event.FieldAppID, event.FieldGoodID, event.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = uint32(value.Int64)
		case event.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = uint32(value.Int64)
			}
		case event.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = uint32(value.Int64)
			}
		case event.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				e.DeletedAt = uint32(value.Int64)
			}
		case event.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				e.EntID = *value
			}
		case event.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				e.AppID = *value
			}
		case event.FieldEventType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_type", values[i])
			} else if value.Valid {
				e.EventType = value.String
			}
		case event.FieldCredits:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field credits", values[i])
			} else if value != nil {
				e.Credits = *value
			}
		case event.FieldCreditsPerUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field credits_per_usd", values[i])
			} else if value != nil {
				e.CreditsPerUsd = *value
			}
		case event.FieldMaxConsecutive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_consecutive", values[i])
			} else if value.Valid {
				e.MaxConsecutive = uint32(value.Int64)
			}
		case event.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				e.GoodID = *value
			}
		case event.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				e.AppGoodID = *value
			}
		case event.FieldInviterLayers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field inviter_layers", values[i])
			} else if value.Valid {
				e.InviterLayers = uint32(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Event.
// This includes values selected through modifiers, order, etc.
func (e *Event) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return NewEventClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("generated: Event is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", e.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", e.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", e.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", e.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", e.AppID))
	builder.WriteString(", ")
	builder.WriteString("event_type=")
	builder.WriteString(e.EventType)
	builder.WriteString(", ")
	builder.WriteString("credits=")
	builder.WriteString(fmt.Sprintf("%v", e.Credits))
	builder.WriteString(", ")
	builder.WriteString("credits_per_usd=")
	builder.WriteString(fmt.Sprintf("%v", e.CreditsPerUsd))
	builder.WriteString(", ")
	builder.WriteString("max_consecutive=")
	builder.WriteString(fmt.Sprintf("%v", e.MaxConsecutive))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", e.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", e.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("inviter_layers=")
	builder.WriteString(fmt.Sprintf("%v", e.InviterLayers))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event
