// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/smstemplate"
	"github.com/google/uuid"
)

// SMSTemplate is the model entity for the SMSTemplate schema.
type SMSTemplate struct {
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
	// LangID holds the value of the "lang_id" field.
	LangID uuid.UUID `json:"lang_id,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Message holds the value of the "message" field.
	Message      string `json:"message,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SMSTemplate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case smstemplate.FieldID, smstemplate.FieldCreatedAt, smstemplate.FieldUpdatedAt, smstemplate.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case smstemplate.FieldUsedFor, smstemplate.FieldSubject, smstemplate.FieldMessage:
			values[i] = new(sql.NullString)
		case smstemplate.FieldEntID, smstemplate.FieldAppID, smstemplate.FieldLangID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SMSTemplate fields.
func (st *SMSTemplate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case smstemplate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			st.ID = uint32(value.Int64)
		case smstemplate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				st.CreatedAt = uint32(value.Int64)
			}
		case smstemplate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				st.UpdatedAt = uint32(value.Int64)
			}
		case smstemplate.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				st.DeletedAt = uint32(value.Int64)
			}
		case smstemplate.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				st.EntID = *value
			}
		case smstemplate.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				st.AppID = *value
			}
		case smstemplate.FieldLangID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field lang_id", values[i])
			} else if value != nil {
				st.LangID = *value
			}
		case smstemplate.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				st.UsedFor = value.String
			}
		case smstemplate.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				st.Subject = value.String
			}
		case smstemplate.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				st.Message = value.String
			}
		default:
			st.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SMSTemplate.
// This includes values selected through modifiers, order, etc.
func (st *SMSTemplate) Value(name string) (ent.Value, error) {
	return st.selectValues.Get(name)
}

// Update returns a builder for updating this SMSTemplate.
// Note that you need to call SMSTemplate.Unwrap() before calling this method if this SMSTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *SMSTemplate) Update() *SMSTemplateUpdateOne {
	return NewSMSTemplateClient(st.config).UpdateOne(st)
}

// Unwrap unwraps the SMSTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *SMSTemplate) Unwrap() *SMSTemplate {
	_tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("generated: SMSTemplate is not a transactional entity")
	}
	st.config.driver = _tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *SMSTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("SMSTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", st.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", st.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", st.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", st.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", st.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", st.AppID))
	builder.WriteString(", ")
	builder.WriteString("lang_id=")
	builder.WriteString(fmt.Sprintf("%v", st.LangID))
	builder.WriteString(", ")
	builder.WriteString("used_for=")
	builder.WriteString(st.UsedFor)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(st.Subject)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(st.Message)
	builder.WriteByte(')')
	return builder.String()
}

// SMSTemplates is a parsable slice of SMSTemplate.
type SMSTemplates []*SMSTemplate
