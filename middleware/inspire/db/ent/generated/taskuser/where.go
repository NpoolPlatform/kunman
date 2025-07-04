// Code generated by ent, DO NOT EDIT.

package taskuser

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldEntID, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldAppID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldUserID, v))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldTaskID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldEventID, v))
}

// TaskState applies equality check predicate on the "task_state" field. It's identical to TaskStateEQ.
func TaskState(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldTaskState, v))
}

// RewardState applies equality check predicate on the "reward_state" field. It's identical to RewardStateEQ.
func RewardState(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldRewardState, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldEntID, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldAppID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldUserID))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldTaskID, vs...))
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldTaskID, v))
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldTaskID, v))
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldTaskID, v))
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldTaskID, v))
}

// TaskIDIsNil applies the IsNil predicate on the "task_id" field.
func TaskIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldTaskID))
}

// TaskIDNotNil applies the NotNil predicate on the "task_id" field.
func TaskIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldTaskID))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldEventID, vs...))
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldEventID, v))
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldEventID, v))
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldEventID, v))
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v uuid.UUID) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldEventID, v))
}

// EventIDIsNil applies the IsNil predicate on the "event_id" field.
func EventIDIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldEventID))
}

// EventIDNotNil applies the NotNil predicate on the "event_id" field.
func EventIDNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldEventID))
}

// TaskStateEQ applies the EQ predicate on the "task_state" field.
func TaskStateEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldTaskState, v))
}

// TaskStateNEQ applies the NEQ predicate on the "task_state" field.
func TaskStateNEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldTaskState, v))
}

// TaskStateIn applies the In predicate on the "task_state" field.
func TaskStateIn(vs ...string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldTaskState, vs...))
}

// TaskStateNotIn applies the NotIn predicate on the "task_state" field.
func TaskStateNotIn(vs ...string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldTaskState, vs...))
}

// TaskStateGT applies the GT predicate on the "task_state" field.
func TaskStateGT(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldTaskState, v))
}

// TaskStateGTE applies the GTE predicate on the "task_state" field.
func TaskStateGTE(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldTaskState, v))
}

// TaskStateLT applies the LT predicate on the "task_state" field.
func TaskStateLT(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldTaskState, v))
}

// TaskStateLTE applies the LTE predicate on the "task_state" field.
func TaskStateLTE(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldTaskState, v))
}

// TaskStateContains applies the Contains predicate on the "task_state" field.
func TaskStateContains(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldContains(FieldTaskState, v))
}

// TaskStateHasPrefix applies the HasPrefix predicate on the "task_state" field.
func TaskStateHasPrefix(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldHasPrefix(FieldTaskState, v))
}

// TaskStateHasSuffix applies the HasSuffix predicate on the "task_state" field.
func TaskStateHasSuffix(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldHasSuffix(FieldTaskState, v))
}

// TaskStateIsNil applies the IsNil predicate on the "task_state" field.
func TaskStateIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldTaskState))
}

// TaskStateNotNil applies the NotNil predicate on the "task_state" field.
func TaskStateNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldTaskState))
}

// TaskStateEqualFold applies the EqualFold predicate on the "task_state" field.
func TaskStateEqualFold(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEqualFold(FieldTaskState, v))
}

// TaskStateContainsFold applies the ContainsFold predicate on the "task_state" field.
func TaskStateContainsFold(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldContainsFold(FieldTaskState, v))
}

// RewardStateEQ applies the EQ predicate on the "reward_state" field.
func RewardStateEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEQ(FieldRewardState, v))
}

// RewardStateNEQ applies the NEQ predicate on the "reward_state" field.
func RewardStateNEQ(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNEQ(FieldRewardState, v))
}

// RewardStateIn applies the In predicate on the "reward_state" field.
func RewardStateIn(vs ...string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIn(FieldRewardState, vs...))
}

// RewardStateNotIn applies the NotIn predicate on the "reward_state" field.
func RewardStateNotIn(vs ...string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotIn(FieldRewardState, vs...))
}

// RewardStateGT applies the GT predicate on the "reward_state" field.
func RewardStateGT(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGT(FieldRewardState, v))
}

// RewardStateGTE applies the GTE predicate on the "reward_state" field.
func RewardStateGTE(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldGTE(FieldRewardState, v))
}

// RewardStateLT applies the LT predicate on the "reward_state" field.
func RewardStateLT(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLT(FieldRewardState, v))
}

// RewardStateLTE applies the LTE predicate on the "reward_state" field.
func RewardStateLTE(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldLTE(FieldRewardState, v))
}

// RewardStateContains applies the Contains predicate on the "reward_state" field.
func RewardStateContains(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldContains(FieldRewardState, v))
}

// RewardStateHasPrefix applies the HasPrefix predicate on the "reward_state" field.
func RewardStateHasPrefix(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldHasPrefix(FieldRewardState, v))
}

// RewardStateHasSuffix applies the HasSuffix predicate on the "reward_state" field.
func RewardStateHasSuffix(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldHasSuffix(FieldRewardState, v))
}

// RewardStateIsNil applies the IsNil predicate on the "reward_state" field.
func RewardStateIsNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldIsNull(FieldRewardState))
}

// RewardStateNotNil applies the NotNil predicate on the "reward_state" field.
func RewardStateNotNil() predicate.TaskUser {
	return predicate.TaskUser(sql.FieldNotNull(FieldRewardState))
}

// RewardStateEqualFold applies the EqualFold predicate on the "reward_state" field.
func RewardStateEqualFold(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldEqualFold(FieldRewardState, v))
}

// RewardStateContainsFold applies the ContainsFold predicate on the "reward_state" field.
func RewardStateContainsFold(v string) predicate.TaskUser {
	return predicate.TaskUser(sql.FieldContainsFold(FieldRewardState, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskUser) predicate.TaskUser {
	return predicate.TaskUser(sql.NotPredicates(p))
}
