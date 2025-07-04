// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: g11n/gateway/v1/country/country.proto

package country

import (
	reflect "reflect"
	sync "sync"

	country "github.com/NpoolPlatform/kunman/message/g11n/middleware/v1/country"
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

type CreateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID   *string `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Country string  `protobuf:"bytes,20,opt,name=Country,proto3" json:"Country,omitempty"`
	Flag    string  `protobuf:"bytes,30,opt,name=Flag,proto3" json:"Flag,omitempty"`
	Code    string  `protobuf:"bytes,40,opt,name=Code,proto3" json:"Code,omitempty"`
	Short   string  `protobuf:"bytes,50,opt,name=Short,proto3" json:"Short,omitempty"`
}

func (x *CreateCountryRequest) Reset() {
	*x = CreateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryRequest) ProtoMessage() {}

func (x *CreateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryRequest.ProtoReflect.Descriptor instead.
func (*CreateCountryRequest) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCountryRequest) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *CreateCountryRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateCountryRequest) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *CreateCountryRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateCountryRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type CreateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *country.Country `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCountryResponse) Reset() {
	*x = CreateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryResponse) ProtoMessage() {}

func (x *CreateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryResponse.ProtoReflect.Descriptor instead.
func (*CreateCountryResponse) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCountryResponse) GetInfo() *country.Country {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*country.CountryReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateCountriesRequest) Reset() {
	*x = CreateCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountriesRequest) ProtoMessage() {}

func (x *CreateCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountriesRequest.ProtoReflect.Descriptor instead.
func (*CreateCountriesRequest) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCountriesRequest) GetInfos() []*country.CountryReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*country.Country `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateCountriesResponse) Reset() {
	*x = CreateCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountriesResponse) ProtoMessage() {}

func (x *CreateCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountriesResponse.ProtoReflect.Descriptor instead.
func (*CreateCountriesResponse) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCountriesResponse) GetInfos() []*country.Country {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Country *string `protobuf:"bytes,20,opt,name=Country,proto3,oneof" json:"Country,omitempty"`
	Flag    *string `protobuf:"bytes,30,opt,name=Flag,proto3,oneof" json:"Flag,omitempty"`
	Code    *string `protobuf:"bytes,40,opt,name=Code,proto3,oneof" json:"Code,omitempty"`
	Short   *string `protobuf:"bytes,50,opt,name=Short,proto3,oneof" json:"Short,omitempty"`
}

func (x *UpdateCountryRequest) Reset() {
	*x = UpdateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryRequest) ProtoMessage() {}

func (x *UpdateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCountryRequest) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCountryRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCountryRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *UpdateCountryRequest) GetFlag() string {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return ""
}

func (x *UpdateCountryRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *UpdateCountryRequest) GetShort() string {
	if x != nil && x.Short != nil {
		return *x.Short
	}
	return ""
}

type UpdateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *country.Country `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateCountryResponse) Reset() {
	*x = UpdateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryResponse) ProtoMessage() {}

func (x *UpdateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryResponse.ProtoReflect.Descriptor instead.
func (*UpdateCountryResponse) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCountryResponse) GetInfo() *country.Country {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCountriesRequest) Reset() {
	*x = GetCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesRequest) ProtoMessage() {}

func (x *GetCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesRequest.ProtoReflect.Descriptor instead.
func (*GetCountriesRequest) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{6}
}

