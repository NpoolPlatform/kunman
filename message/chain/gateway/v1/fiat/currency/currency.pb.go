// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: chain/gateway/v1/fiat/currency/currency.proto

package currency

import (
	reflect "reflect"
	sync "sync"

	currency "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
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

type GetCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FiatName string `protobuf:"bytes,10,opt,name=FiatName,proto3" json:"FiatName,omitempty"`
}

func (x *GetCurrencyRequest) Reset() {
	*x = GetCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyRequest) ProtoMessage() {}

func (x *GetCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyRequest.ProtoReflect.Descriptor instead.
func (*GetCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_chain_gateway_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrencyRequest) GetFiatName() string {
	if x != nil {
		return x.FiatName
	}
	return ""
}

type GetCurrencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *currency.Currency `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetCurrencyResponse) Reset() {
	*x = GetCurrencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyResponse) ProtoMessage() {}

func (x *GetCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyResponse.ProtoReflect.Descriptor instead.
func (*GetCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_chain_gateway_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrencyResponse) GetInfo() *currency.Currency {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCurrenciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FiatIDs []string `protobuf:"bytes,10,rep,name=FiatIDs,proto3" json:"FiatIDs,omitempty"`
	Offset  int32    `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32    `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCurrenciesRequest) Reset() {
	*x = GetCurrenciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesRequest) ProtoMessage() {}

func (x *GetCurrenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*GetCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_chain_gateway_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrenciesRequest) GetFiatIDs() []string {
	if x != nil {
		return x.FiatIDs
	}
	return nil
}

func (x *GetCurrenciesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCurrenciesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*currency.Currency `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_chain_gateway_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrenciesResponse) GetInfos() []*currency.Currency {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCurrenciesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_chain_gateway_v1_fiat_currency_currency_proto protoreflect.FileDescriptor

var file_chain_gateway_v1_fiat_currency_currency_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66,
	0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc5, 0x02, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x97, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x32, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x9f, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_gateway_v1_fiat_currency_currency_proto_rawDescOnce sync.Once
	file_chain_gateway_v1_fiat_currency_currency_proto_rawDescData = file_chain_gateway_v1_fiat_currency_currency_proto_rawDesc
)

func file_chain_gateway_v1_fiat_currency_currency_proto_rawDescGZIP() []byte {
	file_chain_gateway_v1_fiat_currency_currency_proto_rawDescOnce.Do(func() {
		file_chain_gateway_v1_fiat_currency_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_gateway_v1_fiat_currency_currency_proto_rawDescData)
	})
	return file_chain_gateway_v1_fiat_currency_currency_proto_rawDescData
}

var (
	file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_chain_gateway_v1_fiat_currency_currency_proto_goTypes  = []interface{}{
		(*GetCurrencyRequest)(nil),    // 0: chain.gateway.fiat.currency.v1.GetCurrencyRequest
		(*GetCurrencyResponse)(nil),   // 1: chain.gateway.fiat.currency.v1.GetCurrencyResponse
		(*GetCurrenciesRequest)(nil),  // 2: chain.gateway.fiat.currency.v1.GetCurrenciesRequest
		(*GetCurrenciesResponse)(nil), // 3: chain.gateway.fiat.currency.v1.GetCurrenciesResponse
		(*currency.Currency)(nil),     // 4: chain.middleware.fiat.currency.v1.Currency
	}
)
var file_chain_gateway_v1_fiat_currency_currency_proto_depIdxs = []int32{
	4, // 0: chain.gateway.fiat.currency.v1.GetCurrencyResponse.Info:type_name -> chain.middleware.fiat.currency.v1.Currency
	4, // 1: chain.gateway.fiat.currency.v1.GetCurrenciesResponse.Infos:type_name -> chain.middleware.fiat.currency.v1.Currency
	0, // 2: chain.gateway.fiat.currency.v1.Gateway.GetCurrency:input_type -> chain.gateway.fiat.currency.v1.GetCurrencyRequest
	2, // 3: chain.gateway.fiat.currency.v1.Gateway.GetCurrencies:input_type -> chain.gateway.fiat.currency.v1.GetCurrenciesRequest
	1, // 4: chain.gateway.fiat.currency.v1.Gateway.GetCurrency:output_type -> chain.gateway.fiat.currency.v1.GetCurrencyResponse
	3, // 5: chain.gateway.fiat.currency.v1.Gateway.GetCurrencies:output_type -> chain.gateway.fiat.currency.v1.GetCurrenciesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chain_gateway_v1_fiat_currency_currency_proto_init() }
func file_chain_gateway_v1_fiat_currency_currency_proto_init() {
	if File_chain_gateway_v1_fiat_currency_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyRequest); i {
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
		file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyResponse); i {
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
		file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesRequest); i {
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
		file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesResponse); i {
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
			RawDescriptor: file_chain_gateway_v1_fiat_currency_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_gateway_v1_fiat_currency_currency_proto_goTypes,
		DependencyIndexes: file_chain_gateway_v1_fiat_currency_currency_proto_depIdxs,
		MessageInfos:      file_chain_gateway_v1_fiat_currency_currency_proto_msgTypes,
	}.Build()
	File_chain_gateway_v1_fiat_currency_currency_proto = out.File
	file_chain_gateway_v1_fiat_currency_currency_proto_rawDesc = nil
	file_chain_gateway_v1_fiat_currency_currency_proto_goTypes = nil
	file_chain_gateway_v1_fiat_currency_currency_proto_depIdxs = nil
}
