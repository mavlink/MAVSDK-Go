// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: gripper.proto

package gripper

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

// Gripper Actions.
//
// Available gripper actions are defined in mavlink under
// https://mavlink.io/en/messages/common.html#GRIPPER_ACTIONS
type GripperAction int32

const (
	GripperAction_GRIPPER_ACTION_RELEASE GripperAction = 0 // Open the gripper to release the cargo
	GripperAction_GRIPPER_ACTION_GRAB    GripperAction = 1 // Close the gripper and grab onto cargo
)

// Enum value maps for GripperAction.
var (
	GripperAction_name = map[int32]string{
		0: "GRIPPER_ACTION_RELEASE",
		1: "GRIPPER_ACTION_GRAB",
	}
	GripperAction_value = map[string]int32{
		"GRIPPER_ACTION_RELEASE": 0,
		"GRIPPER_ACTION_GRAB":    1,
	}
)

func (x GripperAction) Enum() *GripperAction {
	p := new(GripperAction)
	*p = x
	return p
}

func (x GripperAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GripperAction) Descriptor() protoreflect.EnumDescriptor {
	return file_gripper_proto_enumTypes[0].Descriptor()
}

func (GripperAction) Type() protoreflect.EnumType {
	return &file_gripper_proto_enumTypes[0]
}

func (x GripperAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GripperAction.Descriptor instead.
func (GripperAction) EnumDescriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{0}
}

// Possible results returned for gripper action requests.
type GripperResult_Result int32

const (
	GripperResult_RESULT_UNKNOWN     GripperResult_Result = 0 // Unknown result
	GripperResult_RESULT_SUCCESS     GripperResult_Result = 1 // Request was successful
	GripperResult_RESULT_NO_SYSTEM   GripperResult_Result = 2 // No system is connected
	GripperResult_RESULT_BUSY        GripperResult_Result = 3 // Temporarily rejected
	GripperResult_RESULT_TIMEOUT     GripperResult_Result = 4 // Request timed out
	GripperResult_RESULT_UNSUPPORTED GripperResult_Result = 5 // Action not supported
	GripperResult_RESULT_FAILED      GripperResult_Result = 6 // Action failed
)

// Enum value maps for GripperResult_Result.
var (
	GripperResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_NO_SYSTEM",
		3: "RESULT_BUSY",
		4: "RESULT_TIMEOUT",
		5: "RESULT_UNSUPPORTED",
		6: "RESULT_FAILED",
	}
	GripperResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":     0,
		"RESULT_SUCCESS":     1,
		"RESULT_NO_SYSTEM":   2,
		"RESULT_BUSY":        3,
		"RESULT_TIMEOUT":     4,
		"RESULT_UNSUPPORTED": 5,
		"RESULT_FAILED":      6,
	}
)

func (x GripperResult_Result) Enum() *GripperResult_Result {
	p := new(GripperResult_Result)
	*p = x
	return p
}

func (x GripperResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GripperResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_gripper_proto_enumTypes[1].Descriptor()
}

func (GripperResult_Result) Type() protoreflect.EnumType {
	return &file_gripper_proto_enumTypes[1]
}

func (x GripperResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GripperResult_Result.Descriptor instead.
func (GripperResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{4, 0}
}

type GrabRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance uint32 `protobuf:"varint,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *GrabRequest) Reset() {
	*x = GrabRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gripper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrabRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrabRequest) ProtoMessage() {}

func (x *GrabRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gripper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrabRequest.ProtoReflect.Descriptor instead.
func (*GrabRequest) Descriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{0}
}

func (x *GrabRequest) GetInstance() uint32 {
	if x != nil {
		return x.Instance
	}
	return 0
}

type GrabResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GripperResult *GripperResult `protobuf:"bytes,1,opt,name=gripper_result,json=gripperResult,proto3" json:"gripper_result,omitempty"`
}

func (x *GrabResponse) Reset() {
	*x = GrabResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gripper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrabResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrabResponse) ProtoMessage() {}

func (x *GrabResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gripper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrabResponse.ProtoReflect.Descriptor instead.
func (*GrabResponse) Descriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{1}
}

func (x *GrabResponse) GetGripperResult() *GripperResult {
	if x != nil {
		return x.GripperResult
	}
	return nil
}

type ReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance uint32 `protobuf:"varint,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *ReleaseRequest) Reset() {
	*x = ReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gripper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseRequest) ProtoMessage() {}

func (x *ReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gripper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseRequest.ProtoReflect.Descriptor instead.
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseRequest) GetInstance() uint32 {
	if x != nil {
		return x.Instance
	}
	return 0
}

type ReleaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GripperResult *GripperResult `protobuf:"bytes,1,opt,name=gripper_result,json=gripperResult,proto3" json:"gripper_result,omitempty"`
}

func (x *ReleaseResponse) Reset() {
	*x = ReleaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gripper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseResponse) ProtoMessage() {}

