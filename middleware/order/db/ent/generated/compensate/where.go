// Code generated by ent, DO NOT EDIT.

package compensate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldDeletedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldOrderID, v))
}

// CompensateFromID applies equality check predicate on the "compensate_from_id" field. It's identical to CompensateFromIDEQ.
func CompensateFromID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateFromID, v))
}

// CompensateType applies equality check predicate on the "compensate_type" field. It's identical to CompensateTypeEQ.
func CompensateType(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateType, v))
}

// CompensateSeconds applies equality check predicate on the "compensate_seconds" field. It's identical to CompensateSecondsEQ.
func CompensateSeconds(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateSeconds, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldDeletedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldOrderID, v))
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldOrderID, v))
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldOrderID, v))
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldOrderID, v))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldNotNull(FieldOrderID))
}

// CompensateFromIDEQ applies the EQ predicate on the "compensate_from_id" field.
func CompensateFromIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateFromID, v))
}

// CompensateFromIDNEQ applies the NEQ predicate on the "compensate_from_id" field.
func CompensateFromIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldCompensateFromID, v))
}

// CompensateFromIDIn applies the In predicate on the "compensate_from_id" field.
func CompensateFromIDIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldCompensateFromID, vs...))
}

// CompensateFromIDNotIn applies the NotIn predicate on the "compensate_from_id" field.
func CompensateFromIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldCompensateFromID, vs...))
}

// CompensateFromIDGT applies the GT predicate on the "compensate_from_id" field.
func CompensateFromIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldCompensateFromID, v))
}

// CompensateFromIDGTE applies the GTE predicate on the "compensate_from_id" field.
func CompensateFromIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldCompensateFromID, v))
}

// CompensateFromIDLT applies the LT predicate on the "compensate_from_id" field.
func CompensateFromIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldCompensateFromID, v))
}

// CompensateFromIDLTE applies the LTE predicate on the "compensate_from_id" field.
func CompensateFromIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldCompensateFromID, v))
}

// CompensateFromIDIsNil applies the IsNil predicate on the "compensate_from_id" field.
func CompensateFromIDIsNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldIsNull(FieldCompensateFromID))
}

// CompensateFromIDNotNil applies the NotNil predicate on the "compensate_from_id" field.
func CompensateFromIDNotNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldNotNull(FieldCompensateFromID))
}

// CompensateTypeEQ applies the EQ predicate on the "compensate_type" field.
func CompensateTypeEQ(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateType, v))
}

// CompensateTypeNEQ applies the NEQ predicate on the "compensate_type" field.
func CompensateTypeNEQ(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldCompensateType, v))
}

// CompensateTypeIn applies the In predicate on the "compensate_type" field.
func CompensateTypeIn(vs ...string) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldCompensateType, vs...))
}

// CompensateTypeNotIn applies the NotIn predicate on the "compensate_type" field.
func CompensateTypeNotIn(vs ...string) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldCompensateType, vs...))
}

// CompensateTypeGT applies the GT predicate on the "compensate_type" field.
func CompensateTypeGT(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldCompensateType, v))
}

// CompensateTypeGTE applies the GTE predicate on the "compensate_type" field.
func CompensateTypeGTE(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldCompensateType, v))
}

// CompensateTypeLT applies the LT predicate on the "compensate_type" field.
func CompensateTypeLT(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldCompensateType, v))
}

// CompensateTypeLTE applies the LTE predicate on the "compensate_type" field.
func CompensateTypeLTE(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldCompensateType, v))
}

// CompensateTypeContains applies the Contains predicate on the "compensate_type" field.
func CompensateTypeContains(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldContains(FieldCompensateType, v))
}

// CompensateTypeHasPrefix applies the HasPrefix predicate on the "compensate_type" field.
func CompensateTypeHasPrefix(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldHasPrefix(FieldCompensateType, v))
}

// CompensateTypeHasSuffix applies the HasSuffix predicate on the "compensate_type" field.
func CompensateTypeHasSuffix(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldHasSuffix(FieldCompensateType, v))
}

// CompensateTypeIsNil applies the IsNil predicate on the "compensate_type" field.
func CompensateTypeIsNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldIsNull(FieldCompensateType))
}

// CompensateTypeNotNil applies the NotNil predicate on the "compensate_type" field.
func CompensateTypeNotNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldNotNull(FieldCompensateType))
}

// CompensateTypeEqualFold applies the EqualFold predicate on the "compensate_type" field.
func CompensateTypeEqualFold(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldEqualFold(FieldCompensateType, v))
}

// CompensateTypeContainsFold applies the ContainsFold predicate on the "compensate_type" field.
func CompensateTypeContainsFold(v string) predicate.Compensate {
	return predicate.Compensate(sql.FieldContainsFold(FieldCompensateType, v))
}

// CompensateSecondsEQ applies the EQ predicate on the "compensate_seconds" field.
func CompensateSecondsEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldEQ(FieldCompensateSeconds, v))
}

// CompensateSecondsNEQ applies the NEQ predicate on the "compensate_seconds" field.
func CompensateSecondsNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNEQ(FieldCompensateSeconds, v))
}

// CompensateSecondsIn applies the In predicate on the "compensate_seconds" field.
func CompensateSecondsIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldIn(FieldCompensateSeconds, vs...))
}

// CompensateSecondsNotIn applies the NotIn predicate on the "compensate_seconds" field.
func CompensateSecondsNotIn(vs ...uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldNotIn(FieldCompensateSeconds, vs...))
}

// CompensateSecondsGT applies the GT predicate on the "compensate_seconds" field.
func CompensateSecondsGT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGT(FieldCompensateSeconds, v))
}

// CompensateSecondsGTE applies the GTE predicate on the "compensate_seconds" field.
func CompensateSecondsGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldGTE(FieldCompensateSeconds, v))
}

// CompensateSecondsLT applies the LT predicate on the "compensate_seconds" field.
func CompensateSecondsLT(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLT(FieldCompensateSeconds, v))
}

// CompensateSecondsLTE applies the LTE predicate on the "compensate_seconds" field.
func CompensateSecondsLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(sql.FieldLTE(FieldCompensateSeconds, v))
}

// CompensateSecondsIsNil applies the IsNil predicate on the "compensate_seconds" field.
func CompensateSecondsIsNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldIsNull(FieldCompensateSeconds))
}

// CompensateSecondsNotNil applies the NotNil predicate on the "compensate_seconds" field.
func CompensateSecondsNotNil() predicate.Compensate {
	return predicate.Compensate(sql.FieldNotNull(FieldCompensateSeconds))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(sql.NotPredicates(p))
}
