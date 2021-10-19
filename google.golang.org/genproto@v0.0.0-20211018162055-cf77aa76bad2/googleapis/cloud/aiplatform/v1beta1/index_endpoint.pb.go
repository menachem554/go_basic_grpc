// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1beta1/index_endpoint.proto

package aiplatform

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Indexes are deployed into it. An IndexEndpoint can have multiple
// DeployedIndexes.
type IndexEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the IndexEndpoint.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the IndexEndpoint.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the IndexEndpoint.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The indexes deployed in this endpoint.
	DeployedIndexes []*DeployedIndex `protobuf:"bytes,4,rep,name=deployed_indexes,json=deployedIndexes,proto3" json:"deployed_indexes,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,5,opt,name=etag,proto3" json:"etag,omitempty"`
	// The labels with user-defined metadata to organize your IndexEndpoints.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	//
	// See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Timestamp when this IndexEndpoint was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this IndexEndpoint was last updated.
	// This timestamp is not updated when the endpoint's DeployedIndexes are
	// updated, e.g. due to updates of the original Indexes they are the
	// deployments of.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. Immutable. The full name of the Google Compute Engine
	// [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks)
	// to which the IndexEndpoint should be peered.
	//
	// Private services access must already be configured for the network. If left
	// unspecified, the Endpoint is not peered with any network.
	//
	// [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert):
	// projects/{project}/global/networks/{network}.
	// Where {project} is a project number, as in '12345', and {network} is
	// network name.
	Network string `protobuf:"bytes,9,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *IndexEndpoint) Reset() {
	*x = IndexEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEndpoint) ProtoMessage() {}

func (x *IndexEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEndpoint.ProtoReflect.Descriptor instead.
func (*IndexEndpoint) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *IndexEndpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IndexEndpoint) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *IndexEndpoint) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IndexEndpoint) GetDeployedIndexes() []*DeployedIndex {
	if x != nil {
		return x.DeployedIndexes
	}
	return nil
}

func (x *IndexEndpoint) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *IndexEndpoint) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *IndexEndpoint) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *IndexEndpoint) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *IndexEndpoint) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

// A deployment of an Index. IndexEndpoints contain one or more DeployedIndexes.
type DeployedIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The user specified ID of the DeployedIndex.
	// The ID can be up to 128 characters long and must start with a letter and
	// only contain letters, numbers, and underscores.
	// The ID must be unique within the project it is created in.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. The name of the Index this is the deployment of.
	// We may refer to this Index as the DeployedIndex's "original" Index.
	Index string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	// The display name of the DeployedIndex. If not provided upon creation,
	// the Index's display_name is used.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Timestamp when the DeployedIndex was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Provides paths for users to send requests directly to the deployed index
	// services running on Cloud via private services access. This field is
	// populated if [network][google.cloud.aiplatform.v1beta1.IndexEndpoint.network] is configured.
	PrivateEndpoints *IndexPrivateEndpoints `protobuf:"bytes,5,opt,name=private_endpoints,json=privateEndpoints,proto3" json:"private_endpoints,omitempty"`
	// Output only. The DeployedIndex may depend on various data on its original Index.
	// Additionally when certain changes to the original Index are being done
	// (e.g. when what the Index contains is being changed) the DeployedIndex may
	// be asynchronously updated in the background to reflect this changes.
	// If this timestamp's value is at least the [Index.update_time][google.cloud.aiplatform.v1beta1.Index.update_time] of the
	// original Index, it means that this DeployedIndex and the original Index are
	// in sync. If this timestamp is older, then to see which updates this
	// DeployedIndex already contains (and which not), one must
	// [list][Operations.ListOperations] [Operations][Operation]
	// [working][Operation.name] on the original Index. Only
	// the successfully completed Operations with
	// [Operations.metadata.generic_metadata.update_time]
	// [google.cloud.aiplatform.v1beta1.GenericOperationMetadata.update_time]
	// equal or before this sync time are contained in this DeployedIndex.
	IndexSyncTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=index_sync_time,json=indexSyncTime,proto3" json:"index_sync_time,omitempty"`
	// Optional. A description of resources that the DeployedIndex uses, which to large
	// degree are decided by Vertex AI, and optionally allows only a modest
	// additional configuration.
	// If min_replica_count is not set, the default value is 1. If
	// max_replica_count is not set, the default value is min_replica_count. The
	// max allowed replica count is 1000.
	// The user is billed for the resources (at least their minimal amount)
	// even if the DeployedIndex receives no traffic.
	AutomaticResources *AutomaticResources `protobuf:"bytes,7,opt,name=automatic_resources,json=automaticResources,proto3" json:"automatic_resources,omitempty"`
	// Optional. If true, private endpoint's access logs are sent to StackDriver Logging.
	//
	// These logs are like standard server access logs, containing
	// information like timestamp and latency for each MatchRequest.
	//
	// Note that Stackdriver logs may incur a cost, especially if the deployed
	// index receives a high queries per second rate (QPS).
	// Estimate your costs before enabling this option.
	EnableAccessLogging bool `protobuf:"varint,8,opt,name=enable_access_logging,json=enableAccessLogging,proto3" json:"enable_access_logging,omitempty"`
	// Optional. If set, the authentication is enabled for the private endpoint.
	DeployedIndexAuthConfig *DeployedIndexAuthConfig `protobuf:"bytes,9,opt,name=deployed_index_auth_config,json=deployedIndexAuthConfig,proto3" json:"deployed_index_auth_config,omitempty"`
	// Optional. A list of reserved ip ranges under the VPC network that can be
	// used for this DeployedIndex.
	//
	// If set, we will deploy the index within the provided ip ranges. Otherwise,
	// the index might be deployed to any ip ranges under the provided VPC
	// network.
	//
	// The value sohuld be the name of the address
	// (https://cloud.google.com/compute/docs/reference/rest/v1/addresses)
	// Example: 'vertex-ai-ip-range'.
	ReservedIpRanges []string `protobuf:"bytes,10,rep,name=reserved_ip_ranges,json=reservedIpRanges,proto3" json:"reserved_ip_ranges,omitempty"`
	// Optional. The deployment group can be no longer than 64 characters (eg:
	// 'test', 'prod'). If not set, we will use the 'default' deployment group.
	//
	// Creating `deployment_groups` with `reserved_ip_ranges` is a recommended
	// practice when the peered network has multiple peering ranges. This creates
	// your deployments from predictable IP spaces for easier traffic
	// administration. Also, one deployment_group (except 'default') can only be
	// used with the same reserved_ip_ranges which means if the deployment_group
	// has been used with reserved_ip_ranges: [a, b, c], using it with [a, b] or
	// [d, e] is disallowed.
	//
	// Note: we only support up to 5 deployment groups(not including 'default').
	DeploymentGroup string `protobuf:"bytes,11,opt,name=deployment_group,json=deploymentGroup,proto3" json:"deployment_group,omitempty"`
}

