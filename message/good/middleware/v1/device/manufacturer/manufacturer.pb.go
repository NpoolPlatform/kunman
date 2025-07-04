// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: good/middleware/v1/device/manufacturer/manufacturer.proto

package manufacturer

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

type ManufacturerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Name  *string `protobuf:"bytes,30,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
	Logo  *string `protobuf:"bytes,40,opt,name=Logo,proto3,oneof" json:"Logo,omitempty"`
}

func (x *ManufacturerReq) Reset() {
	*x = ManufacturerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManufacturerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManufacturerReq) ProtoMessage() {}

func (x *ManufacturerReq) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManufacturerReq.ProtoReflect.Descriptor instead.
func (*ManufacturerReq) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescGZIP(), []int{0}
}

func (x *ManufacturerReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *ManufacturerReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *ManufacturerReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ManufacturerReq) GetLogo() string {
	if x != nil && x.Logo != nil {
		return *x.Logo
	}
	return ""
}

type Manufacturer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"name"
	Name string `protobuf:"bytes,30,opt,name=Name,proto3" json:"Name,omitempty" sql:"name"`
	// @inject_tag: sql:"logo"
	Logo string `protobuf:"bytes,40,opt,name=Logo,proto3" json:"Logo,omitempty" sql:"logo"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Manufacturer) Reset() {
	*x = Manufacturer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manufacturer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manufacturer) ProtoMessage() {}

func (x *Manufacturer) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manufacturer.ProtoReflect.Descriptor instead.
func (*Manufacturer) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescGZIP(), []int{1}
}

func (x *Manufacturer) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Manufacturer) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Manufacturer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Manufacturer) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Manufacturer) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Manufacturer) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    *v1.Uint32Val `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID *v1.StringVal `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Name  *v1.StringVal `protobuf:"bytes,30,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[2]
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
	return file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescGZIP(), []int{2}
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

func (x *Conds) GetName() *v1.StringVal {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_good_middleware_v1_device_manufacturer_manufacturer_proto protoreflect.FileDescriptor

var file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01,
	0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4c, 0x6f,
	0x67, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f,
	0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x4c, 0x6f, 0x67, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescOnce sync.Once
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescData = file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDesc
)

func file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescGZIP() []byte {
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescOnce.Do(func() {
		file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescData = protoimpl.X.CompressGZIP(file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescData)
	})
	return file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDescData
}

var (
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_goTypes  = []interface{}{
		(*ManufacturerReq)(nil), // 0: good.middleware.device.manufacturer.v1.ManufacturerReq
		(*Manufacturer)(nil),    // 1: good.middleware.device.manufacturer.v1.Manufacturer
		(*Conds)(nil),           // 2: good.middleware.device.manufacturer.v1.Conds
		(*v1.Uint32Val)(nil),    // 3: basetypes.v1.Uint32Val
		(*v1.StringVal)(nil),    // 4: basetypes.v1.StringVal
	}
)
var file_good_middleware_v1_device_manufacturer_manufacturer_proto_depIdxs = []int32{
	3, // 0: good.middleware.device.manufacturer.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	4, // 1: good.middleware.device.manufacturer.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	4, // 2: good.middleware.device.manufacturer.v1.Conds.Name:type_name -> basetypes.v1.StringVal
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_good_middleware_v1_device_manufacturer_manufacturer_proto_init() }
func file_good_middleware_v1_device_manufacturer_manufacturer_proto_init() {
	if File_good_middleware_v1_device_manufacturer_manufacturer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManufacturerReq); i {
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
		file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Manufacturer); i {
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
		file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_good_middleware_v1_device_manufacturer_manufacturer_proto_goTypes,
		DependencyIndexes: file_good_middleware_v1_device_manufacturer_manufacturer_proto_depIdxs,
		MessageInfos:      file_good_middleware_v1_device_manufacturer_manufacturer_proto_msgTypes,
	}.Build()
	File_good_middleware_v1_device_manufacturer_manufacturer_proto = out.File
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_rawDesc = nil
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_goTypes = nil
	file_good_middleware_v1_device_manufacturer_manufacturer_proto_depIdxs = nil
}
