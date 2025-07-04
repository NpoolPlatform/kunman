// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: inspire/middleware/v1/app/config/config.proto

package config

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	v11 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               *uint32              `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID            *string              `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID            *string              `protobuf:"bytes,30,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	SettleMode       *v1.SettleMode       `protobuf:"varint,40,opt,name=SettleMode,proto3,enum=basetypes.inspire.v1.SettleMode,oneof" json:"SettleMode,omitempty"`
	SettleAmountType *v1.SettleAmountType `protobuf:"varint,50,opt,name=SettleAmountType,proto3,enum=basetypes.inspire.v1.SettleAmountType,oneof" json:"SettleAmountType,omitempty"`
	SettleInterval   *v1.SettleInterval   `protobuf:"varint,60,opt,name=SettleInterval,proto3,enum=basetypes.inspire.v1.SettleInterval,oneof" json:"SettleInterval,omitempty"`
	CommissionType   *v1.CommissionType   `protobuf:"varint,70,opt,name=CommissionType,proto3,enum=basetypes.inspire.v1.CommissionType,oneof" json:"CommissionType,omitempty"`
	SettleBenefit    *bool                `protobuf:"varint,80,opt,name=SettleBenefit,proto3,oneof" json:"SettleBenefit,omitempty"`
	StartAt          *uint32              `protobuf:"varint,90,opt,name=StartAt,proto3,oneof" json:"StartAt,omitempty"`
	MaxLevel         *uint32              `protobuf:"varint,100,opt,name=MaxLevel,proto3,oneof" json:"MaxLevel,omitempty"`
}

func (x *AppConfigReq) Reset() {
	*x = AppConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfigReq) ProtoMessage() {}

func (x *AppConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfigReq.ProtoReflect.Descriptor instead.
func (*AppConfigReq) Descriptor() ([]byte, []int) {
	return file_inspire_middleware_v1_app_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfigReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *AppConfigReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *AppConfigReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *AppConfigReq) GetSettleMode() v1.SettleMode {
	if x != nil && x.SettleMode != nil {
		return *x.SettleMode
	}
	return v1.SettleMode(0)
}

func (x *AppConfigReq) GetSettleAmountType() v1.SettleAmountType {
	if x != nil && x.SettleAmountType != nil {
		return *x.SettleAmountType
	}
	return v1.SettleAmountType(0)
}

func (x *AppConfigReq) GetSettleInterval() v1.SettleInterval {
	if x != nil && x.SettleInterval != nil {
		return *x.SettleInterval
	}
	return v1.SettleInterval(0)
}

func (x *AppConfigReq) GetCommissionType() v1.CommissionType {
	if x != nil && x.CommissionType != nil {
		return *x.CommissionType
	}
	return v1.CommissionType(0)
}

func (x *AppConfigReq) GetSettleBenefit() bool {
	if x != nil && x.SettleBenefit != nil {
		return *x.SettleBenefit
	}
	return false
}

func (x *AppConfigReq) GetStartAt() uint32 {
	if x != nil && x.StartAt != nil {
		return *x.StartAt
	}
	return 0
}

func (x *AppConfigReq) GetMaxLevel() uint32 {
	if x != nil && x.MaxLevel != nil {
		return *x.MaxLevel
	}
	return 0
}

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"settle_mode"
	SettleModeStr string        `protobuf:"bytes,40,opt,name=SettleModeStr,proto3" json:"SettleModeStr,omitempty" sql:"settle_mode"`
	SettleMode    v1.SettleMode `protobuf:"varint,50,opt,name=SettleMode,proto3,enum=basetypes.inspire.v1.SettleMode" json:"SettleMode,omitempty"`
	// @inject_tag: sql:"settle_amount_type"
	SettleAmountTypeStr string              `protobuf:"bytes,60,opt,name=SettleAmountTypeStr,proto3" json:"SettleAmountTypeStr,omitempty" sql:"settle_amount_type"`
	SettleAmountType    v1.SettleAmountType `protobuf:"varint,70,opt,name=SettleAmountType,proto3,enum=basetypes.inspire.v1.SettleAmountType" json:"SettleAmountType,omitempty"`
	// @inject_tag: sql:"settle_interval"
	SettleIntervalStr string            `protobuf:"bytes,80,opt,name=SettleIntervalStr,proto3" json:"SettleIntervalStr,omitempty" sql:"settle_interval"`
	SettleInterval    v1.SettleInterval `protobuf:"varint,90,opt,name=SettleInterval,proto3,enum=basetypes.inspire.v1.SettleInterval" json:"SettleInterval,omitempty"`
	// @inject_tag: sql:"commission_type"
	CommissionTypeStr string            `protobuf:"bytes,100,opt,name=CommissionTypeStr,proto3" json:"CommissionTypeStr,omitempty" sql:"commission_type"`
	CommissionType    v1.CommissionType `protobuf:"varint,110,opt,name=CommissionType,proto3,enum=basetypes.inspire.v1.CommissionType" json:"CommissionType,omitempty"`
	// @inject_tag: sql:"settle_benefit"
	SettleBenefit bool `protobuf:"varint,120,opt,name=SettleBenefit,proto3" json:"SettleBenefit,omitempty" sql:"settle_benefit"`
	// @inject_tag: sql:"start_at"
	StartAt uint32 `protobuf:"varint,130,opt,name=StartAt,proto3" json:"StartAt,omitempty" sql:"start_at"`
	// @inject_tag: sql:"end_at"
	EndAt uint32 `protobuf:"varint,140,opt,name=EndAt,proto3" json:"EndAt,omitempty" sql:"end_at"`
	// @inject_tag: sql:"max_level"
	MaxLevel uint32 `protobuf:"varint,150,opt,name=MaxLevel,proto3" json:"MaxLevel,omitempty" sql:"max_level"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_inspire_middleware_v1_app_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *AppConfig) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AppConfig) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *AppConfig) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AppConfig) GetSettleModeStr() string {
	if x != nil {
		return x.SettleModeStr
	}
	return ""
}

