// Code generated by ent, DO NOT EDIT.

package appconfig

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppID, v))
}

// EnableSimulateOrder applies equality check predicate on the "enable_simulate_order" field. It's identical to EnableSimulateOrderEQ.
func EnableSimulateOrder(v bool) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEnableSimulateOrder, v))
}

// SimulateOrderCouponMode applies equality check predicate on the "simulate_order_coupon_mode" field. It's identical to SimulateOrderCouponModeEQ.
func SimulateOrderCouponMode(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponProbability applies equality check predicate on the "simulate_order_coupon_probability" field. It's identical to SimulateOrderCouponProbabilityEQ.
func SimulateOrderCouponProbability(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCashableProfitProbability applies equality check predicate on the "simulate_order_cashable_profit_probability" field. It's identical to SimulateOrderCashableProfitProbabilityEQ.
func SimulateOrderCashableProfitProbability(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCashableProfitProbability, v))
}

// MaxUnpaidOrders applies equality check predicate on the "max_unpaid_orders" field. It's identical to MaxUnpaidOrdersEQ.
func MaxUnpaidOrders(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldMaxUnpaidOrders, v))
}

// MaxTypedCouponsPerOrder applies equality check predicate on the "max_typed_coupons_per_order" field. It's identical to MaxTypedCouponsPerOrderEQ.
func MaxTypedCouponsPerOrder(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldMaxTypedCouponsPerOrder, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldDeletedAt, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldAppID))
}

// EnableSimulateOrderEQ applies the EQ predicate on the "enable_simulate_order" field.
func EnableSimulateOrderEQ(v bool) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldEnableSimulateOrder, v))
}

// EnableSimulateOrderNEQ applies the NEQ predicate on the "enable_simulate_order" field.
func EnableSimulateOrderNEQ(v bool) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldEnableSimulateOrder, v))
}

// EnableSimulateOrderIsNil applies the IsNil predicate on the "enable_simulate_order" field.
func EnableSimulateOrderIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldEnableSimulateOrder))
}

// EnableSimulateOrderNotNil applies the NotNil predicate on the "enable_simulate_order" field.
func EnableSimulateOrderNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldEnableSimulateOrder))
}

// SimulateOrderCouponModeEQ applies the EQ predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeNEQ applies the NEQ predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeIn applies the In predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldSimulateOrderCouponMode, vs...))
}

// SimulateOrderCouponModeNotIn applies the NotIn predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNotIn(vs ...string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldSimulateOrderCouponMode, vs...))
}

// SimulateOrderCouponModeGT applies the GT predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeGT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeGTE applies the GTE predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeLT applies the LT predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeLT(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeLTE applies the LTE predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeContains applies the Contains predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeContains(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContains(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeHasPrefix applies the HasPrefix predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasPrefix(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeHasSuffix applies the HasSuffix predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldHasSuffix(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeIsNil applies the IsNil predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldSimulateOrderCouponMode))
}

// SimulateOrderCouponModeNotNil applies the NotNil predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldSimulateOrderCouponMode))
}

// SimulateOrderCouponModeEqualFold applies the EqualFold predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEqualFold(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponModeContainsFold applies the ContainsFold predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldContainsFold(FieldSimulateOrderCouponMode, v))
}

// SimulateOrderCouponProbabilityEQ applies the EQ predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityNEQ applies the NEQ predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityIn applies the In predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityIn(vs ...decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldSimulateOrderCouponProbability, vs...))
}

// SimulateOrderCouponProbabilityNotIn applies the NotIn predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNotIn(vs ...decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldSimulateOrderCouponProbability, vs...))
}

// SimulateOrderCouponProbabilityGT applies the GT predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityGT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityGTE applies the GTE predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityGTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityLT applies the LT predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityLT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityLTE applies the LTE predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityLTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldSimulateOrderCouponProbability, v))
}

// SimulateOrderCouponProbabilityIsNil applies the IsNil predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldSimulateOrderCouponProbability))
}

// SimulateOrderCouponProbabilityNotNil applies the NotNil predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldSimulateOrderCouponProbability))
}

// SimulateOrderCashableProfitProbabilityEQ applies the EQ predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityNEQ applies the NEQ predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityIn applies the In predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityIn(vs ...decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldSimulateOrderCashableProfitProbability, vs...))
}

// SimulateOrderCashableProfitProbabilityNotIn applies the NotIn predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNotIn(vs ...decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldSimulateOrderCashableProfitProbability, vs...))
}

// SimulateOrderCashableProfitProbabilityGT applies the GT predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityGT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityGTE applies the GTE predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityGTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityLT applies the LT predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityLT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityLTE applies the LTE predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityLTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldSimulateOrderCashableProfitProbability, v))
}

// SimulateOrderCashableProfitProbabilityIsNil applies the IsNil predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldSimulateOrderCashableProfitProbability))
}

// SimulateOrderCashableProfitProbabilityNotNil applies the NotNil predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldSimulateOrderCashableProfitProbability))
}

// MaxUnpaidOrdersEQ applies the EQ predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersNEQ applies the NEQ predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersIn applies the In predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldMaxUnpaidOrders, vs...))
}

// MaxUnpaidOrdersNotIn applies the NotIn predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNotIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldMaxUnpaidOrders, vs...))
}

// MaxUnpaidOrdersGT applies the GT predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersGTE applies the GTE predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersLT applies the LT predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersLTE applies the LTE predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldMaxUnpaidOrders, v))
}

// MaxUnpaidOrdersIsNil applies the IsNil predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldMaxUnpaidOrders))
}

// MaxUnpaidOrdersNotNil applies the NotNil predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldMaxUnpaidOrders))
}

// MaxTypedCouponsPerOrderEQ applies the EQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldEQ(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderNEQ applies the NEQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNEQ(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderIn applies the In predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIn(FieldMaxTypedCouponsPerOrder, vs...))
}

// MaxTypedCouponsPerOrderNotIn applies the NotIn predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotIn(vs ...uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotIn(FieldMaxTypedCouponsPerOrder, vs...))
}

// MaxTypedCouponsPerOrderGT applies the GT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGT(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderGTE applies the GTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldGTE(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderLT applies the LT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLT(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderLTE applies the LTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(sql.FieldLTE(FieldMaxTypedCouponsPerOrder, v))
}

// MaxTypedCouponsPerOrderIsNil applies the IsNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIsNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldIsNull(FieldMaxTypedCouponsPerOrder))
}

// MaxTypedCouponsPerOrderNotNil applies the NotNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotNil() predicate.AppConfig {
	return predicate.AppConfig(sql.FieldNotNull(FieldMaxTypedCouponsPerOrder))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(sql.NotPredicates(p))
}
