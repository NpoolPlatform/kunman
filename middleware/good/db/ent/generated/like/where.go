// Code generated by ent, DO NOT EDIT.

package like

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldAppGoodID, v))
}

// Like applies equality check predicate on the "like" field. It's identical to LikeEQ.
func Like(v bool) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldLike, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldDeletedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Like {
	return predicate.Like(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Like {
	return predicate.Like(sql.FieldNotNull(FieldUserID))
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldAppGoodID, v))
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldAppGoodID, v))
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldAppGoodID, vs...))
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldAppGoodID, vs...))
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldAppGoodID, v))
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldAppGoodID, v))
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldAppGoodID, v))
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldAppGoodID, v))
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.Like {
	return predicate.Like(sql.FieldIsNull(FieldAppGoodID))
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.Like {
	return predicate.Like(sql.FieldNotNull(FieldAppGoodID))
}

// LikeEQ applies the EQ predicate on the "like" field.
func LikeEQ(v bool) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldLike, v))
}

// LikeNEQ applies the NEQ predicate on the "like" field.
func LikeNEQ(v bool) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldLike, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Like) predicate.Like {
	return predicate.Like(sql.NotPredicates(p))
}
