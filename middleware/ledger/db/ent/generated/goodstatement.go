// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/goodstatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodStatement is the model entity for the GoodStatement schema.
type GoodStatement struct {
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
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// ToPlatform holds the value of the "to_platform" field.
	ToPlatform decimal.Decimal `json:"to_platform,omitempty"`
	// ToUser holds the value of the "to_user" field.
	ToUser decimal.Decimal `json:"to_user,omitempty"`
	// TechniqueServiceFeeAmount holds the value of the "technique_service_fee_amount" field.
	TechniqueServiceFeeAmount decimal.Decimal `json:"technique_service_fee_amount,omitempty"`
	// BenefitDate holds the value of the "benefit_date" field.
	BenefitDate  uint32 `json:"benefit_date,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodStatement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodstatement.FieldAmount, goodstatement.FieldToPlatform, goodstatement.FieldToUser, goodstatement.FieldTechniqueServiceFeeAmount:
			values[i] = new(decimal.Decimal)
		case goodstatement.FieldID, goodstatement.FieldCreatedAt, goodstatement.FieldUpdatedAt, goodstatement.FieldDeletedAt, goodstatement.FieldBenefitDate:
			values[i] = new(sql.NullInt64)
		case goodstatement.FieldEntID, goodstatement.FieldGoodID, goodstatement.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodStatement fields.
func (gs *GoodStatement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodstatement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gs.ID = uint32(value.Int64)
		case goodstatement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gs.CreatedAt = uint32(value.Int64)
			}
		case goodstatement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gs.UpdatedAt = uint32(value.Int64)
			}
		case goodstatement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gs.DeletedAt = uint32(value.Int64)
			}
		case goodstatement.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				gs.EntID = *value
			}
		case goodstatement.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gs.GoodID = *value
			}
		case goodstatement.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				gs.CoinTypeID = *value
			}
		case goodstatement.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				gs.Amount = *value
			}
		case goodstatement.FieldToPlatform:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_platform", values[i])
			} else if value != nil {
				gs.ToPlatform = *value
			}
		case goodstatement.FieldToUser:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field to_user", values[i])
			} else if value != nil {
				gs.ToUser = *value
			}
		case goodstatement.FieldTechniqueServiceFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field technique_service_fee_amount", values[i])
			} else if value != nil {
				gs.TechniqueServiceFeeAmount = *value
			}
		case goodstatement.FieldBenefitDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field benefit_date", values[i])
			} else if value.Valid {
				gs.BenefitDate = uint32(value.Int64)
			}
		default:
			gs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GoodStatement.
// This includes values selected through modifiers, order, etc.
func (gs *GoodStatement) Value(name string) (ent.Value, error) {
	return gs.selectValues.Get(name)
}

// Update returns a builder for updating this GoodStatement.
// Note that you need to call GoodStatement.Unwrap() before calling this method if this GoodStatement
// was returned from a transaction, and the transaction was committed or rolled back.
func (gs *GoodStatement) Update() *GoodStatementUpdateOne {
	return NewGoodStatementClient(gs.config).UpdateOne(gs)
}

// Unwrap unwraps the GoodStatement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gs *GoodStatement) Unwrap() *GoodStatement {
	_tx, ok := gs.config.driver.(*txDriver)
	if !ok {
		panic("generated: GoodStatement is not a transactional entity")
	}
	gs.config.driver = _tx.drv
	return gs
}

// String implements the fmt.Stringer.
func (gs *GoodStatement) String() string {
	var builder strings.Builder
	builder.WriteString("GoodStatement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", gs.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", gs.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", gs.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", gs.EntID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", gs.GoodID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", gs.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", gs.Amount))
	builder.WriteString(", ")
	builder.WriteString("to_platform=")
	builder.WriteString(fmt.Sprintf("%v", gs.ToPlatform))
	builder.WriteString(", ")
	builder.WriteString("to_user=")
	builder.WriteString(fmt.Sprintf("%v", gs.ToUser))
	builder.WriteString(", ")
	builder.WriteString("technique_service_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", gs.TechniqueServiceFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("benefit_date=")
	builder.WriteString(fmt.Sprintf("%v", gs.BenefitDate))
	builder.WriteByte(')')
	return builder.String()
}

// GoodStatements is a parsable slice of GoodStatement.
type GoodStatements []*GoodStatement
