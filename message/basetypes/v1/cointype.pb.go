// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: basetypes/v1/cointype.proto

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

type CoinType int32

const (
	CoinType_DefaultCoinType CoinType = 0
	// mainnet
	CoinType_CoinTypeFileCoin    CoinType = 1
	CoinType_CoinTypeBitCoin     CoinType = 2
	CoinType_CoinTypeEthereum    CoinType = 3
	CoinType_CoinTypeUSDTerc20   CoinType = 4
	CoinType_CoinTypeSpacemesh   CoinType = 5
	CoinType_CoinTypeSolana      CoinType = 6
	CoinType_CoinTypeUSDTtrc20   CoinType = 7
	CoinType_CoinTypeBinanceCoin CoinType = 8
	CoinType_CoinTypeTron        CoinType = 9
	CoinType_CoinTypeBinanceUSD  CoinType = 10
	CoinType_CoinTypeUSDCerc20   CoinType = 11
	CoinType_CoinTypeAleo        CoinType = 12
	CoinType_CoinTypeIronFish    CoinType = 13
	CoinType_CoinTypeUSDTbep20   CoinType = 14
	// testnet
	CoinType_CoinTypeTFileCoin    CoinType = 1001
	CoinType_CoinTypeTBitCoin     CoinType = 1002
	CoinType_CoinTypeTEthereum    CoinType = 1003
	CoinType_CoinTypeTUSDTerc20   CoinType = 1004
	CoinType_CoinTypeTSpacemesh   CoinType = 1005
	CoinType_CoinTypeTSolana      CoinType = 1006
	CoinType_CoinTypeTUSDTtrc20   CoinType = 1007
	CoinType_CoinTypeTBinanceCoin CoinType = 1008
	CoinType_CoinTypeTTron        CoinType = 1009
	CoinType_CoinTypeTBinanceUSD  CoinType = 1010
	CoinType_CoinTypeTUSDCerc20   CoinType = 1011
	CoinType_CoinTypeTAleo        CoinType = 1012
	CoinType_CoinTypeTIronFish    CoinType = 1013
	CoinType_CoinTypeTUSDTbep20   CoinType = 1014
)

// Enum value maps for CoinType.
var (
	CoinType_name = map[int32]string{
		0:    "DefaultCoinType",
		1:    "CoinTypeFileCoin",
		2:    "CoinTypeBitCoin",
		3:    "CoinTypeEthereum",
		4:    "CoinTypeUSDTerc20",
		5:    "CoinTypeSpacemesh",
		6:    "CoinTypeSolana",
		7:    "CoinTypeUSDTtrc20",
		8:    "CoinTypeBinanceCoin",
		9:    "CoinTypeTron",
		10:   "CoinTypeBinanceUSD",
		11:   "CoinTypeUSDCerc20",
		12:   "CoinTypeAleo",
		13:   "CoinTypeIronFish",
		14:   "CoinTypeUSDTbep20",
		1001: "CoinTypeTFileCoin",
		1002: "CoinTypeTBitCoin",
		1003: "CoinTypeTEthereum",
		1004: "CoinTypeTUSDTerc20",
		1005: "CoinTypeTSpacemesh",
		1006: "CoinTypeTSolana",
		1007: "CoinTypeTUSDTtrc20",
		1008: "CoinTypeTBinanceCoin",
		1009: "CoinTypeTTron",
		1010: "CoinTypeTBinanceUSD",
		1011: "CoinTypeTUSDCerc20",
		1012: "CoinTypeTAleo",
		1013: "CoinTypeTIronFish",
		1014: "CoinTypeTUSDTbep20",
	}
	CoinType_value = map[string]int32{
		"DefaultCoinType":      0,
		"CoinTypeFileCoin":     1,
		"CoinTypeBitCoin":      2,
		"CoinTypeEthereum":     3,
		"CoinTypeUSDTerc20":    4,
		"CoinTypeSpacemesh":    5,
		"CoinTypeSolana":       6,
		"CoinTypeUSDTtrc20":    7,
		"CoinTypeBinanceCoin":  8,
		"CoinTypeTron":         9,
		"CoinTypeBinanceUSD":   10,
		"CoinTypeUSDCerc20":    11,
		"CoinTypeAleo":         12,
		"CoinTypeIronFish":     13,
		"CoinTypeUSDTbep20":    14,
		"CoinTypeTFileCoin":    1001,
		"CoinTypeTBitCoin":     1002,
		"CoinTypeTEthereum":    1003,
		"CoinTypeTUSDTerc20":   1004,
		"CoinTypeTSpacemesh":   1005,
		"CoinTypeTSolana":      1006,
		"CoinTypeTUSDTtrc20":   1007,
		"CoinTypeTBinanceCoin": 1008,
		"CoinTypeTTron":        1009,
		"CoinTypeTBinanceUSD":  1010,
		"CoinTypeTUSDCerc20":   1011,
		"CoinTypeTAleo":        1012,
		"CoinTypeTIronFish":    1013,
		"CoinTypeTUSDTbep20":   1014,
	}
)

