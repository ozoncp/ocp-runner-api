// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.7
// source: api/ocp-runner-api/ocp-runner-api.proto

package ocp_runner_api

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

// CREATE
type CreateRunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os   string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Arch string `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
}

func (x *CreateRunnerRequest) Reset() {
	*x = CreateRunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunnerRequest) ProtoMessage() {}

func (x *CreateRunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunnerRequest.ProtoReflect.Descriptor instead.
func (*CreateRunnerRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRunnerRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *CreateRunnerRequest) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

type CreateRunnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRunnerResponse) Reset() {
	*x = CreateRunnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRunnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRunnerResponse) ProtoMessage() {}

func (x *CreateRunnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRunnerResponse.ProtoReflect.Descriptor instead.
func (*CreateRunnerResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{1}
}

// DESCRIBE
type DescribeRunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Os   string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Arch string `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
}

func (x *DescribeRunnerRequest) Reset() {
	*x = DescribeRunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRunnerRequest) ProtoMessage() {}

func (x *DescribeRunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRunnerRequest.ProtoReflect.Descriptor instead.
func (*DescribeRunnerRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeRunnerRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *DescribeRunnerRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *DescribeRunnerRequest) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

type DescribeRunnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DescribeRunnerResponse) Reset() {
	*x = DescribeRunnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRunnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRunnerResponse) ProtoMessage() {}

func (x *DescribeRunnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRunnerResponse.ProtoReflect.Descriptor instead.
func (*DescribeRunnerResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{3}
}

// REMOVE
type RemoveRunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *RemoveRunnerRequest) Reset() {
	*x = RemoveRunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRunnerRequest) ProtoMessage() {}

func (x *RemoveRunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRunnerRequest.ProtoReflect.Descriptor instead.
func (*RemoveRunnerRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRunnerRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

type RemoveRunnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRunnerResponse) Reset() {
	*x = RemoveRunnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRunnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRunnerResponse) ProtoMessage() {}

func (x *RemoveRunnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRunnerResponse.ProtoReflect.Descriptor instead.
func (*RemoveRunnerResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{5}
}

// LIST
type ListFiltersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guids []string `protobuf:"bytes,1,rep,name=guids,proto3" json:"guids,omitempty"`
	Os    []string `protobuf:"bytes,2,rep,name=os,proto3" json:"os,omitempty"`
	Arch  []string `protobuf:"bytes,3,rep,name=arch,proto3" json:"arch,omitempty"`
}

func (x *ListFiltersRequest) Reset() {
	*x = ListFiltersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFiltersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFiltersRequest) ProtoMessage() {}

func (x *ListFiltersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFiltersRequest.ProtoReflect.Descriptor instead.
func (*ListFiltersRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListFiltersRequest) GetGuids() []string {
	if x != nil {
		return x.Guids
	}
	return nil
}

func (x *ListFiltersRequest) GetOs() []string {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *ListFiltersRequest) GetArch() []string {
	if x != nil {
		return x.Arch
	}
	return nil
}

type Runner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guid string `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Os   string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Arch string `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
}

func (x *Runner) Reset() {
	*x = Runner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runner) ProtoMessage() {}

func (x *Runner) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runner.ProtoReflect.Descriptor instead.
func (*Runner) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{7}
}

func (x *Runner) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *Runner) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *Runner) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

type RunnersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runners []*Runner `protobuf:"bytes,1,rep,name=runners,proto3" json:"runners,omitempty"`
}

func (x *RunnersListResponse) Reset() {
	*x = RunnersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnersListResponse) ProtoMessage() {}

func (x *RunnersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnersListResponse.ProtoReflect.Descriptor instead.
func (*RunnersListResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP(), []int{8}
}

func (x *RunnersListResponse) GetRunners() []*Runner {
	if x != nil {
		return x.Runners
	}
	return nil
}

var File_api_ocp_runner_api_ocp_runner_api_proto protoreflect.FileDescriptor

var file_api_ocp_runner_api_ocp_runner_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2d,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x63, 0x70, 0x2e, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x0a, 0x15,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x22, 0x18, 0x0a,
	0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75,
	0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x75, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x22, 0x40, 0x0a, 0x06, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x22, 0x47, 0x0a, 0x13,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x73, 0x32, 0x81, 0x03, 0x0a, 0x10, 0x4f, 0x63, 0x70, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f,
	0x63, 0x70, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69,
	0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_runner_api_ocp_runner_api_proto_rawDescOnce sync.Once
	file_api_ocp_runner_api_ocp_runner_api_proto_rawDescData = file_api_ocp_runner_api_ocp_runner_api_proto_rawDesc
)

