// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: order/middleware/v1/powerrental/poolorderuser/poolorderuser.proto

package miningpoolorderuser

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

type PoolOrderUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID           *string `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	OrderID         *string `protobuf:"bytes,20,opt,name=OrderID,proto3,oneof" json:"OrderID,omitempty"`
	PoolOrderUserID *string `protobuf:"bytes,30,opt,name=PoolOrderUserID,proto3,oneof" json:"PoolOrderUserID,omitempty"`
}

func (x *PoolOrderUserReq) Reset() {
	*x = PoolOrderUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolOrderUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolOrderUserReq) ProtoMessage() {}

func (x *PoolOrderUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolOrderUserReq.ProtoReflect.Descriptor instead.
func (*PoolOrderUserReq) Descriptor() ([]byte, []int) {
	return file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescGZIP(), []int{0}
}

func (x *PoolOrderUserReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *PoolOrderUserReq) GetOrderID() string {
	if x != nil && x.OrderID != nil {
		return *x.OrderID
	}
	return ""
}

func (x *PoolOrderUserReq) GetPoolOrderUserID() string {
	if x != nil && x.PoolOrderUserID != nil {
		return *x.PoolOrderUserID
	}
	return ""
}

type PoolOrderUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"order_id"
	OrderID string `protobuf:"bytes,30,opt,name=OrderID,proto3" json:"OrderID,omitempty" sql:"order_id"`
	// @inject_tag: sql:"order_user_id"
	PoolOrderUserID string `protobuf:"bytes,40,opt,name=PoolOrderUserID,proto3" json:"PoolOrderUserID,omitempty" sql:"order_user_id"`
}

func (x *PoolOrderUser) Reset() {
	*x = PoolOrderUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoolOrderUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolOrderUser) ProtoMessage() {}

func (x *PoolOrderUser) ProtoReflect() protoreflect.Message {
	mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolOrderUser.ProtoReflect.Descriptor instead.
func (*PoolOrderUser) Descriptor() ([]byte, []int) {
	return file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescGZIP(), []int{1}
}

func (x *PoolOrderUser) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PoolOrderUser) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *PoolOrderUser) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *PoolOrderUser) GetPoolOrderUserID() string {
	if x != nil {
		return x.PoolOrderUserID
	}
	return ""
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              *v1.Uint32Val `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID           *v1.StringVal `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	OrderID         *v1.StringVal `protobuf:"bytes,30,opt,name=OrderID,proto3,oneof" json:"OrderID,omitempty"`
	PoolOrderUserID *v1.StringVal `protobuf:"bytes,40,opt,name=PoolOrderUserID,proto3,oneof" json:"PoolOrderUserID,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[2]
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
	return file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescGZIP(), []int{2}
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

func (x *Conds) GetOrderID() *v1.StringVal {
	if x != nil {
		return x.OrderID
	}
	return nil
}

func (x *Conds) GetPoolOrderUserID() *v1.StringVal {
	if x != nil {
		return x.PoolOrderUserID
	}
	return nil
}

var File_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto protoreflect.FileDescriptor

var file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDesc = []byte{
	0x0a, 0x41, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x6f, 0x6f, 0x6c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x50, 0x6f, 0x6f, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x50,
	0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x79,
	0x0a, 0x0d, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x9a, 0x02, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x02, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a,
	0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x03, 0x52, 0x0f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescOnce sync.Once
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescData = file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDesc
)

func file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescGZIP() []byte {
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescOnce.Do(func() {
		file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescData)
	})
	return file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDescData
}

var (
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_goTypes  = []interface{}{
		(*PoolOrderUserReq)(nil), // 0: order.middleware.miningpoolorderuser.v1.PoolOrderUserReq
		(*PoolOrderUser)(nil),    // 1: order.middleware.miningpoolorderuser.v1.PoolOrderUser
		(*Conds)(nil),            // 2: order.middleware.miningpoolorderuser.v1.Conds
		(*v1.Uint32Val)(nil),     // 3: basetypes.v1.Uint32Val
		(*v1.StringVal)(nil),     // 4: basetypes.v1.StringVal
	}
)
var file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_depIdxs = []int32{
	3, // 0: order.middleware.miningpoolorderuser.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	4, // 1: order.middleware.miningpoolorderuser.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	4, // 2: order.middleware.miningpoolorderuser.v1.Conds.OrderID:type_name -> basetypes.v1.StringVal
	4, // 3: order.middleware.miningpoolorderuser.v1.Conds.PoolOrderUserID:type_name -> basetypes.v1.StringVal
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_init() }
func file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_init() {
	if File_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolOrderUserReq); i {
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
		file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoolOrderUser); i {
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
		file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_goTypes,
		DependencyIndexes: file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_depIdxs,
		MessageInfos:      file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_msgTypes,
	}.Build()
	File_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto = out.File
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_rawDesc = nil
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_goTypes = nil
	file_order_middleware_v1_powerrental_poolorderuser_poolorderuser_proto_depIdxs = nil
}
