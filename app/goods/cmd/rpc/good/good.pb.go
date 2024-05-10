// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: good.proto

package good

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

type GoodData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodId  string `protobuf:"bytes,1,opt,name=good_id,json=goodId,proto3" json:"good_id,omitempty"`
	Preface string `protobuf:"bytes,2,opt,name=preface,proto3" json:"preface,omitempty"`
	Img     string `protobuf:"bytes,3,opt,name=img,proto3" json:"img,omitempty"`
	Price   string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GoodData) Reset() {
	*x = GoodData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodData) ProtoMessage() {}

func (x *GoodData) ProtoReflect() protoreflect.Message {
	mi := &file_good_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodData.ProtoReflect.Descriptor instead.
func (*GoodData) Descriptor() ([]byte, []int) {
	return file_good_proto_rawDescGZIP(), []int{0}
}

func (x *GoodData) GetGoodId() string {
	if x != nil {
		return x.GoodId
	}
	return ""
}

func (x *GoodData) GetPreface() string {
	if x != nil {
		return x.Preface
	}
	return ""
}

func (x *GoodData) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *GoodData) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type CreateGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img      string `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	Preface  string `protobuf:"bytes,2,opt,name=preface,proto3" json:"preface,omitempty"`
	Price    string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	SellerId string `protobuf:"bytes,4,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
}

func (x *CreateGoodRequest) Reset() {
	*x = CreateGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodRequest) ProtoMessage() {}

func (x *CreateGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_good_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodRequest) Descriptor() ([]byte, []int) {
	return file_good_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGoodRequest) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *CreateGoodRequest) GetPreface() string {
	if x != nil {
		return x.Preface
	}
	return ""
}

func (x *CreateGoodRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *CreateGoodRequest) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

type CreateGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateGoodResponse) Reset() {
	*x = CreateGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodResponse) ProtoMessage() {}

func (x *CreateGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_good_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodResponse.ProtoReflect.Descriptor instead.
func (*CreateGoodResponse) Descriptor() ([]byte, []int) {
	return file_good_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGoodResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateGoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 商品的分页加载
type FindGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindGoodRequest) Reset() {
	*x = FindGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGoodRequest) ProtoMessage() {}

func (x *FindGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_good_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGoodRequest.ProtoReflect.Descriptor instead.
func (*FindGoodRequest) Descriptor() ([]byte, []int) {
	return file_good_proto_rawDescGZIP(), []int{3}
}

func (x *FindGoodRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindGoodRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*GoodData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindGoodResponse) Reset() {
	*x = FindGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_good_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGoodResponse) ProtoMessage() {}

func (x *FindGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_good_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGoodResponse.ProtoReflect.Descriptor instead.
func (*FindGoodResponse) Descriptor() ([]byte, []int) {
	return file_good_proto_rawDescGZIP(), []int{4}
}

func (x *FindGoodResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FindGoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FindGoodResponse) GetData() []*GoodData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_good_proto protoreflect.FileDescriptor

var file_good_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x6f,
	0x6f, 0x64, 0x22, 0x65, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x66, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6d, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x72, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x39, 0x0a, 0x0f, 0x46, 0x69, 0x6e,
	0x64, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x5c, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x22,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x86, 0x01, 0x0a, 0x04, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x46, 0x69, 0x6e, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_good_proto_rawDescOnce sync.Once
	file_good_proto_rawDescData = file_good_proto_rawDesc
)

func file_good_proto_rawDescGZIP() []byte {
	file_good_proto_rawDescOnce.Do(func() {
		file_good_proto_rawDescData = protoimpl.X.CompressGZIP(file_good_proto_rawDescData)
	})
	return file_good_proto_rawDescData
}

var file_good_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_good_proto_goTypes = []interface{}{
	(*GoodData)(nil),           // 0: good.GoodData
	(*CreateGoodRequest)(nil),  // 1: good.CreateGoodRequest
	(*CreateGoodResponse)(nil), // 2: good.CreateGoodResponse
	(*FindGoodRequest)(nil),    // 3: good.FindGoodRequest
	(*FindGoodResponse)(nil),   // 4: good.FindGoodResponse
}
var file_good_proto_depIdxs = []int32{
	0, // 0: good.FindGoodResponse.data:type_name -> good.GoodData
	1, // 1: good.Good.CreateGood:input_type -> good.CreateGoodRequest
	3, // 2: good.Good.FindGoodPage:input_type -> good.FindGoodRequest
	2, // 3: good.Good.CreateGood:output_type -> good.CreateGoodResponse
	4, // 4: good.Good.FindGoodPage:output_type -> good.FindGoodResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_good_proto_init() }
func file_good_proto_init() {
	if File_good_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_good_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodData); i {
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
		file_good_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodRequest); i {
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
		file_good_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodResponse); i {
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
		file_good_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGoodRequest); i {
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
		file_good_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGoodResponse); i {
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
			RawDescriptor: file_good_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_good_proto_goTypes,
		DependencyIndexes: file_good_proto_depIdxs,
		MessageInfos:      file_good_proto_msgTypes,
	}.Build()
	File_good_proto = out.File
	file_good_proto_rawDesc = nil
	file_good_proto_goTypes = nil
	file_good_proto_depIdxs = nil
}
