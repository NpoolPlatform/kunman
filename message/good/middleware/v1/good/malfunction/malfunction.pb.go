// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: good/middleware/v1/good/malfunction/malfunction.proto

package malfunction

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	v11 "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MalfunctionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID             *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	GoodID            *string `protobuf:"bytes,30,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	Title             *string `protobuf:"bytes,40,opt,name=Title,proto3,oneof" json:"Title,omitempty"`
	Message           *string `protobuf:"bytes,50,opt,name=Message,proto3,oneof" json:"Message,omitempty"`
	StartAt           *uint32 `protobuf:"varint,60,opt,name=StartAt,proto3,oneof" json:"StartAt,omitempty"`
	DurationSeconds   *uint32 `protobuf:"varint,70,opt,name=DurationSeconds,proto3,oneof" json:"DurationSeconds,omitempty"`
	CompensateSeconds *uint32 `protobuf:"varint,80,opt,name=CompensateSeconds,proto3,oneof" json:"CompensateSeconds,omitempty"`
}

func (x *MalfunctionReq) Reset() {
	*x = MalfunctionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MalfunctionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MalfunctionReq) ProtoMessage() {}

func (x *MalfunctionReq) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MalfunctionReq.ProtoReflect.Descriptor instead.
func (*MalfunctionReq) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescGZIP(), []int{0}
}

func (x *MalfunctionReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *MalfunctionReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *MalfunctionReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *MalfunctionReq) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *MalfunctionReq) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *MalfunctionReq) GetStartAt() uint32 {
	if x != nil && x.StartAt != nil {
		return *x.StartAt
	}
	return 0
}

func (x *MalfunctionReq) GetDurationSeconds() uint32 {
	if x != nil && x.DurationSeconds != nil {
		return *x.DurationSeconds
	}
	return 0
}

func (x *MalfunctionReq) GetCompensateSeconds() uint32 {
	if x != nil && x.CompensateSeconds != nil {
		return *x.CompensateSeconds
	}
	return 0
}

type Malfunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"good_type"
	GoodTypeStr string      `protobuf:"bytes,40,opt,name=GoodTypeStr,proto3" json:"GoodTypeStr,omitempty" sql:"good_type"`
	GoodType    v1.GoodType `protobuf:"varint,50,opt,name=GoodType,proto3,enum=basetypes.good.v1.GoodType" json:"GoodType,omitempty"`
	// @inject_tag: sql:"good_name"
	GoodName string `protobuf:"bytes,60,opt,name=GoodName,proto3" json:"GoodName,omitempty" sql:"good_name"`
	// @inject_tag: sql:"title"
	Title string `protobuf:"bytes,70,opt,name=Title,proto3" json:"Title,omitempty" sql:"title"`
	// @inject_tag: sql:"message"
	Message string `protobuf:"bytes,80,opt,name=Message,proto3" json:"Message,omitempty" sql:"message"`
	// @inject_tag: sql:"start_at"
	StartAt uint32 `protobuf:"varint,90,opt,name=StartAt,proto3" json:"StartAt,omitempty" sql:"start_at"`
	// @inject_tag: sql:"duration_seconds"
	DurationSeconds uint32 `protobuf:"varint,100,opt,name=DurationSeconds,proto3" json:"DurationSeconds,omitempty" sql:"duration_seconds"`
	// @inject_tag: sql:"compensate_seconds"
	CompensateSeconds uint32 `protobuf:"varint,110,opt,name=CompensateSeconds,proto3" json:"CompensateSeconds,omitempty" sql:"compensate_seconds"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Malfunction) Reset() {
	*x = Malfunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Malfunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Malfunction) ProtoMessage() {}

func (x *Malfunction) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Malfunction.ProtoReflect.Descriptor instead.
func (*Malfunction) Descriptor() ([]byte, []int) {
	return file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescGZIP(), []int{1}
}

func (x *Malfunction) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Malfunction) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Malfunction) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Malfunction) GetGoodTypeStr() string {
	if x != nil {
		return x.GoodTypeStr
	}
	return ""
}

func (x *Malfunction) GetGoodType() v1.GoodType {
	if x != nil {
		return x.GoodType
	}
	return v1.GoodType(0)
}

func (x *Malfunction) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Malfunction) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Malfunction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Malfunction) GetStartAt() uint32 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Malfunction) GetDurationSeconds() uint32 {
	if x != nil {
		return x.DurationSeconds
	}
	return 0
}

func (x *Malfunction) GetCompensateSeconds() uint32 {
	if x != nil {
		return x.CompensateSeconds
	}
	return 0
}

func (x *Malfunction) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Malfunction) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        *v11.Uint32Val      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID     *v11.StringVal      `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	EntIDs    *v11.StringSliceVal `protobuf:"bytes,30,opt,name=EntIDs,proto3,oneof" json:"EntIDs,omitempty"`
	GoodID    *v11.StringVal      `protobuf:"bytes,40,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	AppGoodID *v11.StringVal      `protobuf:"bytes,50,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[2]
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
	return file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v11.Uint32Val {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetEntID() *v11.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetEntIDs() *v11.StringSliceVal {
	if x != nil {
		return x.EntIDs
	}
	return nil
}

func (x *Conds) GetGoodID() *v11.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetAppGoodID() *v11.StringVal {
	if x != nil {
		return x.AppGoodID
	}
	return nil
}

var File_good_middleware_v1_good_malfunction_malfunction_proto protoreflect.FileDescriptor

var file_good_middleware_v1_good_malfunction_malfunction_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x61, 0x6c, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6d,
	0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x0e, 0x4d, 0x61, 0x6c, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x05, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x2d, 0x0a, 0x0f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x0f, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x31,
	0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xa2, 0x03, 0x0a, 0x0b, 0x4d, 0x61,
	0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x6f,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x43,
	0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcb,
	0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03,
	0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f,
	0x6d, 0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescOnce sync.Once
	file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescData = file_good_middleware_v1_good_malfunction_malfunction_proto_rawDesc
)

