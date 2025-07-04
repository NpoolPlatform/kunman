// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceinfo"
	"github.com/google/uuid"
)

// DeviceInfoCreate is the builder for creating a DeviceInfo entity.
type DeviceInfoCreate struct {
	config
	mutation *DeviceInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEntID sets the "ent_id" field.
func (dic *DeviceInfoCreate) SetEntID(u uuid.UUID) *DeviceInfoCreate {
	dic.mutation.SetEntID(u)
	return dic
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableEntID(u *uuid.UUID) *DeviceInfoCreate {
	if u != nil {
		dic.SetEntID(*u)
	}
	return dic
}

// SetCreatedAt sets the "created_at" field.
func (dic *DeviceInfoCreate) SetCreatedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetCreatedAt(u)
	return dic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableCreatedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetCreatedAt(*u)
	}
	return dic
}

// SetUpdatedAt sets the "updated_at" field.
func (dic *DeviceInfoCreate) SetUpdatedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetUpdatedAt(u)
	return dic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableUpdatedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetUpdatedAt(*u)
	}
	return dic
}

// SetDeletedAt sets the "deleted_at" field.
func (dic *DeviceInfoCreate) SetDeletedAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetDeletedAt(u)
	return dic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableDeletedAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetDeletedAt(*u)
	}
	return dic
}

// SetType sets the "type" field.
func (dic *DeviceInfoCreate) SetType(s string) *DeviceInfoCreate {
	dic.mutation.SetType(s)
	return dic
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableType(s *string) *DeviceInfoCreate {
	if s != nil {
		dic.SetType(*s)
	}
	return dic
}

// SetManufacturerID sets the "manufacturer_id" field.
func (dic *DeviceInfoCreate) SetManufacturerID(u uuid.UUID) *DeviceInfoCreate {
	dic.mutation.SetManufacturerID(u)
	return dic
}

// SetNillableManufacturerID sets the "manufacturer_id" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableManufacturerID(u *uuid.UUID) *DeviceInfoCreate {
	if u != nil {
		dic.SetManufacturerID(*u)
	}
	return dic
}

// SetPowerConsumption sets the "power_consumption" field.
func (dic *DeviceInfoCreate) SetPowerConsumption(u uint32) *DeviceInfoCreate {
	dic.mutation.SetPowerConsumption(u)
	return dic
}

// SetNillablePowerConsumption sets the "power_consumption" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillablePowerConsumption(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetPowerConsumption(*u)
	}
	return dic
}

// SetShipmentAt sets the "shipment_at" field.
func (dic *DeviceInfoCreate) SetShipmentAt(u uint32) *DeviceInfoCreate {
	dic.mutation.SetShipmentAt(u)
	return dic
}

// SetNillableShipmentAt sets the "shipment_at" field if the given value is not nil.
func (dic *DeviceInfoCreate) SetNillableShipmentAt(u *uint32) *DeviceInfoCreate {
	if u != nil {
		dic.SetShipmentAt(*u)
	}
	return dic
}

