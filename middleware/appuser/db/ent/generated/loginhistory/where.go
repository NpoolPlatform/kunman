// Code generated by ent, DO NOT EDIT.

package loginhistory

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldEntID, v))
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldAppID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUserID, v))
}

// ClientIP applies equality check predicate on the "client_ip" field. It's identical to ClientIPEQ.
func ClientIP(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldClientIP, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUserAgent, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldLocation, v))
}

// LoginType applies equality check predicate on the "login_type" field. It's identical to LoginTypeEQ.
func LoginType(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldLoginType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldDeletedAt, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldEntID, v))
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldAppID, v))
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldAppID, v))
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldAppID, vs...))
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldAppID, vs...))
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldAppID, v))
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldAppID, v))
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldAppID, v))
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldAppID, v))
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldAppID))
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldAppID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldUserID))
}

// ClientIPEQ applies the EQ predicate on the "client_ip" field.
func ClientIPEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldClientIP, v))
}

// ClientIPNEQ applies the NEQ predicate on the "client_ip" field.
func ClientIPNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldClientIP, v))
}

// ClientIPIn applies the In predicate on the "client_ip" field.
func ClientIPIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldClientIP, vs...))
}

// ClientIPNotIn applies the NotIn predicate on the "client_ip" field.
func ClientIPNotIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldClientIP, vs...))
}

// ClientIPGT applies the GT predicate on the "client_ip" field.
func ClientIPGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldClientIP, v))
}

// ClientIPGTE applies the GTE predicate on the "client_ip" field.
func ClientIPGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldClientIP, v))
}

// ClientIPLT applies the LT predicate on the "client_ip" field.
func ClientIPLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldClientIP, v))
}

// ClientIPLTE applies the LTE predicate on the "client_ip" field.
func ClientIPLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldClientIP, v))
}

// ClientIPContains applies the Contains predicate on the "client_ip" field.
func ClientIPContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContains(FieldClientIP, v))
}

// ClientIPHasPrefix applies the HasPrefix predicate on the "client_ip" field.
func ClientIPHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasPrefix(FieldClientIP, v))
}

// ClientIPHasSuffix applies the HasSuffix predicate on the "client_ip" field.
func ClientIPHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasSuffix(FieldClientIP, v))
}

// ClientIPIsNil applies the IsNil predicate on the "client_ip" field.
func ClientIPIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldClientIP))
}

// ClientIPNotNil applies the NotNil predicate on the "client_ip" field.
func ClientIPNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldClientIP))
}

// ClientIPEqualFold applies the EqualFold predicate on the "client_ip" field.
func ClientIPEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEqualFold(FieldClientIP, v))
}

// ClientIPContainsFold applies the ContainsFold predicate on the "client_ip" field.
func ClientIPContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContainsFold(FieldClientIP, v))
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentIsNil applies the IsNil predicate on the "user_agent" field.
func UserAgentIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldUserAgent))
}

// UserAgentNotNil applies the NotNil predicate on the "user_agent" field.
func UserAgentNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldUserAgent))
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContainsFold(FieldUserAgent, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldLocation))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContainsFold(FieldLocation, v))
}

// LoginTypeEQ applies the EQ predicate on the "login_type" field.
func LoginTypeEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEQ(FieldLoginType, v))
}

// LoginTypeNEQ applies the NEQ predicate on the "login_type" field.
func LoginTypeNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNEQ(FieldLoginType, v))
}

// LoginTypeIn applies the In predicate on the "login_type" field.
func LoginTypeIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIn(FieldLoginType, vs...))
}

// LoginTypeNotIn applies the NotIn predicate on the "login_type" field.
func LoginTypeNotIn(vs ...string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotIn(FieldLoginType, vs...))
}

// LoginTypeGT applies the GT predicate on the "login_type" field.
func LoginTypeGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGT(FieldLoginType, v))
}

// LoginTypeGTE applies the GTE predicate on the "login_type" field.
func LoginTypeGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldGTE(FieldLoginType, v))
}

// LoginTypeLT applies the LT predicate on the "login_type" field.
func LoginTypeLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLT(FieldLoginType, v))
}

// LoginTypeLTE applies the LTE predicate on the "login_type" field.
func LoginTypeLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldLTE(FieldLoginType, v))
}

// LoginTypeContains applies the Contains predicate on the "login_type" field.
func LoginTypeContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContains(FieldLoginType, v))
}

// LoginTypeHasPrefix applies the HasPrefix predicate on the "login_type" field.
func LoginTypeHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasPrefix(FieldLoginType, v))
}

// LoginTypeHasSuffix applies the HasSuffix predicate on the "login_type" field.
func LoginTypeHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldHasSuffix(FieldLoginType, v))
}

// LoginTypeIsNil applies the IsNil predicate on the "login_type" field.
func LoginTypeIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldIsNull(FieldLoginType))
}

// LoginTypeNotNil applies the NotNil predicate on the "login_type" field.
func LoginTypeNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldNotNull(FieldLoginType))
}

// LoginTypeEqualFold applies the EqualFold predicate on the "login_type" field.
func LoginTypeEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldEqualFold(FieldLoginType, v))
}

// LoginTypeContainsFold applies the ContainsFold predicate on the "login_type" field.
func LoginTypeContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(sql.FieldContainsFold(FieldLoginType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(sql.NotPredicates(p))
}
