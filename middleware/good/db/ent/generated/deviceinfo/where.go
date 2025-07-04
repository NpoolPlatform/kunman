// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldType, v))
}

// ManufacturerID applies equality check predicate on the "manufacturer_id" field. It's identical to ManufacturerIDEQ.
func ManufacturerID(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldManufacturerID, v))
}

// PowerConsumption applies equality check predicate on the "power_consumption" field. It's identical to PowerConsumptionEQ.
func PowerConsumption(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldPowerConsumption, v))
}

// ShipmentAt applies equality check predicate on the "shipment_at" field. It's identical to ShipmentAtEQ.
func ShipmentAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldShipmentAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldDeletedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldType, v))
}

// ManufacturerIDEQ applies the EQ predicate on the "manufacturer_id" field.
func ManufacturerIDEQ(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldManufacturerID, v))
}

// ManufacturerIDNEQ applies the NEQ predicate on the "manufacturer_id" field.
func ManufacturerIDNEQ(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldManufacturerID, v))
}

// ManufacturerIDIn applies the In predicate on the "manufacturer_id" field.
func ManufacturerIDIn(vs ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldManufacturerID, vs...))
}

// ManufacturerIDNotIn applies the NotIn predicate on the "manufacturer_id" field.
func ManufacturerIDNotIn(vs ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldManufacturerID, vs...))
}

// ManufacturerIDGT applies the GT predicate on the "manufacturer_id" field.
func ManufacturerIDGT(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldManufacturerID, v))
}

// ManufacturerIDGTE applies the GTE predicate on the "manufacturer_id" field.
func ManufacturerIDGTE(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldManufacturerID, v))
}

// ManufacturerIDLT applies the LT predicate on the "manufacturer_id" field.
func ManufacturerIDLT(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldManufacturerID, v))
}

// ManufacturerIDLTE applies the LTE predicate on the "manufacturer_id" field.
func ManufacturerIDLTE(v uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldManufacturerID, v))
}

// ManufacturerIDIsNil applies the IsNil predicate on the "manufacturer_id" field.
func ManufacturerIDIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIsNull(FieldManufacturerID))
}

// ManufacturerIDNotNil applies the NotNil predicate on the "manufacturer_id" field.
func ManufacturerIDNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotNull(FieldManufacturerID))
}

// PowerConsumptionEQ applies the EQ predicate on the "power_consumption" field.
func PowerConsumptionEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldPowerConsumption, v))
}

// PowerConsumptionNEQ applies the NEQ predicate on the "power_consumption" field.
func PowerConsumptionNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldPowerConsumption, v))
}

// PowerConsumptionIn applies the In predicate on the "power_consumption" field.
func PowerConsumptionIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldPowerConsumption, vs...))
}

// PowerConsumptionNotIn applies the NotIn predicate on the "power_consumption" field.
func PowerConsumptionNotIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldPowerConsumption, vs...))
}

// PowerConsumptionGT applies the GT predicate on the "power_consumption" field.
func PowerConsumptionGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldPowerConsumption, v))
}

// PowerConsumptionGTE applies the GTE predicate on the "power_consumption" field.
func PowerConsumptionGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldPowerConsumption, v))
}

// PowerConsumptionLT applies the LT predicate on the "power_consumption" field.
func PowerConsumptionLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldPowerConsumption, v))
}

// PowerConsumptionLTE applies the LTE predicate on the "power_consumption" field.
func PowerConsumptionLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldPowerConsumption, v))
}

// PowerConsumptionIsNil applies the IsNil predicate on the "power_consumption" field.
func PowerConsumptionIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIsNull(FieldPowerConsumption))
}

// PowerConsumptionNotNil applies the NotNil predicate on the "power_consumption" field.
func PowerConsumptionNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotNull(FieldPowerConsumption))
}

// ShipmentAtEQ applies the EQ predicate on the "shipment_at" field.
func ShipmentAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldShipmentAt, v))
}

// ShipmentAtNEQ applies the NEQ predicate on the "shipment_at" field.
func ShipmentAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldShipmentAt, v))
}

// ShipmentAtIn applies the In predicate on the "shipment_at" field.
func ShipmentAtIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldShipmentAt, vs...))
}

// ShipmentAtNotIn applies the NotIn predicate on the "shipment_at" field.
func ShipmentAtNotIn(vs ...uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldShipmentAt, vs...))
}

// ShipmentAtGT applies the GT predicate on the "shipment_at" field.
func ShipmentAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldShipmentAt, v))
}

// ShipmentAtGTE applies the GTE predicate on the "shipment_at" field.
func ShipmentAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldShipmentAt, v))
}

// ShipmentAtLT applies the LT predicate on the "shipment_at" field.
func ShipmentAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldShipmentAt, v))
}

// ShipmentAtLTE applies the LTE predicate on the "shipment_at" field.
func ShipmentAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldShipmentAt, v))
}

// ShipmentAtIsNil applies the IsNil predicate on the "shipment_at" field.
func ShipmentAtIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIsNull(FieldShipmentAt))
}

// ShipmentAtNotNil applies the NotNil predicate on the "shipment_at" field.
func ShipmentAtNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotNull(FieldShipmentAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.NotPredicates(p))
}