func (x CoinType) Enum() *CoinType {
	p := new(CoinType)
	*p = x
	return p
}

func (x CoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_basetypes_v1_cointype_proto_enumTypes[0].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_basetypes_v1_cointype_proto_enumTypes[0]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_basetypes_v1_cointype_proto_rawDescGZIP(), []int{0}
}

var File_basetypes_v1_cointype_proto protoreflect.FileDescriptor

var file_basetypes_v1_cointype_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0xa1, 0x05, 0x0a, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x69, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x10, 0x03, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x53, 0x44, 0x54, 0x65, 0x72,
	0x63, 0x32, 0x30, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x10, 0x06,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x53, 0x44, 0x54,
	0x74, 0x72, 0x63, 0x32, 0x30, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x08,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x72, 0x6f, 0x6e,
	0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x53, 0x44, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x53, 0x44, 0x43, 0x65, 0x72, 0x63, 0x32, 0x30, 0x10,
	0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6c, 0x65,
	0x6f, 0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x72, 0x6f, 0x6e, 0x46, 0x69, 0x73, 0x68, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x55, 0x53, 0x44, 0x54, 0x62, 0x65, 0x70, 0x32, 0x30, 0x10, 0x0e,
	0x12, 0x16, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0xe9, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x54, 0x42, 0x69, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0xea, 0x07, 0x12,
	0x16, 0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x10, 0xeb, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x54, 0x55, 0x53, 0x44, 0x54, 0x65, 0x72, 0x63, 0x32, 0x30, 0x10, 0xec, 0x07,
	0x12, 0x17, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x10, 0xed, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x53, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x10, 0xee, 0x07, 0x12,
	0x17, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x55, 0x53, 0x44, 0x54,
	0x74, 0x72, 0x63, 0x32, 0x30, 0x10, 0xef, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x54, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x10, 0xf0, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54,
	0x54, 0x72, 0x6f, 0x6e, 0x10, 0xf1, 0x07, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x54, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x53, 0x44, 0x10, 0xf2,
	0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x55, 0x53,
	0x44, 0x43, 0x65, 0x72, 0x63, 0x32, 0x30, 0x10, 0xf3, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x41, 0x6c, 0x65, 0x6f, 0x10, 0xf4, 0x07, 0x12, 0x16,
	0x0a, 0x11, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x54, 0x49, 0x72, 0x6f, 0x6e, 0x46,
	0x69, 0x73, 0x68, 0x10, 0xf5, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x54, 0x55, 0x53, 0x44, 0x54, 0x62, 0x65, 0x70, 0x32, 0x30, 0x10, 0xf6, 0x07, 0x42,
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d,
	0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basetypes_v1_cointype_proto_rawDescOnce sync.Once
	file_basetypes_v1_cointype_proto_rawDescData = file_basetypes_v1_cointype_proto_rawDesc
)

func file_basetypes_v1_cointype_proto_rawDescGZIP() []byte {
	file_basetypes_v1_cointype_proto_rawDescOnce.Do(func() {
		file_basetypes_v1_cointype_proto_rawDescData = protoimpl.X.CompressGZIP(file_basetypes_v1_cointype_proto_rawDescData)
	})
	return file_basetypes_v1_cointype_proto_rawDescData
}

var (
	file_basetypes_v1_cointype_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_basetypes_v1_cointype_proto_goTypes   = []interface{}{
		(CoinType)(0), // 0: basetypes.v1.CoinType
	}
)
var file_basetypes_v1_cointype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basetypes_v1_cointype_proto_init() }
func file_basetypes_v1_cointype_proto_init() {
	if File_basetypes_v1_cointype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basetypes_v1_cointype_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basetypes_v1_cointype_proto_goTypes,
		DependencyIndexes: file_basetypes_v1_cointype_proto_depIdxs,
		EnumInfos:         file_basetypes_v1_cointype_proto_enumTypes,
	}.Build()
	File_basetypes_v1_cointype_proto = out.File
	file_basetypes_v1_cointype_proto_rawDesc = nil
	file_basetypes_v1_cointype_proto_goTypes = nil
	file_basetypes_v1_cointype_proto_depIdxs = nil
}
