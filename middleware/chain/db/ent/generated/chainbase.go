// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/chainbase"
	"github.com/google/uuid"
)

// ChainBase is the model entity for the ChainBase schema.
type ChainBase struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Logo holds the value of the "logo" field.
	Logo string `json:"logo,omitempty"`
	// NativeUnit holds the value of the "native_unit" field.
	NativeUnit string `json:"native_unit,omitempty"`
	// AtomicUnit holds the value of the "atomic_unit" field.
	AtomicUnit string `json:"atomic_unit,omitempty"`
	// UnitExp holds the value of the "unit_exp" field.
	UnitExp uint32 `json:"unit_exp,omitempty"`
	// Env holds the value of the "env" field.
	Env string `json:"env,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID string `json:"chain_id,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// GasType holds the value of the "gas_type" field.
	GasType      string `json:"gas_type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChainBase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chainbase.FieldID, chainbase.FieldCreatedAt, chainbase.FieldUpdatedAt, chainbase.FieldDeletedAt, chainbase.FieldUnitExp:
			values[i] = new(sql.NullInt64)
		case chainbase.FieldName, chainbase.FieldLogo, chainbase.FieldNativeUnit, chainbase.FieldAtomicUnit, chainbase.FieldEnv, chainbase.FieldChainID, chainbase.FieldNickname, chainbase.FieldGasType:
			values[i] = new(sql.NullString)
		case chainbase.FieldEntID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChainBase fields.
func (cb *ChainBase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chainbase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cb.ID = uint32(value.Int64)
		case chainbase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cb.CreatedAt = uint32(value.Int64)
			}
		case chainbase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cb.UpdatedAt = uint32(value.Int64)
			}
		case chainbase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cb.DeletedAt = uint32(value.Int64)
			}
		case chainbase.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				cb.EntID = *value
			}
		case chainbase.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cb.Name = value.String
			}
		case chainbase.FieldLogo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logo", values[i])
			} else if value.Valid {
				cb.Logo = value.String
			}
		case chainbase.FieldNativeUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field native_unit", values[i])
			} else if value.Valid {
				cb.NativeUnit = value.String
			}
		case chainbase.FieldAtomicUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field atomic_unit", values[i])
			} else if value.Valid {
				cb.AtomicUnit = value.String
			}
		case chainbase.FieldUnitExp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_exp", values[i])
			} else if value.Valid {
				cb.UnitExp = uint32(value.Int64)
			}
		case chainbase.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				cb.Env = value.String
			}
		case chainbase.FieldChainID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				cb.ChainID = value.String
			}
		case chainbase.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				cb.Nickname = value.String
			}
		case chainbase.FieldGasType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gas_type", values[i])
			} else if value.Valid {
				cb.GasType = value.String
			}
		default:
			cb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChainBase.
// This includes values selected through modifiers, order, etc.
func (cb *ChainBase) Value(name string) (ent.Value, error) {
	return cb.selectValues.Get(name)
}

// Update returns a builder for updating this ChainBase.
// Note that you need to call ChainBase.Unwrap() before calling this method if this ChainBase
// was returned from a transaction, and the transaction was committed or rolled back.
func (cb *ChainBase) Update() *ChainBaseUpdateOne {
	return NewChainBaseClient(cb.config).UpdateOne(cb)
}

// Unwrap unwraps the ChainBase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cb *ChainBase) Unwrap() *ChainBase {
	_tx, ok := cb.config.driver.(*txDriver)
	if !ok {
		panic("generated: ChainBase is not a transactional entity")
	}
	cb.config.driver = _tx.drv
	return cb
}

// String implements the fmt.Stringer.
func (cb *ChainBase) String() string {
	var builder strings.Builder
	builder.WriteString("ChainBase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cb.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", cb.EntID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cb.Name)
	builder.WriteString(", ")
	builder.WriteString("logo=")
	builder.WriteString(cb.Logo)
	builder.WriteString(", ")
	builder.WriteString("native_unit=")
	builder.WriteString(cb.NativeUnit)
	builder.WriteString(", ")
	builder.WriteString("atomic_unit=")
	builder.WriteString(cb.AtomicUnit)
	builder.WriteString(", ")
	builder.WriteString("unit_exp=")
	builder.WriteString(fmt.Sprintf("%v", cb.UnitExp))
	builder.WriteString(", ")
	builder.WriteString("env=")
	builder.WriteString(cb.Env)
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(cb.ChainID)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(cb.Nickname)
	builder.WriteString(", ")
	builder.WriteString("gas_type=")
	builder.WriteString(cb.GasType)
	builder.WriteByte(')')
	return builder.String()
}

// ChainBases is a parsable slice of ChainBase.
type ChainBases []*ChainBase
