// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: notif/middleware/v1/template/email/email.proto

package email

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

type EmailTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"lang_id"
	LangID string `protobuf:"bytes,30,opt,name=LangID,proto3" json:"LangID,omitempty" sql:"lang_id"`
	// @inject_tag: sql:"used_for"
	UsedForStr string     `protobuf:"bytes,40,opt,name=UsedForStr,proto3" json:"UsedForStr,omitempty" sql:"used_for"`
	UsedFor    v1.UsedFor `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
	// @inject_tag: sql:"sender"
	Sender string `protobuf:"bytes,60,opt,name=Sender,proto3" json:"Sender,omitempty" sql:"sender"`
	// @inject_tag: sql:"reply_tos"
	ReplyTosStr string   `protobuf:"bytes,69,opt,name=ReplyTosStr,proto3" json:"ReplyTosStr,omitempty" sql:"reply_tos"`
	ReplyTos    []string `protobuf:"bytes,70,rep,name=ReplyTos,proto3" json:"ReplyTos,omitempty"`
	// @inject_tag: sql:"cc_tos"
	CCTosStr string   `protobuf:"bytes,79,opt,name=CCTosStr,proto3" json:"CCTosStr,omitempty" sql:"cc_tos"`
	CCTos    []string `protobuf:"bytes,80,rep,name=CCTos,proto3" json:"CCTos,omitempty"`
	// @inject_tag: sql:"subject"
	Subject string `protobuf:"bytes,90,opt,name=Subject,proto3" json:"Subject,omitempty" sql:"subject"`
	// @inject_tag: sql:"body"
	Body string `protobuf:"bytes,100,opt,name=Body,proto3" json:"Body,omitempty" sql:"body"`
	// @inject_tag: sql:"default_to_username"
	DefaultToUsername string `protobuf:"bytes,110,opt,name=DefaultToUsername,proto3" json:"DefaultToUsername,omitempty" sql:"default_to_username"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *EmailTemplate) Reset() {
	*x = EmailTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailTemplate) ProtoMessage() {}

func (x *EmailTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailTemplate.ProtoReflect.Descriptor instead.
func (*EmailTemplate) Descriptor() ([]byte, []int) {
	return file_notif_middleware_v1_template_email_email_proto_rawDescGZIP(), []int{0}
}

func (x *EmailTemplate) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EmailTemplate) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *EmailTemplate) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *EmailTemplate) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *EmailTemplate) GetUsedForStr() string {
	if x != nil {
		return x.UsedForStr
	}
	return ""
}

func (x *EmailTemplate) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *EmailTemplate) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *EmailTemplate) GetReplyTosStr() string {
	if x != nil {
		return x.ReplyTosStr
	}
	return ""
}

func (x *EmailTemplate) GetReplyTos() []string {
	if x != nil {
		return x.ReplyTos
	}
	return nil
}

func (x *EmailTemplate) GetCCTosStr() string {
	if x != nil {
		return x.CCTosStr
	}
	return ""
}

func (x *EmailTemplate) GetCCTos() []string {
	if x != nil {
		return x.CCTos
	}
	return nil
}

func (x *EmailTemplate) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailTemplate) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailTemplate) GetDefaultToUsername() string {
	if x != nil {
		return x.DefaultToUsername
	}
	return ""
}

