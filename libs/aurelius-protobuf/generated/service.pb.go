// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: service.proto

package aurelius_protobuf

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

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName    string `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName,omitempty"`
	JobVersion string `protobuf:"bytes,2,opt,name=jobVersion,proto3" json:"jobVersion,omitempty"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *RunJobRequest) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *RunJobRequest) GetJobVersion() string {
	if x != nil {
		return x.JobVersion
	}
	return ""
}

type JobChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*JobChatRequest_RunJobRequest
	Payload isJobChatRequest_Payload `protobuf_oneof:"payload"`
}

func (x *JobChatRequest) Reset() {
	*x = JobChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobChatRequest) ProtoMessage() {}

func (x *JobChatRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use JobChatRequest.ProtoReflect.Descriptor instead.
func (*JobChatRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (m *JobChatRequest) GetPayload() isJobChatRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *JobChatRequest) GetRunJobRequest() *RunJobRequest {
	if x, ok := x.GetPayload().(*JobChatRequest_RunJobRequest); ok {
		return x.RunJobRequest
	}
	return nil
}

type isJobChatRequest_Payload interface {
	isJobChatRequest_Payload()
}

type JobChatRequest_RunJobRequest struct {
	RunJobRequest *RunJobRequest `protobuf:"bytes,1,opt,name=runJobRequest,proto3,oneof"`
}

func (*JobChatRequest_RunJobRequest) isJobChatRequest_Payload() {}

type RunJobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId    string `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	JobName    string `protobuf:"bytes,2,opt,name=jobName,proto3" json:"jobName,omitempty"`
	JobVersion string `protobuf:"bytes,3,opt,name=jobVersion,proto3" json:"jobVersion,omitempty"`
	Message    string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RunJobStatus) Reset() {
	*x = RunJobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobStatus) ProtoMessage() {}

func (x *RunJobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobStatus.ProtoReflect.Descriptor instead.
func (*RunJobStatus) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *RunJobStatus) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *RunJobStatus) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *RunJobStatus) GetJobVersion() string {
	if x != nil {
		return x.JobVersion
	}
	return ""
}

func (x *RunJobStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RunJobResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId    string `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	JobName    string `protobuf:"bytes,2,opt,name=jobName,proto3" json:"jobName,omitempty"`
	JobVersion string `protobuf:"bytes,3,opt,name=jobVersion,proto3" json:"jobVersion,omitempty"`
	Success    bool   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	Result     string `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RunJobResult) Reset() {
	*x = RunJobResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobResult) ProtoMessage() {}

func (x *RunJobResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobResult.ProtoReflect.Descriptor instead.
func (*RunJobResult) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *RunJobResult) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *RunJobResult) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *RunJobResult) GetJobVersion() string {
	if x != nil {
		return x.JobVersion
	}
	return ""
}

func (x *RunJobResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RunJobResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type JobChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*JobChatResponse_RunJobStatus
	//	*JobChatResponse_RunJobResult
	Payload isJobChatResponse_Payload `protobuf_oneof:"payload"`
}

func (x *JobChatResponse) Reset() {
	*x = JobChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobChatResponse) ProtoMessage() {}