func file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescGZIP() []byte {
	file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescOnce.Do(func() {
		file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescData = protoimpl.X.CompressGZIP(file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescData)
	})
	return file_good_middleware_v1_good_malfunction_malfunction_proto_rawDescData
}

var (
	file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_good_middleware_v1_good_malfunction_malfunction_proto_goTypes  = []interface{}{
		(*MalfunctionReq)(nil),     // 0: good.middleware.good1.malfunction.v1.MalfunctionReq
		(*Malfunction)(nil),        // 1: good.middleware.good1.malfunction.v1.Malfunction
		(*Conds)(nil),              // 2: good.middleware.good1.malfunction.v1.Conds
		(v1.GoodType)(0),           // 3: basetypes.good.v1.GoodType
		(*v11.Uint32Val)(nil),      // 4: basetypes.v1.Uint32Val
		(*v11.StringVal)(nil),      // 5: basetypes.v1.StringVal
		(*v11.StringSliceVal)(nil), // 6: basetypes.v1.StringSliceVal
	}
)
var file_good_middleware_v1_good_malfunction_malfunction_proto_depIdxs = []int32{
	3, // 0: good.middleware.good1.malfunction.v1.Malfunction.GoodType:type_name -> basetypes.good.v1.GoodType
	4, // 1: good.middleware.good1.malfunction.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	5, // 2: good.middleware.good1.malfunction.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	6, // 3: good.middleware.good1.malfunction.v1.Conds.EntIDs:type_name -> basetypes.v1.StringSliceVal
	5, // 4: good.middleware.good1.malfunction.v1.Conds.GoodID:type_name -> basetypes.v1.StringVal
	5, // 5: good.middleware.good1.malfunction.v1.Conds.AppGoodID:type_name -> basetypes.v1.StringVal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_good_middleware_v1_good_malfunction_malfunction_proto_init() }
func file_good_middleware_v1_good_malfunction_malfunction_proto_init() {
	if File_good_middleware_v1_good_malfunction_malfunction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MalfunctionReq); i {
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
		file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Malfunction); i {
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
		file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_good_middleware_v1_good_malfunction_malfunction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_good_middleware_v1_good_malfunction_malfunction_proto_goTypes,
		DependencyIndexes: file_good_middleware_v1_good_malfunction_malfunction_proto_depIdxs,
		MessageInfos:      file_good_middleware_v1_good_malfunction_malfunction_proto_msgTypes,
	}.Build()
	File_good_middleware_v1_good_malfunction_malfunction_proto = out.File
	file_good_middleware_v1_good_malfunction_malfunction_proto_rawDesc = nil
	file_good_middleware_v1_good_malfunction_malfunction_proto_goTypes = nil
	file_good_middleware_v1_good_malfunction_malfunction_proto_depIdxs = nil
}
