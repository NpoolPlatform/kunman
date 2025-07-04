// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldEntID, v))
}

// PoolID applies equality check predicate on the "pool_id" field. It's identical to PoolIDEQ.
func PoolID(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldPoolID, v))
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCoinTypeID, v))
}

// CoinType applies equality check predicate on the "coin_type" field. It's identical to CoinTypeEQ.
func CoinType(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCoinType, v))
}

// FeeRatio applies equality check predicate on the "fee_ratio" field. It's identical to FeeRatioEQ.
func FeeRatio(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldFeeRatio, v))
}

// FixedRevenueAble applies equality check predicate on the "fixed_revenue_able" field. It's identical to FixedRevenueAbleEQ.
func FixedRevenueAble(v bool) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldFixedRevenueAble, v))
}

// LeastTransferAmount applies equality check predicate on the "least_transfer_amount" field. It's identical to LeastTransferAmountEQ.
func LeastTransferAmount(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldLeastTransferAmount, v))
}

// BenefitIntervalSeconds applies equality check predicate on the "benefit_interval_seconds" field. It's identical to BenefitIntervalSecondsEQ.
func BenefitIntervalSeconds(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldBenefitIntervalSeconds, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldEntID, v))
}

// PoolIDEQ applies the EQ predicate on the "pool_id" field.
func PoolIDEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldPoolID, v))
}

// PoolIDNEQ applies the NEQ predicate on the "pool_id" field.
func PoolIDNEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldPoolID, v))
}

// PoolIDIn applies the In predicate on the "pool_id" field.
func PoolIDIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldPoolID, vs...))
}

// PoolIDNotIn applies the NotIn predicate on the "pool_id" field.
func PoolIDNotIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldPoolID, vs...))
}

// PoolIDGT applies the GT predicate on the "pool_id" field.
func PoolIDGT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldPoolID, v))
}

// PoolIDGTE applies the GTE predicate on the "pool_id" field.
func PoolIDGTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldPoolID, v))
}

// PoolIDLT applies the LT predicate on the "pool_id" field.
func PoolIDLT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldPoolID, v))
}

// PoolIDLTE applies the LTE predicate on the "pool_id" field.
func PoolIDLTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldPoolID, v))
}

// PoolIDIsNil applies the IsNil predicate on the "pool_id" field.
func PoolIDIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldPoolID))
}

// PoolIDNotNil applies the NotNil predicate on the "pool_id" field.
func PoolIDNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldPoolID))
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCoinTypeID, v))
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldCoinTypeID, v))
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldCoinTypeID, v))
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldCoinTypeID, v))
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldCoinTypeID, v))
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldCoinTypeID, v))
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldCoinTypeID))
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldCoinTypeID))
}

// CoinTypeEQ applies the EQ predicate on the "coin_type" field.
func CoinTypeEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldCoinType, v))
}

// CoinTypeNEQ applies the NEQ predicate on the "coin_type" field.
func CoinTypeNEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldCoinType, v))
}

// CoinTypeIn applies the In predicate on the "coin_type" field.
func CoinTypeIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldCoinType, vs...))
}

// CoinTypeNotIn applies the NotIn predicate on the "coin_type" field.
func CoinTypeNotIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldCoinType, vs...))
}

// CoinTypeGT applies the GT predicate on the "coin_type" field.
func CoinTypeGT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldCoinType, v))
}

// CoinTypeGTE applies the GTE predicate on the "coin_type" field.
func CoinTypeGTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldCoinType, v))
}

// CoinTypeLT applies the LT predicate on the "coin_type" field.
func CoinTypeLT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldCoinType, v))
}

// CoinTypeLTE applies the LTE predicate on the "coin_type" field.
func CoinTypeLTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldCoinType, v))
}

// CoinTypeContains applies the Contains predicate on the "coin_type" field.
func CoinTypeContains(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContains(FieldCoinType, v))
}

// CoinTypeHasPrefix applies the HasPrefix predicate on the "coin_type" field.
func CoinTypeHasPrefix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasPrefix(FieldCoinType, v))
}

// CoinTypeHasSuffix applies the HasSuffix predicate on the "coin_type" field.
func CoinTypeHasSuffix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasSuffix(FieldCoinType, v))
}

// CoinTypeIsNil applies the IsNil predicate on the "coin_type" field.
func CoinTypeIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldCoinType))
}

// CoinTypeNotNil applies the NotNil predicate on the "coin_type" field.
func CoinTypeNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldCoinType))
}

// CoinTypeEqualFold applies the EqualFold predicate on the "coin_type" field.
func CoinTypeEqualFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEqualFold(FieldCoinType, v))
}

// CoinTypeContainsFold applies the ContainsFold predicate on the "coin_type" field.
func CoinTypeContainsFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContainsFold(FieldCoinType, v))
}