func (x *AppConfig) GetSettleMode() v1.SettleMode {
	if x != nil {
		return x.SettleMode
	}
	return v1.SettleMode(0)
}

func (x *AppConfig) GetSettleAmountTypeStr() string {
	if x != nil {
		return x.SettleAmountTypeStr
	}
	return ""
}

func (x *AppConfig) GetSettleAmountType() v1.SettleAmountType {
	if x != nil {
		return x.SettleAmountType
	}
	return v1.SettleAmountType(0)
}

func (x *AppConfig) GetSettleIntervalStr() string {
	if x != nil {
		return x.SettleIntervalStr
	}
	return ""
}

func (x *AppConfig) GetSettleInterval() v1.SettleInterval {
	if x != nil {
		return x.SettleInterval
	}
	return v1.SettleInterval(0)
}

func (x *AppConfig) GetCommissionTypeStr() string {
	if x != nil {
		return x.CommissionTypeStr
	}
	return ""
}

func (x *AppConfig) GetCommissionType() v1.CommissionType {
	if x != nil {
		return x.CommissionType
	}
	return v1.CommissionType(0)
}

func (x *AppConfig) GetSettleBenefit() bool {
	if x != nil {
		return x.SettleBenefit
	}
	return false
}

func (x *AppConfig) GetStartAt() uint32 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *AppConfig) GetEndAt() uint32 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *AppConfig) GetMaxLevel() uint32 {
	if x != nil {
		return x.MaxLevel
	}
	return 0
}

func (x *AppConfig) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AppConfig) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID            *v11.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID            *v11.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	SettleMode       *v11.Uint32Val      `protobuf:"bytes,30,opt,name=SettleMode,proto3,oneof" json:"SettleMode,omitempty"`
	SettleAmountType *v11.Uint32Val      `protobuf:"bytes,40,opt,name=SettleAmountType,proto3,oneof" json:"SettleAmountType,omitempty"`
	SettleInterval   *v11.Uint32Val      `protobuf:"bytes,50,opt,name=SettleInterval,proto3,oneof" json:"SettleInterval,omitempty"`
	CommissionType   *v11.Uint32Val      `protobuf:"bytes,60,opt,name=CommissionType,proto3,oneof" json:"CommissionType,omitempty"`
	SettleBenefit    *v11.BoolVal        `protobuf:"bytes,70,opt,name=SettleBenefit,proto3,oneof" json:"SettleBenefit,omitempty"`
	StartAt          *v11.Uint32Val      `protobuf:"bytes,80,opt,name=StartAt,proto3,oneof" json:"StartAt,omitempty"`
	EndAt            *v11.Uint32Val      `protobuf:"bytes,90,opt,name=EndAt,proto3,oneof" json:"EndAt,omitempty"`
	EntIDs           *v11.StringSliceVal `protobuf:"bytes,100,opt,name=EntIDs,proto3,oneof" json:"EntIDs,omitempty"`
	ID               *v11.Uint32Val      `protobuf:"bytes,110,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	MaxLevel         *v11.Uint32Val      `protobuf:"bytes,120,opt,name=MaxLevel,proto3,oneof" json:"MaxLevel,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_app_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_inspire_middleware_v1_app_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v11.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v11.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetSettleMode() *v11.Uint32Val {
	if x != nil {
		return x.SettleMode
	}
	return nil
}

