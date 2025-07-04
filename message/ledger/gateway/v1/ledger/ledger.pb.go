// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: ledger/gateway/v1/ledger/ledger.proto

package ledger

import (
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ledger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyID       string          `protobuf:"bytes,10,opt,name=CurrencyID,proto3" json:"CurrencyID,omitempty"`
	CurrencyType     v1.CurrencyType `protobuf:"varint,11,opt,name=CurrencyType,proto3,enum=basetypes.ledger.v1.CurrencyType" json:"CurrencyType,omitempty"`
	CurrencyName     string          `protobuf:"bytes,20,opt,name=CurrencyName,proto3" json:"CurrencyName,omitempty"`
	DisplayNames     []string        `protobuf:"bytes,21,rep,name=DisplayNames,proto3" json:"DisplayNames,omitempty"`
	CurrencyLogo     string          `protobuf:"bytes,30,opt,name=CurrencyLogo,proto3" json:"CurrencyLogo,omitempty"`
	CurrencyUnit     string          `protobuf:"bytes,40,opt,name=CurrencyUnit,proto3" json:"CurrencyUnit,omitempty"`
	CurrencyDisabled bool            `protobuf:"varint,41,opt,name=CurrencyDisabled,proto3" json:"CurrencyDisabled,omitempty"`
	CurrencyDisplay  bool            `protobuf:"varint,42,opt,name=CurrencyDisplay,proto3" json:"CurrencyDisplay,omitempty"`
	Incoming         string          `protobuf:"bytes,50,opt,name=Incoming,proto3" json:"Incoming,omitempty"`
	Locked           string          `protobuf:"bytes,60,opt,name=Locked,proto3" json:"Locked,omitempty"`
	Outcoming        string          `protobuf:"bytes,70,opt,name=Outcoming,proto3" json:"Outcoming,omitempty"`
	Spendable        string          `protobuf:"bytes,80,opt,name=Spendable,proto3" json:"Spendable,omitempty"`
	UserID           string          `protobuf:"bytes,90,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PhoneNO          string          `protobuf:"bytes,100,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	EmailAddress     string          `protobuf:"bytes,110,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	AppID            string          `protobuf:"bytes,120,opt,name=AppID,proto3" json:"AppID,omitempty"`
	ID               uint32          `protobuf:"varint,130,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID            string          `protobuf:"bytes,140,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *Ledger) Reset() {
	*x = Ledger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ledger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ledger) ProtoMessage() {}

func (x *Ledger) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ledger.ProtoReflect.Descriptor instead.
func (*Ledger) Descriptor() ([]byte, []int) {
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *Ledger) GetCurrencyID() string {
	if x != nil {
		return x.CurrencyID
	}
	return ""
}

func (x *Ledger) GetCurrencyType() v1.CurrencyType {
	if x != nil {
		return x.CurrencyType
	}
	return v1.CurrencyType(0)
}

func (x *Ledger) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

func (x *Ledger) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *Ledger) GetCurrencyLogo() string {
	if x != nil {
		return x.CurrencyLogo
	}
	return ""
}

func (x *Ledger) GetCurrencyUnit() string {
	if x != nil {
		return x.CurrencyUnit
	}
	return ""
}

func (x *Ledger) GetCurrencyDisabled() bool {
	if x != nil {
		return x.CurrencyDisabled
	}
	return false
}

func (x *Ledger) GetCurrencyDisplay() bool {
	if x != nil {
		return x.CurrencyDisplay
	}
	return false
}

func (x *Ledger) GetIncoming() string {
	if x != nil {
		return x.Incoming
	}
	return ""
}

func (x *Ledger) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *Ledger) GetOutcoming() string {
	if x != nil {
		return x.Outcoming
	}
	return ""
}

func (x *Ledger) GetSpendable() string {
	if x != nil {
		return x.Spendable
	}
	return ""
}

func (x *Ledger) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Ledger) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Ledger) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Ledger) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Ledger) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Ledger) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type GetLedgersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLedgersRequest) Reset() {
	*x = GetLedgersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgersRequest) ProtoMessage() {}

