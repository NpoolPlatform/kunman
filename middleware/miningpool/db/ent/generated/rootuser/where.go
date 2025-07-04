// Code generated by ent, DO NOT EDIT.

package rootuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldEntID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldName, v))
}

// PoolID applies equality check predicate on the "pool_id" field. It's identical to PoolIDEQ.
func PoolID(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldPoolID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldEmail, v))
}

// AuthToken applies equality check predicate on the "auth_token" field. It's identical to AuthTokenEQ.
func AuthToken(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthToken, v))
}

// AuthTokenSalt applies equality check predicate on the "auth_token_salt" field. It's identical to AuthTokenSaltEQ.
func AuthTokenSalt(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthTokenSalt, v))
}

// Authed applies equality check predicate on the "authed" field. It's identical to AuthedEQ.
func Authed(v bool) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthed, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldEntID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContainsFold(FieldName, v))
}

// PoolIDEQ applies the EQ predicate on the "pool_id" field.
func PoolIDEQ(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldPoolID, v))
}

// PoolIDNEQ applies the NEQ predicate on the "pool_id" field.
func PoolIDNEQ(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldPoolID, v))
}

// PoolIDIn applies the In predicate on the "pool_id" field.
func PoolIDIn(vs ...uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldPoolID, vs...))
}

// PoolIDNotIn applies the NotIn predicate on the "pool_id" field.
func PoolIDNotIn(vs ...uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldPoolID, vs...))
}

// PoolIDGT applies the GT predicate on the "pool_id" field.
func PoolIDGT(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldPoolID, v))
}

// PoolIDGTE applies the GTE predicate on the "pool_id" field.
func PoolIDGTE(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldPoolID, v))
}

// PoolIDLT applies the LT predicate on the "pool_id" field.
func PoolIDLT(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldPoolID, v))
}

// PoolIDLTE applies the LTE predicate on the "pool_id" field.
func PoolIDLTE(v uuid.UUID) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldPoolID, v))
}

// PoolIDIsNil applies the IsNil predicate on the "pool_id" field.
func PoolIDIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldPoolID))
}

// PoolIDNotNil applies the NotNil predicate on the "pool_id" field.
func PoolIDNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldPoolID))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContainsFold(FieldEmail, v))
}

// AuthTokenEQ applies the EQ predicate on the "auth_token" field.
func AuthTokenEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthToken, v))
}

// AuthTokenNEQ applies the NEQ predicate on the "auth_token" field.
func AuthTokenNEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldAuthToken, v))
}

// AuthTokenIn applies the In predicate on the "auth_token" field.
func AuthTokenIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldAuthToken, vs...))
}

// AuthTokenNotIn applies the NotIn predicate on the "auth_token" field.
func AuthTokenNotIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldAuthToken, vs...))
}

// AuthTokenGT applies the GT predicate on the "auth_token" field.
func AuthTokenGT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldAuthToken, v))
}

// AuthTokenGTE applies the GTE predicate on the "auth_token" field.
func AuthTokenGTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldAuthToken, v))
}

// AuthTokenLT applies the LT predicate on the "auth_token" field.
func AuthTokenLT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldAuthToken, v))
}

// AuthTokenLTE applies the LTE predicate on the "auth_token" field.
func AuthTokenLTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldAuthToken, v))
}

// AuthTokenContains applies the Contains predicate on the "auth_token" field.
func AuthTokenContains(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContains(FieldAuthToken, v))
}

// AuthTokenHasPrefix applies the HasPrefix predicate on the "auth_token" field.
func AuthTokenHasPrefix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasPrefix(FieldAuthToken, v))
}

// AuthTokenHasSuffix applies the HasSuffix predicate on the "auth_token" field.
func AuthTokenHasSuffix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasSuffix(FieldAuthToken, v))
}

// AuthTokenIsNil applies the IsNil predicate on the "auth_token" field.
func AuthTokenIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldAuthToken))
}