func (x *DeployedIndex) Reset() {
	*x = DeployedIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedIndex) ProtoMessage() {}

func (x *DeployedIndex) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedIndex.ProtoReflect.Descriptor instead.
func (*DeployedIndex) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *DeployedIndex) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeployedIndex) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *DeployedIndex) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DeployedIndex) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DeployedIndex) GetPrivateEndpoints() *IndexPrivateEndpoints {
	if x != nil {
		return x.PrivateEndpoints
	}
	return nil
}

func (x *DeployedIndex) GetIndexSyncTime() *timestamppb.Timestamp {
	if x != nil {
		return x.IndexSyncTime
	}
	return nil
}

func (x *DeployedIndex) GetAutomaticResources() *AutomaticResources {
	if x != nil {
		return x.AutomaticResources
	}
	return nil
}

func (x *DeployedIndex) GetEnableAccessLogging() bool {
	if x != nil {
		return x.EnableAccessLogging
	}
	return false
}

func (x *DeployedIndex) GetDeployedIndexAuthConfig() *DeployedIndexAuthConfig {
	if x != nil {
		return x.DeployedIndexAuthConfig
	}
	return nil
}

func (x *DeployedIndex) GetReservedIpRanges() []string {
	if x != nil {
		return x.ReservedIpRanges
	}
	return nil
}

func (x *DeployedIndex) GetDeploymentGroup() string {
	if x != nil {
		return x.DeploymentGroup
	}
	return ""
}

// Used to set up the auth on the DeployedIndex's private endpoint.
type DeployedIndexAuthConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the authentication provider that the DeployedIndex uses.
	AuthProvider *DeployedIndexAuthConfig_AuthProvider `protobuf:"bytes,1,opt,name=auth_provider,json=authProvider,proto3" json:"auth_provider,omitempty"`
}

func (x *DeployedIndexAuthConfig) Reset() {
	*x = DeployedIndexAuthConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedIndexAuthConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedIndexAuthConfig) ProtoMessage() {}

func (x *DeployedIndexAuthConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedIndexAuthConfig.ProtoReflect.Descriptor instead.
func (*DeployedIndexAuthConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *DeployedIndexAuthConfig) GetAuthProvider() *DeployedIndexAuthConfig_AuthProvider {
	if x != nil {
		return x.AuthProvider
	}
	return nil
}

// IndexPrivateEndpoints proto is used to provide paths for users to send
// requests via private services access.
type IndexPrivateEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The ip address used to send match gRPC requests.
	MatchGrpcAddress string `protobuf:"bytes,1,opt,name=match_grpc_address,json=matchGrpcAddress,proto3" json:"match_grpc_address,omitempty"`
}

func (x *IndexPrivateEndpoints) Reset() {
	*x = IndexPrivateEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexPrivateEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexPrivateEndpoints) ProtoMessage() {}

func (x *IndexPrivateEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexPrivateEndpoints.ProtoReflect.Descriptor instead.
func (*IndexPrivateEndpoints) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *IndexPrivateEndpoints) GetMatchGrpcAddress() string {
	if x != nil {
		return x.MatchGrpcAddress
	}
	return ""
}

