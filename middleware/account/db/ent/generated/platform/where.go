// Code generated by ent, DO NOT EDIT.

package platform

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldEntID, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldAccountID, v))
}

// UsedFor applies equality check predicate on the "used_for" field. It's identical to UsedForEQ.
func UsedFor(v string) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldUsedFor, v))
}

// Backup applies equality check predicate on the "backup" field. It's identical to BackupEQ.
func Backup(v bool) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldBackup, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldEntID, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v uuid.UUID) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDIsNil applies the IsNil predicate on the "account_id" field.
func AccountIDIsNil() predicate.Platform {
	return predicate.Platform(sql.FieldIsNull(FieldAccountID))
}

// AccountIDNotNil applies the NotNil predicate on the "account_id" field.
func AccountIDNotNil() predicate.Platform {
	return predicate.Platform(sql.FieldNotNull(FieldAccountID))
}

// UsedForEQ applies the EQ predicate on the "used_for" field.
func UsedForEQ(v string) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldUsedFor, v))
}

// UsedForNEQ applies the NEQ predicate on the "used_for" field.
func UsedForNEQ(v string) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldUsedFor, v))
}

// UsedForIn applies the In predicate on the "used_for" field.
func UsedForIn(vs ...string) predicate.Platform {
	return predicate.Platform(sql.FieldIn(FieldUsedFor, vs...))
}

// UsedForNotIn applies the NotIn predicate on the "used_for" field.
func UsedForNotIn(vs ...string) predicate.Platform {
	return predicate.Platform(sql.FieldNotIn(FieldUsedFor, vs...))
}

// UsedForGT applies the GT predicate on the "used_for" field.
func UsedForGT(v string) predicate.Platform {
	return predicate.Platform(sql.FieldGT(FieldUsedFor, v))
}

// UsedForGTE applies the GTE predicate on the "used_for" field.
func UsedForGTE(v string) predicate.Platform {
	return predicate.Platform(sql.FieldGTE(FieldUsedFor, v))
}

// UsedForLT applies the LT predicate on the "used_for" field.
func UsedForLT(v string) predicate.Platform {
	return predicate.Platform(sql.FieldLT(FieldUsedFor, v))
}

// UsedForLTE applies the LTE predicate on the "used_for" field.
func UsedForLTE(v string) predicate.Platform {
	return predicate.Platform(sql.FieldLTE(FieldUsedFor, v))
}

// UsedForContains applies the Contains predicate on the "used_for" field.
func UsedForContains(v string) predicate.Platform {
	return predicate.Platform(sql.FieldContains(FieldUsedFor, v))
}

// UsedForHasPrefix applies the HasPrefix predicate on the "used_for" field.
func UsedForHasPrefix(v string) predicate.Platform {
	return predicate.Platform(sql.FieldHasPrefix(FieldUsedFor, v))
}

// UsedForHasSuffix applies the HasSuffix predicate on the "used_for" field.
func UsedForHasSuffix(v string) predicate.Platform {
	return predicate.Platform(sql.FieldHasSuffix(FieldUsedFor, v))
}

// UsedForIsNil applies the IsNil predicate on the "used_for" field.
func UsedForIsNil() predicate.Platform {
	return predicate.Platform(sql.FieldIsNull(FieldUsedFor))
}

// UsedForNotNil applies the NotNil predicate on the "used_for" field.
func UsedForNotNil() predicate.Platform {
	return predicate.Platform(sql.FieldNotNull(FieldUsedFor))
}

// UsedForEqualFold applies the EqualFold predicate on the "used_for" field.
func UsedForEqualFold(v string) predicate.Platform {
	return predicate.Platform(sql.FieldEqualFold(FieldUsedFor, v))
}

// UsedForContainsFold applies the ContainsFold predicate on the "used_for" field.
func UsedForContainsFold(v string) predicate.Platform {
	return predicate.Platform(sql.FieldContainsFold(FieldUsedFor, v))
}

// BackupEQ applies the EQ predicate on the "backup" field.
func BackupEQ(v bool) predicate.Platform {
	return predicate.Platform(sql.FieldEQ(FieldBackup, v))
}

// BackupNEQ applies the NEQ predicate on the "backup" field.
func BackupNEQ(v bool) predicate.Platform {
	return predicate.Platform(sql.FieldNEQ(FieldBackup, v))
}

// BackupIsNil applies the IsNil predicate on the "backup" field.
func BackupIsNil() predicate.Platform {
	return predicate.Platform(sql.FieldIsNull(FieldBackup))
}

// BackupNotNil applies the NotNil predicate on the "backup" field.
func BackupNotNil() predicate.Platform {
	return predicate.Platform(sql.FieldNotNull(FieldBackup))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Platform) predicate.Platform {
	return predicate.Platform(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Platform) predicate.Platform {
	return predicate.Platform(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Platform) predicate.Platform {
	return predicate.Platform(sql.NotPredicates(p))
}
