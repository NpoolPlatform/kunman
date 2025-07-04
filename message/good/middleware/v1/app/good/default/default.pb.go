// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: good/middleware/v1/app/good/default/default.proto

package _default

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
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

type DefaultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID      *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppGoodID  *string `protobuf:"bytes,30,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	CoinTypeID *string `protobuf:"bytes,40,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
}

func (x *DefaultReq) Reset() {
	*x = DefaultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultReq) ProtoMessage() {}

func (x *DefaultReq) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultReq.ProtoReflect.Descriptor instead.
func (*DefaultReq) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_app_good_default_default_proto_rawDescGZIP(), []int{0}
}

func (x *DefaultReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *DefaultReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *DefaultReq) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

func (x *DefaultReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

type Default struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,40,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"good_name"
	GoodName string `protobuf:"bytes,50,opt,name=GoodName,proto3" json:"GoodName,omitempty" sql:"good_name"`
	// @inject_tag: sql:"app_good_id"
	AppGoodID string `protobuf:"bytes,60,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty" sql:"app_good_id"`
	// @inject_tag: sql:"app_good_name"
	AppGoodName string `protobuf:"bytes,70,opt,name=AppGoodName,proto3" json:"AppGoodName,omitempty" sql:"app_good_name"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,80,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"good_type"
	GoodTypeStr string      `protobuf:"bytes,90,opt,name=GoodTypeStr,proto3" json:"GoodTypeStr,omitempty" sql:"good_type"`
	GoodType    v1.GoodType `protobuf:"varint,100,opt,name=GoodType,proto3,enum=basetypes.good.v1.GoodType" json:"GoodType,omitempty"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Default) Reset() {
	*x = Default{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Default) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Default) ProtoMessage() {}

func (x *Default) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Default.ProtoReflect.Descriptor instead.
func (*Default) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_app_good_default_default_proto_rawDescGZIP(), []int{1}
}

func (x *Default) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Default) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Default) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Default) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Default) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Default) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Default) GetAppGoodName() string {
	if x != nil {
		return x.AppGoodName
	}
	return ""
}

func (x *Default) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Default) GetGoodTypeStr() string {
	if x != nil {
		return x.GoodTypeStr
	}
	return ""
}

func (x *Default) GetGoodType() v1.GoodType {
	if x != nil {
		return x.GoodType
	}
	return v1.GoodType(0)
}

