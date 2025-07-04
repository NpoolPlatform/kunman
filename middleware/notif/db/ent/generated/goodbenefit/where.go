// Code generated by ent, DO NOT EDIT.

package goodbenefit

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldEntID, v))
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodID, v))
}

// GoodType applies equality check predicate on the "good_type" field. It's identical to GoodTypeEQ.
func GoodType(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodType, v))
}

// GoodName applies equality check predicate on the "good_name" field. It's identical to GoodNameEQ.
func GoodName(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodName, v))
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldCoinTypeID, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldAmount, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldState, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldMessage, v))
}

// BenefitDate applies equality check predicate on the "benefit_date" field. It's identical to BenefitDateEQ.
func BenefitDate(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldBenefitDate, v))
}

// TxID applies equality check predicate on the "tx_id" field. It's identical to TxIDEQ.
func TxID(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldTxID, v))
}

// Generated applies equality check predicate on the "generated" field. It's identical to GeneratedEQ.
func Generated(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGenerated, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldEntID, v))
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodID, v))
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldGoodID, v))
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldGoodID, vs...))
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldGoodID, vs...))
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldGoodID, v))
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldGoodID, v))
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldGoodID, v))
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldGoodID, v))
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldGoodID))
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldGoodID))
}

// GoodTypeEQ applies the EQ predicate on the "good_type" field.
func GoodTypeEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodType, v))
}

// GoodTypeNEQ applies the NEQ predicate on the "good_type" field.
func GoodTypeNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldGoodType, v))
}

// GoodTypeIn applies the In predicate on the "good_type" field.
func GoodTypeIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldGoodType, vs...))
}

// GoodTypeNotIn applies the NotIn predicate on the "good_type" field.
func GoodTypeNotIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldGoodType, vs...))
}

// GoodTypeGT applies the GT predicate on the "good_type" field.
func GoodTypeGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldGoodType, v))
}

// GoodTypeGTE applies the GTE predicate on the "good_type" field.
func GoodTypeGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldGoodType, v))
}

// GoodTypeLT applies the LT predicate on the "good_type" field.
func GoodTypeLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldGoodType, v))
}

// GoodTypeLTE applies the LTE predicate on the "good_type" field.
func GoodTypeLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldGoodType, v))
}

// GoodTypeContains applies the Contains predicate on the "good_type" field.
func GoodTypeContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContains(FieldGoodType, v))
}

// GoodTypeHasPrefix applies the HasPrefix predicate on the "good_type" field.
func GoodTypeHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasPrefix(FieldGoodType, v))
}

// GoodTypeHasSuffix applies the HasSuffix predicate on the "good_type" field.
func GoodTypeHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasSuffix(FieldGoodType, v))
}

// GoodTypeIsNil applies the IsNil predicate on the "good_type" field.
func GoodTypeIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldGoodType))
}

// GoodTypeNotNil applies the NotNil predicate on the "good_type" field.
func GoodTypeNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldGoodType))
}

// GoodTypeEqualFold applies the EqualFold predicate on the "good_type" field.
func GoodTypeEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEqualFold(FieldGoodType, v))
}

// GoodTypeContainsFold applies the ContainsFold predicate on the "good_type" field.
func GoodTypeContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContainsFold(FieldGoodType, v))
}

// GoodNameEQ applies the EQ predicate on the "good_name" field.
func GoodNameEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGoodName, v))
}

// GoodNameNEQ applies the NEQ predicate on the "good_name" field.
func GoodNameNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldGoodName, v))
}

// GoodNameIn applies the In predicate on the "good_name" field.
func GoodNameIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldGoodName, vs...))
}

// GoodNameNotIn applies the NotIn predicate on the "good_name" field.
func GoodNameNotIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldGoodName, vs...))
}

// GoodNameGT applies the GT predicate on the "good_name" field.
func GoodNameGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldGoodName, v))
}

// GoodNameGTE applies the GTE predicate on the "good_name" field.
func GoodNameGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldGoodName, v))
}

// GoodNameLT applies the LT predicate on the "good_name" field.
func GoodNameLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldGoodName, v))
}

// GoodNameLTE applies the LTE predicate on the "good_name" field.
func GoodNameLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldGoodName, v))
}

// GoodNameContains applies the Contains predicate on the "good_name" field.
func GoodNameContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContains(FieldGoodName, v))
}

// GoodNameHasPrefix applies the HasPrefix predicate on the "good_name" field.
func GoodNameHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasPrefix(FieldGoodName, v))
}

// GoodNameHasSuffix applies the HasSuffix predicate on the "good_name" field.
func GoodNameHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasSuffix(FieldGoodName, v))
}

