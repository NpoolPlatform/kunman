// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: chain/middleware/v1/fiat/fiat.proto

package fiat

import (
	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

type FiatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID *string `protobuf:"bytes,11,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Name  *string `protobuf:"bytes,20,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
	Logo  *string `protobuf:"bytes,30,opt,name=Logo,proto3,oneof" json:"Logo,omitempty"`
	Unit  *string `protobuf:"bytes,40,opt,name=Unit,proto3,oneof" json:"Unit,omitempty"`
}

func (x *FiatReq) Reset() {
	*x = FiatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiatReq) ProtoMessage() {}

func (x *FiatReq) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiatReq.ProtoReflect.Descriptor instead.
func (*FiatReq) Descriptor() ([]byte, []int) {
	return file_chain_middleware_v1_fiat_fiat_proto_rawDescGZIP(), []int{0}
}

func (x *FiatReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *FiatReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *FiatReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FiatReq) GetLogo() string {
	if x != nil && x.Logo != nil {
		return *x.Logo
	}
	return ""
}

func (x *FiatReq) GetUnit() string {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return ""
}

type Fiat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"name"
	Name string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty" sql:"name"`
	// @inject_tag: sql:"logo"
	Logo string `protobuf:"bytes,30,opt,name=Logo,proto3" json:"Logo,omitempty" sql:"logo"`
	// @inject_tag: sql:"unit"
	Unit string `protobuf:"bytes,40,opt,name=Unit,proto3" json:"Unit,omitempty" sql:"unit"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Fiat) Reset() {
	*x = Fiat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fiat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fiat) ProtoMessage() {}

func (x *Fiat) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fiat.ProtoReflect.Descriptor instead.
func (*Fiat) Descriptor() ([]byte, []int) {
	return file_chain_middleware_v1_fiat_fiat_proto_rawDescGZIP(), []int{1}
}

func (x *Fiat) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Fiat) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Fiat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fiat) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Fiat) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Fiat) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Fiat) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID  *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Name   *v1.StringVal      `protobuf:"bytes,20,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
	EntIDs *v1.StringSliceVal `protobuf:"bytes,30,opt,name=EntIDs,proto3,oneof" json:"EntIDs,omitempty"`
	Unit   *v1.StringVal      `protobuf:"bytes,40,opt,name=Unit,proto3,oneof" json:"Unit,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_fiat_fiat_proto_msgTypes[2]
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
	return file_chain_middleware_v1_fiat_fiat_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetName() *v1.StringVal {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Conds) GetEntIDs() *v1.StringSliceVal {
	if x != nil {
		return x.EntIDs
	}
	return nil
}

func (x *Conds) GetUnit() *v1.StringVal {
	if x != nil {
		return x.Unit
	}
	return nil
}

var File_chain_middleware_v1_fiat_fiat_proto protoreflect.FileDescriptor

var file_chain_middleware_v1_fiat_fiat_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x07, 0x46, 0x69,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04,
	0x4c, 0x6f, 0x67, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4c,
	0x6f, 0x67, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x55, 0x6e, 0x69, 0x74, 0x22, 0xa6, 0x01, 0x0a,
	0x04, 0x46, 0x69, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c,
	0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x81, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x30, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x73,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_middleware_v1_fiat_fiat_proto_rawDescOnce sync.Once
	file_chain_middleware_v1_fiat_fiat_proto_rawDescData = file_chain_middleware_v1_fiat_fiat_proto_rawDesc
)

func file_chain_middleware_v1_fiat_fiat_proto_rawDescGZIP() []byte {
	file_chain_middleware_v1_fiat_fiat_proto_rawDescOnce.Do(func() {
		file_chain_middleware_v1_fiat_fiat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_middleware_v1_fiat_fiat_proto_rawDescData)
	})
	return file_chain_middleware_v1_fiat_fiat_proto_rawDescData
}

var file_chain_middleware_v1_fiat_fiat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chain_middleware_v1_fiat_fiat_proto_goTypes = []interface{}{
	(*FiatReq)(nil),           // 0: chain.middleware.fiat.v1.FiatReq
	(*Fiat)(nil),              // 1: chain.middleware.fiat.v1.Fiat
	(*Conds)(nil),             // 2: chain.middleware.fiat.v1.Conds
	(*v1.StringVal)(nil),      // 3: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil), // 4: basetypes.v1.StringSliceVal
}
var file_chain_middleware_v1_fiat_fiat_proto_depIdxs = []int32{
	3, // 0: chain.middleware.fiat.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	3, // 1: chain.middleware.fiat.v1.Conds.Name:type_name -> basetypes.v1.StringVal
	4, // 2: chain.middleware.fiat.v1.Conds.EntIDs:type_name -> basetypes.v1.StringSliceVal
	3, // 3: chain.middleware.fiat.v1.Conds.Unit:type_name -> basetypes.v1.StringVal
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chain_middleware_v1_fiat_fiat_proto_init() }
func file_chain_middleware_v1_fiat_fiat_proto_init() {
	if File_chain_middleware_v1_fiat_fiat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_middleware_v1_fiat_fiat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiatReq); i {
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
		file_chain_middleware_v1_fiat_fiat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fiat); i {
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
		file_chain_middleware_v1_fiat_fiat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_chain_middleware_v1_fiat_fiat_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_chain_middleware_v1_fiat_fiat_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chain_middleware_v1_fiat_fiat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chain_middleware_v1_fiat_fiat_proto_goTypes,
		DependencyIndexes: file_chain_middleware_v1_fiat_fiat_proto_depIdxs,
		MessageInfos:      file_chain_middleware_v1_fiat_fiat_proto_msgTypes,
	}.Build()
	File_chain_middleware_v1_fiat_fiat_proto = out.File
	file_chain_middleware_v1_fiat_fiat_proto_rawDesc = nil
	file_chain_middleware_v1_fiat_fiat_proto_goTypes = nil
	file_chain_middleware_v1_fiat_fiat_proto_depIdxs = nil
}