func (x *Conds) GetSettleAmountType() *v11.Uint32Val {
	if x != nil {
		return x.SettleAmountType
	}
	return nil
}

func (x *Conds) GetSettleInterval() *v11.Uint32Val {
	if x != nil {
		return x.SettleInterval
	}
	return nil
}

func (x *Conds) GetCommissionType() *v11.Uint32Val {
	if x != nil {
		return x.CommissionType
	}
	return nil
}

func (x *Conds) GetSettleBenefit() *v11.BoolVal {
	if x != nil {
		return x.SettleBenefit
	}
	return nil
}

func (x *Conds) GetStartAt() *v11.Uint32Val {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *Conds) GetEndAt() *v11.Uint32Val {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Conds) GetEntIDs() *v11.StringSliceVal {
	if x != nil {
		return x.EntIDs
	}
	return nil
}

func (x *Conds) GetID() *v11.Uint32Val {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetMaxLevel() *v11.Uint32Val {
	if x != nil {
		return x.MaxLevel
	}
	return nil
}

var File_inspire_middleware_v1_app_config_config_proto protoreflect.FileDescriptor

var file_inspire_middleware_v1_app_config_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x05,
	0x0a, 0x0c, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x13,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x48,
	0x03, 0x52, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x57, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x48, 0x04, 0x52, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x06, 0x52, 0x0e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x08, 0x48, 0x07, 0x52, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x08, 0x52, 0x07, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x4d, 0x61, 0x78,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09, 0x52, 0x08, 0x4d,
	0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49,
	0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xe0, 0x05, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x72, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x52,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x12, 0x4c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x0e,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2c,
	0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x4c, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x12, 0x19, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x05, 0x45,
	0x6e, 0x64, 0x41, 0x74, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x6e, 0x64,
	0x41, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x96,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf9, 0x06,
	0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x0a,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x48, 0x03, 0x52, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x0e, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x05,
	0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x48, 0x06, 0x52, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48,
	0x07, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a,
	0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x08, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48,
	0x09, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x48, 0x0a, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x4d, 0x61,
	0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x0b, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inspire_middleware_v1_app_config_config_proto_rawDescOnce sync.Once
	file_inspire_middleware_v1_app_config_config_proto_rawDescData = file_inspire_middleware_v1_app_config_config_proto_rawDesc
)

func file_inspire_middleware_v1_app_config_config_proto_rawDescGZIP() []byte {
	file_inspire_middleware_v1_app_config_config_proto_rawDescOnce.Do(func() {
		file_inspire_middleware_v1_app_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_inspire_middleware_v1_app_config_config_proto_rawDescData)
	})
	return file_inspire_middleware_v1_app_config_config_proto_rawDescData
}

