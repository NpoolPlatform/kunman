// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: notif/middleware/v1/contact/contact.proto

package contact

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContactReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          *uint32        `protobuf:"varint,9,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID       *string        `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID       *string        `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	Account     *string        `protobuf:"bytes,30,opt,name=Account,proto3,oneof" json:"Account,omitempty"`
	AccountType *v1.SignMethod `protobuf:"varint,40,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod,oneof" json:"AccountType,omitempty"`
	UsedFor     *v1.UsedFor    `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor,oneof" json:"UsedFor,omitempty"`
	Sender      *string        `protobuf:"bytes,60,opt,name=Sender,proto3,oneof" json:"Sender,omitempty"`
}

func (x *ContactReq) Reset() {
	*x = ContactReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactReq) ProtoMessage() {}

func (x *ContactReq) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactReq.ProtoReflect.Descriptor instead.
func (*ContactReq) Descriptor() ([]byte, []int) {
	return file_notif_middleware_v1_contact_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *ContactReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *ContactReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *ContactReq) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *ContactReq) GetAccountType() v1.SignMethod {
	if x != nil && x.AccountType != nil {
		return *x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *ContactReq) GetUsedFor() v1.UsedFor {
	if x != nil && x.UsedFor != nil {
		return *x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *ContactReq) GetSender() string {
	if x != nil && x.Sender != nil {
		return *x.Sender
	}
	return ""
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"account"
	Account string `protobuf:"bytes,30,opt,name=Account,proto3" json:"Account,omitempty" sql:"account"`
	// @inject_tag: sql:"account_type"
	AccountTypeStr string        `protobuf:"bytes,39,opt,name=AccountTypeStr,proto3" json:"AccountTypeStr,omitempty" sql:"account_type"`
	AccountType    v1.SignMethod `protobuf:"varint,40,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	// @inject_tag: sql:"used_for"
	UsedForStr string     `protobuf:"bytes,49,opt,name=UsedForStr,proto3" json:"UsedForStr,omitempty" sql:"used_for"`
	UsedFor    v1.UsedFor `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
	// @inject_tag: sql:"sender"
	Sender string `protobuf:"bytes,60,opt,name=Sender,proto3" json:"Sender,omitempty" sql:"sender"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_notif_middleware_v1_contact_contact_proto_rawDescGZIP(), []int{1}
}

func (x *Contact) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Contact) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Contact) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Contact) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Contact) GetAccountTypeStr() string {
	if x != nil {
		return x.AccountTypeStr
	}
	return ""
}

func (x *Contact) GetAccountType() v1.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *Contact) GetUsedForStr() string {
	if x != nil {
		return x.UsedForStr
	}
	return ""
}

func (x *Contact) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *Contact) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Contact) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Contact) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          *v1.Uint32Val `protobuf:"bytes,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID       *v1.StringVal `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID       *v1.StringVal `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AccountType *v1.Uint32Val `protobuf:"bytes,40,opt,name=AccountType,proto3" json:"AccountType,omitempty"`
	UsedFor     *v1.Uint32Val `protobuf:"bytes,50,opt,name=UsedFor,proto3" json:"UsedFor,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[2]
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
	return file_notif_middleware_v1_contact_contact_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.Uint32Val {
	if x != nil {
		return x.ID
	}
	return nil
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

func (x *Conds) GetAccountType() *v1.Uint32Val {
	if x != nil {
		return x.AccountType
	}
	return nil
}

func (x *Conds) GetUsedFor() *v1.Uint32Val {
	if x != nil {
		return x.UsedFor
	}
	return nil
}

type TextInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string   `protobuf:"bytes,10,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Content  string   `protobuf:"bytes,20,opt,name=Content,proto3" json:"Content,omitempty"`
	From     string   `protobuf:"bytes,30,opt,name=From,proto3" json:"From,omitempty"`
	To       string   `protobuf:"bytes,40,opt,name=To,proto3" json:"To,omitempty"`
	ToCCs    []string `protobuf:"bytes,50,rep,name=ToCCs,proto3" json:"ToCCs,omitempty"`
	ReplyTos []string `protobuf:"bytes,60,rep,name=ReplyTos,proto3" json:"ReplyTos,omitempty"`
}

func (x *TextInfo) Reset() {
	*x = TextInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextInfo) ProtoMessage() {}

func (x *TextInfo) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_contact_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextInfo.ProtoReflect.Descriptor instead.
func (*TextInfo) Descriptor() ([]byte, []int) {
	return file_notif_middleware_v1_contact_contact_proto_rawDescGZIP(), []int{3}
}

func (x *TextInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *TextInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextInfo) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TextInfo) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TextInfo) GetToCCs() []string {
	if x != nil {
		return x.ToCCs
	}
	return nil
}

