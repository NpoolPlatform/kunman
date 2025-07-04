// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: account/gateway/v1/goodbenefit/goodbenefit.proto

package goodbenefit

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	v11 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32              `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID      string              `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	GoodID     string              `protobuf:"bytes,20,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	GoodName   string              `protobuf:"bytes,30,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	GoodType   v1.GoodType         `protobuf:"varint,40,opt,name=GoodType,proto3,enum=basetypes.good.v1.GoodType" json:"GoodType,omitempty"`
	CoinTypeID string              `protobuf:"bytes,50,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName   string              `protobuf:"bytes,60,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinUnit   string              `protobuf:"bytes,70,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CoinEnv    string              `protobuf:"bytes,80,opt,name=CoinEnv,proto3" json:"CoinEnv,omitempty"`
	CoinLogo   string              `protobuf:"bytes,90,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	AccountID  string              `protobuf:"bytes,100,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Backup     bool                `protobuf:"varint,110,opt,name=Backup,proto3" json:"Backup,omitempty"`
	Address    string              `protobuf:"bytes,120,opt,name=Address,proto3" json:"Address,omitempty"`
	Active     bool                `protobuf:"varint,130,opt,name=Active,proto3" json:"Active,omitempty"`
	Locked     bool                `protobuf:"varint,140,opt,name=Locked,proto3" json:"Locked,omitempty"`
	LockedBy   v11.AccountLockedBy `protobuf:"varint,150,opt,name=LockedBy,proto3,enum=basetypes.v1.AccountLockedBy" json:"LockedBy,omitempty"`
	Blocked    bool                `protobuf:"varint,160,opt,name=Blocked,proto3" json:"Blocked,omitempty"`
	CreatedAt  uint32              `protobuf:"varint,170,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  uint32              `protobuf:"varint,180,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Account) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Account) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Account) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Account) GetGoodType() v1.GoodType {
	if x != nil {
		return x.GoodType
	}
	return v1.GoodType(0)
}

func (x *Account) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Account) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Account) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Account) GetCoinEnv() string {
	if x != nil {
		return x.CoinEnv
	}
	return ""
}

func (x *Account) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Account) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Account) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Account) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Account) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Account) GetLockedBy() v11.AccountLockedBy {
	if x != nil {
		return x.LockedBy
	}
	return v11.AccountLockedBy(0)
}

func (x *Account) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

func (x *Account) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Account) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodID     string `protobuf:"bytes,10,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	CoinTypeID string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *CreateAccountRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Account `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountResponse) GetInfo() *Account {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID   string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty"`
	Backup  *bool  `protobuf:"varint,20,opt,name=Backup,proto3,oneof" json:"Backup,omitempty"`
	Active  *bool  `protobuf:"varint,30,opt,name=Active,proto3,oneof" json:"Active,omitempty"`
	Blocked *bool  `protobuf:"varint,40,opt,name=Blocked,proto3,oneof" json:"Blocked,omitempty"`
	Locked  *bool  `protobuf:"varint,50,opt,name=Locked,proto3,oneof" json:"Locked,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAccountRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateAccountRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UpdateAccountRequest) GetBackup() bool {
	if x != nil && x.Backup != nil {
		return *x.Backup
	}
	return false
}

func (x *UpdateAccountRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateAccountRequest) GetBlocked() bool {
	if x != nil && x.Blocked != nil {
		return *x.Blocked
	}
	return false
}

func (x *UpdateAccountRequest) GetLocked() bool {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return false
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Account `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAccountResponse) GetInfo() *Account {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAccountsRequest) Reset() {
	*x = GetAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsRequest) ProtoMessage() {}

