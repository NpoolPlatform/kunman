// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderpaymentstatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPaymentStatement is the model entity for the OrderPaymentStatement schema.
type OrderPaymentStatement struct {
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
	// StatementID holds the value of the "statement_id" field.
	StatementID uuid.UUID `json:"statement_id,omitempty"`
	// PaymentCoinTypeID holds the value of the "payment_coin_type_id" field.
	PaymentCoinTypeID uuid.UUID `json:"payment_coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// CommissionAmount holds the value of the "commission_amount" field.
	CommissionAmount decimal.Decimal `json:"commission_amount,omitempty"`
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderPaymentStatement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderpaymentstatement.FieldAmount, orderpaymentstatement.FieldCommissionAmount:
			values[i] = new(decimal.Decimal)
		case orderpaymentstatement.FieldID, orderpaymentstatement.FieldCreatedAt, orderpaymentstatement.FieldUpdatedAt, orderpaymentstatement.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case orderpaymentstatement.FieldEntID, orderpaymentstatement.FieldStatementID, orderpaymentstatement.FieldPaymentCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderPaymentStatement fields.
func (ops *OrderPaymentStatement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderpaymentstatement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ops.ID = uint32(value.Int64)
		case orderpaymentstatement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ops.CreatedAt = uint32(value.Int64)
			}
		case orderpaymentstatement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ops.UpdatedAt = uint32(value.Int64)
			}
		case orderpaymentstatement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ops.DeletedAt = uint32(value.Int64)
			}
		case orderpaymentstatement.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ops.EntID = *value
			}
		case orderpaymentstatement.FieldStatementID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field statement_id", values[i])
			} else if value != nil {
				ops.StatementID = *value
			}
		case orderpaymentstatement.FieldPaymentCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field payment_coin_type_id", values[i])
			} else if value != nil {
				ops.PaymentCoinTypeID = *value
			}
		case orderpaymentstatement.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				ops.Amount = *value
			}
		case orderpaymentstatement.FieldCommissionAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field commission_amount", values[i])
			} else if value != nil {
				ops.CommissionAmount = *value
			}
		default:
			ops.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderPaymentStatement.
// This includes values selected through modifiers, order, etc.
func (ops *OrderPaymentStatement) Value(name string) (ent.Value, error) {
	return ops.selectValues.Get(name)
}

// Update returns a builder for updating this OrderPaymentStatement.
// Note that you need to call OrderPaymentStatement.Unwrap() before calling this method if this OrderPaymentStatement
// was returned from a transaction, and the transaction was committed or rolled back.
func (ops *OrderPaymentStatement) Update() *OrderPaymentStatementUpdateOne {
	return NewOrderPaymentStatementClient(ops.config).UpdateOne(ops)
}

// Unwrap unwraps the OrderPaymentStatement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ops *OrderPaymentStatement) Unwrap() *OrderPaymentStatement {
	_tx, ok := ops.config.driver.(*txDriver)
	if !ok {
		panic("generated: OrderPaymentStatement is not a transactional entity")
	}
	ops.config.driver = _tx.drv
	return ops
}

// String implements the fmt.Stringer.
func (ops *OrderPaymentStatement) String() string {
	var builder strings.Builder
	builder.WriteString("OrderPaymentStatement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ops.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ops.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ops.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ops.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ops.EntID))
	builder.WriteString(", ")
	builder.WriteString("statement_id=")
	builder.WriteString(fmt.Sprintf("%v", ops.StatementID))
	builder.WriteString(", ")
	builder.WriteString("payment_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ops.PaymentCoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", ops.Amount))
	builder.WriteString(", ")
	builder.WriteString("commission_amount=")
	builder.WriteString(fmt.Sprintf("%v", ops.CommissionAmount))
	builder.WriteByte(')')
	return builder.String()
}

// OrderPaymentStatements is a parsable slice of OrderPaymentStatement.
type OrderPaymentStatements []*OrderPaymentStatement
