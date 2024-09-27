// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: log_streaming.proto

package log_streaming

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

// Possible results returned for logging requests
type LogStreamingResult_Result int32

const (
	LogStreamingResult_RESULT_SUCCESS          LogStreamingResult_Result = 0 // Request succeeded
	LogStreamingResult_RESULT_NO_SYSTEM        LogStreamingResult_Result = 1 // No system connected
	LogStreamingResult_RESULT_CONNECTION_ERROR LogStreamingResult_Result = 2 // Connection error
	LogStreamingResult_RESULT_BUSY             LogStreamingResult_Result = 3 // System busy
	LogStreamingResult_RESULT_COMMAND_DENIED   LogStreamingResult_Result = 4 // Command denied
	LogStreamingResult_RESULT_TIMEOUT          LogStreamingResult_Result = 5 // Timeout
	LogStreamingResult_RESULT_UNSUPPORTED      LogStreamingResult_Result = 6 // Unsupported
	LogStreamingResult_RESULT_UNKNOWN          LogStreamingResult_Result = 7 // Unknown error
)

// Enum value maps for LogStreamingResult_Result.
var (
	LogStreamingResult_Result_name = map[int32]string{
		0: "RESULT_SUCCESS",
		1: "RESULT_NO_SYSTEM",
		2: "RESULT_CONNECTION_ERROR",
		3: "RESULT_BUSY",
		4: "RESULT_COMMAND_DENIED",
		5: "RESULT_TIMEOUT",
		6: "RESULT_UNSUPPORTED",
		7: "RESULT_UNKNOWN",
	}
	LogStreamingResult_Result_value = map[string]int32{
		"RESULT_SUCCESS":          0,
		"RESULT_NO_SYSTEM":        1,
		"RESULT_CONNECTION_ERROR": 2,
		"RESULT_BUSY":             3,
		"RESULT_COMMAND_DENIED":   4,
		"RESULT_TIMEOUT":          5,
		"RESULT_UNSUPPORTED":      6,
		"RESULT_UNKNOWN":          7,
	}
)

func (x LogStreamingResult_Result) Enum() *LogStreamingResult_Result {
	p := new(LogStreamingResult_Result)
	*p = x
	return p
}

func (x LogStreamingResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogStreamingResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_log_streaming_proto_enumTypes[0].Descriptor()
}

func (LogStreamingResult_Result) Type() protoreflect.EnumType {
	return &file_log_streaming_proto_enumTypes[0]
}

func (x LogStreamingResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogStreamingResult_Result.Descriptor instead.
func (LogStreamingResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{7, 0}
}

type StartLogStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartLogStreamingRequest) Reset() {
	*x = StartLogStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLogStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLogStreamingRequest) ProtoMessage() {}

func (x *StartLogStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLogStreamingRequest.ProtoReflect.Descriptor instead.
func (*StartLogStreamingRequest) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{0}
}

type StartLogStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogStreamingResult *LogStreamingResult `protobuf:"bytes,1,opt,name=log_streaming_result,json=logStreamingResult,proto3" json:"log_streaming_result,omitempty"`
}

func (x *StartLogStreamingResponse) Reset() {
	*x = StartLogStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLogStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLogStreamingResponse) ProtoMessage() {}

func (x *StartLogStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLogStreamingResponse.ProtoReflect.Descriptor instead.
func (*StartLogStreamingResponse) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *StartLogStreamingResponse) GetLogStreamingResult() *LogStreamingResult {
	if x != nil {
		return x.LogStreamingResult
	}
	return nil
}

type StopLogStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopLogStreamingRequest) Reset() {
	*x = StopLogStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopLogStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopLogStreamingRequest) ProtoMessage() {}

func (x *StopLogStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopLogStreamingRequest.ProtoReflect.Descriptor instead.
func (*StopLogStreamingRequest) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{2}
}

type StopLogStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogStreamingResult *LogStreamingResult `protobuf:"bytes,1,opt,name=log_streaming_result,json=logStreamingResult,proto3" json:"log_streaming_result,omitempty"`
}

func (x *StopLogStreamingResponse) Reset() {
	*x = StopLogStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopLogStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopLogStreamingResponse) ProtoMessage() {}