func (x *EmailTemplate) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *EmailTemplate) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type EmailTemplateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                *uint32     `protobuf:"varint,9,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID             *string     `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID             *string     `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	LangID            *string     `protobuf:"bytes,30,opt,name=LangID,proto3,oneof" json:"LangID,omitempty"`
	UsedFor           *v1.UsedFor `protobuf:"varint,40,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor,oneof" json:"UsedFor,omitempty"`
	Sender            *string     `protobuf:"bytes,50,opt,name=Sender,proto3,oneof" json:"Sender,omitempty"`
	ReplyTos          []string    `protobuf:"bytes,60,rep,name=ReplyTos,proto3" json:"ReplyTos,omitempty"`
	CCTos             []string    `protobuf:"bytes,70,rep,name=CCTos,proto3" json:"CCTos,omitempty"`
	Subject           *string     `protobuf:"bytes,80,opt,name=Subject,proto3,oneof" json:"Subject,omitempty"`
	Body              *string     `protobuf:"bytes,90,opt,name=Body,proto3,oneof" json:"Body,omitempty"`
	DefaultToUsername *string     `protobuf:"bytes,100,opt,name=DefaultToUsername,proto3,oneof" json:"DefaultToUsername,omitempty"`
}

func (x *EmailTemplateReq) Reset() {
	*x = EmailTemplateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailTemplateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailTemplateReq) ProtoMessage() {}

func (x *EmailTemplateReq) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailTemplateReq.ProtoReflect.Descriptor instead.
func (*EmailTemplateReq) Descriptor() ([]byte, []int) {
	return file_notif_middleware_v1_template_email_email_proto_rawDescGZIP(), []int{1}
}

func (x *EmailTemplateReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *EmailTemplateReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *EmailTemplateReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *EmailTemplateReq) GetLangID() string {
	if x != nil && x.LangID != nil {
		return *x.LangID
	}
	return ""
}

func (x *EmailTemplateReq) GetUsedFor() v1.UsedFor {
	if x != nil && x.UsedFor != nil {
		return *x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *EmailTemplateReq) GetSender() string {
	if x != nil && x.Sender != nil {
		return *x.Sender
	}
	return ""
}

func (x *EmailTemplateReq) GetReplyTos() []string {
	if x != nil {
		return x.ReplyTos
	}
	return nil
}

func (x *EmailTemplateReq) GetCCTos() []string {
	if x != nil {
		return x.CCTos
	}
	return nil
}

func (x *EmailTemplateReq) GetSubject() string {
	if x != nil && x.Subject != nil {
		return *x.Subject
	}
	return ""
}

func (x *EmailTemplateReq) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *EmailTemplateReq) GetDefaultToUsername() string {
	if x != nil && x.DefaultToUsername != nil {
		return *x.DefaultToUsername
	}
	return ""
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       *v1.Uint32Val      `protobuf:"bytes,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID    *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID    *v1.StringVal      `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	LangID   *v1.StringVal      `protobuf:"bytes,40,opt,name=LangID,proto3" json:"LangID,omitempty"`
	UsedFor  *v1.Uint32Val      `protobuf:"bytes,50,opt,name=UsedFor,proto3" json:"UsedFor,omitempty"`
	Sender   *v1.StringVal      `protobuf:"bytes,60,opt,name=Sender,proto3" json:"Sender,omitempty"`
	AppIDs   *v1.StringSliceVal `protobuf:"bytes,70,opt,name=AppIDs,proto3" json:"AppIDs,omitempty"`
	LangIDs  *v1.StringSliceVal `protobuf:"bytes,80,opt,name=LangIDs,proto3" json:"LangIDs,omitempty"`
	UsedFors *v1.Uint32SliceVal `protobuf:"bytes,90,opt,name=UsedFors,proto3" json:"UsedFors,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_notif_middleware_v1_template_email_email_proto_msgTypes[2]
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
	return file_notif_middleware_v1_template_email_email_proto_rawDescGZIP(), []int{2}
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

func (x *Conds) GetLangID() *v1.StringVal {
	if x != nil {
		return x.LangID
	}
	return nil
}

func (x *Conds) GetUsedFor() *v1.Uint32Val {
	if x != nil {
		return x.UsedFor
	}
	return nil
}

func (x *Conds) GetSender() *v1.StringVal {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Conds) GetAppIDs() *v1.StringSliceVal {
	if x != nil {
		return x.AppIDs
	}
	return nil
}

func (x *Conds) GetLangIDs() *v1.StringSliceVal {
	if x != nil {
		return x.LangIDs
	}
	return nil
}

func (x *Conds) GetUsedFors() *v1.Uint32SliceVal {
	if x != nil {
		return x.UsedFors
	}
	return nil
}

var File_notif_middleware_v1_template_email_email_proto protoreflect.FileDescriptor

var file_notif_middleware_v1_template_email_email_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03, 0x0a, 0x0d, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x12, 0x2f,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x54, 0x6f, 0x73, 0x53, 0x74, 0x72, 0x18, 0x45, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x53, 0x74, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x54, 0x6f, 0x73, 0x18, 0x46, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x54, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x43, 0x54, 0x6f, 0x73, 0x53, 0x74,
	0x72, 0x18, 0x4f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x43, 0x54, 0x6f, 0x73, 0x53, 0x74,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x43, 0x54, 0x6f, 0x73, 0x18, 0x50, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x43, 0x43, 0x54, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xd2, 0x03, 0x0a, 0x10, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x03, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x48, 0x04, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x18, 0x3c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x43, 0x54, 0x6f, 0x73, 0x18, 0x46, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x43, 0x43,
	0x54, 0x6f, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x07, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4c, 0x61,
	0x6e, 0x67, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x42, 0x6f, 0x64, 0x79,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x27, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49,
	0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x52, 0x06, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x73, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61,
	0x6c, 0x52, 0x07, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x52, 0x08, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x73, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notif_middleware_v1_template_email_email_proto_rawDescOnce sync.Once
	file_notif_middleware_v1_template_email_email_proto_rawDescData = file_notif_middleware_v1_template_email_email_proto_rawDesc
)

func file_notif_middleware_v1_template_email_email_proto_rawDescGZIP() []byte {
	file_notif_middleware_v1_template_email_email_proto_rawDescOnce.Do(func() {
		file_notif_middleware_v1_template_email_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_notif_middleware_v1_template_email_email_proto_rawDescData)
	})
	return file_notif_middleware_v1_template_email_email_proto_rawDescData
}

var (
	file_notif_middleware_v1_template_email_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_notif_middleware_v1_template_email_email_proto_goTypes  = []interface{}{
		(*EmailTemplate)(nil),     // 0: notif.middleware.template.email.v1.EmailTemplate
		(*EmailTemplateReq)(nil),  // 1: notif.middleware.template.email.v1.EmailTemplateReq
		(*Conds)(nil),             // 2: notif.middleware.template.email.v1.Conds
		(v1.UsedFor)(0),           // 3: basetypes.v1.UsedFor
		(*v1.Uint32Val)(nil),      // 4: basetypes.v1.Uint32Val
		(*v1.StringVal)(nil),      // 5: basetypes.v1.StringVal
		(*v1.StringSliceVal)(nil), // 6: basetypes.v1.StringSliceVal
		(*v1.Uint32SliceVal)(nil), // 7: basetypes.v1.Uint32SliceVal
	}
)
var file_notif_middleware_v1_template_email_email_proto_depIdxs = []int32{
	3,  // 0: notif.middleware.template.email.v1.EmailTemplate.UsedFor:type_name -> basetypes.v1.UsedFor
	3,  // 1: notif.middleware.template.email.v1.EmailTemplateReq.UsedFor:type_name -> basetypes.v1.UsedFor
	4,  // 2: notif.middleware.template.email.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	5,  // 3: notif.middleware.template.email.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	5,  // 4: notif.middleware.template.email.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	5,  // 5: notif.middleware.template.email.v1.Conds.LangID:type_name -> basetypes.v1.StringVal
	4,  // 6: notif.middleware.template.email.v1.Conds.UsedFor:type_name -> basetypes.v1.Uint32Val
	5,  // 7: notif.middleware.template.email.v1.Conds.Sender:type_name -> basetypes.v1.StringVal
	6,  // 8: notif.middleware.template.email.v1.Conds.AppIDs:type_name -> basetypes.v1.StringSliceVal
	6,  // 9: notif.middleware.template.email.v1.Conds.LangIDs:type_name -> basetypes.v1.StringSliceVal
	7,  // 10: notif.middleware.template.email.v1.Conds.UsedFors:type_name -> basetypes.v1.Uint32SliceVal
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_notif_middleware_v1_template_email_email_proto_init() }
func file_notif_middleware_v1_template_email_email_proto_init() {
	if File_notif_middleware_v1_template_email_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notif_middleware_v1_template_email_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailTemplate); i {
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
		file_notif_middleware_v1_template_email_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailTemplateReq); i {
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
		file_notif_middleware_v1_template_email_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_notif_middleware_v1_template_email_email_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notif_middleware_v1_template_email_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notif_middleware_v1_template_email_email_proto_goTypes,
		DependencyIndexes: file_notif_middleware_v1_template_email_email_proto_depIdxs,
		MessageInfos:      file_notif_middleware_v1_template_email_email_proto_msgTypes,
	}.Build()
	File_notif_middleware_v1_template_email_email_proto = out.File
	file_notif_middleware_v1_template_email_email_proto_rawDesc = nil
	file_notif_middleware_v1_template_email_email_proto_goTypes = nil
	file_notif_middleware_v1_template_email_email_proto_depIdxs = nil
}
