// Code generated by ent, DO NOT EDIT.

package appgoodlabel

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldID, id))
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldEntID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldDeletedAt, v))
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldAppGoodID, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIcon, v))
}

// IconBgColor applies equality check predicate on the "icon_bg_color" field. It's identical to IconBgColorEQ.
func IconBgColor(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIconBgColor, v))
}

// LabelBgColor applies equality check predicate on the "label_bg_color" field. It's identical to LabelBgColorEQ.
func LabelBgColor(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldLabelBgColor, v))
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIndex, v))
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldEntID, v))
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldEntID, v))
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldEntID, vs...))
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldEntID, vs...))
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldEntID, v))
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldEntID, v))
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldEntID, v))
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldEntID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldDeletedAt, v))
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldAppGoodID, v))
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldAppGoodID, v))
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldAppGoodID, vs...))
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldAppGoodID, vs...))
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldAppGoodID, v))
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldAppGoodID, v))
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldAppGoodID, v))
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldAppGoodID, v))
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldAppGoodID))
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldAppGoodID))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasSuffix(FieldIcon, v))
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldIcon))
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldIcon))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContainsFold(FieldIcon, v))
}

// IconBgColorEQ applies the EQ predicate on the "icon_bg_color" field.
func IconBgColorEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIconBgColor, v))
}

// IconBgColorNEQ applies the NEQ predicate on the "icon_bg_color" field.
func IconBgColorNEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldIconBgColor, v))
}

// IconBgColorIn applies the In predicate on the "icon_bg_color" field.
func IconBgColorIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldIconBgColor, vs...))
}

// IconBgColorNotIn applies the NotIn predicate on the "icon_bg_color" field.
func IconBgColorNotIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldIconBgColor, vs...))
}

// IconBgColorGT applies the GT predicate on the "icon_bg_color" field.
func IconBgColorGT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldIconBgColor, v))
}

// IconBgColorGTE applies the GTE predicate on the "icon_bg_color" field.
func IconBgColorGTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldIconBgColor, v))
}

// IconBgColorLT applies the LT predicate on the "icon_bg_color" field.
func IconBgColorLT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldIconBgColor, v))
}

// IconBgColorLTE applies the LTE predicate on the "icon_bg_color" field.
func IconBgColorLTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldIconBgColor, v))
}

// IconBgColorContains applies the Contains predicate on the "icon_bg_color" field.
func IconBgColorContains(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContains(FieldIconBgColor, v))
}

// IconBgColorHasPrefix applies the HasPrefix predicate on the "icon_bg_color" field.
func IconBgColorHasPrefix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasPrefix(FieldIconBgColor, v))
}

// IconBgColorHasSuffix applies the HasSuffix predicate on the "icon_bg_color" field.
func IconBgColorHasSuffix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasSuffix(FieldIconBgColor, v))
}

// IconBgColorIsNil applies the IsNil predicate on the "icon_bg_color" field.
func IconBgColorIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldIconBgColor))
}

// IconBgColorNotNil applies the NotNil predicate on the "icon_bg_color" field.
func IconBgColorNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldIconBgColor))
}

// IconBgColorEqualFold applies the EqualFold predicate on the "icon_bg_color" field.
func IconBgColorEqualFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEqualFold(FieldIconBgColor, v))
}

// IconBgColorContainsFold applies the ContainsFold predicate on the "icon_bg_color" field.
func IconBgColorContainsFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContainsFold(FieldIconBgColor, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelIsNil applies the IsNil predicate on the "label" field.
func LabelIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldLabel))
}

// LabelNotNil applies the NotNil predicate on the "label" field.
func LabelNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldLabel))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContainsFold(FieldLabel, v))
}

// LabelBgColorEQ applies the EQ predicate on the "label_bg_color" field.
func LabelBgColorEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldLabelBgColor, v))
}

// LabelBgColorNEQ applies the NEQ predicate on the "label_bg_color" field.
func LabelBgColorNEQ(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldLabelBgColor, v))
}

// LabelBgColorIn applies the In predicate on the "label_bg_color" field.
func LabelBgColorIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldLabelBgColor, vs...))
}

// LabelBgColorNotIn applies the NotIn predicate on the "label_bg_color" field.
func LabelBgColorNotIn(vs ...string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldLabelBgColor, vs...))
}

// LabelBgColorGT applies the GT predicate on the "label_bg_color" field.
func LabelBgColorGT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldLabelBgColor, v))
}

// LabelBgColorGTE applies the GTE predicate on the "label_bg_color" field.
func LabelBgColorGTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldLabelBgColor, v))
}

// LabelBgColorLT applies the LT predicate on the "label_bg_color" field.
func LabelBgColorLT(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldLabelBgColor, v))
}

// LabelBgColorLTE applies the LTE predicate on the "label_bg_color" field.
func LabelBgColorLTE(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldLabelBgColor, v))
}

// LabelBgColorContains applies the Contains predicate on the "label_bg_color" field.
func LabelBgColorContains(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContains(FieldLabelBgColor, v))
}

// LabelBgColorHasPrefix applies the HasPrefix predicate on the "label_bg_color" field.
func LabelBgColorHasPrefix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasPrefix(FieldLabelBgColor, v))
}

// LabelBgColorHasSuffix applies the HasSuffix predicate on the "label_bg_color" field.
func LabelBgColorHasSuffix(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldHasSuffix(FieldLabelBgColor, v))
}

// LabelBgColorIsNil applies the IsNil predicate on the "label_bg_color" field.
func LabelBgColorIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldLabelBgColor))
}

// LabelBgColorNotNil applies the NotNil predicate on the "label_bg_color" field.
func LabelBgColorNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldLabelBgColor))
}

// LabelBgColorEqualFold applies the EqualFold predicate on the "label_bg_color" field.
func LabelBgColorEqualFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEqualFold(FieldLabelBgColor, v))
}

// LabelBgColorContainsFold applies the ContainsFold predicate on the "label_bg_color" field.
func LabelBgColorContainsFold(v string) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldContainsFold(FieldLabelBgColor, v))
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldEQ(FieldIndex, v))
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNEQ(FieldIndex, v))
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIn(FieldIndex, vs...))
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotIn(FieldIndex, vs...))
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGT(FieldIndex, v))
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldGTE(FieldIndex, v))
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLT(FieldIndex, v))
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v uint8) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldLTE(FieldIndex, v))
}

// IndexIsNil applies the IsNil predicate on the "index" field.
func IndexIsNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldIsNull(FieldIndex))
}

// IndexNotNil applies the NotNil predicate on the "index" field.
func IndexNotNil() predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.FieldNotNull(FieldIndex))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppGoodLabel) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppGoodLabel) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppGoodLabel) predicate.AppGoodLabel {
	return predicate.AppGoodLabel(sql.NotPredicates(p))
}
