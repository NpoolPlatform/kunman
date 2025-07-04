// Code generated by ent, DO NOT EDIT.

package appstock

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldDeletedAt, v))
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldAppGoodID, v))
}

// Reserved applies equality check predicate on the "reserved" field. It's identical to ReservedEQ.
func Reserved(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldReserved, v))
}

// SpotQuantity applies equality check predicate on the "spot_quantity" field. It's identical to SpotQuantityEQ.
func SpotQuantity(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldSpotQuantity, v))
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldLocked, v))
}

// InService applies equality check predicate on the "in_service" field. It's identical to InServiceEQ.
func InService(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldInService, v))
}

// WaitStart applies equality check predicate on the "wait_start" field. It's identical to WaitStartEQ.
func WaitStart(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldWaitStart, v))
}

// Sold applies equality check predicate on the "sold" field. It's identical to SoldEQ.
func Sold(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldSold, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldDeletedAt, v))
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldAppGoodID, v))
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldAppGoodID, v))
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldAppGoodID, vs...))
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldAppGoodID, vs...))
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldAppGoodID, v))
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldAppGoodID, v))
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldAppGoodID, v))
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldAppGoodID, v))
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldAppGoodID))
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldAppGoodID))
}

// ReservedEQ applies the EQ predicate on the "reserved" field.
func ReservedEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldReserved, v))
}

// ReservedNEQ applies the NEQ predicate on the "reserved" field.
func ReservedNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldReserved, v))
}

// ReservedIn applies the In predicate on the "reserved" field.
func ReservedIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldReserved, vs...))
}

// ReservedNotIn applies the NotIn predicate on the "reserved" field.
func ReservedNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldReserved, vs...))
}

// ReservedGT applies the GT predicate on the "reserved" field.
func ReservedGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldReserved, v))
}

// ReservedGTE applies the GTE predicate on the "reserved" field.
func ReservedGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldReserved, v))
}

// ReservedLT applies the LT predicate on the "reserved" field.
func ReservedLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldReserved, v))
}

// ReservedLTE applies the LTE predicate on the "reserved" field.
func ReservedLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldReserved, v))
}

// ReservedIsNil applies the IsNil predicate on the "reserved" field.
func ReservedIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldReserved))
}

// ReservedNotNil applies the NotNil predicate on the "reserved" field.
func ReservedNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldReserved))
}

// SpotQuantityEQ applies the EQ predicate on the "spot_quantity" field.
func SpotQuantityEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldSpotQuantity, v))
}

// SpotQuantityNEQ applies the NEQ predicate on the "spot_quantity" field.
func SpotQuantityNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldSpotQuantity, v))
}

// SpotQuantityIn applies the In predicate on the "spot_quantity" field.
func SpotQuantityIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldSpotQuantity, vs...))
}

// SpotQuantityNotIn applies the NotIn predicate on the "spot_quantity" field.
func SpotQuantityNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldSpotQuantity, vs...))
}

// SpotQuantityGT applies the GT predicate on the "spot_quantity" field.
func SpotQuantityGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldSpotQuantity, v))
}

// SpotQuantityGTE applies the GTE predicate on the "spot_quantity" field.
func SpotQuantityGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldSpotQuantity, v))
}

// SpotQuantityLT applies the LT predicate on the "spot_quantity" field.
func SpotQuantityLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldSpotQuantity, v))
}

// SpotQuantityLTE applies the LTE predicate on the "spot_quantity" field.
func SpotQuantityLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldSpotQuantity, v))
}

// SpotQuantityIsNil applies the IsNil predicate on the "spot_quantity" field.
func SpotQuantityIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldSpotQuantity))
}

// SpotQuantityNotNil applies the NotNil predicate on the "spot_quantity" field.
func SpotQuantityNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldSpotQuantity))
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldLocked, v))
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldLocked, v))
}

// LockedIn applies the In predicate on the "locked" field.
func LockedIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldLocked, vs...))
}

// LockedNotIn applies the NotIn predicate on the "locked" field.
func LockedNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldLocked, vs...))
}

// LockedGT applies the GT predicate on the "locked" field.
func LockedGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldLocked, v))
}

// LockedGTE applies the GTE predicate on the "locked" field.
func LockedGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldLocked, v))
}

