// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: inspire/middleware/v1/invitation/invitationcode/invitationcode.proto

package invitationcode

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InvitationCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID          *string `protobuf:"bytes,11,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID          *string `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID         *string `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	InvitationCode *string `protobuf:"bytes,40,opt,name=InvitationCode,proto3,oneof" json:"InvitationCode,omitempty"`
	Disabled       *bool   `protobuf:"varint,50,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
}

func (x *InvitationCodeReq) Reset() {
	*x = InvitationCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationCodeReq) ProtoMessage() {}

func (x *InvitationCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationCodeReq.ProtoReflect.Descriptor instead.
func (*InvitationCodeReq) Descriptor() ([]byte, []int) {
	return file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescGZIP(), []int{0}
}

func (x *InvitationCodeReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *InvitationCodeReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *InvitationCodeReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *InvitationCodeReq) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *InvitationCodeReq) GetInvitationCode() string {
	if x != nil && x.InvitationCode != nil {
		return *x.InvitationCode
	}
	return ""
}

func (x *InvitationCodeReq) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

type InvitationCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"invitation_code"
	InvitationCode string `protobuf:"bytes,40,opt,name=InvitationCode,proto3" json:"InvitationCode,omitempty" sql:"invitation_code"`
	// @inject_tag: sql:"disabled"
	Disabled bool `protobuf:"varint,50,opt,name=Disabled,proto3" json:"Disabled,omitempty" sql:"disabled"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,60,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,70,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *InvitationCode) Reset() {
	*x = InvitationCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationCode) ProtoMessage() {}

func (x *InvitationCode) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationCode.ProtoReflect.Descriptor instead.
func (*InvitationCode) Descriptor() ([]byte, []int) {
	return file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescGZIP(), []int{1}
}

func (x *InvitationCode) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *InvitationCode) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *InvitationCode) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *InvitationCode) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *InvitationCode) GetInvitationCode() string {
	if x != nil {
		return x.InvitationCode
	}
	return ""
}

func (x *InvitationCode) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *InvitationCode) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *InvitationCode) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID          *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID          *v1.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID         *v1.StringVal      `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	InvitationCode *v1.StringVal      `protobuf:"bytes,40,opt,name=InvitationCode,proto3,oneof" json:"InvitationCode,omitempty"`
	Disabled       *v1.BoolVal        `protobuf:"bytes,50,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
	UserIDs        *v1.StringSliceVal `protobuf:"bytes,60,opt,name=UserIDs,proto3,oneof" json:"UserIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[2]
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
	return file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetInvitationCode() *v1.StringVal {
	if x != nil {
		return x.InvitationCode
	}
	return nil
}

func (x *Conds) GetDisabled() *v1.BoolVal {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *Conds) GetUserIDs() *v1.StringSliceVal {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

var File_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto protoreflect.FileDescriptor

var file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDesc = []byte{
	0x0a, 0x44, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05,
	0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xab, 0x03, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x48, 0x04, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x3b, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48,
	0x05, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescOnce sync.Once
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescData = file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDesc
)

func file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescGZIP() []byte {
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescOnce.Do(func() {
		file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescData)
	})
	return file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDescData
}

var (
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_goTypes  = []interface{}{
		(*InvitationCodeReq)(nil), // 0: inspire.middleware.invitation.invitationcode.v1.InvitationCodeReq
		(*InvitationCode)(nil),    // 1: inspire.middleware.invitation.invitationcode.v1.InvitationCode
		(*Conds)(nil),             // 2: inspire.middleware.invitation.invitationcode.v1.Conds
		(*v1.StringVal)(nil),      // 3: basetypes.v1.StringVal
		(*v1.BoolVal)(nil),        // 4: basetypes.v1.BoolVal
		(*v1.StringSliceVal)(nil), // 5: basetypes.v1.StringSliceVal
	}
)
var file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_depIdxs = []int32{
	3, // 0: inspire.middleware.invitation.invitationcode.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	3, // 1: inspire.middleware.invitation.invitationcode.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	3, // 2: inspire.middleware.invitation.invitationcode.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	3, // 3: inspire.middleware.invitation.invitationcode.v1.Conds.InvitationCode:type_name -> basetypes.v1.StringVal
	4, // 4: inspire.middleware.invitation.invitationcode.v1.Conds.Disabled:type_name -> basetypes.v1.BoolVal
	5, // 5: inspire.middleware.invitation.invitationcode.v1.Conds.UserIDs:type_name -> basetypes.v1.StringSliceVal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_init() }
func file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_init() {
	if File_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationCodeReq); i {
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
		file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvitationCode); i {
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
		file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_goTypes,
		DependencyIndexes: file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_depIdxs,
		MessageInfos:      file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_msgTypes,
	}.Build()
	File_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto = out.File
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_rawDesc = nil
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_goTypes = nil
	file_inspire_middleware_v1_invitation_invitationcode_invitationcode_proto_depIdxs = nil
}
