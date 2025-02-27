// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: protos/dynamodb.proto

package dynamodb_service

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

type StoreDataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreDataRequest) Reset() {
	*x = StoreDataRequest{}
	mi := &file_protos_dynamodb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDataRequest) ProtoMessage() {}

func (x *StoreDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dynamodb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDataRequest.ProtoReflect.Descriptor instead.
func (*StoreDataRequest) Descriptor() ([]byte, []int) {
	return file_protos_dynamodb_proto_rawDescGZIP(), []int{0}
}

func (x *StoreDataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoreDataRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type StoreDataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StoreDataResponse) Reset() {
	*x = StoreDataResponse{}
	mi := &file_protos_dynamodb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDataResponse) ProtoMessage() {}

func (x *StoreDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dynamodb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDataResponse.ProtoReflect.Descriptor instead.
func (*StoreDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_dynamodb_proto_rawDescGZIP(), []int{1}
}

func (x *StoreDataResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_protos_dynamodb_proto protoreflect.FileDescriptor

var file_protos_dynamodb_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64,
	0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0x67, 0x0a, 0x0f, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x22, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_dynamodb_proto_rawDescOnce sync.Once
	file_protos_dynamodb_proto_rawDescData = file_protos_dynamodb_proto_rawDesc
)

func file_protos_dynamodb_proto_rawDescGZIP() []byte {
	file_protos_dynamodb_proto_rawDescOnce.Do(func() {
		file_protos_dynamodb_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_dynamodb_proto_rawDescData)
	})
	return file_protos_dynamodb_proto_rawDescData
}

var file_protos_dynamodb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_dynamodb_proto_goTypes = []any{
	(*StoreDataRequest)(nil),  // 0: dynamodb_service.StoreDataRequest
	(*StoreDataResponse)(nil), // 1: dynamodb_service.StoreDataResponse
}
var file_protos_dynamodb_proto_depIdxs = []int32{
	0, // 0: dynamodb_service.DynamoDBService.StoreData:input_type -> dynamodb_service.StoreDataRequest
	1, // 1: dynamodb_service.DynamoDBService.StoreData:output_type -> dynamodb_service.StoreDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_dynamodb_proto_init() }
func file_protos_dynamodb_proto_init() {
	if File_protos_dynamodb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_dynamodb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_dynamodb_proto_goTypes,
		DependencyIndexes: file_protos_dynamodb_proto_depIdxs,
		MessageInfos:      file_protos_dynamodb_proto_msgTypes,
	}.Build()
	File_protos_dynamodb_proto = out.File
	file_protos_dynamodb_proto_rawDesc = nil
	file_protos_dynamodb_proto_goTypes = nil
	file_protos_dynamodb_proto_depIdxs = nil
}
