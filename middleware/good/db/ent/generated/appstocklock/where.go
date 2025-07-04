// Code generated by ent, DO NOT EDIT.

package appstocklock

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldDeletedAt, v))
}

// AppStockID applies equality check predicate on the "app_stock_id" field. It's identical to AppStockIDEQ.
func AppStockID(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppStockID, v))
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppGoodID, v))
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldUnits, v))
}

// AppSpotUnits applies equality check predicate on the "app_spot_units" field. It's identical to AppSpotUnitsEQ.
func AppSpotUnits(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppSpotUnits, v))
}

// LockState applies equality check predicate on the "lock_state" field. It's identical to LockStateEQ.
func LockState(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldLockState, v))
}

// ChargeBackState applies equality check predicate on the "charge_back_state" field. It's identical to ChargeBackStateEQ.
func ChargeBackState(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldChargeBackState, v))
}

// ExLockID applies equality check predicate on the "ex_lock_id" field. It's identical to ExLockIDEQ.
func ExLockID(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldExLockID, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldDeletedAt, v))
}

// AppStockIDEQ applies the EQ predicate on the "app_stock_id" field.
func AppStockIDEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppStockID, v))
}

// AppStockIDNEQ applies the NEQ predicate on the "app_stock_id" field.
func AppStockIDNEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldAppStockID, v))
}

// AppStockIDIn applies the In predicate on the "app_stock_id" field.
func AppStockIDIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldAppStockID, vs...))
}

// AppStockIDNotIn applies the NotIn predicate on the "app_stock_id" field.
func AppStockIDNotIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldAppStockID, vs...))
}

// AppStockIDGT applies the GT predicate on the "app_stock_id" field.
func AppStockIDGT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldAppStockID, v))
}

// AppStockIDGTE applies the GTE predicate on the "app_stock_id" field.
func AppStockIDGTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldAppStockID, v))
}

// AppStockIDLT applies the LT predicate on the "app_stock_id" field.
func AppStockIDLT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldAppStockID, v))
}

// AppStockIDLTE applies the LTE predicate on the "app_stock_id" field.
func AppStockIDLTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldAppStockID, v))
}

// AppStockIDIsNil applies the IsNil predicate on the "app_stock_id" field.
func AppStockIDIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldAppStockID))
}

// AppStockIDNotNil applies the NotNil predicate on the "app_stock_id" field.
func AppStockIDNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldAppStockID))
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppGoodID, v))
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldAppGoodID, v))
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldAppGoodID, vs...))
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldAppGoodID, vs...))
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldAppGoodID, v))
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldAppGoodID, v))
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldAppGoodID, v))
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldAppGoodID, v))
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldAppGoodID))
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldAppGoodID))
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldUnits, v))
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldUnits, v))
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldUnits, vs...))
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldUnits, vs...))
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldUnits, v))
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldUnits, v))
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldUnits, v))
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldUnits, v))
}

// UnitsIsNil applies the IsNil predicate on the "units" field.
func UnitsIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldUnits))
}

// UnitsNotNil applies the NotNil predicate on the "units" field.
func UnitsNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldUnits))
}

// AppSpotUnitsEQ applies the EQ predicate on the "app_spot_units" field.
func AppSpotUnitsEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldAppSpotUnits, v))
}

// AppSpotUnitsNEQ applies the NEQ predicate on the "app_spot_units" field.
func AppSpotUnitsNEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldAppSpotUnits, v))
}

// AppSpotUnitsIn applies the In predicate on the "app_spot_units" field.
func AppSpotUnitsIn(vs ...decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldAppSpotUnits, vs...))
}

// AppSpotUnitsNotIn applies the NotIn predicate on the "app_spot_units" field.
func AppSpotUnitsNotIn(vs ...decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldAppSpotUnits, vs...))
}

// AppSpotUnitsGT applies the GT predicate on the "app_spot_units" field.
func AppSpotUnitsGT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldAppSpotUnits, v))
}

// AppSpotUnitsGTE applies the GTE predicate on the "app_spot_units" field.
func AppSpotUnitsGTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldAppSpotUnits, v))
}

// AppSpotUnitsLT applies the LT predicate on the "app_spot_units" field.
func AppSpotUnitsLT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldAppSpotUnits, v))
}

// AppSpotUnitsLTE applies the LTE predicate on the "app_spot_units" field.
func AppSpotUnitsLTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldAppSpotUnits, v))
}

// AppSpotUnitsIsNil applies the IsNil predicate on the "app_spot_units" field.
func AppSpotUnitsIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldAppSpotUnits))
}

// AppSpotUnitsNotNil applies the NotNil predicate on the "app_spot_units" field.
func AppSpotUnitsNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldAppSpotUnits))
}

// LockStateEQ applies the EQ predicate on the "lock_state" field.
func LockStateEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldLockState, v))
}

// LockStateNEQ applies the NEQ predicate on the "lock_state" field.
func LockStateNEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldLockState, v))
}

// LockStateIn applies the In predicate on the "lock_state" field.
func LockStateIn(vs ...string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldLockState, vs...))
}

// LockStateNotIn applies the NotIn predicate on the "lock_state" field.
func LockStateNotIn(vs ...string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldLockState, vs...))
}

// LockStateGT applies the GT predicate on the "lock_state" field.
func LockStateGT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldLockState, v))
}