func (x *GetLedgersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgersRequest.ProtoReflect.Descriptor instead.
func (*GetLedgersRequest) Descriptor() ([]byte, []int) {
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *GetLedgersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetLedgersRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetLedgersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetLedgersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLedgersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Ledger `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32    `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLedgersResponse) Reset() {
	*x = GetLedgersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgersResponse) ProtoMessage() {}

func (x *GetLedgersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgersResponse.ProtoReflect.Descriptor instead.
func (*GetLedgersResponse) Descriptor() ([]byte, []int) {
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP(), []int{2}
}

func (x *GetLedgersResponse) GetInfos() []*Ledger {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetLedgersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppLedgersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppLedgersRequest) Reset() {
	*x = GetAppLedgersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppLedgersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppLedgersRequest) ProtoMessage() {}

func (x *GetAppLedgersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppLedgersRequest.ProtoReflect.Descriptor instead.
func (*GetAppLedgersRequest) Descriptor() ([]byte, []int) {
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppLedgersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppLedgersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppLedgersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppLedgersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Ledger `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32    `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppLedgersResponse) Reset() {
	*x = GetAppLedgersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppLedgersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppLedgersResponse) ProtoMessage() {}

func (x *GetAppLedgersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppLedgersResponse.ProtoReflect.Descriptor instead.
func (*GetAppLedgersResponse) Descriptor() ([]byte, []int) {
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppLedgersResponse) GetInfos() []*Ledger {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppLedgersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_ledger_gateway_v1_ledger_ledger_proto protoreflect.FileDescriptor

var file_ledger_gateway_v1_ledger_ledger_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd9, 0x04, 0x0a, 0x06, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x12, 0x45, 0x0a, 0x0c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x22,
	0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x6e,
	0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x0f, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x8c,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x6f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x62, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0xa4, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x84, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01,
	0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01,
	0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ledger_gateway_v1_ledger_ledger_proto_rawDescOnce sync.Once
	file_ledger_gateway_v1_ledger_ledger_proto_rawDescData = file_ledger_gateway_v1_ledger_ledger_proto_rawDesc
)

func file_ledger_gateway_v1_ledger_ledger_proto_rawDescGZIP() []byte {
	file_ledger_gateway_v1_ledger_ledger_proto_rawDescOnce.Do(func() {
		file_ledger_gateway_v1_ledger_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_gateway_v1_ledger_ledger_proto_rawDescData)
	})
	return file_ledger_gateway_v1_ledger_ledger_proto_rawDescData
}

var file_ledger_gateway_v1_ledger_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ledger_gateway_v1_ledger_ledger_proto_goTypes = []interface{}{
	(*Ledger)(nil),                // 0: ledger.gateway.ledger.v1.Ledger
	(*GetLedgersRequest)(nil),     // 1: ledger.gateway.ledger.v1.GetLedgersRequest
	(*GetLedgersResponse)(nil),    // 2: ledger.gateway.ledger.v1.GetLedgersResponse
	(*GetAppLedgersRequest)(nil),  // 3: ledger.gateway.ledger.v1.GetAppLedgersRequest
	(*GetAppLedgersResponse)(nil), // 4: ledger.gateway.ledger.v1.GetAppLedgersResponse
	(v1.CurrencyType)(0),          // 5: basetypes.ledger.v1.CurrencyType
}
var file_ledger_gateway_v1_ledger_ledger_proto_depIdxs = []int32{
	5, // 0: ledger.gateway.ledger.v1.Ledger.CurrencyType:type_name -> basetypes.ledger.v1.CurrencyType
	0, // 1: ledger.gateway.ledger.v1.GetLedgersResponse.Infos:type_name -> ledger.gateway.ledger.v1.Ledger
	0, // 2: ledger.gateway.ledger.v1.GetAppLedgersResponse.Infos:type_name -> ledger.gateway.ledger.v1.Ledger
	1, // 3: ledger.gateway.ledger.v1.Gateway.GetLedgers:input_type -> ledger.gateway.ledger.v1.GetLedgersRequest
	3, // 4: ledger.gateway.ledger.v1.Gateway.GetAppLedgers:input_type -> ledger.gateway.ledger.v1.GetAppLedgersRequest
	2, // 5: ledger.gateway.ledger.v1.Gateway.GetLedgers:output_type -> ledger.gateway.ledger.v1.GetLedgersResponse
	4, // 6: ledger.gateway.ledger.v1.Gateway.GetAppLedgers:output_type -> ledger.gateway.ledger.v1.GetAppLedgersResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ledger_gateway_v1_ledger_ledger_proto_init() }
func file_ledger_gateway_v1_ledger_ledger_proto_init() {
	if File_ledger_gateway_v1_ledger_ledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ledger); i {
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
		file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedgersRequest); i {
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
		file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLedgersResponse); i {
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
		file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppLedgersRequest); i {
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
		file_ledger_gateway_v1_ledger_ledger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppLedgersResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ledger_gateway_v1_ledger_ledger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ledger_gateway_v1_ledger_ledger_proto_goTypes,
		DependencyIndexes: file_ledger_gateway_v1_ledger_ledger_proto_depIdxs,
		MessageInfos:      file_ledger_gateway_v1_ledger_ledger_proto_msgTypes,
	}.Build()
	File_ledger_gateway_v1_ledger_ledger_proto = out.File
	file_ledger_gateway_v1_ledger_ledger_proto_rawDesc = nil
	file_ledger_gateway_v1_ledger_ledger_proto_goTypes = nil
	file_ledger_gateway_v1_ledger_ledger_proto_depIdxs = nil
}
