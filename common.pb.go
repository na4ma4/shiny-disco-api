// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: github.com/na4ma4/shiny-disco-api/common.proto

package shiny_disco_api

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*Message_ServiceRegister
	//	*Message_ServiceDeregister
	Message isMessage_Message `protobuf_oneof:"message"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetMessage() isMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Message) GetServiceRegister() *ServiceRegister {
	if x, ok := x.GetMessage().(*Message_ServiceRegister); ok {
		return x.ServiceRegister
	}
	return nil
}

func (x *Message) GetServiceDeregister() *ServiceDeregister {
	if x, ok := x.GetMessage().(*Message_ServiceDeregister); ok {
		return x.ServiceDeregister
	}
	return nil
}

type isMessage_Message interface {
	isMessage_Message()
}

type Message_ServiceRegister struct {
	ServiceRegister *ServiceRegister `protobuf:"bytes,100,opt,name=service_register,json=serviceRegister,proto3,oneof"`
}

type Message_ServiceDeregister struct {
	ServiceDeregister *ServiceDeregister `protobuf:"bytes,101,opt,name=service_deregister,json=serviceDeregister,proto3,oneof"`
}

func (*Message_ServiceRegister) isMessage_Message() {}

func (*Message_ServiceDeregister) isMessage_Message() {}

type ServiceRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag     string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ServiceRegister) Reset() {
	*x = ServiceRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRegister) ProtoMessage() {}

func (x *ServiceRegister) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRegister.ProtoReflect.Descriptor instead.
func (*ServiceRegister) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceRegister) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ServiceRegister) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ServiceDeregister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceTag []string `protobuf:"bytes,1,rep,name=service_tag,json=serviceTag,proto3" json:"service_tag,omitempty"`
}

func (x *ServiceDeregister) Reset() {
	*x = ServiceDeregister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDeregister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDeregister) ProtoMessage() {}

func (x *ServiceDeregister) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDeregister.ProtoReflect.Descriptor instead.
func (*ServiceDeregister) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceDeregister) GetServiceTag() []string {
	if x != nil {
		return x.ServiceTag
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload  string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescGZIP(), []int{3}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Service) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_github_com_na4ma4_shiny_disco_api_common_proto protoreflect.FileDescriptor

var file_github_com_na4ma4_shiny_disco_api_common_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x34,
	0x6d, 0x61, 0x34, 0x2f, 0x73, 0x68, 0x69, 0x6e, 0x79, 0x2d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x68, 0x69, 0x6e, 0x79, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4f, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x69, 0x6e,
	0x79, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x55, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73,
	0x68, 0x69, 0x6e, 0x79, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x34, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x67, 0x22, 0x4f, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x34, 0x6d, 0x61, 0x34, 0x2f, 0x73, 0x68,
	0x69, 0x6e, 0x79, 0x2d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x73, 0x68,
	0x69, 0x6e, 0x79, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescOnce sync.Once
	file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescData = file_github_com_na4ma4_shiny_disco_api_common_proto_rawDesc
)

func file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescGZIP() []byte {
	file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescOnce.Do(func() {
		file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescData)
	})
	return file_github_com_na4ma4_shiny_disco_api_common_proto_rawDescData
}

var file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_na4ma4_shiny_disco_api_common_proto_goTypes = []interface{}{
	(*Message)(nil),           // 0: shinydisco.api.v1.Message
	(*ServiceRegister)(nil),   // 1: shinydisco.api.v1.ServiceRegister
	(*ServiceDeregister)(nil), // 2: shinydisco.api.v1.ServiceDeregister
	(*Service)(nil),           // 3: shinydisco.api.v1.Service
}
var file_github_com_na4ma4_shiny_disco_api_common_proto_depIdxs = []int32{
	1, // 0: shinydisco.api.v1.Message.service_register:type_name -> shinydisco.api.v1.ServiceRegister
	2, // 1: shinydisco.api.v1.Message.service_deregister:type_name -> shinydisco.api.v1.ServiceDeregister
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_na4ma4_shiny_disco_api_common_proto_init() }
func file_github_com_na4ma4_shiny_disco_api_common_proto_init() {
	if File_github_com_na4ma4_shiny_disco_api_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRegister); i {
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
		file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDeregister); i {
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
		file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
	file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_ServiceRegister)(nil),
		(*Message_ServiceDeregister)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_na4ma4_shiny_disco_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_na4ma4_shiny_disco_api_common_proto_goTypes,
		DependencyIndexes: file_github_com_na4ma4_shiny_disco_api_common_proto_depIdxs,
		MessageInfos:      file_github_com_na4ma4_shiny_disco_api_common_proto_msgTypes,
	}.Build()
	File_github_com_na4ma4_shiny_disco_api_common_proto = out.File
	file_github_com_na4ma4_shiny_disco_api_common_proto_rawDesc = nil
	file_github_com_na4ma4_shiny_disco_api_common_proto_goTypes = nil
	file_github_com_na4ma4_shiny_disco_api_common_proto_depIdxs = nil
}