// SetID sets the "id" field.
func (dic *DeviceInfoCreate) SetID(u uint32) *DeviceInfoCreate {
	dic.mutation.SetID(u)
	return dic
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (dic *DeviceInfoCreate) Mutation() *DeviceInfoMutation {
	return dic.mutation
}

// Save creates the DeviceInfo in the database.
func (dic *DeviceInfoCreate) Save(ctx context.Context) (*DeviceInfo, error) {
	dic.defaults()
	return withHooks(ctx, dic.sqlSave, dic.mutation, dic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dic *DeviceInfoCreate) SaveX(ctx context.Context) *DeviceInfo {
	v, err := dic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dic *DeviceInfoCreate) Exec(ctx context.Context) error {
	_, err := dic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dic *DeviceInfoCreate) ExecX(ctx context.Context) {
	if err := dic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dic *DeviceInfoCreate) defaults() {
	if _, ok := dic.mutation.EntID(); !ok {
		v := deviceinfo.DefaultEntID()
		dic.mutation.SetEntID(v)
	}
	if _, ok := dic.mutation.CreatedAt(); !ok {
		v := deviceinfo.DefaultCreatedAt()
		dic.mutation.SetCreatedAt(v)
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		v := deviceinfo.DefaultUpdatedAt()
		dic.mutation.SetUpdatedAt(v)
	}
	if _, ok := dic.mutation.DeletedAt(); !ok {
		v := deviceinfo.DefaultDeletedAt()
		dic.mutation.SetDeletedAt(v)
	}
	if _, ok := dic.mutation.GetType(); !ok {
		v := deviceinfo.DefaultType
		dic.mutation.SetType(v)
	}
	if _, ok := dic.mutation.ManufacturerID(); !ok {
		v := deviceinfo.DefaultManufacturerID()
		dic.mutation.SetManufacturerID(v)
	}
	if _, ok := dic.mutation.PowerConsumption(); !ok {
		v := deviceinfo.DefaultPowerConsumption
		dic.mutation.SetPowerConsumption(v)
	}
	if _, ok := dic.mutation.ShipmentAt(); !ok {
		v := deviceinfo.DefaultShipmentAt
		dic.mutation.SetShipmentAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dic *DeviceInfoCreate) check() error {
	if _, ok := dic.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "DeviceInfo.ent_id"`)}
	}
	if _, ok := dic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "DeviceInfo.created_at"`)}
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "DeviceInfo.updated_at"`)}
	}
	if _, ok := dic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "DeviceInfo.deleted_at"`)}
	}
	if v, ok := dic.mutation.GetType(); ok {
		if err := deviceinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`generated: validator failed for field "DeviceInfo.type": %w`, err)}
		}
	}
	return nil
}

func (dic *DeviceInfoCreate) sqlSave(ctx context.Context) (*DeviceInfo, error) {
	if err := dic.check(); err != nil {
		return nil, err
	}
	_node, _spec := dic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	dic.mutation.id = &_node.ID
	dic.mutation.done = true
	return _node, nil
}

func (dic *DeviceInfoCreate) createSpec() (*DeviceInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceInfo{config: dic.config}
		_spec = sqlgraph.NewCreateSpec(deviceinfo.Table, sqlgraph.NewFieldSpec(deviceinfo.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = dic.conflict
	if id, ok := dic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dic.mutation.EntID(); ok {
		_spec.SetField(deviceinfo.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := dic.mutation.CreatedAt(); ok {
		_spec.SetField(deviceinfo.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := dic.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceinfo.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := dic.mutation.DeletedAt(); ok {
		_spec.SetField(deviceinfo.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := dic.mutation.GetType(); ok {
		_spec.SetField(deviceinfo.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := dic.mutation.ManufacturerID(); ok {
		_spec.SetField(deviceinfo.FieldManufacturerID, field.TypeUUID, value)
		_node.ManufacturerID = value
	}
	if value, ok := dic.mutation.PowerConsumption(); ok {
		_spec.SetField(deviceinfo.FieldPowerConsumption, field.TypeUint32, value)
		_node.PowerConsumption = value
	}
	if value, ok := dic.mutation.ShipmentAt(); ok {
		_spec.SetField(deviceinfo.FieldShipmentAt, field.TypeUint32, value)
		_node.ShipmentAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceInfo.Create().
//		SetEntID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceInfoUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (dic *DeviceInfoCreate) OnConflict(opts ...sql.ConflictOption) *DeviceInfoUpsertOne {
	dic.conflict = opts
	return &DeviceInfoUpsertOne{
		create: dic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dic *DeviceInfoCreate) OnConflictColumns(columns ...string) *DeviceInfoUpsertOne {
	dic.conflict = append(dic.conflict, sql.ConflictColumns(columns...))
	return &DeviceInfoUpsertOne{
		create: dic,
	}
}

type (
	// DeviceInfoUpsertOne is the builder for "upsert"-ing
	//  one DeviceInfo node.
	DeviceInfoUpsertOne struct {
		create *DeviceInfoCreate
	}

	// DeviceInfoUpsert is the "OnConflict" setter.
	DeviceInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetEntID sets the "ent_id" field.
func (u *DeviceInfoUpsert) SetEntID(v uuid.UUID) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateEntID() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldEntID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsert) SetCreatedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateCreatedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsert) AddCreatedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsert) SetUpdatedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateUpdatedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsert) AddUpdatedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsert) SetDeletedAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateDeletedAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsert) AddDeletedAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldDeletedAt, v)
	return u
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsert) SetType(v string) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateType() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsert) ClearType() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldType)
	return u
}

// SetManufacturerID sets the "manufacturer_id" field.
func (u *DeviceInfoUpsert) SetManufacturerID(v uuid.UUID) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldManufacturerID, v)
	return u
}

