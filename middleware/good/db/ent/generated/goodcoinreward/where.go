// Code generated by ent, DO NOT EDIT.

package goodcoinreward

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldDeletedAt, v))
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldGoodID, v))
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldCoinTypeID, v))
}

// RewardTid applies equality check predicate on the "reward_tid" field. It's identical to RewardTidEQ.
func RewardTid(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldRewardTid, v))
}

// NextRewardStartAmount applies equality check predicate on the "next_reward_start_amount" field. It's identical to NextRewardStartAmountEQ.
func NextRewardStartAmount(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldNextRewardStartAmount, v))
}

// LastRewardAmount applies equality check predicate on the "last_reward_amount" field. It's identical to LastRewardAmountEQ.
func LastRewardAmount(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldLastRewardAmount, v))
}

// LastUnitRewardAmount applies equality check predicate on the "last_unit_reward_amount" field. It's identical to LastUnitRewardAmountEQ.
func LastUnitRewardAmount(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldLastUnitRewardAmount, v))
}

// TotalRewardAmount applies equality check predicate on the "total_reward_amount" field. It's identical to TotalRewardAmountEQ.
func TotalRewardAmount(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldTotalRewardAmount, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldDeletedAt, v))
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldGoodID, v))
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldGoodID, v))
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldGoodID, vs...))
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldGoodID, vs...))
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldGoodID, v))
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldGoodID, v))
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldGoodID, v))
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldGoodID, v))
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldGoodID))
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldGoodID))
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldCoinTypeID, v))
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldCoinTypeID, v))
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldCoinTypeID, vs...))
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldCoinTypeID, v))
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldCoinTypeID, v))
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldCoinTypeID, v))
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldCoinTypeID, v))
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldCoinTypeID))
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldCoinTypeID))
}

// RewardTidEQ applies the EQ predicate on the "reward_tid" field.
func RewardTidEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldRewardTid, v))
}

// RewardTidNEQ applies the NEQ predicate on the "reward_tid" field.
func RewardTidNEQ(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldRewardTid, v))
}

// RewardTidIn applies the In predicate on the "reward_tid" field.
func RewardTidIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldRewardTid, vs...))
}

// RewardTidNotIn applies the NotIn predicate on the "reward_tid" field.
func RewardTidNotIn(vs ...uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldRewardTid, vs...))
}

// RewardTidGT applies the GT predicate on the "reward_tid" field.
func RewardTidGT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldRewardTid, v))
}

// RewardTidGTE applies the GTE predicate on the "reward_tid" field.
func RewardTidGTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldRewardTid, v))
}

// RewardTidLT applies the LT predicate on the "reward_tid" field.
func RewardTidLT(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldRewardTid, v))
}

// RewardTidLTE applies the LTE predicate on the "reward_tid" field.
func RewardTidLTE(v uuid.UUID) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldRewardTid, v))
}

// RewardTidIsNil applies the IsNil predicate on the "reward_tid" field.
func RewardTidIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldRewardTid))
}

// RewardTidNotNil applies the NotNil predicate on the "reward_tid" field.
func RewardTidNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldRewardTid))
}

// NextRewardStartAmountEQ applies the EQ predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountNEQ applies the NEQ predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountNEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountIn applies the In predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldNextRewardStartAmount, vs...))
}

// NextRewardStartAmountNotIn applies the NotIn predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountNotIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldNextRewardStartAmount, vs...))
}

// NextRewardStartAmountGT applies the GT predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountGT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountGTE applies the GTE predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountGTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountLT applies the LT predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountLT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountLTE applies the LTE predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountLTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldNextRewardStartAmount, v))
}

// NextRewardStartAmountIsNil applies the IsNil predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldNextRewardStartAmount))
}

// NextRewardStartAmountNotNil applies the NotNil predicate on the "next_reward_start_amount" field.
func NextRewardStartAmountNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldNextRewardStartAmount))
}

// LastRewardAmountEQ applies the EQ predicate on the "last_reward_amount" field.
func LastRewardAmountEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldLastRewardAmount, v))
}

