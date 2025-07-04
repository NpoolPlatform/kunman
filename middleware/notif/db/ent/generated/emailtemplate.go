// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/emailtemplate"
	"github.com/google/uuid"
)

// EmailTemplate is the model entity for the EmailTemplate schema.
type EmailTemplate struct {
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
	// DefaultToUsername holds the value of the "default_to_username" field.
	DefaultToUsername string `json:"default_to_username,omitempty"`
	// UsedFor holds the value of the "used_for" field.
	UsedFor string `json:"used_for,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// ReplyTos holds the value of the "reply_tos" field.
	ReplyTos []string `json:"reply_tos,omitempty"`
	// CcTos holds the value of the "cc_tos" field.
	CcTos []string `json:"cc_tos,omitempty"`
	// Subject holds the value of the "subject" field.
	Subject string `json:"subject,omitempty"`
	// Body holds the value of the "body" field.
	Body         string `json:"body,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmailTemplate) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case emailtemplate.FieldReplyTos, emailtemplate.FieldCcTos:
			values[i] = new([]byte)
		case emailtemplate.FieldID, emailtemplate.FieldCreatedAt, emailtemplate.FieldUpdatedAt, emailtemplate.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case emailtemplate.FieldDefaultToUsername, emailtemplate.FieldUsedFor, emailtemplate.FieldSender, emailtemplate.FieldSubject, emailtemplate.FieldBody:
			values[i] = new(sql.NullString)
		case emailtemplate.FieldEntID, emailtemplate.FieldAppID, emailtemplate.FieldLangID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmailTemplate fields.
func (et *EmailTemplate) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emailtemplate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			et.ID = uint32(value.Int64)
		case emailtemplate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				et.CreatedAt = uint32(value.Int64)
			}
		case emailtemplate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				et.UpdatedAt = uint32(value.Int64)
			}
		case emailtemplate.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				et.DeletedAt = uint32(value.Int64)
			}
		case emailtemplate.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				et.EntID = *value
			}
		case emailtemplate.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				et.AppID = *value
			}
		case emailtemplate.FieldLangID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field lang_id", values[i])
			} else if value != nil {
				et.LangID = *value
			}
		case emailtemplate.FieldDefaultToUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_to_username", values[i])
			} else if value.Valid {
				et.DefaultToUsername = value.String
			}
		case emailtemplate.FieldUsedFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field used_for", values[i])
			} else if value.Valid {
				et.UsedFor = value.String
			}
		case emailtemplate.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				et.Sender = value.String
			}
		case emailtemplate.FieldReplyTos:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field reply_tos", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &et.ReplyTos); err != nil {
					return fmt.Errorf("unmarshal field reply_tos: %w", err)
				}
			}
		case emailtemplate.FieldCcTos:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cc_tos", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &et.CcTos); err != nil {
					return fmt.Errorf("unmarshal field cc_tos: %w", err)
				}
			}
		case emailtemplate.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				et.Subject = value.String
			}
		case emailtemplate.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				et.Body = value.String
			}
		default:
			et.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmailTemplate.
// This includes values selected through modifiers, order, etc.
func (et *EmailTemplate) Value(name string) (ent.Value, error) {
	return et.selectValues.Get(name)
}

// Update returns a builder for updating this EmailTemplate.
// Note that you need to call EmailTemplate.Unwrap() before calling this method if this EmailTemplate
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *EmailTemplate) Update() *EmailTemplateUpdateOne {
	return NewEmailTemplateClient(et.config).UpdateOne(et)
}

// Unwrap unwraps the EmailTemplate entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (et *EmailTemplate) Unwrap() *EmailTemplate {
	_tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("generated: EmailTemplate is not a transactional entity")
	}
	et.config.driver = _tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *EmailTemplate) String() string {
	var builder strings.Builder
	builder.WriteString("EmailTemplate(")
	builder.WriteString(fmt.Sprintf("id=%v, ", et.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", et.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", et.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", et.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", et.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", et.AppID))
	builder.WriteString(", ")
	builder.WriteString("lang_id=")
	builder.WriteString(fmt.Sprintf("%v", et.LangID))
	builder.WriteString(", ")
	builder.WriteString("default_to_username=")
	builder.WriteString(et.DefaultToUsername)
	builder.WriteString(", ")
	builder.WriteString("used_for=")
	builder.WriteString(et.UsedFor)
	builder.WriteString(", ")
	builder.WriteString("sender=")
	builder.WriteString(et.Sender)
	builder.WriteString(", ")
	builder.WriteString("reply_tos=")
	builder.WriteString(fmt.Sprintf("%v", et.ReplyTos))
	builder.WriteString(", ")
	builder.WriteString("cc_tos=")
	builder.WriteString(fmt.Sprintf("%v", et.CcTos))
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(et.Subject)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(et.Body)
	builder.WriteByte(')')
	return builder.String()
}

// EmailTemplates is a parsable slice of EmailTemplate.
type EmailTemplates []*EmailTemplate
