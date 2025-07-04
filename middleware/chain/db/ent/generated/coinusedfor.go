// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinusedfor"
	"github.com/google/uuid"
)

// CoinUsedFor is the model entity for the CoinUsedFor schema.
type CoinUsedFor struct {
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
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority     uint32 `json:"priority,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinUsedFor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinusedfor.FieldID, coinusedfor.FieldCreatedAt, coinusedfor.FieldUpdatedAt, coinusedfor.FieldDeletedAt, coinusedfor.FieldPriority:
			values[i] = new(sql.NullInt64)
		case coinusedfor.FieldUsedFor:
			values[i] = new(sql.NullString)
		case coinusedfor.FieldEntID, coinusedfor.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinUsedFor fields.
func (cuf *CoinUsedFor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinusedfor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cuf.ID = uint32(value.Int64)
		case coinusedfor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cuf.CreatedAt = uint32(value.Int64)
			}
		case coinusedfor.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cuf.UpdatedAt = uint32(value.Int64)
			}
		case coinusedfor.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cuf.DeletedAt = uint32(value.Int64)
			}
		case coinusedfor.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				cuf.EntID = *value
			}
		case coinusedfor.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cuf.CoinTypeID = *value
			}
		case coinusedfor.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				cuf.UsedFor = value.String
			}
		case coinusedfor.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				cuf.Priority = uint32(value.Int64)
			}
		default:
			cuf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CoinUsedFor.
// This includes values selected through modifiers, order, etc.
func (cuf *CoinUsedFor) Value(name string) (ent.Value, error) {
	return cuf.selectValues.Get(name)
}

// Update returns a builder for updating this CoinUsedFor.
// Note that you need to call CoinUsedFor.Unwrap() before calling this method if this CoinUsedFor
// was returned from a transaction, and the transaction was committed or rolled back.
func (cuf *CoinUsedFor) Update() *CoinUsedForUpdateOne {
	return NewCoinUsedForClient(cuf.config).UpdateOne(cuf)
}

// Unwrap unwraps the CoinUsedFor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cuf *CoinUsedFor) Unwrap() *CoinUsedFor {
	_tx, ok := cuf.config.driver.(*txDriver)
	if !ok {
		panic("generated: CoinUsedFor is not a transactional entity")
	}
	cuf.config.driver = _tx.drv
	return cuf
}

// String implements the fmt.Stringer.
func (cuf *CoinUsedFor) String() string {
	var builder strings.Builder
	builder.WriteString("CoinUsedFor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cuf.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cuf.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cuf.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cuf.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", cuf.EntID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cuf.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("used_for=")
	builder.WriteString(cuf.UsedFor)
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", cuf.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// CoinUsedFors is a parsable slice of CoinUsedFor.
type CoinUsedFors []*CoinUsedFor