func file_api_ocp_runner_api_ocp_runner_api_proto_rawDescGZIP() []byte {
	file_api_ocp_runner_api_ocp_runner_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_runner_api_ocp_runner_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_runner_api_ocp_runner_api_proto_rawDescData)
	})
	return file_api_ocp_runner_api_ocp_runner_api_proto_rawDescData
}

var file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_runner_api_ocp_runner_api_proto_goTypes = []interface{}{
	(*CreateRunnerRequest)(nil),    // 0: ocp.runner.api.CreateRunnerRequest
	(*CreateRunnerResponse)(nil),   // 1: ocp.runner.api.CreateRunnerResponse
	(*DescribeRunnerRequest)(nil),  // 2: ocp.runner.api.DescribeRunnerRequest
	(*DescribeRunnerResponse)(nil), // 3: ocp.runner.api.DescribeRunnerResponse
	(*RemoveRunnerRequest)(nil),    // 4: ocp.runner.api.RemoveRunnerRequest
	(*RemoveRunnerResponse)(nil),   // 5: ocp.runner.api.RemoveRunnerResponse
	(*ListFiltersRequest)(nil),     // 6: ocp.runner.api.ListFiltersRequest
	(*Runner)(nil),                 // 7: ocp.runner.api.Runner
	(*RunnersListResponse)(nil),    // 8: ocp.runner.api.RunnersListResponse
}
var file_api_ocp_runner_api_ocp_runner_api_proto_depIdxs = []int32{
	7, // 0: ocp.runner.api.RunnersListResponse.runners:type_name -> ocp.runner.api.Runner
	0, // 1: ocp.runner.api.OcpRunnerService.CreateRunner:input_type -> ocp.runner.api.CreateRunnerRequest
	2, // 2: ocp.runner.api.OcpRunnerService.DescribeRunner:input_type -> ocp.runner.api.DescribeRunnerRequest
	4, // 3: ocp.runner.api.OcpRunnerService.RemoveRunner:input_type -> ocp.runner.api.RemoveRunnerRequest
	6, // 4: ocp.runner.api.OcpRunnerService.ListRunners:input_type -> ocp.runner.api.ListFiltersRequest
	1, // 5: ocp.runner.api.OcpRunnerService.CreateRunner:output_type -> ocp.runner.api.CreateRunnerResponse
	3, // 6: ocp.runner.api.OcpRunnerService.DescribeRunner:output_type -> ocp.runner.api.DescribeRunnerResponse
	5, // 7: ocp.runner.api.OcpRunnerService.RemoveRunner:output_type -> ocp.runner.api.RemoveRunnerResponse
	8, // 8: ocp.runner.api.OcpRunnerService.ListRunners:output_type -> ocp.runner.api.RunnersListResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ocp_runner_api_ocp_runner_api_proto_init() }
func file_api_ocp_runner_api_ocp_runner_api_proto_init() {
	if File_api_ocp_runner_api_ocp_runner_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRunnerRequest); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRunnerResponse); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRunnerRequest); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRunnerResponse); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRunnerRequest); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRunnerResponse); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFiltersRequest); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runner); i {
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
		file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnersListResponse); i {
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
			RawDescriptor: file_api_ocp_runner_api_ocp_runner_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_runner_api_ocp_runner_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_runner_api_ocp_runner_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_runner_api_ocp_runner_api_proto_msgTypes,
	}.Build()
	File_api_ocp_runner_api_ocp_runner_api_proto = out.File
	file_api_ocp_runner_api_ocp_runner_api_proto_rawDesc = nil
	file_api_ocp_runner_api_ocp_runner_api_proto_goTypes = nil
	file_api_ocp_runner_api_ocp_runner_api_proto_depIdxs = nil
}
