// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: good/gateway/v1/app/good/good.proto

package good

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	good "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
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

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_good_gateway_v1_app_good_good_proto_rawDescGZIP(), []int{0}
}

func (x *GetGoodsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetGoodsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetGoodsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*good.Good `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetGoodsResponse) Reset() {
	*x = GetGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsResponse) ProtoMessage() {}

func (x *GetGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsResponse) Descriptor() ([]byte, []int) {
	return file_good_gateway_v1_app_good_good_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsResponse) GetInfos() []*good.Good {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetGoodsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetGoodsRequest) Reset() {
	*x = AdminGetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetGoodsRequest) ProtoMessage() {}

func (x *AdminGetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetGoodsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_good_gateway_v1_app_good_good_proto_rawDescGZIP(), []int{2}
}

func (x *AdminGetGoodsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminGetGoodsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetGoodsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*good.Good `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetGoodsResponse) Reset() {
	*x = AdminGetGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetGoodsResponse) ProtoMessage() {}

func (x *AdminGetGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_good_gateway_v1_app_good_good_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetGoodsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetGoodsResponse) Descriptor() ([]byte, []int) {
	return file_good_gateway_v1_app_good_good_proto_rawDescGZIP(), []int{3}
}

func (x *AdminGetGoodsResponse) GetInfos() []*good.Good {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetGoodsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_good_gateway_v1_app_good_good_proto protoreflect.FileDescriptor

var file_good_gateway_v1_app_good_good_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31,
	0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67,
	0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x62, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x6f, 0x6f, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x66, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x67, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x32, 0xa4, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x80, 0x01,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x95, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a,
	0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_good_gateway_v1_app_good_good_proto_rawDescOnce sync.Once
	file_good_gateway_v1_app_good_good_proto_rawDescData = file_good_gateway_v1_app_good_good_proto_rawDesc
)

func file_good_gateway_v1_app_good_good_proto_rawDescGZIP() []byte {
	file_good_gateway_v1_app_good_good_proto_rawDescOnce.Do(func() {
		file_good_gateway_v1_app_good_good_proto_rawDescData = protoimpl.X.CompressGZIP(file_good_gateway_v1_app_good_good_proto_rawDescData)
	})
	return file_good_gateway_v1_app_good_good_proto_rawDescData
}

var (
	file_good_gateway_v1_app_good_good_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
	file_good_gateway_v1_app_good_good_proto_goTypes  = []interface{}{
		(*GetGoodsRequest)(nil),       // 0: good.gateway.app.good1.v1.GetGoodsRequest
		(*GetGoodsResponse)(nil),      // 1: good.gateway.app.good1.v1.GetGoodsResponse
		(*AdminGetGoodsRequest)(nil),  // 2: good.gateway.app.good1.v1.AdminGetGoodsRequest
		(*AdminGetGoodsResponse)(nil), // 3: good.gateway.app.good1.v1.AdminGetGoodsResponse
		(*good.Good)(nil),             // 4: good.middleware.app.good1.v1.Good
	}
)
var file_good_gateway_v1_app_good_good_proto_depIdxs = []int32{
	4, // 0: good.gateway.app.good1.v1.GetGoodsResponse.Infos:type_name -> good.middleware.app.good1.v1.Good
	4, // 1: good.gateway.app.good1.v1.AdminGetGoodsResponse.Infos:type_name -> good.middleware.app.good1.v1.Good
	0, // 2: good.gateway.app.good1.v1.Gateway.GetGoods:input_type -> good.gateway.app.good1.v1.GetGoodsRequest
	2, // 3: good.gateway.app.good1.v1.Gateway.AdminGetGoods:input_type -> good.gateway.app.good1.v1.AdminGetGoodsRequest
	1, // 4: good.gateway.app.good1.v1.Gateway.GetGoods:output_type -> good.gateway.app.good1.v1.GetGoodsResponse
	3, // 5: good.gateway.app.good1.v1.Gateway.AdminGetGoods:output_type -> good.gateway.app.good1.v1.AdminGetGoodsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_good_gateway_v1_app_good_good_proto_init() }
func file_good_gateway_v1_app_good_good_proto_init() {
	if File_good_gateway_v1_app_good_good_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_good_gateway_v1_app_good_good_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_good_gateway_v1_app_good_good_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsResponse); i {
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
		file_good_gateway_v1_app_good_good_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetGoodsRequest); i {
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
		file_good_gateway_v1_app_good_good_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetGoodsResponse); i {
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
			RawDescriptor: file_good_gateway_v1_app_good_good_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_good_gateway_v1_app_good_good_proto_goTypes,
		DependencyIndexes: file_good_gateway_v1_app_good_good_proto_depIdxs,
		MessageInfos:      file_good_gateway_v1_app_good_good_proto_msgTypes,
	}.Build()
	File_good_gateway_v1_app_good_good_proto = out.File
	file_good_gateway_v1_app_good_good_proto_rawDesc = nil
	file_good_gateway_v1_app_good_good_proto_goTypes = nil
	file_good_gateway_v1_app_good_good_proto_depIdxs = nil
}
