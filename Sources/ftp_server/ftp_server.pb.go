// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: ftp_server.proto

package ftp_server

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

// Possible results returned for FTP server requests.
type FtpServerResult_Result int32

const (
	FtpServerResult_RESULT_UNKNOWN        FtpServerResult_Result = 0 // Unknown result
	FtpServerResult_RESULT_SUCCESS        FtpServerResult_Result = 1 // Request succeeded
	FtpServerResult_RESULT_DOES_NOT_EXIST FtpServerResult_Result = 2 // Directory does not exist
	FtpServerResult_RESULT_BUSY           FtpServerResult_Result = 3 // Operations in progress
)

// Enum value maps for FtpServerResult_Result.
var (
	FtpServerResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_DOES_NOT_EXIST",
		3: "RESULT_BUSY",
	}
	FtpServerResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":        0,
		"RESULT_SUCCESS":        1,
		"RESULT_DOES_NOT_EXIST": 2,
		"RESULT_BUSY":           3,
	}
)

func (x FtpServerResult_Result) Enum() *FtpServerResult_Result {
	p := new(FtpServerResult_Result)
	*p = x
	return p
}

func (x FtpServerResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FtpServerResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_ftp_server_proto_enumTypes[0].Descriptor()
}

func (FtpServerResult_Result) Type() protoreflect.EnumType {
	return &file_ftp_server_proto_enumTypes[0]
}

func (x FtpServerResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FtpServerResult_Result.Descriptor instead.
func (FtpServerResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_ftp_server_proto_rawDescGZIP(), []int{2, 0}
}

type SetRootDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // Absolute path of folder
}

func (x *SetRootDirRequest) Reset() {
	*x = SetRootDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRootDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRootDirRequest) ProtoMessage() {}

func (x *SetRootDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRootDirRequest.ProtoReflect.Descriptor instead.
func (*SetRootDirRequest) Descriptor() ([]byte, []int) {
	return file_ftp_server_proto_rawDescGZIP(), []int{0}
}

func (x *SetRootDirRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type SetRootDirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FtpServerResult *FtpServerResult `protobuf:"bytes,1,opt,name=ftp_server_result,json=ftpServerResult,proto3" json:"ftp_server_result,omitempty"`
}

func (x *SetRootDirResponse) Reset() {
	*x = SetRootDirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRootDirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRootDirResponse) ProtoMessage() {}

func (x *SetRootDirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRootDirResponse.ProtoReflect.Descriptor instead.
func (*SetRootDirResponse) Descriptor() ([]byte, []int) {
	return file_ftp_server_proto_rawDescGZIP(), []int{1}
}

func (x *SetRootDirResponse) GetFtpServerResult() *FtpServerResult {
	if x != nil {
		return x.FtpServerResult
	}
	return nil
}

// Result type.
type FtpServerResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    FtpServerResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.ftp_server.FtpServerResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string                 `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                             // Human-readable English string describing the result
}

func (x *FtpServerResult) Reset() {
	*x = FtpServerResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ftp_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FtpServerResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FtpServerResult) ProtoMessage() {}

func (x *FtpServerResult) ProtoReflect() protoreflect.Message {
	mi := &file_ftp_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FtpServerResult.ProtoReflect.Descriptor instead.
func (*FtpServerResult) Descriptor() ([]byte, []int) {
	return file_ftp_server_proto_rawDescGZIP(), []int{2}
}

func (x *FtpServerResult) GetResult() FtpServerResult_Result {
	if x != nil {
		return x.Result
	}
	return FtpServerResult_RESULT_UNKNOWN
}

func (x *FtpServerResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_ftp_server_proto protoreflect.FileDescriptor

var file_ftp_server_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66,
	0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x14, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x27, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x68, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x11, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x61, 0x76, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x46, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0f, 0x66, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x46, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46,
	0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x22, 0x5c, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45,
	0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x03, 0x32, 0x7b, 0x0a, 0x10, 0x46, 0x74,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x12, 0x28, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x66, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x42, 0x1e, 0x42, 0x0e, 0x46, 0x74, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0c, 0x2e, 0x3b, 0x66, 0x74, 0x70,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ftp_server_proto_rawDescOnce sync.Once
	file_ftp_server_proto_rawDescData = file_ftp_server_proto_rawDesc
)

func file_ftp_server_proto_rawDescGZIP() []byte {
	file_ftp_server_proto_rawDescOnce.Do(func() {
		file_ftp_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_ftp_server_proto_rawDescData)
	})
	return file_ftp_server_proto_rawDescData
}

var file_ftp_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ftp_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ftp_server_proto_goTypes = []any{
	(FtpServerResult_Result)(0), // 0: mavsdk.rpc.ftp_server.FtpServerResult.Result
	(*SetRootDirRequest)(nil),   // 1: mavsdk.rpc.ftp_server.SetRootDirRequest
	(*SetRootDirResponse)(nil),  // 2: mavsdk.rpc.ftp_server.SetRootDirResponse
	(*FtpServerResult)(nil),     // 3: mavsdk.rpc.ftp_server.FtpServerResult
}
var file_ftp_server_proto_depIdxs = []int32{
	3, // 0: mavsdk.rpc.ftp_server.SetRootDirResponse.ftp_server_result:type_name -> mavsdk.rpc.ftp_server.FtpServerResult
	0, // 1: mavsdk.rpc.ftp_server.FtpServerResult.result:type_name -> mavsdk.rpc.ftp_server.FtpServerResult.Result
	1, // 2: mavsdk.rpc.ftp_server.FtpServerService.SetRootDir:input_type -> mavsdk.rpc.ftp_server.SetRootDirRequest
	2, // 3: mavsdk.rpc.ftp_server.FtpServerService.SetRootDir:output_type -> mavsdk.rpc.ftp_server.SetRootDirResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ftp_server_proto_init() }
func file_ftp_server_proto_init() {
	if File_ftp_server_proto != nil {
		return
	}
	
	if !protoimpl.UnsafeEnabled {
		file_ftp_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetRootDirRequest); i {
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
		file_ftp_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetRootDirResponse); i {
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
		file_ftp_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FtpServerResult); i {
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
			RawDescriptor: file_ftp_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ftp_server_proto_goTypes,
		DependencyIndexes: file_ftp_server_proto_depIdxs,
		EnumInfos:         file_ftp_server_proto_enumTypes,
		MessageInfos:      file_ftp_server_proto_msgTypes,
	}.Build()
	File_ftp_server_proto = out.File
	file_ftp_server_proto_rawDesc = nil
	file_ftp_server_proto_goTypes = nil
	file_ftp_server_proto_depIdxs = nil
}
