// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscription"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSubscription is the model entity for the AppSubscription schema.
type AppSubscription struct {
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
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// UsdPrice holds the value of the "usd_price" field.
	UsdPrice decimal.Decimal `json:"usd_price,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID string `json:"product_id,omitempty"`
	// PlanID holds the value of the "plan_id" field.
	PlanID string `json:"plan_id,omitempty"`
	// TrialUnits holds the value of the "trial_units" field.
	TrialUnits uint32 `json:"trial_units,omitempty"`
	// TrialUsdPrice holds the value of the "trial_usd_price" field.
	TrialUsdPrice decimal.Decimal `json:"trial_usd_price,omitempty"`
	// PriceFiatID holds the value of the "price_fiat_id" field.
	PriceFiatID uuid.UUID `json:"price_fiat_id,omitempty"`
	// FiatPrice holds the value of the "fiat_price" field.
	FiatPrice decimal.Decimal `json:"fiat_price,omitempty"`
	// TrialFiatPrice holds the value of the "trial_fiat_price" field.
	TrialFiatPrice decimal.Decimal `json:"trial_fiat_price,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppSubscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appsubscription.FieldUsdPrice, appsubscription.FieldTrialUsdPrice, appsubscription.FieldFiatPrice, appsubscription.FieldTrialFiatPrice:
			values[i] = new(decimal.Decimal)
		case appsubscription.FieldID, appsubscription.FieldCreatedAt, appsubscription.FieldUpdatedAt, appsubscription.FieldDeletedAt, appsubscription.FieldTrialUnits:
			values[i] = new(sql.NullInt64)
		case appsubscription.FieldProductID, appsubscription.FieldPlanID:
			values[i] = new(sql.NullString)
		case appsubscription.FieldEntID, appsubscription.FieldAppGoodID, appsubscription.FieldPriceFiatID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppSubscription fields.
func (as *AppSubscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appsubscription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = uint32(value.Int64)
		case appsubscription.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				as.EntID = *value
			}
		case appsubscription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = uint32(value.Int64)
			}
		case appsubscription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = uint32(value.Int64)
			}
		case appsubscription.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				as.DeletedAt = uint32(value.Int64)
			}
		case appsubscription.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				as.AppGoodID = *value
			}
		case appsubscription.FieldUsdPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field usd_price", values[i])
			} else if value != nil {
				as.UsdPrice = *value
			}
		case appsubscription.FieldProductID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				as.ProductID = value.String
			}
		case appsubscription.FieldPlanID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plan_id", values[i])
			} else if value.Valid {
				as.PlanID = value.String
			}
		case appsubscription.FieldTrialUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field trial_units", values[i])
			} else if value.Valid {
				as.TrialUnits = uint32(value.Int64)
			}
		case appsubscription.FieldTrialUsdPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field trial_usd_price", values[i])
			} else if value != nil {
				as.TrialUsdPrice = *value
			}
		case appsubscription.FieldPriceFiatID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field price_fiat_id", values[i])
			} else if value != nil {
				as.PriceFiatID = *value
			}
		case appsubscription.FieldFiatPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field fiat_price", values[i])
			} else if value != nil {
				as.FiatPrice = *value
			}
		case appsubscription.FieldTrialFiatPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field trial_fiat_price", values[i])
			} else if value != nil {
				as.TrialFiatPrice = *value
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppSubscription.
// This includes values selected through modifiers, order, etc.
func (as *AppSubscription) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// Update returns a builder for updating this AppSubscription.
// Note that you need to call AppSubscription.Unwrap() before calling this method if this AppSubscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AppSubscription) Update() *AppSubscriptionUpdateOne {
	return NewAppSubscriptionClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AppSubscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AppSubscription) Unwrap() *AppSubscription {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("generated: AppSubscription is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AppSubscription) String() string {
	var builder strings.Builder
	builder.WriteString("AppSubscription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", as.EntID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", as.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", as.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", as.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", as.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("usd_price=")
	builder.WriteString(fmt.Sprintf("%v", as.UsdPrice))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(as.ProductID)
	builder.WriteString(", ")
	builder.WriteString("plan_id=")
	builder.WriteString(as.PlanID)
	builder.WriteString(", ")
	builder.WriteString("trial_units=")
	builder.WriteString(fmt.Sprintf("%v", as.TrialUnits))
	builder.WriteString(", ")
	builder.WriteString("trial_usd_price=")
	builder.WriteString(fmt.Sprintf("%v", as.TrialUsdPrice))
	builder.WriteString(", ")
	builder.WriteString("price_fiat_id=")
	builder.WriteString(fmt.Sprintf("%v", as.PriceFiatID))
	builder.WriteString(", ")
	builder.WriteString("fiat_price=")
	builder.WriteString(fmt.Sprintf("%v", as.FiatPrice))
	builder.WriteString(", ")
	builder.WriteString("trial_fiat_price=")
	builder.WriteString(fmt.Sprintf("%v", as.TrialFiatPrice))
	builder.WriteByte(')')
	return builder.String()
}

// AppSubscriptions is a parsable slice of AppSubscription.
type AppSubscriptions []*AppSubscription
