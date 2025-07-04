// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/achievement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Achievement is the model entity for the Achievement schema.
type Achievement struct {
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
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// TotalUnitsV1 holds the value of the "total_units_v1" field.
	TotalUnitsV1 decimal.Decimal `json:"total_units_v1,omitempty"`
	// SelfUnitsV1 holds the value of the "self_units_v1" field.
	SelfUnitsV1 decimal.Decimal `json:"self_units_v1,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount decimal.Decimal `json:"total_amount,omitempty"`
	// SelfAmount holds the value of the "self_amount" field.
	SelfAmount decimal.Decimal `json:"self_amount,omitempty"`
	// TotalCommission holds the value of the "total_commission" field.
	TotalCommission decimal.Decimal `json:"total_commission,omitempty"`
	// SelfCommission holds the value of the "self_commission" field.
	SelfCommission decimal.Decimal `json:"self_commission,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Achievement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case achievement.FieldTotalUnitsV1, achievement.FieldSelfUnitsV1, achievement.FieldTotalAmount, achievement.FieldSelfAmount, achievement.FieldTotalCommission, achievement.FieldSelfCommission:
			values[i] = new(decimal.Decimal)
		case achievement.FieldID, achievement.FieldCreatedAt, achievement.FieldUpdatedAt, achievement.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case achievement.FieldEntID, achievement.FieldAppID, achievement.FieldUserID, achievement.FieldGoodID, achievement.FieldAppGoodID, achievement.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Achievement fields.
func (a *Achievement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case achievement.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint32(value.Int64)
		case achievement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case achievement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case achievement.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case achievement.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				a.EntID = *value
			}
		case achievement.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				a.AppID = *value
			}
		case achievement.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				a.UserID = *value
			}
		case achievement.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				a.GoodID = *value
			}
		case achievement.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				a.AppGoodID = *value
			}
		case achievement.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				a.CoinTypeID = *value
			}
		case achievement.FieldTotalUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_units_v1", values[i])
			} else if value != nil {
				a.TotalUnitsV1 = *value
			}
		case achievement.FieldSelfUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_units_v1", values[i])
			} else if value != nil {
				a.SelfUnitsV1 = *value
			}
		case achievement.FieldTotalAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value != nil {
				a.TotalAmount = *value
			}
		case achievement.FieldSelfAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_amount", values[i])
			} else if value != nil {
				a.SelfAmount = *value
			}
		case achievement.FieldTotalCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_commission", values[i])
			} else if value != nil {
				a.TotalCommission = *value
			}
		case achievement.FieldSelfCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_commission", values[i])
			} else if value != nil {
				a.SelfCommission = *value
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Achievement.
// This includes values selected through modifiers, order, etc.
func (a *Achievement) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Achievement.
// Note that you need to call Achievement.Unwrap() before calling this method if this Achievement
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Achievement) Update() *AchievementUpdateOne {
	return NewAchievementClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Achievement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Achievement) Unwrap() *Achievement {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("generated: Achievement is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Achievement) String() string {
	var builder strings.Builder
	builder.WriteString("Achievement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", a.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("total_units_v1=")
	builder.WriteString(fmt.Sprintf("%v", a.TotalUnitsV1))
	builder.WriteString(", ")
	builder.WriteString("self_units_v1=")
	builder.WriteString(fmt.Sprintf("%v", a.SelfUnitsV1))
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", a.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("self_amount=")
	builder.WriteString(fmt.Sprintf("%v", a.SelfAmount))
	builder.WriteString(", ")
	builder.WriteString("total_commission=")
	builder.WriteString(fmt.Sprintf("%v", a.TotalCommission))
	builder.WriteString(", ")
	builder.WriteString("self_commission=")
	builder.WriteString(fmt.Sprintf("%v", a.SelfCommission))
	builder.WriteByte(')')
	return builder.String()
}

// Achievements is a parsable slice of Achievement.
type Achievements []*Achievement