func (x *Default) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Default) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          *v11.Uint32Val      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID       *v11.StringVal      `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID       *v11.StringVal      `protobuf:"bytes,30,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	GoodID      *v11.StringVal      `protobuf:"bytes,40,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	AppGoodID   *v11.StringVal      `protobuf:"bytes,50,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	CoinTypeID  *v11.StringVal      `protobuf:"bytes,60,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	GoodIDs     *v11.StringSliceVal `protobuf:"bytes,70,opt,name=GoodIDs,proto3,oneof" json:"GoodIDs,omitempty"`
	CoinTypeIDs *v11.StringSliceVal `protobuf:"bytes,80,opt,name=CoinTypeIDs,proto3,oneof" json:"CoinTypeIDs,omitempty"`
	GoodType    *v11.Uint32Val      `protobuf:"bytes,90,opt,name=GoodType,proto3,oneof" json:"GoodType,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_app_good_default_default_proto_msgTypes[2]
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
	return file_good_middleware_v1_app_good_default_default_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v11.Uint32Val {
	if x != nil {
		return x.ID
	}
	return nil
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

func (x *Conds) GetGoodID() *v11.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetAppGoodID() *v11.StringVal {
	if x != nil {
		return x.AppGoodID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v11.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetGoodIDs() *v11.StringSliceVal {
	if x != nil {
		return x.GoodIDs
	}
	return nil
}

func (x *Conds) GetCoinTypeIDs() *v11.StringSliceVal {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

func (x *Conds) GetGoodType() *v11.Uint32Val {
	if x != nil {
		return x.GoodType
	}
	return nil
}

var File_good_middleware_v1_app_good_default_default_proto protoreflect.FileDescriptor

var file_good_middleware_v1_app_good_default_default_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x31, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49,
	0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0xf2, 0x02, 0x0a, 0x07, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf5, 0x04,
	0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x04, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a,
	0x07, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x06, 0x52, 0x07,
	0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x07, 0x52,
	0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x38, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x08, 0x52, 0x08, 0x47, 0x6f,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x47, 0x6f, 0x6f,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_good_middleware_v1_app_good_default_default_proto_rawDescOnce sync.Once
	file_good_middleware_v1_app_good_default_default_proto_rawDescData = file_good_middleware_v1_app_good_default_default_proto_rawDesc
)

func file_good_middleware_v1_app_good_default_default_proto_rawDescGZIP() []byte {
	file_good_middleware_v1_app_good_default_default_proto_rawDescOnce.Do(func() {
		file_good_middleware_v1_app_good_default_default_proto_rawDescData = protoimpl.X.CompressGZIP(file_good_middleware_v1_app_good_default_default_proto_rawDescData)
	})
	return file_good_middleware_v1_app_good_default_default_proto_rawDescData
}

var (
	file_good_middleware_v1_app_good_default_default_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_good_middleware_v1_app_good_default_default_proto_goTypes  = []interface{}{
		(*DefaultReq)(nil),         // 0: good.middleware.app.good1.default1.v1.DefaultReq
		(*Default)(nil),            // 1: good.middleware.app.good1.default1.v1.Default
		(*Conds)(nil),              // 2: good.middleware.app.good1.default1.v1.Conds
		(v1.GoodType)(0),           // 3: basetypes.good.v1.GoodType
		(*v11.Uint32Val)(nil),      // 4: basetypes.v1.Uint32Val
		(*v11.StringVal)(nil),      // 5: basetypes.v1.StringVal
		(*v11.StringSliceVal)(nil), // 6: basetypes.v1.StringSliceVal
	}
)
var file_good_middleware_v1_app_good_default_default_proto_depIdxs = []int32{
	3,  // 0: good.middleware.app.good1.default1.v1.Default.GoodType:type_name -> basetypes.good.v1.GoodType
	4,  // 1: good.middleware.app.good1.default1.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	5,  // 2: good.middleware.app.good1.default1.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	5,  // 3: good.middleware.app.good1.default1.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	5,  // 4: good.middleware.app.good1.default1.v1.Conds.GoodID:type_name -> basetypes.v1.StringVal
	5,  // 5: good.middleware.app.good1.default1.v1.Conds.AppGoodID:type_name -> basetypes.v1.StringVal
	5,  // 6: good.middleware.app.good1.default1.v1.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	6,  // 7: good.middleware.app.good1.default1.v1.Conds.GoodIDs:type_name -> basetypes.v1.StringSliceVal
	6,  // 8: good.middleware.app.good1.default1.v1.Conds.CoinTypeIDs:type_name -> basetypes.v1.StringSliceVal
	4,  // 9: good.middleware.app.good1.default1.v1.Conds.GoodType:type_name -> basetypes.v1.Uint32Val
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_good_middleware_v1_app_good_default_default_proto_init() }
func file_good_middleware_v1_app_good_default_default_proto_init() {
	if File_good_middleware_v1_app_good_default_default_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_good_middleware_v1_app_good_default_default_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultReq); i {
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
		file_good_middleware_v1_app_good_default_default_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Default); i {
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
		file_good_middleware_v1_app_good_default_default_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_good_middleware_v1_app_good_default_default_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_good_middleware_v1_app_good_default_default_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_good_middleware_v1_app_good_default_default_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_good_middleware_v1_app_good_default_default_proto_goTypes,
		DependencyIndexes: file_good_middleware_v1_app_good_default_default_proto_depIdxs,
		MessageInfos:      file_good_middleware_v1_app_good_default_default_proto_msgTypes,
	}.Build()
	File_good_middleware_v1_app_good_default_default_proto = out.File
	file_good_middleware_v1_app_good_default_default_proto_rawDesc = nil
	file_good_middleware_v1_app_good_default_default_proto_goTypes = nil
	file_good_middleware_v1_app_good_default_default_proto_depIdxs = nil
}
