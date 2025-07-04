// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: notif/gateway/v1/usercode/usercode.proto

package usercode

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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

type SendCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID       string        `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	LangID      string        `protobuf:"bytes,20,opt,name=LangID,proto3" json:"LangID,omitempty"`
	UserID      *string       `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	Account     *string       `protobuf:"bytes,40,opt,name=Account,proto3,oneof" json:"Account,omitempty"`
	AccountType v1.SignMethod `protobuf:"varint,50,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	UsedFor     v1.UsedFor    `protobuf:"varint,60,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
	ToUsername  *string       `protobuf:"bytes,70,opt,name=ToUsername,proto3,oneof" json:"ToUsername,omitempty"`
}

func (x *SendCodeRequest) Reset() {
	*x = SendCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_gateway_v1_usercode_usercode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeRequest) ProtoMessage() {}

func (x *SendCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notif_gateway_v1_usercode_usercode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeRequest.ProtoReflect.Descriptor instead.
func (*SendCodeRequest) Descriptor() ([]byte, []int) {
	return file_notif_gateway_v1_usercode_usercode_proto_rawDescGZIP(), []int{0}
}

func (x *SendCodeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *SendCodeRequest) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *SendCodeRequest) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *SendCodeRequest) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *SendCodeRequest) GetAccountType() v1.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *SendCodeRequest) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *SendCodeRequest) GetToUsername() string {
	if x != nil && x.ToUsername != nil {
		return *x.ToUsername
	}
	return ""
}

type SendCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendCodeResponse) Reset() {
	*x = SendCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notif_gateway_v1_usercode_usercode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeResponse) ProtoMessage() {}

func (x *SendCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notif_gateway_v1_usercode_usercode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeResponse.ProtoReflect.Descriptor instead.
func (*SendCodeResponse) Descriptor() ([]byte, []int) {
	return file_notif_gateway_v1_usercode_usercode_proto_rawDescGZIP(), []int{1}
}

var File_notif_gateway_v1_usercode_usercode_proto protoreflect.FileDescriptor

var file_notif_gateway_v1_usercode_usercode_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3,
	0x02, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44,
	0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0a, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0a, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8d, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notif_gateway_v1_usercode_usercode_proto_rawDescOnce sync.Once
	file_notif_gateway_v1_usercode_usercode_proto_rawDescData = file_notif_gateway_v1_usercode_usercode_proto_rawDesc
)

func file_notif_gateway_v1_usercode_usercode_proto_rawDescGZIP() []byte {
	file_notif_gateway_v1_usercode_usercode_proto_rawDescOnce.Do(func() {
		file_notif_gateway_v1_usercode_usercode_proto_rawDescData = protoimpl.X.CompressGZIP(file_notif_gateway_v1_usercode_usercode_proto_rawDescData)
	})
	return file_notif_gateway_v1_usercode_usercode_proto_rawDescData
}

var (
	file_notif_gateway_v1_usercode_usercode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_notif_gateway_v1_usercode_usercode_proto_goTypes  = []interface{}{
		(*SendCodeRequest)(nil),  // 0: notif.gateway.usercode.v1.SendCodeRequest
		(*SendCodeResponse)(nil), // 1: notif.gateway.usercode.v1.SendCodeResponse
		(v1.SignMethod)(0),       // 2: basetypes.v1.SignMethod
		(v1.UsedFor)(0),          // 3: basetypes.v1.UsedFor
	}
)
var file_notif_gateway_v1_usercode_usercode_proto_depIdxs = []int32{
	2, // 0: notif.gateway.usercode.v1.SendCodeRequest.AccountType:type_name -> basetypes.v1.SignMethod
	3, // 1: notif.gateway.usercode.v1.SendCodeRequest.UsedFor:type_name -> basetypes.v1.UsedFor
	0, // 2: notif.gateway.usercode.v1.Gateway.SendCode:input_type -> notif.gateway.usercode.v1.SendCodeRequest
	1, // 3: notif.gateway.usercode.v1.Gateway.SendCode:output_type -> notif.gateway.usercode.v1.SendCodeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notif_gateway_v1_usercode_usercode_proto_init() }
func file_notif_gateway_v1_usercode_usercode_proto_init() {
	if File_notif_gateway_v1_usercode_usercode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notif_gateway_v1_usercode_usercode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeRequest); i {
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
		file_notif_gateway_v1_usercode_usercode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeResponse); i {
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
	file_notif_gateway_v1_usercode_usercode_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notif_gateway_v1_usercode_usercode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notif_gateway_v1_usercode_usercode_proto_goTypes,
		DependencyIndexes: file_notif_gateway_v1_usercode_usercode_proto_depIdxs,
		MessageInfos:      file_notif_gateway_v1_usercode_usercode_proto_msgTypes,
	}.Build()
	File_notif_gateway_v1_usercode_usercode_proto = out.File
	file_notif_gateway_v1_usercode_usercode_proto_rawDesc = nil
	file_notif_gateway_v1_usercode_usercode_proto_goTypes = nil
	file_notif_gateway_v1_usercode_usercode_proto_depIdxs = nil
}