func (x *TextInfo) GetReplyTos() []string {
	if x != nil {
		return x.ReplyTos
	}
	return nil
}

var File_notif_middleware_v1_contact_contact_proto protoreflect.FileDescriptor

var file_notif_middleware_v1_contact_contact_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x48, 0x04, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x48, 0x05, 0x52,
	0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xe8, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x18, 0x31, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x12, 0x2f, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xfc, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x39, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x22, 0x94, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x43, 0x43, 0x73, 0x18, 0x32,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x43, 0x43, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notif_middleware_v1_contact_contact_proto_rawDescOnce sync.Once
	file_notif_middleware_v1_contact_contact_proto_rawDescData = file_notif_middleware_v1_contact_contact_proto_rawDesc
)

func file_notif_middleware_v1_contact_contact_proto_rawDescGZIP() []byte {
	file_notif_middleware_v1_contact_contact_proto_rawDescOnce.Do(func() {
		file_notif_middleware_v1_contact_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_notif_middleware_v1_contact_contact_proto_rawDescData)
	})
	return file_notif_middleware_v1_contact_contact_proto_rawDescData
}

var (
	file_notif_middleware_v1_contact_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_notif_middleware_v1_contact_contact_proto_goTypes  = []interface{}{
		(*ContactReq)(nil),   // 0: notif.middleware.contact.v1.ContactReq
		(*Contact)(nil),      // 1: notif.middleware.contact.v1.Contact
		(*Conds)(nil),        // 2: notif.middleware.contact.v1.Conds
		(*TextInfo)(nil),     // 3: notif.middleware.contact.v1.TextInfo
		(v1.SignMethod)(0),   // 4: basetypes.v1.SignMethod
		(v1.UsedFor)(0),      // 5: basetypes.v1.UsedFor
		(*v1.Uint32Val)(nil), // 6: basetypes.v1.Uint32Val
		(*v1.StringVal)(nil), // 7: basetypes.v1.StringVal
	}
)
var file_notif_middleware_v1_contact_contact_proto_depIdxs = []int32{
	4, // 0: notif.middleware.contact.v1.ContactReq.AccountType:type_name -> basetypes.v1.SignMethod
	5, // 1: notif.middleware.contact.v1.ContactReq.UsedFor:type_name -> basetypes.v1.UsedFor
	4, // 2: notif.middleware.contact.v1.Contact.AccountType:type_name -> basetypes.v1.SignMethod
	5, // 3: notif.middleware.contact.v1.Contact.UsedFor:type_name -> basetypes.v1.UsedFor
	6, // 4: notif.middleware.contact.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	7, // 5: notif.middleware.contact.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	7, // 6: notif.middleware.contact.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	6, // 7: notif.middleware.contact.v1.Conds.AccountType:type_name -> basetypes.v1.Uint32Val
	6, // 8: notif.middleware.contact.v1.Conds.UsedFor:type_name -> basetypes.v1.Uint32Val
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_notif_middleware_v1_contact_contact_proto_init() }
func file_notif_middleware_v1_contact_contact_proto_init() {
	if File_notif_middleware_v1_contact_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notif_middleware_v1_contact_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactReq); i {
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
		file_notif_middleware_v1_contact_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_notif_middleware_v1_contact_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_notif_middleware_v1_contact_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextInfo); i {
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
	file_notif_middleware_v1_contact_contact_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notif_middleware_v1_contact_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notif_middleware_v1_contact_contact_proto_goTypes,
		DependencyIndexes: file_notif_middleware_v1_contact_contact_proto_depIdxs,
		MessageInfos:      file_notif_middleware_v1_contact_contact_proto_msgTypes,
	}.Build()
	File_notif_middleware_v1_contact_contact_proto = out.File
	file_notif_middleware_v1_contact_contact_proto_rawDesc = nil
	file_notif_middleware_v1_contact_contact_proto_goTypes = nil
	file_notif_middleware_v1_contact_contact_proto_depIdxs = nil
}