// LockStateGTE applies the GTE predicate on the "lock_state" field.
func LockStateGTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldLockState, v))
}

// LockStateLT applies the LT predicate on the "lock_state" field.
func LockStateLT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldLockState, v))
}

// LockStateLTE applies the LTE predicate on the "lock_state" field.
func LockStateLTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldLockState, v))
}

// LockStateContains applies the Contains predicate on the "lock_state" field.
func LockStateContains(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldContains(FieldLockState, v))
}

// LockStateHasPrefix applies the HasPrefix predicate on the "lock_state" field.
func LockStateHasPrefix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldHasPrefix(FieldLockState, v))
}

// LockStateHasSuffix applies the HasSuffix predicate on the "lock_state" field.
func LockStateHasSuffix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldHasSuffix(FieldLockState, v))
}

// LockStateIsNil applies the IsNil predicate on the "lock_state" field.
func LockStateIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldLockState))
}

// LockStateNotNil applies the NotNil predicate on the "lock_state" field.
func LockStateNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldLockState))
}

// LockStateEqualFold applies the EqualFold predicate on the "lock_state" field.
func LockStateEqualFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEqualFold(FieldLockState, v))
}

// LockStateContainsFold applies the ContainsFold predicate on the "lock_state" field.
func LockStateContainsFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldContainsFold(FieldLockState, v))
}

// ChargeBackStateEQ applies the EQ predicate on the "charge_back_state" field.
func ChargeBackStateEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldChargeBackState, v))
}

// ChargeBackStateNEQ applies the NEQ predicate on the "charge_back_state" field.
func ChargeBackStateNEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldChargeBackState, v))
}

// ChargeBackStateIn applies the In predicate on the "charge_back_state" field.
func ChargeBackStateIn(vs ...string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldChargeBackState, vs...))
}

// ChargeBackStateNotIn applies the NotIn predicate on the "charge_back_state" field.
func ChargeBackStateNotIn(vs ...string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldChargeBackState, vs...))
}

// ChargeBackStateGT applies the GT predicate on the "charge_back_state" field.
func ChargeBackStateGT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldChargeBackState, v))
}

// ChargeBackStateGTE applies the GTE predicate on the "charge_back_state" field.
func ChargeBackStateGTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldChargeBackState, v))
}

// ChargeBackStateLT applies the LT predicate on the "charge_back_state" field.
func ChargeBackStateLT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldChargeBackState, v))
}

// ChargeBackStateLTE applies the LTE predicate on the "charge_back_state" field.
func ChargeBackStateLTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldChargeBackState, v))
}

// ChargeBackStateContains applies the Contains predicate on the "charge_back_state" field.
func ChargeBackStateContains(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldContains(FieldChargeBackState, v))
}

// ChargeBackStateHasPrefix applies the HasPrefix predicate on the "charge_back_state" field.
func ChargeBackStateHasPrefix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldHasPrefix(FieldChargeBackState, v))
}

// ChargeBackStateHasSuffix applies the HasSuffix predicate on the "charge_back_state" field.
func ChargeBackStateHasSuffix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldHasSuffix(FieldChargeBackState, v))
}

// ChargeBackStateIsNil applies the IsNil predicate on the "charge_back_state" field.
func ChargeBackStateIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldChargeBackState))
}

// ChargeBackStateNotNil applies the NotNil predicate on the "charge_back_state" field.
func ChargeBackStateNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldChargeBackState))
}

// ChargeBackStateEqualFold applies the EqualFold predicate on the "charge_back_state" field.
func ChargeBackStateEqualFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEqualFold(FieldChargeBackState, v))
}

// ChargeBackStateContainsFold applies the ContainsFold predicate on the "charge_back_state" field.
func ChargeBackStateContainsFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldContainsFold(FieldChargeBackState, v))
}

// ExLockIDEQ applies the EQ predicate on the "ex_lock_id" field.
func ExLockIDEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldEQ(FieldExLockID, v))
}

// ExLockIDNEQ applies the NEQ predicate on the "ex_lock_id" field.
func ExLockIDNEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNEQ(FieldExLockID, v))
}

// ExLockIDIn applies the In predicate on the "ex_lock_id" field.
func ExLockIDIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIn(FieldExLockID, vs...))
}

// ExLockIDNotIn applies the NotIn predicate on the "ex_lock_id" field.
func ExLockIDNotIn(vs ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotIn(FieldExLockID, vs...))
}

// ExLockIDGT applies the GT predicate on the "ex_lock_id" field.
func ExLockIDGT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGT(FieldExLockID, v))
}

// ExLockIDGTE applies the GTE predicate on the "ex_lock_id" field.
func ExLockIDGTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldGTE(FieldExLockID, v))
}

// ExLockIDLT applies the LT predicate on the "ex_lock_id" field.
func ExLockIDLT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLT(FieldExLockID, v))
}

// ExLockIDLTE applies the LTE predicate on the "ex_lock_id" field.
func ExLockIDLTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldLTE(FieldExLockID, v))
}

// ExLockIDIsNil applies the IsNil predicate on the "ex_lock_id" field.
func ExLockIDIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldIsNull(FieldExLockID))
}

// ExLockIDNotNil applies the NotNil predicate on the "ex_lock_id" field.
func ExLockIDNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(sql.FieldNotNull(FieldExLockID))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(sql.NotPredicates(p))
}
