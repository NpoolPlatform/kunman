// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentcontract"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentContract is the model entity for the PaymentContract schema.
type PaymentContract struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount       decimal.Decimal `json:"amount,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentContract) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymentcontract.FieldAmount:
			values[i] = new(decimal.Decimal)
		case paymentcontract.FieldID, paymentcontract.FieldCreatedAt, paymentcontract.FieldUpdatedAt, paymentcontract.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case paymentcontract.FieldEntID, paymentcontract.FieldOrderID, paymentcontract.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentContract fields.
func (pc *PaymentContract) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymentcontract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = uint32(value.Int64)
		case paymentcontract.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				pc.EntID = *value
			}
		case paymentcontract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pc.CreatedAt = uint32(value.Int64)
			}
		case paymentcontract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pc.UpdatedAt = uint32(value.Int64)
			}
		case paymentcontract.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pc.DeletedAt = uint32(value.Int64)
			}
		case paymentcontract.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				pc.OrderID = *value
			}
		case paymentcontract.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				pc.CoinTypeID = *value
			}
		case paymentcontract.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				pc.Amount = *value
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PaymentContract.
// This includes values selected through modifiers, order, etc.
func (pc *PaymentContract) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// Update returns a builder for updating this PaymentContract.
// Note that you need to call PaymentContract.Unwrap() before calling this method if this PaymentContract
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PaymentContract) Update() *PaymentContractUpdateOne {
	return NewPaymentContractClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PaymentContract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PaymentContract) Unwrap() *PaymentContract {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("generated: PaymentContract is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PaymentContract) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentContract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.EntID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.OrderID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pc.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// PaymentContracts is a parsable slice of PaymentContract.
type PaymentContracts []*PaymentContract