// AuthTokenNotNil applies the NotNil predicate on the "auth_token" field.
func AuthTokenNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldAuthToken))
}

// AuthTokenEqualFold applies the EqualFold predicate on the "auth_token" field.
func AuthTokenEqualFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEqualFold(FieldAuthToken, v))
}

// AuthTokenContainsFold applies the ContainsFold predicate on the "auth_token" field.
func AuthTokenContainsFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContainsFold(FieldAuthToken, v))
}

// AuthTokenSaltEQ applies the EQ predicate on the "auth_token_salt" field.
func AuthTokenSaltEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthTokenSalt, v))
}

// AuthTokenSaltNEQ applies the NEQ predicate on the "auth_token_salt" field.
func AuthTokenSaltNEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldAuthTokenSalt, v))
}

// AuthTokenSaltIn applies the In predicate on the "auth_token_salt" field.
func AuthTokenSaltIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldAuthTokenSalt, vs...))
}

// AuthTokenSaltNotIn applies the NotIn predicate on the "auth_token_salt" field.
func AuthTokenSaltNotIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldAuthTokenSalt, vs...))
}

// AuthTokenSaltGT applies the GT predicate on the "auth_token_salt" field.
func AuthTokenSaltGT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldAuthTokenSalt, v))
}

// AuthTokenSaltGTE applies the GTE predicate on the "auth_token_salt" field.
func AuthTokenSaltGTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldAuthTokenSalt, v))
}

// AuthTokenSaltLT applies the LT predicate on the "auth_token_salt" field.
func AuthTokenSaltLT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldAuthTokenSalt, v))
}

// AuthTokenSaltLTE applies the LTE predicate on the "auth_token_salt" field.
func AuthTokenSaltLTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldAuthTokenSalt, v))
}

// AuthTokenSaltContains applies the Contains predicate on the "auth_token_salt" field.
func AuthTokenSaltContains(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContains(FieldAuthTokenSalt, v))
}

// AuthTokenSaltHasPrefix applies the HasPrefix predicate on the "auth_token_salt" field.
func AuthTokenSaltHasPrefix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasPrefix(FieldAuthTokenSalt, v))
}

// AuthTokenSaltHasSuffix applies the HasSuffix predicate on the "auth_token_salt" field.
func AuthTokenSaltHasSuffix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasSuffix(FieldAuthTokenSalt, v))
}

// AuthTokenSaltIsNil applies the IsNil predicate on the "auth_token_salt" field.
func AuthTokenSaltIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldAuthTokenSalt))
}

// AuthTokenSaltNotNil applies the NotNil predicate on the "auth_token_salt" field.
func AuthTokenSaltNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldAuthTokenSalt))
}

// AuthTokenSaltEqualFold applies the EqualFold predicate on the "auth_token_salt" field.
func AuthTokenSaltEqualFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEqualFold(FieldAuthTokenSalt, v))
}

// AuthTokenSaltContainsFold applies the ContainsFold predicate on the "auth_token_salt" field.
func AuthTokenSaltContainsFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContainsFold(FieldAuthTokenSalt, v))
}

// AuthedEQ applies the EQ predicate on the "authed" field.
func AuthedEQ(v bool) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldAuthed, v))
}

// AuthedNEQ applies the NEQ predicate on the "authed" field.
func AuthedNEQ(v bool) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldAuthed, v))
}

// AuthedIsNil applies the IsNil predicate on the "authed" field.
func AuthedIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldAuthed))
}

// AuthedNotNil applies the NotNil predicate on the "authed" field.
func AuthedNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldAuthed))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.RootUser {
	return predicate.RootUser(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.RootUser {
	return predicate.RootUser(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.RootUser {
	return predicate.RootUser(sql.FieldContainsFold(FieldRemark, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RootUser) predicate.RootUser {
	return predicate.RootUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RootUser) predicate.RootUser {
	return predicate.RootUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RootUser) predicate.RootUser {
	return predicate.RootUser(sql.NotPredicates(p))
}