var (
	file_inspire_middleware_v1_app_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_inspire_middleware_v1_app_config_config_proto_goTypes  = []interface{}{
		(*AppConfigReq)(nil),       // 0: inspire.middleware.app.config.v1.AppConfigReq
		(*AppConfig)(nil),          // 1: inspire.middleware.app.config.v1.AppConfig
		(*Conds)(nil),              // 2: inspire.middleware.app.config.v1.Conds
		(v1.SettleMode)(0),         // 3: basetypes.inspire.v1.SettleMode
		(v1.SettleAmountType)(0),   // 4: basetypes.inspire.v1.SettleAmountType
		(v1.SettleInterval)(0),     // 5: basetypes.inspire.v1.SettleInterval
		(v1.CommissionType)(0),     // 6: basetypes.inspire.v1.CommissionType
		(*v11.StringVal)(nil),      // 7: basetypes.v1.StringVal
		(*v11.Uint32Val)(nil),      // 8: basetypes.v1.Uint32Val
		(*v11.BoolVal)(nil),        // 9: basetypes.v1.BoolVal
		(*v11.StringSliceVal)(nil), // 10: basetypes.v1.StringSliceVal
	}
)
var file_inspire_middleware_v1_app_config_config_proto_depIdxs = []int32{
	3,  // 0: inspire.middleware.app.config.v1.AppConfigReq.SettleMode:type_name -> basetypes.inspire.v1.SettleMode
	4,  // 1: inspire.middleware.app.config.v1.AppConfigReq.SettleAmountType:type_name -> basetypes.inspire.v1.SettleAmountType
	5,  // 2: inspire.middleware.app.config.v1.AppConfigReq.SettleInterval:type_name -> basetypes.inspire.v1.SettleInterval
	6,  // 3: inspire.middleware.app.config.v1.AppConfigReq.CommissionType:type_name -> basetypes.inspire.v1.CommissionType
	3,  // 4: inspire.middleware.app.config.v1.AppConfig.SettleMode:type_name -> basetypes.inspire.v1.SettleMode
	4,  // 5: inspire.middleware.app.config.v1.AppConfig.SettleAmountType:type_name -> basetypes.inspire.v1.SettleAmountType
	5,  // 6: inspire.middleware.app.config.v1.AppConfig.SettleInterval:type_name -> basetypes.inspire.v1.SettleInterval
	6,  // 7: inspire.middleware.app.config.v1.AppConfig.CommissionType:type_name -> basetypes.inspire.v1.CommissionType
	7,  // 8: inspire.middleware.app.config.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	7,  // 9: inspire.middleware.app.config.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	8,  // 10: inspire.middleware.app.config.v1.Conds.SettleMode:type_name -> basetypes.v1.Uint32Val
	8,  // 11: inspire.middleware.app.config.v1.Conds.SettleAmountType:type_name -> basetypes.v1.Uint32Val
	8,  // 12: inspire.middleware.app.config.v1.Conds.SettleInterval:type_name -> basetypes.v1.Uint32Val
	8,  // 13: inspire.middleware.app.config.v1.Conds.CommissionType:type_name -> basetypes.v1.Uint32Val
	9,  // 14: inspire.middleware.app.config.v1.Conds.SettleBenefit:type_name -> basetypes.v1.BoolVal
	8,  // 15: inspire.middleware.app.config.v1.Conds.StartAt:type_name -> basetypes.v1.Uint32Val
	8,  // 16: inspire.middleware.app.config.v1.Conds.EndAt:type_name -> basetypes.v1.Uint32Val
	10, // 17: inspire.middleware.app.config.v1.Conds.EntIDs:type_name -> basetypes.v1.StringSliceVal
	8,  // 18: inspire.middleware.app.config.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	8,  // 19: inspire.middleware.app.config.v1.Conds.MaxLevel:type_name -> basetypes.v1.Uint32Val
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_inspire_middleware_v1_app_config_config_proto_init() }
func file_inspire_middleware_v1_app_config_config_proto_init() {
	if File_inspire_middleware_v1_app_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inspire_middleware_v1_app_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfigReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inspire_middleware_v1_app_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_inspire_middleware_v1_app_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_inspire_middleware_v1_app_config_config_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_inspire_middleware_v1_app_config_config_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inspire_middleware_v1_app_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inspire_middleware_v1_app_config_config_proto_goTypes,
		DependencyIndexes: file_inspire_middleware_v1_app_config_config_proto_depIdxs,
		MessageInfos:      file_inspire_middleware_v1_app_config_config_proto_msgTypes,
	}.Build()
	File_inspire_middleware_v1_app_config_config_proto = out.File
	file_inspire_middleware_v1_app_config_config_proto_rawDesc = nil
	file_inspire_middleware_v1_app_config_config_proto_goTypes = nil
	file_inspire_middleware_v1_app_config_config_proto_depIdxs = nil
}