func (x *JobChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobChatResponse.ProtoReflect.Descriptor instead.
func (*JobChatResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (m *JobChatResponse) GetPayload() isJobChatResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *JobChatResponse) GetRunJobStatus() *RunJobStatus {
	if x, ok := x.GetPayload().(*JobChatResponse_RunJobStatus); ok {
		return x.RunJobStatus
	}
	return nil
}

func (x *JobChatResponse) GetRunJobResult() *RunJobResult {
	if x, ok := x.GetPayload().(*JobChatResponse_RunJobResult); ok {
		return x.RunJobResult
	}
	return nil
}

type isJobChatResponse_Payload interface {
	isJobChatResponse_Payload()
}

type JobChatResponse_RunJobStatus struct {
	RunJobStatus *RunJobStatus `protobuf:"bytes,1,opt,name=runJobStatus,proto3,oneof"`
}

type JobChatResponse_RunJobResult struct {
	RunJobResult *RunJobResult `protobuf:"bytes,2,opt,name=runJobResult,proto3,oneof"`
}

func (*JobChatResponse_RunJobStatus) isJobChatResponse_Payload() {}

func (*JobChatResponse_RunJobResult) isJobChatResponse_Payload() {}

type RegisterAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Instance string           `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	Stack    string           `protobuf:"bytes,3,opt,name=stack,proto3" json:"stack,omitempty"`
	Version  string           `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Jobs     []*JobDefinition `protobuf:"bytes,5,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *RegisterAgentRequest) Reset() {
	*x = RegisterAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentRequest) ProtoMessage() {}

func (x *RegisterAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentRequest.ProtoReflect.Descriptor instead.
func (*RegisterAgentRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterAgentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterAgentRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *RegisterAgentRequest) GetStack() string {
	if x != nil {
		return x.Stack
	}
	return ""
}

func (x *RegisterAgentRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegisterAgentRequest) GetJobs() []*JobDefinition {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type JobDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *JobDefinition) Reset() {
	*x = JobDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDefinition) ProtoMessage() {}

func (x *JobDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDefinition.ProtoReflect.Descriptor instead.
func (*JobDefinition) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *JobDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JobDefinition) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RegisterAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string `protobuf:"bytes,1,opt,name=agentId,proto3" json:"agentId,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterAgentResponse) Reset() {
	*x = RegisterAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentResponse) ProtoMessage() {}

func (x *RegisterAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentResponse.ProtoReflect.Descriptor instead.
func (*RegisterAgentResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterAgentResponse) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *RegisterAgentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f,
	0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6a, 0x6f, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0e, 0x4a, 0x6f,
	0x62, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d,
	0x72, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x72, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x7c, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x94, 0x01,
	0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52,
	0x0c, 0x72, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a,
	0x0c, 0x72, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9a, 0x01,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x3d, 0x0a, 0x0d, 0x4a, 0x6f,
	0x62, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x15, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x14, 0x41, 0x75, 0x72, 0x65, 0x6c,
	0x69, 0x75, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x4a, 0x6f, 0x62,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4a, 0x6f,
	0x62, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x1a, 0x61, 0x75, 0x72, 0x65, 0x6c, 0x69, 0x75, 0x73, 0x2f,
	0x61, 0x75, 0x72, 0x65, 0x6c, 0x69, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0xaa, 0x02, 0x1c, 0x41, 0x75, 0x72, 0x65, 0x6c, 0x69, 0x75, 0x73, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x44, 0x6f, 0x74, 0x6e, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_proto_goTypes = []interface{}{
	(*RunJobRequest)(nil),         // 0: RunJobRequest
	(*JobChatRequest)(nil),        // 1: JobChatRequest
	(*RunJobStatus)(nil),          // 2: RunJobStatus
	(*RunJobResult)(nil),          // 3: RunJobResult
	(*JobChatResponse)(nil),       // 4: JobChatResponse
	(*RegisterAgentRequest)(nil),  // 5: RegisterAgentRequest
	(*JobDefinition)(nil),         // 6: JobDefinition
	(*RegisterAgentResponse)(nil), // 7: RegisterAgentResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: JobChatRequest.runJobRequest:type_name -> RunJobRequest
	2, // 1: JobChatResponse.runJobStatus:type_name -> RunJobStatus
	3, // 2: JobChatResponse.runJobResult:type_name -> RunJobResult
	6, // 3: RegisterAgentRequest.jobs:type_name -> JobDefinition
	1, // 4: AureliusAgentManager.JobChat:input_type -> JobChatRequest
	5, // 5: AureliusAgentManager.RegisterAgent:input_type -> RegisterAgentRequest
	4, // 6: AureliusAgentManager.JobChat:output_type -> JobChatResponse
	7, // 7: AureliusAgentManager.RegisterAgent:output_type -> RegisterAgentResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
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
			switch v := v.(*JobChatRequest); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobStatus); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobResult); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobChatResponse); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentRequest); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDefinition); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAgentResponse); i {
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
	file_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JobChatRequest_RunJobRequest)(nil),
	}
	file_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*JobChatResponse_RunJobStatus)(nil),
		(*JobChatResponse_RunJobResult)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
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
