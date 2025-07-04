// Code generated by ent, DO NOT EDIT.

package deviceposter

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldDeletedAt, v))
}

// DeviceTypeID applies equality check predicate on the "device_type_id" field. It's identical to DeviceTypeIDEQ.
func DeviceTypeID(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldDeviceTypeID, v))
}

// Poster applies equality check predicate on the "poster" field. It's identical to PosterEQ.
func Poster(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldPoster, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldIndex, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldDeletedAt, v))
}

// DeviceTypeIDEQ applies the EQ predicate on the "device_type_id" field.
func DeviceTypeIDEQ(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldDeviceTypeID, v))
}

// DeviceTypeIDNEQ applies the NEQ predicate on the "device_type_id" field.
func DeviceTypeIDNEQ(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldDeviceTypeID, v))
}

// DeviceTypeIDIn applies the In predicate on the "device_type_id" field.
func DeviceTypeIDIn(vs ...uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldDeviceTypeID, vs...))
}

// DeviceTypeIDNotIn applies the NotIn predicate on the "device_type_id" field.
func DeviceTypeIDNotIn(vs ...uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldDeviceTypeID, vs...))
}

// DeviceTypeIDGT applies the GT predicate on the "device_type_id" field.
func DeviceTypeIDGT(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldDeviceTypeID, v))
}

// DeviceTypeIDGTE applies the GTE predicate on the "device_type_id" field.
func DeviceTypeIDGTE(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldDeviceTypeID, v))
}

// DeviceTypeIDLT applies the LT predicate on the "device_type_id" field.
func DeviceTypeIDLT(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldDeviceTypeID, v))
}

// DeviceTypeIDLTE applies the LTE predicate on the "device_type_id" field.
func DeviceTypeIDLTE(v uuid.UUID) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldDeviceTypeID, v))
}

// DeviceTypeIDIsNil applies the IsNil predicate on the "device_type_id" field.
func DeviceTypeIDIsNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIsNull(FieldDeviceTypeID))
}

// DeviceTypeIDNotNil applies the NotNil predicate on the "device_type_id" field.
func DeviceTypeIDNotNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotNull(FieldDeviceTypeID))
}

// PosterEQ applies the EQ predicate on the "poster" field.
func PosterEQ(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldPoster, v))
}

// PosterNEQ applies the NEQ predicate on the "poster" field.
func PosterNEQ(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldPoster, v))
}

// PosterIn applies the In predicate on the "poster" field.
func PosterIn(vs ...string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldPoster, vs...))
}

// PosterNotIn applies the NotIn predicate on the "poster" field.
func PosterNotIn(vs ...string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldPoster, vs...))
}

// PosterGT applies the GT predicate on the "poster" field.
func PosterGT(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldPoster, v))
}

// PosterGTE applies the GTE predicate on the "poster" field.
func PosterGTE(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldPoster, v))
}

// PosterLT applies the LT predicate on the "poster" field.
func PosterLT(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldPoster, v))
}

// PosterLTE applies the LTE predicate on the "poster" field.
func PosterLTE(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldPoster, v))
}

// PosterContains applies the Contains predicate on the "poster" field.
func PosterContains(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldContains(FieldPoster, v))
}

// PosterHasPrefix applies the HasPrefix predicate on the "poster" field.
func PosterHasPrefix(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldHasPrefix(FieldPoster, v))
}

// PosterHasSuffix applies the HasSuffix predicate on the "poster" field.
func PosterHasSuffix(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldHasSuffix(FieldPoster, v))
}

// PosterIsNil applies the IsNil predicate on the "poster" field.
func PosterIsNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIsNull(FieldPoster))
}

// PosterNotNil applies the NotNil predicate on the "poster" field.
func PosterNotNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotNull(FieldPoster))
}

// PosterEqualFold applies the EqualFold predicate on the "poster" field.
func PosterEqualFold(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEqualFold(FieldPoster, v))
}

// PosterContainsFold applies the ContainsFold predicate on the "poster" field.
func PosterContainsFold(v string) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldContainsFold(FieldPoster, v))
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldEQ(FieldIndex, v))
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNEQ(FieldIndex, v))
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIn(FieldIndex, vs...))
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotIn(FieldIndex, vs...))
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGT(FieldIndex, v))
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldGTE(FieldIndex, v))
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLT(FieldIndex, v))
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v uint8) predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldLTE(FieldIndex, v))
}

// IndexIsNil applies the IsNil predicate on the "index" field.
func IndexIsNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldIsNull(FieldIndex))
}

// IndexNotNil applies the NotNil predicate on the "index" field.
func IndexNotNil() predicate.DevicePoster {
	return predicate.DevicePoster(sql.FieldNotNull(FieldIndex))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DevicePoster) predicate.DevicePoster {
	return predicate.DevicePoster(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DevicePoster) predicate.DevicePoster {
	return predicate.DevicePoster(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DevicePoster) predicate.DevicePoster {
	return predicate.DevicePoster(sql.NotPredicates(p))
}
