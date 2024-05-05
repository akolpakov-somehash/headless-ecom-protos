// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: sale/quote.proto

package salev1

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_sale_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_sale_quote_proto_rawDescGZIP(), []int{0}
}

type CustomerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CustomerId) Reset() {
	*x = CustomerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerId) ProtoMessage() {}

func (x *CustomerId) ProtoReflect() protoreflect.Message {
	mi := &file_sale_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerId.ProtoReflect.Descriptor instead.
func (*CustomerId) Descriptor() ([]byte, []int) {
	return file_sale_quote_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QuoteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32   `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *QuoteItem) Reset() {
	*x = QuoteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteItem) ProtoMessage() {}

func (x *QuoteItem) ProtoReflect() protoreflect.Message {
	mi := &file_sale_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteItem.ProtoReflect.Descriptor instead.
func (*QuoteItem) Descriptor() ([]byte, []int) {
	return file_sale_quote_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *QuoteItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *QuoteItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32        `protobuf:"varint,2,opt,name=customerId,proto3" json:"customerId,omitempty"`
	Items      []*QuoteItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_sale_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_sale_quote_proto_rawDescGZIP(), []int{3}
}

func (x *Quote) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Quote) GetItems() []*QuoteItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32 `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	ProductId  int32 `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity   int32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_quote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sale_quote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_sale_quote_proto_rawDescGZIP(), []int{4}
}

func (x *ProductRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ProductRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_sale_quote_proto protoreflect.FileDescriptor

var file_sale_quote_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x61, 0x6c, 0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x73, 0x61, 0x6c, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1c, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5b, 0x0a, 0x09, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x05,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6a, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xdb, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x73, 0x61, 0x6c,
	0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0d, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x61, 0x6c,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sale_quote_proto_rawDescOnce sync.Once
	file_sale_quote_proto_rawDescData = file_sale_quote_proto_rawDesc
)

func file_sale_quote_proto_rawDescGZIP() []byte {
	file_sale_quote_proto_rawDescOnce.Do(func() {
		file_sale_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_sale_quote_proto_rawDescData)
	})
	return file_sale_quote_proto_rawDescData
}

var file_sale_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sale_quote_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: sale.Empty
	(*CustomerId)(nil),     // 1: sale.CustomerId
	(*QuoteItem)(nil),      // 2: sale.QuoteItem
	(*Quote)(nil),          // 3: sale.Quote
	(*ProductRequest)(nil), // 4: sale.ProductRequest
}
var file_sale_quote_proto_depIdxs = []int32{
	2, // 0: sale.Quote.items:type_name -> sale.QuoteItem
	1, // 1: sale.QuoteService.GetQuote:input_type -> sale.CustomerId
	4, // 2: sale.QuoteService.AddProduct:input_type -> sale.ProductRequest
	4, // 3: sale.QuoteService.RemoveProduct:input_type -> sale.ProductRequest
	4, // 4: sale.QuoteService.UpdateQuantity:input_type -> sale.ProductRequest
	3, // 5: sale.QuoteService.GetQuote:output_type -> sale.Quote
	3, // 6: sale.QuoteService.AddProduct:output_type -> sale.Quote
	3, // 7: sale.QuoteService.RemoveProduct:output_type -> sale.Quote
	3, // 8: sale.QuoteService.UpdateQuantity:output_type -> sale.Quote
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sale_quote_proto_init() }
func file_sale_quote_proto_init() {
	if File_sale_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sale_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_sale_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerId); i {
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
		file_sale_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteItem); i {
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
		file_sale_quote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
		file_sale_quote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
			RawDescriptor: file_sale_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sale_quote_proto_goTypes,
		DependencyIndexes: file_sale_quote_proto_depIdxs,
		MessageInfos:      file_sale_quote_proto_msgTypes,
	}.Build()
	File_sale_quote_proto = out.File
	file_sale_quote_proto_rawDesc = nil
	file_sale_quote_proto_goTypes = nil
	file_sale_quote_proto_depIdxs = nil
}