func (x *ReleaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gripper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseResponse.ProtoReflect.Descriptor instead.
func (*ReleaseResponse) Descriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseResponse) GetGripperResult() *GripperResult {
	if x != nil {
		return x.GripperResult
	}
	return nil
}

// Result type.
type GripperResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    GripperResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.gripper.GripperResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr string               `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                        // Human-readable English string describing the result
}

func (x *GripperResult) Reset() {
	*x = GripperResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gripper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GripperResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GripperResult) ProtoMessage() {}

func (x *GripperResult) ProtoReflect() protoreflect.Message {
	mi := &file_gripper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GripperResult.ProtoReflect.Descriptor instead.
func (*GripperResult) Descriptor() ([]byte, []int) {
	return file_gripper_proto_rawDescGZIP(), []int{4}
}

func (x *GripperResult) GetResult() GripperResult_Result {
	if x != nil {
		return x.Result
	}
	return GripperResult_RESULT_UNKNOWN
}

func (x *GripperResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_gripper_proto protoreflect.FileDescriptor

var file_gripper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x72, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x58,
	0x0a, 0x0c, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0e, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x67, 0x72, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x67, 0x72, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x89, 0x02, 0x0a, 0x0d, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x2a,
	0x44, 0x0a, 0x0d, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x16, 0x47, 0x52, 0x49, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x47, 0x52, 0x49, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47,
	0x52, 0x41, 0x42, 0x10, 0x01, 0x32, 0xb3, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x47, 0x72, 0x61, 0x62,
	0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x72,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x61, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x22, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x72,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x76, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x67, 0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x42, 0x0c, 0x47,
	0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x09, 0x2e, 0x3b, 0x67,
	0x72, 0x69, 0x70, 0x70, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gripper_proto_rawDescOnce sync.Once
	file_gripper_proto_rawDescData = file_gripper_proto_rawDesc
)

func file_gripper_proto_rawDescGZIP() []byte {
	file_gripper_proto_rawDescOnce.Do(func() {
		file_gripper_proto_rawDescData = protoimpl.X.CompressGZIP(file_gripper_proto_rawDescData)
	})
	return file_gripper_proto_rawDescData
}

var file_gripper_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gripper_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gripper_proto_goTypes = []any{
	(GripperAction)(0),        // 0: mavsdk.rpc.gripper.GripperAction
	(GripperResult_Result)(0), // 1: mavsdk.rpc.gripper.GripperResult.Result
	(*GrabRequest)(nil),       // 2: mavsdk.rpc.gripper.GrabRequest
	(*GrabResponse)(nil),      // 3: mavsdk.rpc.gripper.GrabResponse
	(*ReleaseRequest)(nil),    // 4: mavsdk.rpc.gripper.ReleaseRequest
	(*ReleaseResponse)(nil),   // 5: mavsdk.rpc.gripper.ReleaseResponse
	(*GripperResult)(nil),     // 6: mavsdk.rpc.gripper.GripperResult
}
var file_gripper_proto_depIdxs = []int32{
	6, // 0: mavsdk.rpc.gripper.GrabResponse.gripper_result:type_name -> mavsdk.rpc.gripper.GripperResult
	6, // 1: mavsdk.rpc.gripper.ReleaseResponse.gripper_result:type_name -> mavsdk.rpc.gripper.GripperResult
	1, // 2: mavsdk.rpc.gripper.GripperResult.result:type_name -> mavsdk.rpc.gripper.GripperResult.Result
	2, // 3: mavsdk.rpc.gripper.GripperService.Grab:input_type -> mavsdk.rpc.gripper.GrabRequest
	4, // 4: mavsdk.rpc.gripper.GripperService.Release:input_type -> mavsdk.rpc.gripper.ReleaseRequest
	3, // 5: mavsdk.rpc.gripper.GripperService.Grab:output_type -> mavsdk.rpc.gripper.GrabResponse
	5, // 6: mavsdk.rpc.gripper.GripperService.Release:output_type -> mavsdk.rpc.gripper.ReleaseResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gripper_proto_init() }
func file_gripper_proto_init() {
	if File_gripper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gripper_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GrabRequest); i {
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
		file_gripper_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GrabResponse); i {
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
		file_gripper_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ReleaseRequest); i {
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
		file_gripper_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReleaseResponse); i {
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
		file_gripper_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GripperResult); i {
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
			RawDescriptor: file_gripper_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gripper_proto_goTypes,
		DependencyIndexes: file_gripper_proto_depIdxs,
		EnumInfos:         file_gripper_proto_enumTypes,
		MessageInfos:      file_gripper_proto_msgTypes,
	}.Build()
	File_gripper_proto = out.File
	file_gripper_proto_rawDesc = nil
	file_gripper_proto_goTypes = nil
	file_gripper_proto_depIdxs = nil
}