// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: component_metadata_server.proto

package component_metadata_server

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

type MetadataType int32

const (
	MetadataType_METADATA_TYPE_PARAMETER MetadataType = 0 // Parameter metadata
	MetadataType_METADATA_TYPE_EVENTS    MetadataType = 1 // Event definitions
	MetadataType_METADATA_TYPE_ACTUATORS MetadataType = 2 // Actuator definitions
)

// Enum value maps for MetadataType.
var (
	MetadataType_name = map[int32]string{
		0: "METADATA_TYPE_PARAMETER",
		1: "METADATA_TYPE_EVENTS",
		2: "METADATA_TYPE_ACTUATORS",
	}
	MetadataType_value = map[string]int32{
		"METADATA_TYPE_PARAMETER": 0,
		"METADATA_TYPE_EVENTS":    1,
		"METADATA_TYPE_ACTUATORS": 2,
	}
)

func (x MetadataType) Enum() *MetadataType {
	p := new(MetadataType)
	*p = x
	return p
}

func (x MetadataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataType) Descriptor() protoreflect.EnumDescriptor {
	return file_component_metadata_server_proto_enumTypes[0].Descriptor()
}

func (MetadataType) Type() protoreflect.EnumType {
	return &file_component_metadata_server_proto_enumTypes[0]
}

func (x MetadataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataType.Descriptor instead.
func (MetadataType) EnumDescriptor() ([]byte, []int) {
	return file_component_metadata_server_proto_rawDescGZIP(), []int{0}
}

type SetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*Metadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"` // List of metadata
}

func (x *SetMetadataRequest) Reset() {
	*x = SetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_metadata_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetadataRequest) ProtoMessage() {}

func (x *SetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_component_metadata_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetadataRequest.ProtoReflect.Descriptor instead.
func (*SetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_component_metadata_server_proto_rawDescGZIP(), []int{0}
}

func (x *SetMetadataRequest) GetMetadata() []*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetMetadataResponse) Reset() {
	*x = SetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_metadata_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetadataResponse) ProtoMessage() {}

func (x *SetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_component_metadata_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetadataResponse.ProtoReflect.Descriptor instead.
func (*SetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_component_metadata_server_proto_rawDescGZIP(), []int{1}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         MetadataType `protobuf:"varint,1,opt,name=type,proto3,enum=mavsdk.rpc.component_metadata_server.MetadataType" json:"type,omitempty"` // The metadata type
	JsonMetadata string       `protobuf:"bytes,2,opt,name=json_metadata,json=jsonMetadata,proto3" json:"json_metadata,omitempty"`                     // The JSON metadata
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_component_metadata_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_component_metadata_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_component_metadata_server_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetType() MetadataType {
	if x != nil {
		return x.Type
	}
	return MetadataType_METADATA_TYPE_PARAMETER
}

func (x *Metadata) GetJsonMetadata() string {
	if x != nil {
		return x.JsonMetadata
	}
	return ""
}

var File_component_metadata_server_proto protoreflect.FileDescriptor

var file_component_metadata_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x24, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6a, 0x73,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2a,
	0x62, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x55, 0x41, 0x54, 0x4f, 0x52,
	0x53, 0x10, 0x02, 0x32, 0xab, 0x01, 0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x80, 0xb5, 0x18,
	0x01, 0x42, 0x3b, 0x42, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x1b, 0x2e, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_component_metadata_server_proto_rawDescOnce sync.Once
	file_component_metadata_server_proto_rawDescData = file_component_metadata_server_proto_rawDesc
)

func file_component_metadata_server_proto_rawDescGZIP() []byte {
	file_component_metadata_server_proto_rawDescOnce.Do(func() {
		file_component_metadata_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_component_metadata_server_proto_rawDescData)
	})
	return file_component_metadata_server_proto_rawDescData
}

var file_component_metadata_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_component_metadata_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_component_metadata_server_proto_goTypes = []any{
	(MetadataType)(0),           // 0: mavsdk.rpc.component_metadata_server.MetadataType
	(*SetMetadataRequest)(nil),  // 1: mavsdk.rpc.component_metadata_server.SetMetadataRequest
	(*SetMetadataResponse)(nil), // 2: mavsdk.rpc.component_metadata_server.SetMetadataResponse
	(*Metadata)(nil),            // 3: mavsdk.rpc.component_metadata_server.Metadata
}
var file_component_metadata_server_proto_depIdxs = []int32{
	3, // 0: mavsdk.rpc.component_metadata_server.SetMetadataRequest.metadata:type_name -> mavsdk.rpc.component_metadata_server.Metadata
	0, // 1: mavsdk.rpc.component_metadata_server.Metadata.type:type_name -> mavsdk.rpc.component_metadata_server.MetadataType
	1, // 2: mavsdk.rpc.component_metadata_server.ComponentMetadataServerService.SetMetadata:input_type -> mavsdk.rpc.component_metadata_server.SetMetadataRequest
	2, // 3: mavsdk.rpc.component_metadata_server.ComponentMetadataServerService.SetMetadata:output_type -> mavsdk.rpc.component_metadata_server.SetMetadataResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_component_metadata_server_proto_init() }
func file_component_metadata_server_proto_init() {
	if File_component_metadata_server_proto != nil {
		return
	}
	
	if !protoimpl.UnsafeEnabled {
		file_component_metadata_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetMetadataRequest); i {
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
		file_component_metadata_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetMetadataResponse); i {
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
		file_component_metadata_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_component_metadata_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_component_metadata_server_proto_goTypes,
		DependencyIndexes: file_component_metadata_server_proto_depIdxs,
		EnumInfos:         file_component_metadata_server_proto_enumTypes,
		MessageInfos:      file_component_metadata_server_proto_msgTypes,
	}.Build()
	File_component_metadata_server_proto = out.File
	file_component_metadata_server_proto_rawDesc = nil
	file_component_metadata_server_proto_goTypes = nil
	file_component_metadata_server_proto_depIdxs = nil
}