// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: chain/middleware/v1/coin/currency/currency.proto

package currency

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

type CurrencyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              *uint32              `protobuf:"varint,9,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID           *string              `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	CoinTypeID      *string              `protobuf:"bytes,20,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	FeedType        *v1.CurrencyFeedType `protobuf:"varint,30,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType,oneof" json:"FeedType,omitempty"`
	MarketValueHigh *string              `protobuf:"bytes,40,opt,name=MarketValueHigh,proto3,oneof" json:"MarketValueHigh,omitempty"`
	MarketValueLow  *string              `protobuf:"bytes,50,opt,name=MarketValueLow,proto3,oneof" json:"MarketValueLow,omitempty"`
}

func (x *CurrencyReq) Reset() {
	*x = CurrencyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReq) ProtoMessage() {}

func (x *CurrencyReq) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReq.ProtoReflect.Descriptor instead.
func (*CurrencyReq) Descriptor() ([]byte, []int) {
	return file_chain_middleware_v1_coin_currency_currency_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *CurrencyReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *CurrencyReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *CurrencyReq) GetFeedType() v1.CurrencyFeedType {
	if x != nil && x.FeedType != nil {
		return *x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *CurrencyReq) GetMarketValueHigh() string {
	if x != nil && x.MarketValueHigh != nil {
		return *x.MarketValueHigh
	}
	return ""
}

func (x *CurrencyReq) GetMarketValueLow() string {
	if x != nil && x.MarketValueLow != nil {
		return *x.MarketValueLow
	}
	return ""
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"coin_name"
	CoinName string `protobuf:"bytes,30,opt,name=CoinName,proto3" json:"CoinName,omitempty" sql:"coin_name"`
	// @inject_tag: sql:"coin_logo"
	CoinLogo string `protobuf:"bytes,40,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty" sql:"coin_logo"`
	// @inject_tag: sql:"coin_unit"
	CoinUnit string `protobuf:"bytes,50,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty" sql:"coin_unit"`
	// @inject_tag: sql:"coin_env"
	CoinENV string `protobuf:"bytes,60,opt,name=CoinENV,proto3" json:"CoinENV,omitempty" sql:"coin_env"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
	// @inject_tag: sql:"market_value_high"
	MarketValueHigh string `protobuf:"bytes,90,opt,name=MarketValueHigh,proto3" json:"MarketValueHigh,omitempty" sql:"market_value_high"`
	// @inject_tag: sql:"market_value_low"
	MarketValueLow string `protobuf:"bytes,100,opt,name=MarketValueLow,proto3" json:"MarketValueLow,omitempty" sql:"market_value_low"`
	// @inject_tag: sql:"feed_type"
	FeedTypeStr string              `protobuf:"bytes,110,opt,name=FeedTypeStr,proto3" json:"FeedTypeStr,omitempty" sql:"feed_type"`
	FeedType    v1.CurrencyFeedType `protobuf:"varint,120,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType" json:"FeedType,omitempty"`
	// @inject_tag: sql:"stable_usd"
	StableUSD bool `protobuf:"varint,130,opt,name=StableUSD,proto3" json:"StableUSD,omitempty" sql:"stable_usd"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_chain_middleware_v1_coin_currency_currency_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Currency) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Currency) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Currency) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Currency) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Currency) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Currency) GetCoinENV() string {
	if x != nil {
		return x.CoinENV
	}
	return ""
}

func (x *Currency) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Currency) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Currency) GetMarketValueHigh() string {
	if x != nil {
		return x.MarketValueHigh
	}
	return ""
}

func (x *Currency) GetMarketValueLow() string {
	if x != nil {
		return x.MarketValueLow
	}
	return ""
}

func (x *Currency) GetFeedTypeStr() string {
	if x != nil {
		return x.FeedTypeStr
	}
	return ""
}

func (x *Currency) GetFeedType() v1.CurrencyFeedType {
	if x != nil {
		return x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *Currency) GetStableUSD() bool {
	if x != nil {
		return x.StableUSD
	}
	return false
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID       *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	CoinTypeID  *v1.StringVal      `protobuf:"bytes,20,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	FeedType    *v1.Uint32Val      `protobuf:"bytes,30,opt,name=FeedType,proto3,oneof" json:"FeedType,omitempty"`
	CoinTypeIDs *v1.StringSliceVal `protobuf:"bytes,40,opt,name=CoinTypeIDs,proto3,oneof" json:"CoinTypeIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[2]
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
	return file_chain_middleware_v1_coin_currency_currency_proto_rawDescGZIP(), []int{2}
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

func (x *Conds) GetFeedType() *v1.Uint32Val {
	if x != nil {
		return x.FeedType
	}
	return nil
}

func (x *Conds) GetCoinTypeIDs() *v1.StringSliceVal {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

var File_chain_middleware_v1_coin_currency_currency_proto protoreflect.FileDescriptor

var file_chain_middleware_v1_coin_currency_currency_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x66, 0x65, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x65, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x03, 0x52, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x05, 0x52, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46, 0x65, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x22, 0xc9, 0x03, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67,
	0x68, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4c, 0x6f, 0x77, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x65, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x53, 0x44, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x55, 0x53, 0x44, 0x22, 0xae, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x01, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52,
	0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x0b,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48,
	0x03, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6b, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_middleware_v1_coin_currency_currency_proto_rawDescOnce sync.Once
	file_chain_middleware_v1_coin_currency_currency_proto_rawDescData = file_chain_middleware_v1_coin_currency_currency_proto_rawDesc
)

func file_chain_middleware_v1_coin_currency_currency_proto_rawDescGZIP() []byte {
	file_chain_middleware_v1_coin_currency_currency_proto_rawDescOnce.Do(func() {
		file_chain_middleware_v1_coin_currency_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_middleware_v1_coin_currency_currency_proto_rawDescData)
	})
	return file_chain_middleware_v1_coin_currency_currency_proto_rawDescData
}

var (
	file_chain_middleware_v1_coin_currency_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_chain_middleware_v1_coin_currency_currency_proto_goTypes  = []interface{}{
		(*CurrencyReq)(nil),       // 0: chain.middleware.coin.currency.v1.CurrencyReq
		(*Currency)(nil),          // 1: chain.middleware.coin.currency.v1.Currency
		(*Conds)(nil),             // 2: chain.middleware.coin.currency.v1.Conds
		(v1.CurrencyFeedType)(0),  // 3: basetypes.v1.CurrencyFeedType
		(*v1.StringVal)(nil),      // 4: basetypes.v1.StringVal
		(*v1.Uint32Val)(nil),      // 5: basetypes.v1.Uint32Val
		(*v1.StringSliceVal)(nil), // 6: basetypes.v1.StringSliceVal
	}
)
var file_chain_middleware_v1_coin_currency_currency_proto_depIdxs = []int32{
	3, // 0: chain.middleware.coin.currency.v1.CurrencyReq.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	3, // 1: chain.middleware.coin.currency.v1.Currency.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	4, // 2: chain.middleware.coin.currency.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	4, // 3: chain.middleware.coin.currency.v1.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	5, // 4: chain.middleware.coin.currency.v1.Conds.FeedType:type_name -> basetypes.v1.Uint32Val
	6, // 5: chain.middleware.coin.currency.v1.Conds.CoinTypeIDs:type_name -> basetypes.v1.StringSliceVal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chain_middleware_v1_coin_currency_currency_proto_init() }
func file_chain_middleware_v1_coin_currency_currency_proto_init() {
	if File_chain_middleware_v1_coin_currency_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReq); i {
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
		file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
		file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_chain_middleware_v1_coin_currency_currency_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chain_middleware_v1_coin_currency_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chain_middleware_v1_coin_currency_currency_proto_goTypes,
		DependencyIndexes: file_chain_middleware_v1_coin_currency_currency_proto_depIdxs,
		MessageInfos:      file_chain_middleware_v1_coin_currency_currency_proto_msgTypes,
	}.Build()
	File_chain_middleware_v1_coin_currency_currency_proto = out.File
	file_chain_middleware_v1_coin_currency_currency_proto_rawDesc = nil
	file_chain_middleware_v1_coin_currency_currency_proto_goTypes = nil
	file_chain_middleware_v1_coin_currency_currency_proto_depIdxs = nil
}
