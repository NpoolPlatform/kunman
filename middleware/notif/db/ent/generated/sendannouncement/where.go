// Code generated by ent, DO NOT EDIT.

package sendannouncement

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldEntID, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldAppID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldUserID, v))
}

// AnnouncementID applies equality check predicate on the "announcement_id" field. It's identical to AnnouncementIDEQ.
func AnnouncementID(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldAnnouncementID, v))
}

// Channel applies equality check predicate on the "channel" field. It's identical to ChannelEQ.
func Channel(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldChannel, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldEntID, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotNull(FieldAppID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotNull(FieldUserID))
}

// AnnouncementIDEQ applies the EQ predicate on the "announcement_id" field.
func AnnouncementIDEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldAnnouncementID, v))
}

// AnnouncementIDNEQ applies the NEQ predicate on the "announcement_id" field.
func AnnouncementIDNEQ(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldAnnouncementID, v))
}

// AnnouncementIDIn applies the In predicate on the "announcement_id" field.
func AnnouncementIDIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldAnnouncementID, vs...))
}

// AnnouncementIDNotIn applies the NotIn predicate on the "announcement_id" field.
func AnnouncementIDNotIn(vs ...uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldAnnouncementID, vs...))
}

// AnnouncementIDGT applies the GT predicate on the "announcement_id" field.
func AnnouncementIDGT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldAnnouncementID, v))
}

// AnnouncementIDGTE applies the GTE predicate on the "announcement_id" field.
func AnnouncementIDGTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldAnnouncementID, v))
}

// AnnouncementIDLT applies the LT predicate on the "announcement_id" field.
func AnnouncementIDLT(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldAnnouncementID, v))
}

// AnnouncementIDLTE applies the LTE predicate on the "announcement_id" field.
func AnnouncementIDLTE(v uuid.UUID) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldAnnouncementID, v))
}

// AnnouncementIDIsNil applies the IsNil predicate on the "announcement_id" field.
func AnnouncementIDIsNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIsNull(FieldAnnouncementID))
}

// AnnouncementIDNotNil applies the NotNil predicate on the "announcement_id" field.
func AnnouncementIDNotNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotNull(FieldAnnouncementID))
}

// ChannelEQ applies the EQ predicate on the "channel" field.
func ChannelEQ(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEQ(FieldChannel, v))
}

// ChannelNEQ applies the NEQ predicate on the "channel" field.
func ChannelNEQ(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNEQ(FieldChannel, v))
}

// ChannelIn applies the In predicate on the "channel" field.
func ChannelIn(vs ...string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIn(FieldChannel, vs...))
}

// ChannelNotIn applies the NotIn predicate on the "channel" field.
func ChannelNotIn(vs ...string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotIn(FieldChannel, vs...))
}

// ChannelGT applies the GT predicate on the "channel" field.
func ChannelGT(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGT(FieldChannel, v))
}

// ChannelGTE applies the GTE predicate on the "channel" field.
func ChannelGTE(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldGTE(FieldChannel, v))
}

// ChannelLT applies the LT predicate on the "channel" field.
func ChannelLT(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLT(FieldChannel, v))
}

// ChannelLTE applies the LTE predicate on the "channel" field.
func ChannelLTE(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldLTE(FieldChannel, v))
}

// ChannelContains applies the Contains predicate on the "channel" field.
func ChannelContains(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldContains(FieldChannel, v))
}

// ChannelHasPrefix applies the HasPrefix predicate on the "channel" field.
func ChannelHasPrefix(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldHasPrefix(FieldChannel, v))
}

// ChannelHasSuffix applies the HasSuffix predicate on the "channel" field.
func ChannelHasSuffix(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldHasSuffix(FieldChannel, v))
}

// ChannelIsNil applies the IsNil predicate on the "channel" field.
func ChannelIsNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldIsNull(FieldChannel))
}

// ChannelNotNil applies the NotNil predicate on the "channel" field.
func ChannelNotNil() predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldNotNull(FieldChannel))
}

// ChannelEqualFold applies the EqualFold predicate on the "channel" field.
func ChannelEqualFold(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldEqualFold(FieldChannel, v))
}

// ChannelContainsFold applies the ContainsFold predicate on the "channel" field.
func ChannelContainsFold(v string) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.FieldContainsFold(FieldChannel, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SendAnnouncement) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SendAnnouncement) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SendAnnouncement) predicate.SendAnnouncement {
	return predicate.SendAnnouncement(sql.NotPredicates(p))
}
