// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: agi/gateway/v1/subscription/subscription.proto

package subscription

import (
	subscription "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription"
	_ "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  *string `protobuf:"bytes,10,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID *string `protobuf:"bytes,20,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	EntID  *string `protobuf:"bytes,30,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
}

func (x *GetSubscriptionRequest) Reset() {
	*x = GetSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionRequest) ProtoMessage() {}

func (x *GetSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *GetSubscriptionRequest) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *GetSubscriptionRequest) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *GetSubscriptionRequest) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

type GetSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *subscription.Subscription `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetSubscriptionResponse) Reset() {
	*x = GetSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionResponse) ProtoMessage() {}

func (x *GetSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *GetSubscriptionResponse) GetInfo() *subscription.Subscription {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID *string `protobuf:"bytes,20,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	Offset    int32   `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32   `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetSubscriptionsRequest) Reset() {
	*x = GetSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsRequest) ProtoMessage() {}

func (x *GetSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubscriptionsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetSubscriptionsRequest) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

func (x *GetSubscriptionsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetSubscriptionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*subscription.Subscription `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetSubscriptionsResponse) Reset() {
	*x = GetSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionsResponse) ProtoMessage() {}

func (x *GetSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubscriptionsResponse) GetInfos() []*subscription.Subscription {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CountSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID *string `protobuf:"bytes,20,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
}

func (x *CountSubscriptionsRequest) Reset() {
	*x = CountSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountSubscriptionsRequest) ProtoMessage() {}

func (x *CountSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*CountSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *CountSubscriptionsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CountSubscriptionsRequest) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

type CountSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32 `protobuf:"varint,10,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *CountSubscriptionsResponse) Reset() {
	*x = CountSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountSubscriptionsResponse) ProtoMessage() {}

func (x *CountSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*CountSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *CountSubscriptionsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID *string `protobuf:"bytes,10,opt,name=TargetAppID,proto3,oneof" json:"TargetAppID,omitempty"`
	Offset      int32   `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32   `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetSubscriptionsRequest) Reset() {
	*x = AdminGetSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetSubscriptionsRequest) ProtoMessage() {}

func (x *AdminGetSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *AdminGetSubscriptionsRequest) GetTargetAppID() string {
	if x != nil && x.TargetAppID != nil {
		return *x.TargetAppID
	}
	return ""
}

func (x *AdminGetSubscriptionsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetSubscriptionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*subscription.Subscription `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *AdminGetSubscriptionsResponse) Reset() {
	*x = AdminGetSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetSubscriptionsResponse) ProtoMessage() {}

func (x *AdminGetSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *AdminGetSubscriptionsResponse) GetInfos() []*subscription.Subscription {
	if x != nil {
		return x.Infos
	}
	return nil
}

type AdminCountSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID *string `protobuf:"bytes,10,opt,name=TargetAppID,proto3,oneof" json:"TargetAppID,omitempty"`
}

func (x *AdminCountSubscriptionsRequest) Reset() {
	*x = AdminCountSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCountSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCountSubscriptionsRequest) ProtoMessage() {}

func (x *AdminCountSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCountSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*AdminCountSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{8}
}

func (x *AdminCountSubscriptionsRequest) GetTargetAppID() string {
	if x != nil && x.TargetAppID != nil {
		return *x.TargetAppID
	}
	return ""
}

type AdminCountSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32 `protobuf:"varint,10,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminCountSubscriptionsResponse) Reset() {
	*x = AdminCountSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCountSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCountSubscriptionsResponse) ProtoMessage() {}

func (x *AdminCountSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agi_gateway_v1_subscription_subscription_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCountSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*AdminCountSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP(), []int{9}
}

func (x *AdminCountSubscriptionsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_agi_gateway_v1_subscription_subscription_proto protoreflect.FileDescriptor

var file_agi_gateway_v1_subscription_subscription_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x67, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x31, 0x61,
	0x67, 0x69, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x67, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x5b, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47,
	0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x62, 0x0a, 0x19, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x1a, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x83, 0x01, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x1d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x57, 0x0a, 0x1e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x22, 0x37, 0x0a, 0x1f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf3, 0x06, 0x0a,
	0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x9d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x61,
	0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa1, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e,
	0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xa9, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x36, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x67,
	0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22,
	0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xb6, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x39, 0x2e, 0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0xbe, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x2e,
	0x61, 0x67, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x61, 0x67, 0x69,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b,
	0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x67,
	0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agi_gateway_v1_subscription_subscription_proto_rawDescOnce sync.Once
	file_agi_gateway_v1_subscription_subscription_proto_rawDescData = file_agi_gateway_v1_subscription_subscription_proto_rawDesc
)

func file_agi_gateway_v1_subscription_subscription_proto_rawDescGZIP() []byte {
	file_agi_gateway_v1_subscription_subscription_proto_rawDescOnce.Do(func() {
		file_agi_gateway_v1_subscription_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_agi_gateway_v1_subscription_subscription_proto_rawDescData)
	})
	return file_agi_gateway_v1_subscription_subscription_proto_rawDescData
}

var file_agi_gateway_v1_subscription_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_agi_gateway_v1_subscription_subscription_proto_goTypes = []interface{}{
	(*GetSubscriptionRequest)(nil),          // 0: agi.gateway.subscription.v1.GetSubscriptionRequest
	(*GetSubscriptionResponse)(nil),         // 1: agi.gateway.subscription.v1.GetSubscriptionResponse
	(*GetSubscriptionsRequest)(nil),         // 2: agi.gateway.subscription.v1.GetSubscriptionsRequest
	(*GetSubscriptionsResponse)(nil),        // 3: agi.gateway.subscription.v1.GetSubscriptionsResponse
	(*CountSubscriptionsRequest)(nil),       // 4: agi.gateway.subscription.v1.CountSubscriptionsRequest
	(*CountSubscriptionsResponse)(nil),      // 5: agi.gateway.subscription.v1.CountSubscriptionsResponse
	(*AdminGetSubscriptionsRequest)(nil),    // 6: agi.gateway.subscription.v1.AdminGetSubscriptionsRequest
	(*AdminGetSubscriptionsResponse)(nil),   // 7: agi.gateway.subscription.v1.AdminGetSubscriptionsResponse
	(*AdminCountSubscriptionsRequest)(nil),  // 8: agi.gateway.subscription.v1.AdminCountSubscriptionsRequest
	(*AdminCountSubscriptionsResponse)(nil), // 9: agi.gateway.subscription.v1.AdminCountSubscriptionsResponse
	(*subscription.Subscription)(nil),       // 10: agi.middleware.subscription.v1.Subscription
}
var file_agi_gateway_v1_subscription_subscription_proto_depIdxs = []int32{
	10, // 0: agi.gateway.subscription.v1.GetSubscriptionResponse.Info:type_name -> agi.middleware.subscription.v1.Subscription
	10, // 1: agi.gateway.subscription.v1.GetSubscriptionsResponse.Infos:type_name -> agi.middleware.subscription.v1.Subscription
	10, // 2: agi.gateway.subscription.v1.AdminGetSubscriptionsResponse.Infos:type_name -> agi.middleware.subscription.v1.Subscription
	0,  // 3: agi.gateway.subscription.v1.Gateway.GetSubscription:input_type -> agi.gateway.subscription.v1.GetSubscriptionRequest
	2,  // 4: agi.gateway.subscription.v1.Gateway.GetSubscriptions:input_type -> agi.gateway.subscription.v1.GetSubscriptionsRequest
	4,  // 5: agi.gateway.subscription.v1.Gateway.CountSubscriptions:input_type -> agi.gateway.subscription.v1.CountSubscriptionsRequest
	6,  // 6: agi.gateway.subscription.v1.Gateway.AdminGetSubscriptions:input_type -> agi.gateway.subscription.v1.AdminGetSubscriptionsRequest
	8,  // 7: agi.gateway.subscription.v1.Gateway.AdminCountSubscriptions:input_type -> agi.gateway.subscription.v1.AdminCountSubscriptionsRequest
	1,  // 8: agi.gateway.subscription.v1.Gateway.GetSubscription:output_type -> agi.gateway.subscription.v1.GetSubscriptionResponse
	3,  // 9: agi.gateway.subscription.v1.Gateway.GetSubscriptions:output_type -> agi.gateway.subscription.v1.GetSubscriptionsResponse
	5,  // 10: agi.gateway.subscription.v1.Gateway.CountSubscriptions:output_type -> agi.gateway.subscription.v1.CountSubscriptionsResponse
	7,  // 11: agi.gateway.subscription.v1.Gateway.AdminGetSubscriptions:output_type -> agi.gateway.subscription.v1.AdminGetSubscriptionsResponse
	9,  // 12: agi.gateway.subscription.v1.Gateway.AdminCountSubscriptions:output_type -> agi.gateway.subscription.v1.AdminCountSubscriptionsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_agi_gateway_v1_subscription_subscription_proto_init() }
func file_agi_gateway_v1_subscription_subscription_proto_init() {
	if File_agi_gateway_v1_subscription_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionRequest); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionResponse); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionsRequest); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionsResponse); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountSubscriptionsRequest); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountSubscriptionsResponse); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetSubscriptionsRequest); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetSubscriptionsResponse); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCountSubscriptionsRequest); i {
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
		file_agi_gateway_v1_subscription_subscription_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCountSubscriptionsResponse); i {
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
	file_agi_gateway_v1_subscription_subscription_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_agi_gateway_v1_subscription_subscription_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_agi_gateway_v1_subscription_subscription_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_agi_gateway_v1_subscription_subscription_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_agi_gateway_v1_subscription_subscription_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agi_gateway_v1_subscription_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agi_gateway_v1_subscription_subscription_proto_goTypes,
		DependencyIndexes: file_agi_gateway_v1_subscription_subscription_proto_depIdxs,
		MessageInfos:      file_agi_gateway_v1_subscription_subscription_proto_msgTypes,
	}.Build()
	File_agi_gateway_v1_subscription_subscription_proto = out.File
	file_agi_gateway_v1_subscription_subscription_proto_rawDesc = nil
	file_agi_gateway_v1_subscription_subscription_proto_goTypes = nil
	file_agi_gateway_v1_subscription_subscription_proto_depIdxs = nil
}
