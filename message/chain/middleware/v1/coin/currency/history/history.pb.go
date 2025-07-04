// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: chain/middleware/v1/coin/currency/history/history.proto

package history

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	_ "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID       *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	CoinTypeID  *v1.StringVal      `protobuf:"bytes,20,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	CoinTypeIDs *v1.StringSliceVal `protobuf:"bytes,30,opt,name=CoinTypeIDs,proto3,oneof" json:"CoinTypeIDs,omitempty"`
	StartAt     *v1.Uint32Val      `protobuf:"bytes,40,opt,name=StartAt,proto3,oneof" json:"StartAt,omitempty"`
	EndAt       *v1.Uint32Val      `protobuf:"bytes,50,opt,name=EndAt,proto3,oneof" json:"EndAt,omitempty"`
	CoinNames   *v1.StringSliceVal `protobuf:"bytes,60,opt,name=CoinNames,proto3,oneof" json:"CoinNames,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes[0]
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
	return file_chain_middleware_v1_coin_currency_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetCoinTypeIDs() *v1.StringSliceVal {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

func (x *Conds) GetStartAt() *v1.Uint32Val {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *Conds) GetEndAt() *v1.Uint32Val {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Conds) GetCoinNames() *v1.StringSliceVal {
	if x != nil {
		return x.CoinNames
	}
	return nil
}

var File_chain_middleware_v1_coin_currency_history_history_proto protoreflect.FileDescriptor

var file_chain_middleware_v1_coin_currency_history_history_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x30, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a,
	0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c,
	0x48, 0x02, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x07,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e,
	0x64, 0x41, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x48, 0x04, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3f,
	0x0a, 0x09, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48,
	0x05, 0x52, 0x09, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x53, 0x5a, 0x51,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_middleware_v1_coin_currency_history_history_proto_rawDescOnce sync.Once
	file_chain_middleware_v1_coin_currency_history_history_proto_rawDescData = file_chain_middleware_v1_coin_currency_history_history_proto_rawDesc
)

func file_chain_middleware_v1_coin_currency_history_history_proto_rawDescGZIP() []byte {
	file_chain_middleware_v1_coin_currency_history_history_proto_rawDescOnce.Do(func() {
		file_chain_middleware_v1_coin_currency_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_middleware_v1_coin_currency_history_history_proto_rawDescData)
	})
	return file_chain_middleware_v1_coin_currency_history_history_proto_rawDescData
}

var (
	file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_chain_middleware_v1_coin_currency_history_history_proto_goTypes  = []interface{}{
		(*Conds)(nil),             // 0: chain.middleware.coin.currency1.history.v1.Conds
		(*v1.StringVal)(nil),      // 1: basetypes.v1.StringVal
		(*v1.StringSliceVal)(nil), // 2: basetypes.v1.StringSliceVal
		(*v1.Uint32Val)(nil),      // 3: basetypes.v1.Uint32Val
	}
)
var file_chain_middleware_v1_coin_currency_history_history_proto_depIdxs = []int32{
	1, // 0: chain.middleware.coin.currency1.history.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	1, // 1: chain.middleware.coin.currency1.history.v1.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	2, // 2: chain.middleware.coin.currency1.history.v1.Conds.CoinTypeIDs:type_name -> basetypes.v1.StringSliceVal
	3, // 3: chain.middleware.coin.currency1.history.v1.Conds.StartAt:type_name -> basetypes.v1.Uint32Val
	3, // 4: chain.middleware.coin.currency1.history.v1.Conds.EndAt:type_name -> basetypes.v1.Uint32Val
	2, // 5: chain.middleware.coin.currency1.history.v1.Conds.CoinNames:type_name -> basetypes.v1.StringSliceVal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chain_middleware_v1_coin_currency_history_history_proto_init() }
func file_chain_middleware_v1_coin_currency_history_history_proto_init() {
	if File_chain_middleware_v1_coin_currency_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chain_middleware_v1_coin_currency_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chain_middleware_v1_coin_currency_history_history_proto_goTypes,
		DependencyIndexes: file_chain_middleware_v1_coin_currency_history_history_proto_depIdxs,
		MessageInfos:      file_chain_middleware_v1_coin_currency_history_history_proto_msgTypes,
	}.Build()
	File_chain_middleware_v1_coin_currency_history_history_proto = out.File
	file_chain_middleware_v1_coin_currency_history_history_proto_rawDesc = nil
	file_chain_middleware_v1_coin_currency_history_history_proto_goTypes = nil
	file_chain_middleware_v1_coin_currency_history_history_proto_depIdxs = nil
}