func (x *StopLogStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopLogStreamingResponse.ProtoReflect.Descriptor instead.
func (*StopLogStreamingResponse) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{3}
}

func (x *StopLogStreamingResponse) GetLogStreamingResult() *LogStreamingResult {
	if x != nil {
		return x.LogStreamingResult
	}
	return nil
}

type SubscribeLogStreamingRawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeLogStreamingRawRequest) Reset() {
	*x = SubscribeLogStreamingRawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeLogStreamingRawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeLogStreamingRawRequest) ProtoMessage() {}

func (x *SubscribeLogStreamingRawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeLogStreamingRawRequest.ProtoReflect.Descriptor instead.
func (*SubscribeLogStreamingRawRequest) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{4}
}

type LogStreamingRawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoggingRaw *LogStreamingRaw `protobuf:"bytes,1,opt,name=logging_raw,json=loggingRaw,proto3" json:"logging_raw,omitempty"` // A message containing logged data
}

func (x *LogStreamingRawResponse) Reset() {
	*x = LogStreamingRawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStreamingRawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStreamingRawResponse) ProtoMessage() {}

func (x *LogStreamingRawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStreamingRawResponse.ProtoReflect.Descriptor instead.
func (*LogStreamingRawResponse) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{5}
}

func (x *LogStreamingRawResponse) GetLoggingRaw() *LogStreamingRaw {
	if x != nil {
		return x.LoggingRaw
	}
	return nil
}

// Raw logging data type
type LogStreamingRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Ulog file stream data encoded as base64
}

func (x *LogStreamingRaw) Reset() {
	*x = LogStreamingRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStreamingRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStreamingRaw) ProtoMessage() {}

func (x *LogStreamingRaw) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStreamingRaw.ProtoReflect.Descriptor instead.
func (*LogStreamingRaw) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{6}
}

