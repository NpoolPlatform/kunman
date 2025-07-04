// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FeeOrder is the model entity for the FeeOrder schema.
type FeeOrder struct {
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
	// GoodValueUsd holds the value of the "good_value_usd" field.
	GoodValueUsd decimal.Decimal `json:"good_value_usd,omitempty"`
	// PaymentAmountUsd holds the value of the "payment_amount_usd" field.
	PaymentAmountUsd decimal.Decimal `json:"payment_amount_usd,omitempty"`
	// DiscountAmountUsd holds the value of the "discount_amount_usd" field.
	DiscountAmountUsd decimal.Decimal `json:"discount_amount_usd,omitempty"`
	// PromotionID holds the value of the "promotion_id" field.
	PromotionID uuid.UUID `json:"promotion_id,omitempty"`
	// DurationSeconds holds the value of the "duration_seconds" field.
	DurationSeconds uint32 `json:"duration_seconds,omitempty"`
	selectValues    sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeeOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feeorder.FieldGoodValueUsd, feeorder.FieldPaymentAmountUsd, feeorder.FieldDiscountAmountUsd:
			values[i] = new(decimal.Decimal)
		case feeorder.FieldID, feeorder.FieldCreatedAt, feeorder.FieldUpdatedAt, feeorder.FieldDeletedAt, feeorder.FieldDurationSeconds:
			values[i] = new(sql.NullInt64)
		case feeorder.FieldEntID, feeorder.FieldOrderID, feeorder.FieldPromotionID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeeOrder fields.
func (fo *FeeOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feeorder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fo.ID = uint32(value.Int64)
		case feeorder.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				fo.EntID = *value
			}
		case feeorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fo.CreatedAt = uint32(value.Int64)
			}
		case feeorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fo.UpdatedAt = uint32(value.Int64)
			}
		case feeorder.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fo.DeletedAt = uint32(value.Int64)
			}
		case feeorder.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				fo.OrderID = *value
			}
		case feeorder.FieldGoodValueUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field good_value_usd", values[i])
			} else if value != nil {
				fo.GoodValueUsd = *value
			}
		case feeorder.FieldPaymentAmountUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_amount_usd", values[i])
			} else if value != nil {
				fo.PaymentAmountUsd = *value
			}
		case feeorder.FieldDiscountAmountUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field discount_amount_usd", values[i])
			} else if value != nil {
				fo.DiscountAmountUsd = *value
			}
		case feeorder.FieldPromotionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field promotion_id", values[i])
			} else if value != nil {
				fo.PromotionID = *value
			}
		case feeorder.FieldDurationSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_seconds", values[i])
			} else if value.Valid {
				fo.DurationSeconds = uint32(value.Int64)
			}
		default:
			fo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeeOrder.
// This includes values selected through modifiers, order, etc.
func (fo *FeeOrder) Value(name string) (ent.Value, error) {
	return fo.selectValues.Get(name)
}

// Update returns a builder for updating this FeeOrder.
// Note that you need to call FeeOrder.Unwrap() before calling this method if this FeeOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (fo *FeeOrder) Update() *FeeOrderUpdateOne {
	return NewFeeOrderClient(fo.config).UpdateOne(fo)
}

// Unwrap unwraps the FeeOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fo *FeeOrder) Unwrap() *FeeOrder {
	_tx, ok := fo.config.driver.(*txDriver)
	if !ok {
		panic("generated: FeeOrder is not a transactional entity")
	}
	fo.config.driver = _tx.drv
	return fo
}

// String implements the fmt.Stringer.
func (fo *FeeOrder) String() string {
	var builder strings.Builder
	builder.WriteString("FeeOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fo.ID))
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", fo.EntID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", fo.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fo.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fo.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", fo.OrderID))
	builder.WriteString(", ")
	builder.WriteString("good_value_usd=")
	builder.WriteString(fmt.Sprintf("%v", fo.GoodValueUsd))
	builder.WriteString(", ")
	builder.WriteString("payment_amount_usd=")
	builder.WriteString(fmt.Sprintf("%v", fo.PaymentAmountUsd))
	builder.WriteString(", ")
	builder.WriteString("discount_amount_usd=")
	builder.WriteString(fmt.Sprintf("%v", fo.DiscountAmountUsd))
	builder.WriteString(", ")
	builder.WriteString("promotion_id=")
	builder.WriteString(fmt.Sprintf("%v", fo.PromotionID))
	builder.WriteString(", ")
	builder.WriteString("duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", fo.DurationSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// FeeOrders is a parsable slice of FeeOrder.
type FeeOrders []*FeeOrder