func (x *GetAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountsRequest) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccountsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAccountsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Account `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAccountsResponse) Reset() {
	*x = GetAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsResponse) ProtoMessage() {}

func (x *GetAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountsResponse) Descriptor() ([]byte, []int) {
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccountsResponse) GetInfos() []*Account {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAccountsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_account_gateway_v1_goodbenefit_goodbenefit_proto protoreflect.FileDescriptor

var file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x62, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x04, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x69, 0x6e, 0x45, 0x6e, 0x76, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69,
	0x6e, 0x45, 0x6e, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f,
	0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x17, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x06, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x18, 0x96,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x42, 0x79, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x12, 0x19,
	0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xdf, 0x01,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22,
	0x54, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6a, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xfc, 0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0xa6, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xa6, 0x01, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescOnce sync.Once
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescData = file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDesc
)

func file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescGZIP() []byte {
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescOnce.Do(func() {
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescData)
	})
	return file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDescData
}

var (
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_goTypes  = []interface{}{
		(*Account)(nil),               // 0: account.gateway.goodbenefit.v1.Account
		(*CreateAccountRequest)(nil),  // 1: account.gateway.goodbenefit.v1.CreateAccountRequest
		(*CreateAccountResponse)(nil), // 2: account.gateway.goodbenefit.v1.CreateAccountResponse
		(*UpdateAccountRequest)(nil),  // 3: account.gateway.goodbenefit.v1.UpdateAccountRequest
		(*UpdateAccountResponse)(nil), // 4: account.gateway.goodbenefit.v1.UpdateAccountResponse
		(*GetAccountsRequest)(nil),    // 5: account.gateway.goodbenefit.v1.GetAccountsRequest
		(*GetAccountsResponse)(nil),   // 6: account.gateway.goodbenefit.v1.GetAccountsResponse
		(v1.GoodType)(0),              // 7: basetypes.good.v1.GoodType
		(v11.AccountLockedBy)(0),      // 8: basetypes.v1.AccountLockedBy
	}
)
var file_account_gateway_v1_goodbenefit_goodbenefit_proto_depIdxs = []int32{
	7, // 0: account.gateway.goodbenefit.v1.Account.GoodType:type_name -> basetypes.good.v1.GoodType
	8, // 1: account.gateway.goodbenefit.v1.Account.LockedBy:type_name -> basetypes.v1.AccountLockedBy
	0, // 2: account.gateway.goodbenefit.v1.CreateAccountResponse.Info:type_name -> account.gateway.goodbenefit.v1.Account
	0, // 3: account.gateway.goodbenefit.v1.UpdateAccountResponse.Info:type_name -> account.gateway.goodbenefit.v1.Account
	0, // 4: account.gateway.goodbenefit.v1.GetAccountsResponse.Infos:type_name -> account.gateway.goodbenefit.v1.Account
	1, // 5: account.gateway.goodbenefit.v1.Gateway.CreateAccount:input_type -> account.gateway.goodbenefit.v1.CreateAccountRequest
	3, // 6: account.gateway.goodbenefit.v1.Gateway.UpdateAccount:input_type -> account.gateway.goodbenefit.v1.UpdateAccountRequest
	5, // 7: account.gateway.goodbenefit.v1.Gateway.GetAccounts:input_type -> account.gateway.goodbenefit.v1.GetAccountsRequest
	2, // 8: account.gateway.goodbenefit.v1.Gateway.CreateAccount:output_type -> account.gateway.goodbenefit.v1.CreateAccountResponse
	4, // 9: account.gateway.goodbenefit.v1.Gateway.UpdateAccount:output_type -> account.gateway.goodbenefit.v1.UpdateAccountResponse
	6, // 10: account.gateway.goodbenefit.v1.Gateway.GetAccounts:output_type -> account.gateway.goodbenefit.v1.GetAccountsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_account_gateway_v1_goodbenefit_goodbenefit_proto_init() }
func file_account_gateway_v1_goodbenefit_goodbenefit_proto_init() {
	if File_account_gateway_v1_goodbenefit_goodbenefit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountResponse); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsRequest); i {
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
		file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsResponse); i {
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
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_gateway_v1_goodbenefit_goodbenefit_proto_goTypes,
		DependencyIndexes: file_account_gateway_v1_goodbenefit_goodbenefit_proto_depIdxs,
		MessageInfos:      file_account_gateway_v1_goodbenefit_goodbenefit_proto_msgTypes,
	}.Build()
	File_account_gateway_v1_goodbenefit_goodbenefit_proto = out.File
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_rawDesc = nil
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_goTypes = nil
	file_account_gateway_v1_goodbenefit_goodbenefit_proto_depIdxs = nil
}