// Configuration for an authentication provider, including support for
// [JSON Web Token
// (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type DeployedIndexAuthConfig_AuthProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of JWT
	// [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	// that are allowed to access. A JWT containing any of these audiences will
	// be accepted.
	Audiences []string `protobuf:"bytes,1,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// A list of allowed JWT issuers. Each entry must be a valid Google
	// service account, in the following format:
	//
	// `service-account-name@project-id.iam.gserviceaccount.com`
	AllowedIssuers []string `protobuf:"bytes,2,rep,name=allowed_issuers,json=allowedIssuers,proto3" json:"allowed_issuers,omitempty"`
}

func (x *DeployedIndexAuthConfig_AuthProvider) Reset() {
	*x = DeployedIndexAuthConfig_AuthProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedIndexAuthConfig_AuthProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedIndexAuthConfig_AuthProvider) ProtoMessage() {}

func (x *DeployedIndexAuthConfig_AuthProvider) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedIndexAuthConfig_AuthProvider.ProtoReflect.Descriptor instead.
func (*DeployedIndexAuthConfig_AuthProvider) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DeployedIndexAuthConfig_AuthProvider) GetAudiences() []string {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *DeployedIndexAuthConfig_AuthProvider) GetAllowedIssuers() []string {
	if x != nil {
		return x.AllowedIssuers
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_index_endpoint_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x05, 0x0a, 0x0d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a,
	0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x12, 0x52, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0,
	0x41, 0x05, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x75, 0xea, 0x41, 0x72, 0x0a, 0x27, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x47, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x7d, 0x22, 0xfe, 0x05,
	0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x27, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x47, 0x0a, 0x0f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x69, 0x0a, 0x13, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x7a, 0x0a, 0x1a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x17, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x31, 0x0a, 0x12, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xdc,
	0x01, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6a, 0x0a, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x55, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x22, 0x4a, 0x0a,
	0x15, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x12, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x47, 0x72,
	0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0xef, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x12, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_goTypes = []interface{}{
	(*IndexEndpoint)(nil),           // 0: google.cloud.aiplatform.v1beta1.IndexEndpoint
	(*DeployedIndex)(nil),           // 1: google.cloud.aiplatform.v1beta1.DeployedIndex
	(*DeployedIndexAuthConfig)(nil), // 2: google.cloud.aiplatform.v1beta1.DeployedIndexAuthConfig
	(*IndexPrivateEndpoints)(nil),   // 3: google.cloud.aiplatform.v1beta1.IndexPrivateEndpoints
	nil,                             // 4: google.cloud.aiplatform.v1beta1.IndexEndpoint.LabelsEntry
	(*DeployedIndexAuthConfig_AuthProvider)(nil), // 5: google.cloud.aiplatform.v1beta1.DeployedIndexAuthConfig.AuthProvider
	(*timestamppb.Timestamp)(nil),                // 6: google.protobuf.Timestamp
	(*AutomaticResources)(nil),                   // 7: google.cloud.aiplatform.v1beta1.AutomaticResources
}
var file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_depIdxs = []int32{
	1,  // 0: google.cloud.aiplatform.v1beta1.IndexEndpoint.deployed_indexes:type_name -> google.cloud.aiplatform.v1beta1.DeployedIndex
	4,  // 1: google.cloud.aiplatform.v1beta1.IndexEndpoint.labels:type_name -> google.cloud.aiplatform.v1beta1.IndexEndpoint.LabelsEntry
	6,  // 2: google.cloud.aiplatform.v1beta1.IndexEndpoint.create_time:type_name -> google.protobuf.Timestamp
	6,  // 3: google.cloud.aiplatform.v1beta1.IndexEndpoint.update_time:type_name -> google.protobuf.Timestamp
	6,  // 4: google.cloud.aiplatform.v1beta1.DeployedIndex.create_time:type_name -> google.protobuf.Timestamp
	3,  // 5: google.cloud.aiplatform.v1beta1.DeployedIndex.private_endpoints:type_name -> google.cloud.aiplatform.v1beta1.IndexPrivateEndpoints
	6,  // 6: google.cloud.aiplatform.v1beta1.DeployedIndex.index_sync_time:type_name -> google.protobuf.Timestamp
	7,  // 7: google.cloud.aiplatform.v1beta1.DeployedIndex.automatic_resources:type_name -> google.cloud.aiplatform.v1beta1.AutomaticResources
	2,  // 8: google.cloud.aiplatform.v1beta1.DeployedIndex.deployed_index_auth_config:type_name -> google.cloud.aiplatform.v1beta1.DeployedIndexAuthConfig
	5,  // 9: google.cloud.aiplatform.v1beta1.DeployedIndexAuthConfig.auth_provider:type_name -> google.cloud.aiplatform.v1beta1.DeployedIndexAuthConfig.AuthProvider
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_init() }
func file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_index_endpoint_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_index_proto_init()
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEndpoint); i {
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
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedIndex); i {
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
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedIndexAuthConfig); i {
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
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexPrivateEndpoints); i {
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
		file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedIndexAuthConfig_AuthProvider); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_index_endpoint_proto = out.File
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_index_endpoint_proto_depIdxs = nil
}