// LockedLT applies the LT predicate on the "locked" field.
func LockedLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldLocked, v))
}

// LockedLTE applies the LTE predicate on the "locked" field.
func LockedLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldLocked, v))
}

// LockedIsNil applies the IsNil predicate on the "locked" field.
func LockedIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldLocked))
}

// LockedNotNil applies the NotNil predicate on the "locked" field.
func LockedNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldLocked))
}

// InServiceEQ applies the EQ predicate on the "in_service" field.
func InServiceEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldInService, v))
}

// InServiceNEQ applies the NEQ predicate on the "in_service" field.
func InServiceNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldInService, v))
}

// InServiceIn applies the In predicate on the "in_service" field.
func InServiceIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldInService, vs...))
}

// InServiceNotIn applies the NotIn predicate on the "in_service" field.
func InServiceNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldInService, vs...))
}

// InServiceGT applies the GT predicate on the "in_service" field.
func InServiceGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldInService, v))
}

// InServiceGTE applies the GTE predicate on the "in_service" field.
func InServiceGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldInService, v))
}

// InServiceLT applies the LT predicate on the "in_service" field.
func InServiceLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldInService, v))
}

// InServiceLTE applies the LTE predicate on the "in_service" field.
func InServiceLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldInService, v))
}

// InServiceIsNil applies the IsNil predicate on the "in_service" field.
func InServiceIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldInService))
}

// InServiceNotNil applies the NotNil predicate on the "in_service" field.
func InServiceNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldInService))
}

// WaitStartEQ applies the EQ predicate on the "wait_start" field.
func WaitStartEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldWaitStart, v))
}

// WaitStartNEQ applies the NEQ predicate on the "wait_start" field.
func WaitStartNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldWaitStart, v))
}

// WaitStartIn applies the In predicate on the "wait_start" field.
func WaitStartIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldWaitStart, vs...))
}

// WaitStartNotIn applies the NotIn predicate on the "wait_start" field.
func WaitStartNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldWaitStart, vs...))
}

// WaitStartGT applies the GT predicate on the "wait_start" field.
func WaitStartGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldWaitStart, v))
}

// WaitStartGTE applies the GTE predicate on the "wait_start" field.
func WaitStartGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldWaitStart, v))
}

// WaitStartLT applies the LT predicate on the "wait_start" field.
func WaitStartLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldWaitStart, v))
}

// WaitStartLTE applies the LTE predicate on the "wait_start" field.
func WaitStartLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldWaitStart, v))
}

// WaitStartIsNil applies the IsNil predicate on the "wait_start" field.
func WaitStartIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldWaitStart))
}

// WaitStartNotNil applies the NotNil predicate on the "wait_start" field.
func WaitStartNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldWaitStart))
}

// SoldEQ applies the EQ predicate on the "sold" field.
func SoldEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldEQ(FieldSold, v))
}

// SoldNEQ applies the NEQ predicate on the "sold" field.
func SoldNEQ(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNEQ(FieldSold, v))
}

// SoldIn applies the In predicate on the "sold" field.
func SoldIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldIn(FieldSold, vs...))
}

// SoldNotIn applies the NotIn predicate on the "sold" field.
func SoldNotIn(vs ...decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldNotIn(FieldSold, vs...))
}

// SoldGT applies the GT predicate on the "sold" field.
func SoldGT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGT(FieldSold, v))
}

// SoldGTE applies the GTE predicate on the "sold" field.
func SoldGTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldGTE(FieldSold, v))
}

// SoldLT applies the LT predicate on the "sold" field.
func SoldLT(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLT(FieldSold, v))
}

// SoldLTE applies the LTE predicate on the "sold" field.
func SoldLTE(v decimal.Decimal) predicate.AppStock {
	return predicate.AppStock(sql.FieldLTE(FieldSold, v))
}

// SoldIsNil applies the IsNil predicate on the "sold" field.
func SoldIsNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldIsNull(FieldSold))
}

// SoldNotNil applies the NotNil predicate on the "sold" field.
func SoldNotNil() predicate.AppStock {
	return predicate.AppStock(sql.FieldNotNull(FieldSold))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppStock) predicate.AppStock {
	return predicate.AppStock(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppStock) predicate.AppStock {
	return predicate.AppStock(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppStock) predicate.AppStock {
	return predicate.AppStock(sql.NotPredicates(p))
}
