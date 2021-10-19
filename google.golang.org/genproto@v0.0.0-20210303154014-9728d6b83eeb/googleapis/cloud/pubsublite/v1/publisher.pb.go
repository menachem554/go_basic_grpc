// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/pubsublite/v1/publisher.proto

package pubsublite

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The first request that must be sent on a newly-opened stream.
type InitialPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The topic to which messages will be written.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// The partition within the topic to which messages will be written.
	// Partitions are zero indexed, so `partition` must be in the range [0,
	// topic.num_partitions).
	Partition int64 `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *InitialPublishRequest) Reset() {
	*x = InitialPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitialPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialPublishRequest) ProtoMessage() {}

func (x *InitialPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialPublishRequest.ProtoReflect.Descriptor instead.
func (*InitialPublishRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{0}
}

func (x *InitialPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *InitialPublishRequest) GetPartition() int64 {
	if x != nil {
		return x.Partition
	}
	return 0
}

// Response to an InitialPublishRequest.
type InitialPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitialPublishResponse) Reset() {
	*x = InitialPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitialPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialPublishResponse) ProtoMessage() {}

func (x *InitialPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialPublishResponse.ProtoReflect.Descriptor instead.
func (*InitialPublishResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{1}
}

// Request to publish messages to the topic.
type MessagePublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The messages to publish.
	Messages []*PubSubMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessagePublishRequest) Reset() {
	*x = MessagePublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePublishRequest) ProtoMessage() {}

func (x *MessagePublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePublishRequest.ProtoReflect.Descriptor instead.
func (*MessagePublishRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{2}
}

func (x *MessagePublishRequest) GetMessages() []*PubSubMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

// Response to a MessagePublishRequest.
type MessagePublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cursor of the first published message in the batch. The cursors for any
	// remaining messages in the batch are guaranteed to be sequential.
	StartCursor *Cursor `protobuf:"bytes,1,opt,name=start_cursor,json=startCursor,proto3" json:"start_cursor,omitempty"`
}

func (x *MessagePublishResponse) Reset() {
	*x = MessagePublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePublishResponse) ProtoMessage() {}

func (x *MessagePublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePublishResponse.ProtoReflect.Descriptor instead.
func (*MessagePublishResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{3}
}

func (x *MessagePublishResponse) GetStartCursor() *Cursor {
	if x != nil {
		return x.StartCursor
	}
	return nil
}

// Request sent from the client to the server on a stream.
type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of request this is.
	//
	// Types that are assignable to RequestType:
	//	*PublishRequest_InitialRequest
	//	*PublishRequest_MessagePublishRequest
	RequestType isPublishRequest_RequestType `protobuf_oneof:"request_type"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{4}
}

func (m *PublishRequest) GetRequestType() isPublishRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (x *PublishRequest) GetInitialRequest() *InitialPublishRequest {
	if x, ok := x.GetRequestType().(*PublishRequest_InitialRequest); ok {
		return x.InitialRequest
	}
	return nil
}

func (x *PublishRequest) GetMessagePublishRequest() *MessagePublishRequest {
	if x, ok := x.GetRequestType().(*PublishRequest_MessagePublishRequest); ok {
		return x.MessagePublishRequest
	}
	return nil
}

type isPublishRequest_RequestType interface {
	isPublishRequest_RequestType()
}

type PublishRequest_InitialRequest struct {
	// Initial request on the stream.
	InitialRequest *InitialPublishRequest `protobuf:"bytes,1,opt,name=initial_request,json=initialRequest,proto3,oneof"`
}

type PublishRequest_MessagePublishRequest struct {
	// Request to publish messages.
	MessagePublishRequest *MessagePublishRequest `protobuf:"bytes,2,opt,name=message_publish_request,json=messagePublishRequest,proto3,oneof"`
}

func (*PublishRequest_InitialRequest) isPublishRequest_RequestType() {}

func (*PublishRequest_MessagePublishRequest) isPublishRequest_RequestType() {}

