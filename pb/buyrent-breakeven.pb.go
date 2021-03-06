// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: buyrent-breakeven.proto

package buyrent_breakeven

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyID        int64 `protobuf:"varint,1,opt,name=propertyID,proto3" json:"propertyID,omitempty"`
	TotalPropertyCost int64 `protobuf:"varint,2,opt,name=totalPropertyCost,proto3" json:"totalPropertyCost,omitempty"`
	StayTerm          int64 `protobuf:"varint,3,opt,name=stayTerm,proto3" json:"stayTerm,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyrent_breakeven_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_buyrent_breakeven_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_buyrent_breakeven_proto_rawDescGZIP(), []int{0}
}

func (x *Property) GetPropertyID() int64 {
	if x != nil {
		return x.PropertyID
	}
	return 0
}

func (x *Property) GetTotalPropertyCost() int64 {
	if x != nil {
		return x.TotalPropertyCost
	}
	return 0
}

func (x *Property) GetStayTerm() int64 {
	if x != nil {
		return x.StayTerm
	}
	return 0
}

type BuyProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownPayment           float64 `protobuf:"fixed64,1,opt,name=downPayment,proto3" json:"downPayment,omitempty"`
	InterestRate          float64 `protobuf:"fixed64,2,opt,name=interestRate,proto3" json:"interestRate,omitempty"`
	PropertyTaxes         float64 `protobuf:"fixed64,3,opt,name=propertyTaxes,proto3" json:"propertyTaxes,omitempty"`
	PropertyTransferTaxes float64 `protobuf:"fixed64,4,opt,name=propertyTransferTaxes,proto3" json:"propertyTransferTaxes,omitempty"`
	LoanTerm              int64   `protobuf:"varint,5,opt,name=loanTerm,proto3" json:"loanTerm,omitempty"`
}

func (x *BuyProperty) Reset() {
	*x = BuyProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyrent_breakeven_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyProperty) ProtoMessage() {}

func (x *BuyProperty) ProtoReflect() protoreflect.Message {
	mi := &file_buyrent_breakeven_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyProperty.ProtoReflect.Descriptor instead.
func (*BuyProperty) Descriptor() ([]byte, []int) {
	return file_buyrent_breakeven_proto_rawDescGZIP(), []int{1}
}

func (x *BuyProperty) GetDownPayment() float64 {
	if x != nil {
		return x.DownPayment
	}
	return 0
}

func (x *BuyProperty) GetInterestRate() float64 {
	if x != nil {
		return x.InterestRate
	}
	return 0
}

func (x *BuyProperty) GetPropertyTaxes() float64 {
	if x != nil {
		return x.PropertyTaxes
	}
	return 0
}

func (x *BuyProperty) GetPropertyTransferTaxes() float64 {
	if x != nil {
		return x.PropertyTransferTaxes
	}
	return 0
}

func (x *BuyProperty) GetLoanTerm() int64 {
	if x != nil {
		return x.LoanTerm
	}
	return 0
}

type RentProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonthlyRent int32 `protobuf:"varint,2,opt,name=monthlyRent,proto3" json:"monthlyRent,omitempty"`
}

func (x *RentProperty) Reset() {
	*x = RentProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyrent_breakeven_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RentProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RentProperty) ProtoMessage() {}

func (x *RentProperty) ProtoReflect() protoreflect.Message {
	mi := &file_buyrent_breakeven_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RentProperty.ProtoReflect.Descriptor instead.
func (*RentProperty) Descriptor() ([]byte, []int) {
	return file_buyrent_breakeven_proto_rawDescGZIP(), []int{2}
}

func (x *RentProperty) GetMonthlyRent() int32 {
	if x != nil {
		return x.MonthlyRent
	}
	return 0
}

type BreakEvenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P *Property     `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	B *BuyProperty  `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	R *RentProperty `protobuf:"bytes,3,opt,name=r,proto3" json:"r,omitempty"`
}

func (x *BreakEvenRequest) Reset() {
	*x = BreakEvenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyrent_breakeven_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakEvenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakEvenRequest) ProtoMessage() {}

func (x *BreakEvenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buyrent_breakeven_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakEvenRequest.ProtoReflect.Descriptor instead.
func (*BreakEvenRequest) Descriptor() ([]byte, []int) {
	return file_buyrent_breakeven_proto_rawDescGZIP(), []int{3}
}

func (x *BreakEvenRequest) GetP() *Property {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *BreakEvenRequest) GetB() *BuyProperty {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *BreakEvenRequest) GetR() *RentProperty {
	if x != nil {
		return x.R
	}
	return nil
}

type BreakEvenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result      string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	IsBreakEven bool   `protobuf:"varint,2,opt,name=isBreakEven,proto3" json:"isBreakEven,omitempty"`
}

func (x *BreakEvenResponse) Reset() {
	*x = BreakEvenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyrent_breakeven_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakEvenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakEvenResponse) ProtoMessage() {}

func (x *BreakEvenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buyrent_breakeven_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakEvenResponse.ProtoReflect.Descriptor instead.
func (*BreakEvenResponse) Descriptor() ([]byte, []int) {
	return file_buyrent_breakeven_proto_rawDescGZIP(), []int{4}
}

func (x *BreakEvenResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *BreakEvenResponse) GetIsBreakEven() bool {
	if x != nil {
		return x.IsBreakEven
	}
	return false
}

var File_buyrent_breakeven_proto protoreflect.FileDescriptor

var file_buyrent_breakeven_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x75, 0x79, 0x72, 0x65, 0x6e, 0x74, 0x2d, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65,
	0x76, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x79, 0x54, 0x65, 0x72, 0x6d, 0x22,
	0xcb, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x54, 0x61, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x61, 0x78, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54,
	0x61, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x61, 0x78, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x61, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x61, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x30, 0x0a,
	0x0c, 0x52, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x52, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x52, 0x65, 0x6e, 0x74, 0x22,
	0x64, 0x0a, 0x10, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x01, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x01, 0x70, 0x12, 0x1a, 0x0a, 0x01,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x75, 0x79, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x01, 0x62, 0x12, 0x1b, 0x0a, 0x01, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x01, 0x72, 0x22, 0x4d, 0x0a, 0x11, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x45, 0x76, 0x65, 0x6e, 0x32, 0x59, 0x0a, 0x10, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x1a, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x12, 0x11, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x45, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x3b, 0x62, 0x75, 0x79, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x65, 0x76, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buyrent_breakeven_proto_rawDescOnce sync.Once
	file_buyrent_breakeven_proto_rawDescData = file_buyrent_breakeven_proto_rawDesc
)

func file_buyrent_breakeven_proto_rawDescGZIP() []byte {
	file_buyrent_breakeven_proto_rawDescOnce.Do(func() {
		file_buyrent_breakeven_proto_rawDescData = protoimpl.X.CompressGZIP(file_buyrent_breakeven_proto_rawDescData)
	})
	return file_buyrent_breakeven_proto_rawDescData
}

var file_buyrent_breakeven_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_buyrent_breakeven_proto_goTypes = []interface{}{
	(*Property)(nil),          // 0: Property
	(*BuyProperty)(nil),       // 1: BuyProperty
	(*RentProperty)(nil),      // 2: RentProperty
	(*BreakEvenRequest)(nil),  // 3: BreakEvenRequest
	(*BreakEvenResponse)(nil), // 4: BreakEvenResponse
}
var file_buyrent_breakeven_proto_depIdxs = []int32{
	0, // 0: BreakEvenRequest.p:type_name -> Property
	1, // 1: BreakEvenRequest.b:type_name -> BuyProperty
	2, // 2: BreakEvenRequest.r:type_name -> RentProperty
	3, // 3: BreakEvenService.CalculatePropertyBreakEven:input_type -> BreakEvenRequest
	4, // 4: BreakEvenService.CalculatePropertyBreakEven:output_type -> BreakEvenResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buyrent_breakeven_proto_init() }
func file_buyrent_breakeven_proto_init() {
	if File_buyrent_breakeven_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buyrent_breakeven_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_buyrent_breakeven_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyProperty); i {
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
		file_buyrent_breakeven_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RentProperty); i {
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
		file_buyrent_breakeven_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakEvenRequest); i {
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
		file_buyrent_breakeven_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakEvenResponse); i {
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
			RawDescriptor: file_buyrent_breakeven_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buyrent_breakeven_proto_goTypes,
		DependencyIndexes: file_buyrent_breakeven_proto_depIdxs,
		MessageInfos:      file_buyrent_breakeven_proto_msgTypes,
	}.Build()
	File_buyrent_breakeven_proto = out.File
	file_buyrent_breakeven_proto_rawDesc = nil
	file_buyrent_breakeven_proto_goTypes = nil
	file_buyrent_breakeven_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BreakEvenServiceClient is the client API for BreakEvenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BreakEvenServiceClient interface {
	CalculatePropertyBreakEven(ctx context.Context, in *BreakEvenRequest, opts ...grpc.CallOption) (*BreakEvenResponse, error)
}

type breakEvenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBreakEvenServiceClient(cc grpc.ClientConnInterface) BreakEvenServiceClient {
	return &breakEvenServiceClient{cc}
}

func (c *breakEvenServiceClient) CalculatePropertyBreakEven(ctx context.Context, in *BreakEvenRequest, opts ...grpc.CallOption) (*BreakEvenResponse, error) {
	out := new(BreakEvenResponse)
	err := c.cc.Invoke(ctx, "/BreakEvenService/CalculatePropertyBreakEven", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BreakEvenServiceServer is the server API for BreakEvenService service.
type BreakEvenServiceServer interface {
	CalculatePropertyBreakEven(context.Context, *BreakEvenRequest) (*BreakEvenResponse, error)
}

// UnimplementedBreakEvenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBreakEvenServiceServer struct {
}

func (*UnimplementedBreakEvenServiceServer) CalculatePropertyBreakEven(context.Context, *BreakEvenRequest) (*BreakEvenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculatePropertyBreakEven not implemented")
}

func RegisterBreakEvenServiceServer(s *grpc.Server, srv BreakEvenServiceServer) {
	s.RegisterService(&_BreakEvenService_serviceDesc, srv)
}

func _BreakEvenService_CalculatePropertyBreakEven_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BreakEvenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreakEvenServiceServer).CalculatePropertyBreakEven(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BreakEvenService/CalculatePropertyBreakEven",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreakEvenServiceServer).CalculatePropertyBreakEven(ctx, req.(*BreakEvenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BreakEvenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BreakEvenService",
	HandlerType: (*BreakEvenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculatePropertyBreakEven",
			Handler:    _BreakEvenService_CalculatePropertyBreakEven_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buyrent-breakeven.proto",
}