// UpdateManufacturerID sets the "manufacturer_id" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateManufacturerID() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldManufacturerID)
	return u
}

// ClearManufacturerID clears the value of the "manufacturer_id" field.
func (u *DeviceInfoUpsert) ClearManufacturerID() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldManufacturerID)
	return u
}

// SetPowerConsumption sets the "power_consumption" field.
func (u *DeviceInfoUpsert) SetPowerConsumption(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldPowerConsumption, v)
	return u
}

// UpdatePowerConsumption sets the "power_consumption" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdatePowerConsumption() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldPowerConsumption)
	return u
}

// AddPowerConsumption adds v to the "power_consumption" field.
func (u *DeviceInfoUpsert) AddPowerConsumption(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldPowerConsumption, v)
	return u
}

// ClearPowerConsumption clears the value of the "power_consumption" field.
func (u *DeviceInfoUpsert) ClearPowerConsumption() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldPowerConsumption)
	return u
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsert) SetShipmentAt(v uint32) *DeviceInfoUpsert {
	u.Set(deviceinfo.FieldShipmentAt, v)
	return u
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsert) UpdateShipmentAt() *DeviceInfoUpsert {
	u.SetExcluded(deviceinfo.FieldShipmentAt)
	return u
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsert) AddShipmentAt(v uint32) *DeviceInfoUpsert {
	u.Add(deviceinfo.FieldShipmentAt, v)
	return u
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsert) ClearShipmentAt() *DeviceInfoUpsert {
	u.SetNull(deviceinfo.FieldShipmentAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceInfoUpsertOne) UpdateNewValues() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deviceinfo.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceInfoUpsertOne) Ignore() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceInfoUpsertOne) DoNothing() *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceInfoCreate.OnConflict
// documentation for more info.
func (u *DeviceInfoUpsertOne) Update(set func(*DeviceInfoUpsert)) *DeviceInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *DeviceInfoUpsertOne) SetEntID(v uuid.UUID) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateEntID() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsertOne) SetCreatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsertOne) AddCreatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateCreatedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsertOne) SetUpdatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsertOne) AddUpdatedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateUpdatedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsertOne) SetDeletedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsertOne) AddDeletedAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateDeletedAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsertOne) SetType(v string) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateType() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsertOne) ClearType() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearType()
	})
}

// SetManufacturerID sets the "manufacturer_id" field.
func (u *DeviceInfoUpsertOne) SetManufacturerID(v uuid.UUID) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetManufacturerID(v)
	})
}

// UpdateManufacturerID sets the "manufacturer_id" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateManufacturerID() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateManufacturerID()
	})
}

// ClearManufacturerID clears the value of the "manufacturer_id" field.
func (u *DeviceInfoUpsertOne) ClearManufacturerID() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearManufacturerID()
	})
}

// SetPowerConsumption sets the "power_consumption" field.
func (u *DeviceInfoUpsertOne) SetPowerConsumption(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPowerConsumption(v)
	})
}

// AddPowerConsumption adds v to the "power_consumption" field.
func (u *DeviceInfoUpsertOne) AddPowerConsumption(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddPowerConsumption(v)
	})
}

// UpdatePowerConsumption sets the "power_consumption" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdatePowerConsumption() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePowerConsumption()
	})
}

// ClearPowerConsumption clears the value of the "power_consumption" field.
func (u *DeviceInfoUpsertOne) ClearPowerConsumption() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPowerConsumption()
	})
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsertOne) SetShipmentAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetShipmentAt(v)
	})
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsertOne) AddShipmentAt(v uint32) *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddShipmentAt(v)
	})
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertOne) UpdateShipmentAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateShipmentAt()
	})
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsertOne) ClearShipmentAt() *DeviceInfoUpsertOne {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearShipmentAt()
	})
}

