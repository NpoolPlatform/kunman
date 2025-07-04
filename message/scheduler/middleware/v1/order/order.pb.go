// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: scheduler/middleware/v1/order/order.proto

package order

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentType int32

const (
	PaymentType_DefaultPaymentType PaymentType = 0
	PaymentType_PayWithBalance     PaymentType = 10
	PaymentType_PayWithTransfer    PaymentType = 20
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0:  "DefaultPaymentType",
		10: "PayWithBalance",
		20: "PayWithTransfer",
	}
	PaymentType_value = map[string]int32{
		"DefaultPaymentType": 0,
		"PayWithBalance":     10,
		"PayWithTransfer":    20,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_scheduler_middleware_v1_order_order_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_scheduler_middleware_v1_order_order_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_scheduler_middleware_v1_order_order_proto_rawDescGZIP(), []int{0}
}

type PaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID  string      `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	Amount      string      `protobuf:"bytes,20,opt,name=Amount,proto3" json:"Amount,omitempty"`
	PaymentType PaymentType `protobuf:"varint,30,opt,name=PaymentType,proto3,enum=scheduler.middleware.order.v1.PaymentType" json:"PaymentType,omitempty"`
}

func (x *PaymentInfo) Reset() {
	*x = PaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_middleware_v1_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo) ProtoMessage() {}

func (x *PaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_middleware_v1_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo.ProtoReflect.Descriptor instead.
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return file_scheduler_middleware_v1_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentInfo) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *PaymentInfo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *PaymentInfo) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_DefaultPaymentType
}

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID            string         `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID           string         `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	OrderID          string         `protobuf:"bytes,30,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	GoodType         v1.GoodType    `protobuf:"varint,40,opt,name=GoodType,proto3,enum=basetypes.good.v1.GoodType" json:"GoodType,omitempty"`
	Units            string         `protobuf:"bytes,50,opt,name=Units,proto3" json:"Units,omitempty"`
	PaymentAmountUSD string         `protobuf:"bytes,60,opt,name=PaymentAmountUSD,proto3" json:"PaymentAmountUSD,omitempty"`
	Payments         []*PaymentInfo `protobuf:"bytes,70,rep,name=Payments,proto3" json:"Payments,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_middleware_v1_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_middleware_v1_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_scheduler_middleware_v1_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderInfo) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *OrderInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *OrderInfo) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *OrderInfo) GetGoodType() v1.GoodType {
	if x != nil {
		return x.GoodType
	}
	return v1.GoodType(0)
}

func (x *OrderInfo) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *OrderInfo) GetPaymentAmountUSD() string {
	if x != nil {
		return x.PaymentAmountUSD
	}
	return ""
}

func (x *OrderInfo) GetPayments() []*PaymentInfo {
	if x != nil {
		return x.Payments
	}
	return nil
}

var File_scheduler_middleware_v1_order_order_proto protoreflect.FileDescriptor

var file_scheduler_middleware_v1_order_order_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x96, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44,
	0x12, 0x46, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x46, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x4e, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x14, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_middleware_v1_order_order_proto_rawDescOnce sync.Once
	file_scheduler_middleware_v1_order_order_proto_rawDescData = file_scheduler_middleware_v1_order_order_proto_rawDesc
)

func file_scheduler_middleware_v1_order_order_proto_rawDescGZIP() []byte {
	file_scheduler_middleware_v1_order_order_proto_rawDescOnce.Do(func() {
		file_scheduler_middleware_v1_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_middleware_v1_order_order_proto_rawDescData)
	})
	return file_scheduler_middleware_v1_order_order_proto_rawDescData
}

var (
	file_scheduler_middleware_v1_order_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_scheduler_middleware_v1_order_order_proto_msgTypes  = make([]protoimpl.MessageInfo, 2)
	file_scheduler_middleware_v1_order_order_proto_goTypes   = []interface{}{
		(PaymentType)(0),    // 0: scheduler.middleware.order.v1.PaymentType
		(*PaymentInfo)(nil), // 1: scheduler.middleware.order.v1.PaymentInfo
		(*OrderInfo)(nil),   // 2: scheduler.middleware.order.v1.OrderInfo
		(v1.GoodType)(0),    // 3: basetypes.good.v1.GoodType
	}
)
var file_scheduler_middleware_v1_order_order_proto_depIdxs = []int32{
	0, // 0: scheduler.middleware.order.v1.PaymentInfo.PaymentType:type_name -> scheduler.middleware.order.v1.PaymentType
	3, // 1: scheduler.middleware.order.v1.OrderInfo.GoodType:type_name -> basetypes.good.v1.GoodType
	1, // 2: scheduler.middleware.order.v1.OrderInfo.Payments:type_name -> scheduler.middleware.order.v1.PaymentInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_scheduler_middleware_v1_order_order_proto_init() }
func file_scheduler_middleware_v1_order_order_proto_init() {
	if File_scheduler_middleware_v1_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_middleware_v1_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentInfo); i {
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
		file_scheduler_middleware_v1_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
			RawDescriptor: file_scheduler_middleware_v1_order_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scheduler_middleware_v1_order_order_proto_goTypes,
		DependencyIndexes: file_scheduler_middleware_v1_order_order_proto_depIdxs,
		EnumInfos:         file_scheduler_middleware_v1_order_order_proto_enumTypes,
		MessageInfos:      file_scheduler_middleware_v1_order_order_proto_msgTypes,
	}.Build()
	File_scheduler_middleware_v1_order_order_proto = out.File
	file_scheduler_middleware_v1_order_order_proto_rawDesc = nil
	file_scheduler_middleware_v1_order_order_proto_goTypes = nil
	file_scheduler_middleware_v1_order_order_proto_depIdxs = nil
}
