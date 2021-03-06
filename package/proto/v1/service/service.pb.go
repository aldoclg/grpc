// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: service.proto

package greeting_v1

import (
	message "github.com/aldoclg/grpc/package/proto/v1/message"
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

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *message.Greet `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetMsg() *message.Greet {
	if x != nil {
		return x.Msg
	}
	return nil
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0c, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x23,
	0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72,
	0x65, 0x73, 0x70, 0x32, 0x52, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x19,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x64, 0x6f, 0x63, 0x6c, 0x67, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: greeting.v1.GreetRequest
	(*GreetResponse)(nil), // 1: greeting.v1.GreetResponse
	(*message.Greet)(nil), // 2: greet.v1.Greet
}
var file_service_proto_depIdxs = []int32{
	2, // 0: greeting.v1.GreetRequest.msg:type_name -> greet.v1.Greet
	0, // 1: greeting.v1.GreeterService.Greet:input_type -> greeting.v1.GreetRequest
	1, // 2: greeting.v1.GreeterService.Greet:output_type -> greeting.v1.GreetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
