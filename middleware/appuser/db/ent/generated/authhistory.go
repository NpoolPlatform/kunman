// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/authhistory"
	"github.com/google/uuid"
)

// AuthHistory is the model entity for the AuthHistory schema.
type AuthHistory struct {
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
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Resource holds the value of the "resource" field.
	Resource string `json:"resource,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Allowed holds the value of the "allowed" field.
	Allowed      bool `json:"allowed,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authhistory.FieldAllowed:
			values[i] = new(sql.NullBool)
		case authhistory.FieldID, authhistory.FieldCreatedAt, authhistory.FieldUpdatedAt, authhistory.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case authhistory.FieldResource, authhistory.FieldMethod:
			values[i] = new(sql.NullString)
		case authhistory.FieldEntID, authhistory.FieldAppID, authhistory.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthHistory fields.
func (ah *AuthHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ah.ID = uint32(value.Int64)
		case authhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ah.CreatedAt = uint32(value.Int64)
			}
		case authhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ah.UpdatedAt = uint32(value.Int64)
			}
		case authhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ah.DeletedAt = uint32(value.Int64)
			}
		case authhistory.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ah.EntID = *value
			}
		case authhistory.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ah.AppID = *value
			}
		case authhistory.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ah.UserID = *value
			}
		case authhistory.FieldResource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource", values[i])
			} else if value.Valid {
				ah.Resource = value.String
			}
		case authhistory.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				ah.Method = value.String
			}
		case authhistory.FieldAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field allowed", values[i])
			} else if value.Valid {
				ah.Allowed = value.Bool
			}
		default:
			ah.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AuthHistory.
// This includes values selected through modifiers, order, etc.
func (ah *AuthHistory) Value(name string) (ent.Value, error) {
	return ah.selectValues.Get(name)
}

// Update returns a builder for updating this AuthHistory.
// Note that you need to call AuthHistory.Unwrap() before calling this method if this AuthHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ah *AuthHistory) Update() *AuthHistoryUpdateOne {
	return NewAuthHistoryClient(ah.config).UpdateOne(ah)
}

// Unwrap unwraps the AuthHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ah *AuthHistory) Unwrap() *AuthHistory {
	_tx, ok := ah.config.driver.(*txDriver)
	if !ok {
		panic("generated: AuthHistory is not a transactional entity")
	}
	ah.config.driver = _tx.drv
	return ah
}

// String implements the fmt.Stringer.
func (ah *AuthHistory) String() string {
	var builder strings.Builder
	builder.WriteString("AuthHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ah.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ah.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ah.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ah.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ah.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ah.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ah.UserID))
	builder.WriteString(", ")
	builder.WriteString("resource=")
	builder.WriteString(ah.Resource)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(ah.Method)
	builder.WriteString(", ")
	builder.WriteString("allowed=")
	builder.WriteString(fmt.Sprintf("%v", ah.Allowed))
	builder.WriteByte(')')
	return builder.String()
}

// AuthHistories is a parsable slice of AuthHistory.
type AuthHistories []*AuthHistory