// Response to a PublishRequest.
type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of response this is.
	//
	// Types that are assignable to ResponseType:
	//	*PublishResponse_InitialResponse
	//	*PublishResponse_MessageResponse
	ResponseType isPublishResponse_ResponseType `protobuf_oneof:"response_type"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP(), []int{5}
}

func (m *PublishResponse) GetResponseType() isPublishResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (x *PublishResponse) GetInitialResponse() *InitialPublishResponse {
	if x, ok := x.GetResponseType().(*PublishResponse_InitialResponse); ok {
		return x.InitialResponse
	}
	return nil
}

func (x *PublishResponse) GetMessageResponse() *MessagePublishResponse {
	if x, ok := x.GetResponseType().(*PublishResponse_MessageResponse); ok {
		return x.MessageResponse
	}
	return nil
}

type isPublishResponse_ResponseType interface {
	isPublishResponse_ResponseType()
}

type PublishResponse_InitialResponse struct {
	// Initial response on the stream.
	InitialResponse *InitialPublishResponse `protobuf:"bytes,1,opt,name=initial_response,json=initialResponse,proto3,oneof"`
}

type PublishResponse_MessageResponse struct {
	// Response to publishing messages.
	MessageResponse *MessagePublishResponse `protobuf:"bytes,2,opt,name=message_response,json=messageResponse,proto3,oneof"`
}

func (*PublishResponse_InitialResponse) isPublishResponse_ResponseType() {}

func (*PublishResponse_MessageResponse) isPublishResponse_ResponseType() {}

var File_google_cloud_pubsublite_v1_publisher_proto protoreflect.FileDescriptor

var file_google_cloud_pubsublite_v1_publisher_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x15, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x16, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5e, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x5f, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x22, 0xeb, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x6b, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xe4, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x32, 0xcb, 0x01, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xd8, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c,
	0x69, 0x74, 0x65, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x74, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x74, 0x65, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_pubsublite_v1_publisher_proto_rawDescOnce sync.Once
	file_google_cloud_pubsublite_v1_publisher_proto_rawDescData = file_google_cloud_pubsublite_v1_publisher_proto_rawDesc
)

func file_google_cloud_pubsublite_v1_publisher_proto_rawDescGZIP() []byte {
	file_google_cloud_pubsublite_v1_publisher_proto_rawDescOnce.Do(func() {
		file_google_cloud_pubsublite_v1_publisher_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_pubsublite_v1_publisher_proto_rawDescData)
	})
	return file_google_cloud_pubsublite_v1_publisher_proto_rawDescData
}

var file_google_cloud_pubsublite_v1_publisher_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_pubsublite_v1_publisher_proto_goTypes = []interface{}{
	(*InitialPublishRequest)(nil),  // 0: google.cloud.pubsublite.v1.InitialPublishRequest
	(*InitialPublishResponse)(nil), // 1: google.cloud.pubsublite.v1.InitialPublishResponse
	(*MessagePublishRequest)(nil),  // 2: google.cloud.pubsublite.v1.MessagePublishRequest
	(*MessagePublishResponse)(nil), // 3: google.cloud.pubsublite.v1.MessagePublishResponse
	(*PublishRequest)(nil),         // 4: google.cloud.pubsublite.v1.PublishRequest
	(*PublishResponse)(nil),        // 5: google.cloud.pubsublite.v1.PublishResponse
	(*PubSubMessage)(nil),          // 6: google.cloud.pubsublite.v1.PubSubMessage
	(*Cursor)(nil),                 // 7: google.cloud.pubsublite.v1.Cursor
}
var file_google_cloud_pubsublite_v1_publisher_proto_depIdxs = []int32{
	6, // 0: google.cloud.pubsublite.v1.MessagePublishRequest.messages:type_name -> google.cloud.pubsublite.v1.PubSubMessage
	7, // 1: google.cloud.pubsublite.v1.MessagePublishResponse.start_cursor:type_name -> google.cloud.pubsublite.v1.Cursor
	0, // 2: google.cloud.pubsublite.v1.PublishRequest.initial_request:type_name -> google.cloud.pubsublite.v1.InitialPublishRequest
	2, // 3: google.cloud.pubsublite.v1.PublishRequest.message_publish_request:type_name -> google.cloud.pubsublite.v1.MessagePublishRequest
	1, // 4: google.cloud.pubsublite.v1.PublishResponse.initial_response:type_name -> google.cloud.pubsublite.v1.InitialPublishResponse
	3, // 5: google.cloud.pubsublite.v1.PublishResponse.message_response:type_name -> google.cloud.pubsublite.v1.MessagePublishResponse
	4, // 6: google.cloud.pubsublite.v1.PublisherService.Publish:input_type -> google.cloud.pubsublite.v1.PublishRequest
	5, // 7: google.cloud.pubsublite.v1.PublisherService.Publish:output_type -> google.cloud.pubsublite.v1.PublishResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_pubsublite_v1_publisher_proto_init() }
func file_google_cloud_pubsublite_v1_publisher_proto_init() {
	if File_google_cloud_pubsublite_v1_publisher_proto != nil {
		return
	}
	file_google_cloud_pubsublite_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitialPublishRequest); i {
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
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitialPublishResponse); i {
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
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePublishRequest); i {
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
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePublishResponse); i {
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
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
	file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PublishRequest_InitialRequest)(nil),
		(*PublishRequest_MessagePublishRequest)(nil),
	}
	file_google_cloud_pubsublite_v1_publisher_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*PublishResponse_InitialResponse)(nil),
		(*PublishResponse_MessageResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_pubsublite_v1_publisher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_pubsublite_v1_publisher_proto_goTypes,
		DependencyIndexes: file_google_cloud_pubsublite_v1_publisher_proto_depIdxs,
		MessageInfos:      file_google_cloud_pubsublite_v1_publisher_proto_msgTypes,
	}.Build()
	File_google_cloud_pubsublite_v1_publisher_proto = out.File
	file_google_cloud_pubsublite_v1_publisher_proto_rawDesc = nil
	file_google_cloud_pubsublite_v1_publisher_proto_goTypes = nil
	file_google_cloud_pubsublite_v1_publisher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublisherServiceClient is the client API for PublisherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherServiceClient interface {
	// Establishes a stream with the server for publishing messages. Once the
	// stream is initialized, the client publishes messages by sending publish
	// requests on the stream. The server responds with a PublishResponse for each
	// PublishRequest sent by the client, in the same order that the requests
	// were sent. Note that multiple PublishRequests can be in flight
	// simultaneously, but they will be processed by the server in the order that
	// they are sent by the client on a given stream.
	Publish(ctx context.Context, opts ...grpc.CallOption) (PublisherService_PublishClient, error)
}

type publisherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherServiceClient(cc grpc.ClientConnInterface) PublisherServiceClient {
	return &publisherServiceClient{cc}
}

func (c *publisherServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (PublisherService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PublisherService_serviceDesc.Streams[0], "/google.cloud.pubsublite.v1.PublisherService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherServicePublishClient{stream}
	return x, nil
}

type PublisherService_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type publisherServicePublishClient struct {
	grpc.ClientStream
}

func (x *publisherServicePublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *publisherServicePublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherServiceServer is the server API for PublisherService service.
type PublisherServiceServer interface {
	// Establishes a stream with the server for publishing messages. Once the
	// stream is initialized, the client publishes messages by sending publish
	// requests on the stream. The server responds with a PublishResponse for each
	// PublishRequest sent by the client, in the same order that the requests
	// were sent. Note that multiple PublishRequests can be in flight
	// simultaneously, but they will be processed by the server in the order that
	// they are sent by the client on a given stream.
	Publish(PublisherService_PublishServer) error
}

// UnimplementedPublisherServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherServiceServer struct {
}

func (*UnimplementedPublisherServiceServer) Publish(PublisherService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

func RegisterPublisherServiceServer(s *grpc.Server, srv PublisherServiceServer) {
	s.RegisterService(&_PublisherService_serviceDesc, srv)
}

func _PublisherService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PublisherServiceServer).Publish(&publisherServicePublishServer{stream})
}

type PublisherService_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type publisherServicePublishServer struct {
	grpc.ServerStream
}

func (x *publisherServicePublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *publisherServicePublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PublisherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.pubsublite.v1.PublisherService",
	HandlerType: (*PublisherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _PublisherService_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/cloud/pubsublite/v1/publisher.proto",
}