// LastRewardAmountNEQ applies the NEQ predicate on the "last_reward_amount" field.
func LastRewardAmountNEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldLastRewardAmount, v))
}

// LastRewardAmountIn applies the In predicate on the "last_reward_amount" field.
func LastRewardAmountIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldLastRewardAmount, vs...))
}

// LastRewardAmountNotIn applies the NotIn predicate on the "last_reward_amount" field.
func LastRewardAmountNotIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldLastRewardAmount, vs...))
}

// LastRewardAmountGT applies the GT predicate on the "last_reward_amount" field.
func LastRewardAmountGT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldLastRewardAmount, v))
}

// LastRewardAmountGTE applies the GTE predicate on the "last_reward_amount" field.
func LastRewardAmountGTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldLastRewardAmount, v))
}

// LastRewardAmountLT applies the LT predicate on the "last_reward_amount" field.
func LastRewardAmountLT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldLastRewardAmount, v))
}

// LastRewardAmountLTE applies the LTE predicate on the "last_reward_amount" field.
func LastRewardAmountLTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldLastRewardAmount, v))
}

// LastRewardAmountIsNil applies the IsNil predicate on the "last_reward_amount" field.
func LastRewardAmountIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldLastRewardAmount))
}

// LastRewardAmountNotNil applies the NotNil predicate on the "last_reward_amount" field.
func LastRewardAmountNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldLastRewardAmount))
}

// LastUnitRewardAmountEQ applies the EQ predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountNEQ applies the NEQ predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountNEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountIn applies the In predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldLastUnitRewardAmount, vs...))
}

// LastUnitRewardAmountNotIn applies the NotIn predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountNotIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldLastUnitRewardAmount, vs...))
}

// LastUnitRewardAmountGT applies the GT predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountGT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountGTE applies the GTE predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountGTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountLT applies the LT predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountLT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountLTE applies the LTE predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountLTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldLastUnitRewardAmount, v))
}

// LastUnitRewardAmountIsNil applies the IsNil predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldLastUnitRewardAmount))
}

// LastUnitRewardAmountNotNil applies the NotNil predicate on the "last_unit_reward_amount" field.
func LastUnitRewardAmountNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldLastUnitRewardAmount))
}

// TotalRewardAmountEQ applies the EQ predicate on the "total_reward_amount" field.
func TotalRewardAmountEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldEQ(FieldTotalRewardAmount, v))
}

// TotalRewardAmountNEQ applies the NEQ predicate on the "total_reward_amount" field.
func TotalRewardAmountNEQ(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNEQ(FieldTotalRewardAmount, v))
}

// TotalRewardAmountIn applies the In predicate on the "total_reward_amount" field.
func TotalRewardAmountIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIn(FieldTotalRewardAmount, vs...))
}

// TotalRewardAmountNotIn applies the NotIn predicate on the "total_reward_amount" field.
func TotalRewardAmountNotIn(vs ...decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotIn(FieldTotalRewardAmount, vs...))
}

// TotalRewardAmountGT applies the GT predicate on the "total_reward_amount" field.
func TotalRewardAmountGT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGT(FieldTotalRewardAmount, v))
}

// TotalRewardAmountGTE applies the GTE predicate on the "total_reward_amount" field.
func TotalRewardAmountGTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldGTE(FieldTotalRewardAmount, v))
}

// TotalRewardAmountLT applies the LT predicate on the "total_reward_amount" field.
func TotalRewardAmountLT(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLT(FieldTotalRewardAmount, v))
}

// TotalRewardAmountLTE applies the LTE predicate on the "total_reward_amount" field.
func TotalRewardAmountLTE(v decimal.Decimal) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldLTE(FieldTotalRewardAmount, v))
}

// TotalRewardAmountIsNil applies the IsNil predicate on the "total_reward_amount" field.
func TotalRewardAmountIsNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldIsNull(FieldTotalRewardAmount))
}

// TotalRewardAmountNotNil applies the NotNil predicate on the "total_reward_amount" field.
func TotalRewardAmountNotNil() predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.FieldNotNull(FieldTotalRewardAmount))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodCoinReward) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodCoinReward) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GoodCoinReward) predicate.GoodCoinReward {
	return predicate.GoodCoinReward(sql.NotPredicates(p))
}
