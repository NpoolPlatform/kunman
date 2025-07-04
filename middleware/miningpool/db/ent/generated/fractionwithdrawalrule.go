// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/fractionwithdrawalrule"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FractionWithdrawalRule is the model entity for the FractionWithdrawalRule schema.
type FractionWithdrawalRule struct {
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
	// PoolCoinTypeID holds the value of the "pool_coin_type_id" field.
	PoolCoinTypeID uuid.UUID `json:"pool_coin_type_id,omitempty"`
	// WithdrawInterval holds the value of the "withdraw_interval" field.
	WithdrawInterval uint32 `json:"withdraw_interval,omitempty"`
	// LeastWithdrawalAmount holds the value of the "least_withdrawal_amount" field.
	LeastWithdrawalAmount decimal.Decimal `json:"least_withdrawal_amount,omitempty"`
	// PayoutThreshold holds the value of the "payout_threshold" field.
	PayoutThreshold decimal.Decimal `json:"payout_threshold,omitempty"`
	// WithdrawFee holds the value of the "withdraw_fee" field.
	WithdrawFee  decimal.Decimal `json:"withdraw_fee,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FractionWithdrawalRule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fractionwithdrawalrule.FieldLeastWithdrawalAmount, fractionwithdrawalrule.FieldPayoutThreshold, fractionwithdrawalrule.FieldWithdrawFee:
			values[i] = new(decimal.Decimal)
		case fractionwithdrawalrule.FieldID, fractionwithdrawalrule.FieldCreatedAt, fractionwithdrawalrule.FieldUpdatedAt, fractionwithdrawalrule.FieldDeletedAt, fractionwithdrawalrule.FieldWithdrawInterval:
			values[i] = new(sql.NullInt64)
		case fractionwithdrawalrule.FieldEntID, fractionwithdrawalrule.FieldPoolCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FractionWithdrawalRule fields.
func (fwr *FractionWithdrawalRule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fractionwithdrawalrule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fwr.ID = uint32(value.Int64)
		case fractionwithdrawalrule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fwr.CreatedAt = uint32(value.Int64)
			}
		case fractionwithdrawalrule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fwr.UpdatedAt = uint32(value.Int64)
			}
		case fractionwithdrawalrule.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				fwr.DeletedAt = uint32(value.Int64)
			}
		case fractionwithdrawalrule.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				fwr.EntID = *value
			}
		case fractionwithdrawalrule.FieldPoolCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field pool_coin_type_id", values[i])
			} else if value != nil {
				fwr.PoolCoinTypeID = *value
			}
		case fractionwithdrawalrule.FieldWithdrawInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_interval", values[i])
			} else if value.Valid {
				fwr.WithdrawInterval = uint32(value.Int64)
			}
		case fractionwithdrawalrule.FieldLeastWithdrawalAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field least_withdrawal_amount", values[i])
			} else if value != nil {
				fwr.LeastWithdrawalAmount = *value
			}
		case fractionwithdrawalrule.FieldPayoutThreshold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payout_threshold", values[i])
			} else if value != nil {
				fwr.PayoutThreshold = *value
			}
		case fractionwithdrawalrule.FieldWithdrawFee:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_fee", values[i])
			} else if value != nil {
				fwr.WithdrawFee = *value
			}
		default:
			fwr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FractionWithdrawalRule.
// This includes values selected through modifiers, order, etc.
func (fwr *FractionWithdrawalRule) Value(name string) (ent.Value, error) {
	return fwr.selectValues.Get(name)
}

// Update returns a builder for updating this FractionWithdrawalRule.
// Note that you need to call FractionWithdrawalRule.Unwrap() before calling this method if this FractionWithdrawalRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (fwr *FractionWithdrawalRule) Update() *FractionWithdrawalRuleUpdateOne {
	return NewFractionWithdrawalRuleClient(fwr.config).UpdateOne(fwr)
}

// Unwrap unwraps the FractionWithdrawalRule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fwr *FractionWithdrawalRule) Unwrap() *FractionWithdrawalRule {
	_tx, ok := fwr.config.driver.(*txDriver)
	if !ok {
		panic("generated: FractionWithdrawalRule is not a transactional entity")
	}
	fwr.config.driver = _tx.drv
	return fwr
}

// String implements the fmt.Stringer.
func (fwr *FractionWithdrawalRule) String() string {
	var builder strings.Builder
	builder.WriteString("FractionWithdrawalRule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fwr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", fwr.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", fwr.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", fwr.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", fwr.EntID))
	builder.WriteString(", ")
	builder.WriteString("pool_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", fwr.PoolCoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("withdraw_interval=")
	builder.WriteString(fmt.Sprintf("%v", fwr.WithdrawInterval))
	builder.WriteString(", ")
	builder.WriteString("least_withdrawal_amount=")
	builder.WriteString(fmt.Sprintf("%v", fwr.LeastWithdrawalAmount))
	builder.WriteString(", ")
	builder.WriteString("payout_threshold=")
	builder.WriteString(fmt.Sprintf("%v", fwr.PayoutThreshold))
	builder.WriteString(", ")
	builder.WriteString("withdraw_fee=")
	builder.WriteString(fmt.Sprintf("%v", fwr.WithdrawFee))
	builder.WriteByte(')')
	return builder.String()
}

// FractionWithdrawalRules is a parsable slice of FractionWithdrawalRule.
type FractionWithdrawalRules []*FractionWithdrawalRule
