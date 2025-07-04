// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: basetypes/chain/v1/enums.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CoinUsedFor int32

const (
	CoinUsedFor_DefaultCoinUsedFor    CoinUsedFor = 0
	CoinUsedFor_CoinUsedForCouponCash CoinUsedFor = 10
	CoinUsedFor_CoinUsedForGoodFee    CoinUsedFor = 20
)

// Enum value maps for CoinUsedFor.
var (
	CoinUsedFor_name = map[int32]string{
		0:  "DefaultCoinUsedFor",
		10: "CoinUsedForCouponCash",
		20: "CoinUsedForGoodFee",
	}
	CoinUsedFor_value = map[string]int32{
		"DefaultCoinUsedFor":    0,
		"CoinUsedForCouponCash": 10,
		"CoinUsedForGoodFee":    20,
	}
)

func (x CoinUsedFor) Enum() *CoinUsedFor {
	p := new(CoinUsedFor)
	*p = x
	return p
}

func (x CoinUsedFor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinUsedFor) Descriptor() protoreflect.EnumDescriptor {
	return file_basetypes_chain_v1_enums_proto_enumTypes[0].Descriptor()
}

func (CoinUsedFor) Type() protoreflect.EnumType {
	return &file_basetypes_chain_v1_enums_proto_enumTypes[0]
}

func (x CoinUsedFor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinUsedFor.Descriptor instead.
func (CoinUsedFor) EnumDescriptor() ([]byte, []int) {
	return file_basetypes_chain_v1_enums_proto_rawDescGZIP(), []int{0}
}

var File_basetypes_chain_v1_enums_proto protoreflect.FileDescriptor

var file_basetypes_chain_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x58, 0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x43, 0x61, 0x73, 0x68, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x46, 0x65, 0x65, 0x10, 0x14, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61,
	0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basetypes_chain_v1_enums_proto_rawDescOnce sync.Once
	file_basetypes_chain_v1_enums_proto_rawDescData = file_basetypes_chain_v1_enums_proto_rawDesc
)

func file_basetypes_chain_v1_enums_proto_rawDescGZIP() []byte {
	file_basetypes_chain_v1_enums_proto_rawDescOnce.Do(func() {
		file_basetypes_chain_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_basetypes_chain_v1_enums_proto_rawDescData)
	})
	return file_basetypes_chain_v1_enums_proto_rawDescData
}

var (
	file_basetypes_chain_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_basetypes_chain_v1_enums_proto_goTypes   = []interface{}{
		(CoinUsedFor)(0), // 0: basetypes.chain.v1.CoinUsedFor
	}
)
var file_basetypes_chain_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basetypes_chain_v1_enums_proto_init() }
func file_basetypes_chain_v1_enums_proto_init() {
	if File_basetypes_chain_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basetypes_chain_v1_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basetypes_chain_v1_enums_proto_goTypes,
		DependencyIndexes: file_basetypes_chain_v1_enums_proto_depIdxs,
		EnumInfos:         file_basetypes_chain_v1_enums_proto_enumTypes,
	}.Build()
	File_basetypes_chain_v1_enums_proto = out.File
	file_basetypes_chain_v1_enums_proto_rawDesc = nil
	file_basetypes_chain_v1_enums_proto_goTypes = nil
	file_basetypes_chain_v1_enums_proto_depIdxs = nil
}