func (x *GetCountriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCountriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*country.Country `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCountriesResponse) Reset() {
	*x = GetCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesResponse) ProtoMessage() {}

func (x *GetCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_g11n_gateway_v1_country_country_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesResponse.ProtoReflect.Descriptor instead.
func (*GetCountriesResponse) Descriptor() ([]byte, []int) {
	return file_g11n_gateway_v1_country_country_proto_rawDescGZIP(), []int{7}
}

func (x *GetCountriesResponse) GetInfos() []*country.Country {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCountriesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_g11n_gateway_v1_country_country_proto protoreflect.FileDescriptor

var file_g11n_gateway_v1_country_country_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28,
	0x67, 0x31, 0x31, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x50,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x56, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x31, 0x31, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x54, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0xba,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x46, 0x6c, 0x61, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x50, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x43, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x67, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x31, 0x31, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xcd, 0x04, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x8d, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x67, 0x31, 0x31, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x95, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x31,
	0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x8d, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x2d, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x89, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x2c, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_g11n_gateway_v1_country_country_proto_rawDescOnce sync.Once
	file_g11n_gateway_v1_country_country_proto_rawDescData = file_g11n_gateway_v1_country_country_proto_rawDesc
)

func file_g11n_gateway_v1_country_country_proto_rawDescGZIP() []byte {
	file_g11n_gateway_v1_country_country_proto_rawDescOnce.Do(func() {
		file_g11n_gateway_v1_country_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_g11n_gateway_v1_country_country_proto_rawDescData)
	})
	return file_g11n_gateway_v1_country_country_proto_rawDescData
}

var (
	file_g11n_gateway_v1_country_country_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
	file_g11n_gateway_v1_country_country_proto_goTypes  = []interface{}{
		(*CreateCountryRequest)(nil),    // 0: g11n.gateway.country.v1.CreateCountryRequest
		(*CreateCountryResponse)(nil),   // 1: g11n.gateway.country.v1.CreateCountryResponse
		(*CreateCountriesRequest)(nil),  // 2: g11n.gateway.country.v1.CreateCountriesRequest
		(*CreateCountriesResponse)(nil), // 3: g11n.gateway.country.v1.CreateCountriesResponse
		(*UpdateCountryRequest)(nil),    // 4: g11n.gateway.country.v1.UpdateCountryRequest
		(*UpdateCountryResponse)(nil),   // 5: g11n.gateway.country.v1.UpdateCountryResponse
		(*GetCountriesRequest)(nil),     // 6: g11n.gateway.country.v1.GetCountriesRequest
		(*GetCountriesResponse)(nil),    // 7: g11n.gateway.country.v1.GetCountriesResponse
		(*country.Country)(nil),         // 8: g11n.middleware.country.v1.Country
		(*country.CountryReq)(nil),      // 9: g11n.middleware.country.v1.CountryReq
	}
)
var file_g11n_gateway_v1_country_country_proto_depIdxs = []int32{
	8, // 0: g11n.gateway.country.v1.CreateCountryResponse.Info:type_name -> g11n.middleware.country.v1.Country
	9, // 1: g11n.gateway.country.v1.CreateCountriesRequest.Infos:type_name -> g11n.middleware.country.v1.CountryReq
	8, // 2: g11n.gateway.country.v1.CreateCountriesResponse.Infos:type_name -> g11n.middleware.country.v1.Country
	8, // 3: g11n.gateway.country.v1.UpdateCountryResponse.Info:type_name -> g11n.middleware.country.v1.Country
	8, // 4: g11n.gateway.country.v1.GetCountriesResponse.Infos:type_name -> g11n.middleware.country.v1.Country
	0, // 5: g11n.gateway.country.v1.Gateway.CreateCountry:input_type -> g11n.gateway.country.v1.CreateCountryRequest
	2, // 6: g11n.gateway.country.v1.Gateway.CreateCountries:input_type -> g11n.gateway.country.v1.CreateCountriesRequest
	4, // 7: g11n.gateway.country.v1.Gateway.UpdateCountry:input_type -> g11n.gateway.country.v1.UpdateCountryRequest
	6, // 8: g11n.gateway.country.v1.Gateway.GetCountries:input_type -> g11n.gateway.country.v1.GetCountriesRequest
	1, // 9: g11n.gateway.country.v1.Gateway.CreateCountry:output_type -> g11n.gateway.country.v1.CreateCountryResponse
	3, // 10: g11n.gateway.country.v1.Gateway.CreateCountries:output_type -> g11n.gateway.country.v1.CreateCountriesResponse
	5, // 11: g11n.gateway.country.v1.Gateway.UpdateCountry:output_type -> g11n.gateway.country.v1.UpdateCountryResponse
	7, // 12: g11n.gateway.country.v1.Gateway.GetCountries:output_type -> g11n.gateway.country.v1.GetCountriesResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_g11n_gateway_v1_country_country_proto_init() }
func file_g11n_gateway_v1_country_country_proto_init() {
	if File_g11n_gateway_v1_country_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_g11n_gateway_v1_country_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryRequest); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryResponse); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountriesRequest); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountriesResponse); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryRequest); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryResponse); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesRequest); i {
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
		file_g11n_gateway_v1_country_country_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesResponse); i {
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
	file_g11n_gateway_v1_country_country_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_g11n_gateway_v1_country_country_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_g11n_gateway_v1_country_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_g11n_gateway_v1_country_country_proto_goTypes,
		DependencyIndexes: file_g11n_gateway_v1_country_country_proto_depIdxs,
		MessageInfos:      file_g11n_gateway_v1_country_country_proto_msgTypes,
	}.Build()
	File_g11n_gateway_v1_country_country_proto = out.File
	file_g11n_gateway_v1_country_country_proto_rawDesc = nil
	file_g11n_gateway_v1_country_country_proto_goTypes = nil
	file_g11n_gateway_v1_country_country_proto_depIdxs = nil
}