// GoodNameIsNil applies the IsNil predicate on the "good_name" field.
func GoodNameIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldGoodName))
}

// GoodNameNotNil applies the NotNil predicate on the "good_name" field.
func GoodNameNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldGoodName))
}

// GoodNameEqualFold applies the EqualFold predicate on the "good_name" field.
func GoodNameEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEqualFold(FieldGoodName, v))
}

// GoodNameContainsFold applies the ContainsFold predicate on the "good_name" field.
func GoodNameContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContainsFold(FieldGoodName, v))
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldCoinTypeID, v))
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldCoinTypeID, v))
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldCoinTypeID, v))
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldCoinTypeID, v))
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldCoinTypeID, v))
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldCoinTypeID, v))
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldCoinTypeID))
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldCoinTypeID))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldAmount, v))
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContains(FieldAmount, v))
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasPrefix(FieldAmount, v))
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasSuffix(FieldAmount, v))
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldAmount))
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldAmount))
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEqualFold(FieldAmount, v))
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContainsFold(FieldAmount, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasSuffix(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldState))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContainsFold(FieldState, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldMessage))
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldMessage))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldContainsFold(FieldMessage, v))
}

// BenefitDateEQ applies the EQ predicate on the "benefit_date" field.
func BenefitDateEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldBenefitDate, v))
}

// BenefitDateNEQ applies the NEQ predicate on the "benefit_date" field.
func BenefitDateNEQ(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldBenefitDate, v))
}

// BenefitDateIn applies the In predicate on the "benefit_date" field.
func BenefitDateIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldBenefitDate, vs...))
}

// BenefitDateNotIn applies the NotIn predicate on the "benefit_date" field.
func BenefitDateNotIn(vs ...uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldBenefitDate, vs...))
}

// BenefitDateGT applies the GT predicate on the "benefit_date" field.
func BenefitDateGT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldBenefitDate, v))
}

// BenefitDateGTE applies the GTE predicate on the "benefit_date" field.
func BenefitDateGTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldBenefitDate, v))
}

// BenefitDateLT applies the LT predicate on the "benefit_date" field.
func BenefitDateLT(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldBenefitDate, v))
}

// BenefitDateLTE applies the LTE predicate on the "benefit_date" field.
func BenefitDateLTE(v uint32) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldBenefitDate, v))
}

// BenefitDateIsNil applies the IsNil predicate on the "benefit_date" field.
func BenefitDateIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldBenefitDate))
}

// BenefitDateNotNil applies the NotNil predicate on the "benefit_date" field.
func BenefitDateNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldBenefitDate))
}

// TxIDEQ applies the EQ predicate on the "tx_id" field.
func TxIDEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldTxID, v))
}

// TxIDNEQ applies the NEQ predicate on the "tx_id" field.
func TxIDNEQ(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldTxID, v))
}

// TxIDIn applies the In predicate on the "tx_id" field.
func TxIDIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIn(FieldTxID, vs...))
}

// TxIDNotIn applies the NotIn predicate on the "tx_id" field.
func TxIDNotIn(vs ...uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotIn(FieldTxID, vs...))
}

// TxIDGT applies the GT predicate on the "tx_id" field.
func TxIDGT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGT(FieldTxID, v))
}

// TxIDGTE applies the GTE predicate on the "tx_id" field.
func TxIDGTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldGTE(FieldTxID, v))
}

// TxIDLT applies the LT predicate on the "tx_id" field.
func TxIDLT(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLT(FieldTxID, v))
}

// TxIDLTE applies the LTE predicate on the "tx_id" field.
func TxIDLTE(v uuid.UUID) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldLTE(FieldTxID, v))
}

// TxIDIsNil applies the IsNil predicate on the "tx_id" field.
func TxIDIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldTxID))
}

// TxIDNotNil applies the NotNil predicate on the "tx_id" field.
func TxIDNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldTxID))
}

// GeneratedEQ applies the EQ predicate on the "generated" field.
func GeneratedEQ(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldEQ(FieldGenerated, v))
}

// GeneratedNEQ applies the NEQ predicate on the "generated" field.
func GeneratedNEQ(v bool) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNEQ(FieldGenerated, v))
}

// GeneratedIsNil applies the IsNil predicate on the "generated" field.
func GeneratedIsNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldIsNull(FieldGenerated))
}

// GeneratedNotNil applies the NotNil predicate on the "generated" field.
func GeneratedNotNil() predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.FieldNotNull(FieldGenerated))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GoodBenefit) predicate.GoodBenefit {
	return predicate.GoodBenefit(sql.NotPredicates(p))
}