// Exec executes the query.
func (u *DeviceInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for DeviceInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceInfoUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceInfoUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceInfoCreateBulk is the builder for creating many DeviceInfo entities in bulk.
type DeviceInfoCreateBulk struct {
	config
	err      error
	builders []*DeviceInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceInfo entities in the database.
func (dicb *DeviceInfoCreateBulk) Save(ctx context.Context) ([]*DeviceInfo, error) {
	if dicb.err != nil {
		return nil, dicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dicb.builders))
	nodes := make([]*DeviceInfo, len(dicb.builders))
	mutators := make([]Mutator, len(dicb.builders))
	for i := range dicb.builders {
		func(i int, root context.Context) {
			builder := dicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dicb *DeviceInfoCreateBulk) SaveX(ctx context.Context) []*DeviceInfo {
	v, err := dicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dicb *DeviceInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := dicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dicb *DeviceInfoCreateBulk) ExecX(ctx context.Context) {
	if err := dicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceInfoUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (dicb *DeviceInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceInfoUpsertBulk {
	dicb.conflict = opts
	return &DeviceInfoUpsertBulk{
		create: dicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dicb *DeviceInfoCreateBulk) OnConflictColumns(columns ...string) *DeviceInfoUpsertBulk {
	dicb.conflict = append(dicb.conflict, sql.ConflictColumns(columns...))
	return &DeviceInfoUpsertBulk{
		create: dicb,
	}
}

// DeviceInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceInfo nodes.
type DeviceInfoUpsertBulk struct {
	create *DeviceInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceInfoUpsertBulk) UpdateNewValues() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deviceinfo.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceInfoUpsertBulk) Ignore() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceInfoUpsertBulk) DoNothing() *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceInfoCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceInfoUpsertBulk) Update(set func(*DeviceInfoUpsert)) *DeviceInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *DeviceInfoUpsertBulk) SetEntID(v uuid.UUID) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateEntID() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceInfoUpsertBulk) SetCreatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceInfoUpsertBulk) AddCreatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateCreatedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceInfoUpsertBulk) SetUpdatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceInfoUpsertBulk) AddUpdatedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateUpdatedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceInfoUpsertBulk) SetDeletedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceInfoUpsertBulk) AddDeletedAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateDeletedAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetType sets the "type" field.
func (u *DeviceInfoUpsertBulk) SetType(v string) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateType() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DeviceInfoUpsertBulk) ClearType() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearType()
	})
}

// SetManufacturerID sets the "manufacturer_id" field.
func (u *DeviceInfoUpsertBulk) SetManufacturerID(v uuid.UUID) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetManufacturerID(v)
	})
}

// UpdateManufacturerID sets the "manufacturer_id" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateManufacturerID() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateManufacturerID()
	})
}

// ClearManufacturerID clears the value of the "manufacturer_id" field.
func (u *DeviceInfoUpsertBulk) ClearManufacturerID() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearManufacturerID()
	})
}

// SetPowerConsumption sets the "power_consumption" field.
func (u *DeviceInfoUpsertBulk) SetPowerConsumption(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetPowerConsumption(v)
	})
}

// AddPowerConsumption adds v to the "power_consumption" field.
func (u *DeviceInfoUpsertBulk) AddPowerConsumption(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddPowerConsumption(v)
	})
}

// UpdatePowerConsumption sets the "power_consumption" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdatePowerConsumption() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdatePowerConsumption()
	})
}

// ClearPowerConsumption clears the value of the "power_consumption" field.
func (u *DeviceInfoUpsertBulk) ClearPowerConsumption() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearPowerConsumption()
	})
}

// SetShipmentAt sets the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) SetShipmentAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.SetShipmentAt(v)
	})
}

// AddShipmentAt adds v to the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) AddShipmentAt(v uint32) *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.AddShipmentAt(v)
	})
}

// UpdateShipmentAt sets the "shipment_at" field to the value that was provided on create.
func (u *DeviceInfoUpsertBulk) UpdateShipmentAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.UpdateShipmentAt()
	})
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (u *DeviceInfoUpsertBulk) ClearShipmentAt() *DeviceInfoUpsertBulk {
	return u.Update(func(s *DeviceInfoUpsert) {
		s.ClearShipmentAt()
	})
}

// Exec executes the query.
func (u *DeviceInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the DeviceInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for DeviceInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