func (x *LogStreamingRaw) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Result type.
type LogStreamingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    LogStreamingResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.log_streaming.LogStreamingResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string                    `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                                   // Human-readable English string describing the result
}

func (x *LogStreamingResult) Reset() {
	*x = LogStreamingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_streaming_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStreamingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStreamingResult) ProtoMessage() {}

func (x *LogStreamingResult) ProtoReflect() protoreflect.Message {
	mi := &file_log_streaming_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStreamingResult.ProtoReflect.Descriptor instead.
func (*LogStreamingResult) Descriptor() ([]byte, []int) {
	return file_log_streaming_proto_rawDescGZIP(), []int{7}
}

func (x *LogStreamingResult) GetResult() LogStreamingResult_Result {
	if x != nil {
		return x.Result
	}
	return LogStreamingResult_RESULT_SUCCESS
}

func (x *LogStreamingResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_log_streaming_proto protoreflect.FileDescriptor

var file_log_streaming_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x1a,
	0x14, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f,
	0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x7b, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d,
	0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x19,
	0x0a, 0x17, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x18, 0x53, 0x74, 0x6f,
	0x70, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x0a, 0x1f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x65, 0x0a, 0x17, 0x4c, 0x6f, 0x67, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x61, 0x77, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x77, 0x22,
	0x25, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x02, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4b, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e,
	0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x55,
	0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x44,
	0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c,
	0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x07, 0x32, 0xa5, 0x03, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x7e, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x31, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x90, 0x01, 0x0a,
	0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x77, 0x12, 0x39, 0x2e, 0x6d, 0x61, 0x76, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x6f,
	0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x00, 0x30, 0x01, 0x42,
	0x24, 0x42, 0x11, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0f, 0x2e, 0x3b, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_streaming_proto_rawDescOnce sync.Once
	file_log_streaming_proto_rawDescData = file_log_streaming_proto_rawDesc
)

func file_log_streaming_proto_rawDescGZIP() []byte {
	file_log_streaming_proto_rawDescOnce.Do(func() {
		file_log_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_streaming_proto_rawDescData)
	})
	return file_log_streaming_proto_rawDescData
}

var file_log_streaming_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_log_streaming_proto_goTypes = []any{
	(LogStreamingResult_Result)(0),          // 0: mavsdk.rpc.log_streaming.LogStreamingResult.Result
	(*StartLogStreamingRequest)(nil),        // 1: mavsdk.rpc.log_streaming.StartLogStreamingRequest
	(*StartLogStreamingResponse)(nil),       // 2: mavsdk.rpc.log_streaming.StartLogStreamingResponse
	(*StopLogStreamingRequest)(nil),         // 3: mavsdk.rpc.log_streaming.StopLogStreamingRequest
	(*StopLogStreamingResponse)(nil),        // 4: mavsdk.rpc.log_streaming.StopLogStreamingResponse
	(*SubscribeLogStreamingRawRequest)(nil), // 5: mavsdk.rpc.log_streaming.SubscribeLogStreamingRawRequest
	(*LogStreamingRawResponse)(nil),         // 6: mavsdk.rpc.log_streaming.LogStreamingRawResponse
	(*LogStreamingRaw)(nil),                 // 7: mavsdk.rpc.log_streaming.LogStreamingRaw
	(*LogStreamingResult)(nil),              // 8: mavsdk.rpc.log_streaming.LogStreamingResult
}
var file_log_streaming_proto_depIdxs = []int32{
	8, // 0: mavsdk.rpc.log_streaming.StartLogStreamingResponse.log_streaming_result:type_name -> mavsdk.rpc.log_streaming.LogStreamingResult
	8, // 1: mavsdk.rpc.log_streaming.StopLogStreamingResponse.log_streaming_result:type_name -> mavsdk.rpc.log_streaming.LogStreamingResult
	7, // 2: mavsdk.rpc.log_streaming.LogStreamingRawResponse.logging_raw:type_name -> mavsdk.rpc.log_streaming.LogStreamingRaw
	0, // 3: mavsdk.rpc.log_streaming.LogStreamingResult.result:type_name -> mavsdk.rpc.log_streaming.LogStreamingResult.Result
	1, // 4: mavsdk.rpc.log_streaming.LogStreamingService.StartLogStreaming:input_type -> mavsdk.rpc.log_streaming.StartLogStreamingRequest
	3, // 5: mavsdk.rpc.log_streaming.LogStreamingService.StopLogStreaming:input_type -> mavsdk.rpc.log_streaming.StopLogStreamingRequest
	5, // 6: mavsdk.rpc.log_streaming.LogStreamingService.SubscribeLogStreamingRaw:input_type -> mavsdk.rpc.log_streaming.SubscribeLogStreamingRawRequest
	2, // 7: mavsdk.rpc.log_streaming.LogStreamingService.StartLogStreaming:output_type -> mavsdk.rpc.log_streaming.StartLogStreamingResponse
	4, // 8: mavsdk.rpc.log_streaming.LogStreamingService.StopLogStreaming:output_type -> mavsdk.rpc.log_streaming.StopLogStreamingResponse
	6, // 9: mavsdk.rpc.log_streaming.LogStreamingService.SubscribeLogStreamingRaw:output_type -> mavsdk.rpc.log_streaming.LogStreamingRawResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_log_streaming_proto_init() }
func file_log_streaming_proto_init() {
	if File_log_streaming_proto != nil {
		return
	}
	
	if !protoimpl.UnsafeEnabled {
		file_log_streaming_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StartLogStreamingRequest); i {
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
		file_log_streaming_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StartLogStreamingResponse); i {
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
		file_log_streaming_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StopLogStreamingRequest); i {
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
		file_log_streaming_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StopLogStreamingResponse); i {
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
		file_log_streaming_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SubscribeLogStreamingRawRequest); i {
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
		file_log_streaming_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LogStreamingRawResponse); i {
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
		file_log_streaming_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*LogStreamingRaw); i {
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
		file_log_streaming_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*LogStreamingResult); i {
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
			RawDescriptor: file_log_streaming_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_streaming_proto_goTypes,
		DependencyIndexes: file_log_streaming_proto_depIdxs,
		EnumInfos:         file_log_streaming_proto_enumTypes,
		MessageInfos:      file_log_streaming_proto_msgTypes,
	}.Build()
	File_log_streaming_proto = out.File
	file_log_streaming_proto_rawDesc = nil
	file_log_streaming_proto_goTypes = nil
	file_log_streaming_proto_depIdxs = nil
}