// FeeRatioEQ applies the EQ predicate on the "fee_ratio" field.
func FeeRatioEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldFeeRatio, v))
}

// FeeRatioNEQ applies the NEQ predicate on the "fee_ratio" field.
func FeeRatioNEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldFeeRatio, v))
}

// FeeRatioIn applies the In predicate on the "fee_ratio" field.
func FeeRatioIn(vs ...decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldFeeRatio, vs...))
}

// FeeRatioNotIn applies the NotIn predicate on the "fee_ratio" field.
func FeeRatioNotIn(vs ...decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldFeeRatio, vs...))
}

// FeeRatioGT applies the GT predicate on the "fee_ratio" field.
func FeeRatioGT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldFeeRatio, v))
}

// FeeRatioGTE applies the GTE predicate on the "fee_ratio" field.
func FeeRatioGTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldFeeRatio, v))
}

// FeeRatioLT applies the LT predicate on the "fee_ratio" field.
func FeeRatioLT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldFeeRatio, v))
}

// FeeRatioLTE applies the LTE predicate on the "fee_ratio" field.
func FeeRatioLTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldFeeRatio, v))
}

// FeeRatioIsNil applies the IsNil predicate on the "fee_ratio" field.
func FeeRatioIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldFeeRatio))
}

// FeeRatioNotNil applies the NotNil predicate on the "fee_ratio" field.
func FeeRatioNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldFeeRatio))
}

// FixedRevenueAbleEQ applies the EQ predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleEQ(v bool) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldFixedRevenueAble, v))
}

// FixedRevenueAbleNEQ applies the NEQ predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleNEQ(v bool) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldFixedRevenueAble, v))
}

// FixedRevenueAbleIsNil applies the IsNil predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldFixedRevenueAble))
}

// FixedRevenueAbleNotNil applies the NotNil predicate on the "fixed_revenue_able" field.
func FixedRevenueAbleNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldFixedRevenueAble))
}

// LeastTransferAmountEQ applies the EQ predicate on the "least_transfer_amount" field.
func LeastTransferAmountEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldLeastTransferAmount, v))
}

// LeastTransferAmountNEQ applies the NEQ predicate on the "least_transfer_amount" field.
func LeastTransferAmountNEQ(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldLeastTransferAmount, v))
}

// LeastTransferAmountIn applies the In predicate on the "least_transfer_amount" field.
func LeastTransferAmountIn(vs ...decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldLeastTransferAmount, vs...))
}

// LeastTransferAmountNotIn applies the NotIn predicate on the "least_transfer_amount" field.
func LeastTransferAmountNotIn(vs ...decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldLeastTransferAmount, vs...))
}

// LeastTransferAmountGT applies the GT predicate on the "least_transfer_amount" field.
func LeastTransferAmountGT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldLeastTransferAmount, v))
}

// LeastTransferAmountGTE applies the GTE predicate on the "least_transfer_amount" field.
func LeastTransferAmountGTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldLeastTransferAmount, v))
}

// LeastTransferAmountLT applies the LT predicate on the "least_transfer_amount" field.
func LeastTransferAmountLT(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldLeastTransferAmount, v))
}

// LeastTransferAmountLTE applies the LTE predicate on the "least_transfer_amount" field.
func LeastTransferAmountLTE(v decimal.Decimal) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldLeastTransferAmount, v))
}

// LeastTransferAmountIsNil applies the IsNil predicate on the "least_transfer_amount" field.
func LeastTransferAmountIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldLeastTransferAmount))
}

// LeastTransferAmountNotNil applies the NotNil predicate on the "least_transfer_amount" field.
func LeastTransferAmountNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldLeastTransferAmount))
}

// BenefitIntervalSecondsEQ applies the EQ predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsNEQ applies the NEQ predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsNEQ(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsIn applies the In predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldBenefitIntervalSeconds, vs...))
}

// BenefitIntervalSecondsNotIn applies the NotIn predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsNotIn(vs ...uint32) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldBenefitIntervalSeconds, vs...))
}

// BenefitIntervalSecondsGT applies the GT predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsGT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsGTE applies the GTE predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsGTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsLT applies the LT predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsLT(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsLTE applies the LTE predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsLTE(v uint32) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldBenefitIntervalSeconds, v))
}

// BenefitIntervalSecondsIsNil applies the IsNil predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldBenefitIntervalSeconds))
}

// BenefitIntervalSecondsNotNil applies the NotNil predicate on the "benefit_interval_seconds" field.
func BenefitIntervalSecondsNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldBenefitIntervalSeconds))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContainsFold(FieldRemark, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.NotPredicates(p))
}
