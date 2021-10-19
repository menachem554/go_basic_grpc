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
// source: google/cloud/osconfig/agentendpoint/v1/config_common.proto

package agentendpoint

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Supported OSPolicy compliance states.
type OSPolicyComplianceState int32

const (
	// Default value. This value is unused.
	OSPolicyComplianceState_OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED OSPolicyComplianceState = 0
	// Compliant state.
	OSPolicyComplianceState_COMPLIANT OSPolicyComplianceState = 1
	// Non-compliant state
	OSPolicyComplianceState_NON_COMPLIANT OSPolicyComplianceState = 2
	// Unknown compliance state.
	OSPolicyComplianceState_UNKNOWN OSPolicyComplianceState = 3
	// No applicable OS policies were found for the instance.
	// This state is only applicable to the instance.
	OSPolicyComplianceState_NO_OS_POLICIES_APPLICABLE OSPolicyComplianceState = 4
)

// Enum value maps for OSPolicyComplianceState.
var (
	OSPolicyComplianceState_name = map[int32]string{
		0: "OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED",
		1: "COMPLIANT",
		2: "NON_COMPLIANT",
		3: "UNKNOWN",
		4: "NO_OS_POLICIES_APPLICABLE",
	}
	OSPolicyComplianceState_value = map[string]int32{
		"OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED": 0,
		"COMPLIANT":                              1,
		"NON_COMPLIANT":                          2,
		"UNKNOWN":                                3,
		"NO_OS_POLICIES_APPLICABLE":              4,
	}
)

func (x OSPolicyComplianceState) Enum() *OSPolicyComplianceState {
	p := new(OSPolicyComplianceState)
	*p = x
	return p
}

func (x OSPolicyComplianceState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OSPolicyComplianceState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[0].Descriptor()
}

func (OSPolicyComplianceState) Type() protoreflect.EnumType {
	return &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[0]
}

func (x OSPolicyComplianceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OSPolicyComplianceState.Descriptor instead.
func (OSPolicyComplianceState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{0}
}

// Supported configuration step types
type OSPolicyResourceConfigStep_Type int32

const (
	// Default value. This value is unused.
	OSPolicyResourceConfigStep_TYPE_UNSPECIFIED OSPolicyResourceConfigStep_Type = 0
	// Validation to detect resource conflicts, schema errors, etc.
	OSPolicyResourceConfigStep_VALIDATION OSPolicyResourceConfigStep_Type = 1
	// Check the current desired state status of the resource.
	OSPolicyResourceConfigStep_DESIRED_STATE_CHECK OSPolicyResourceConfigStep_Type = 2
	// Enforce the desired state for a resource that is not in desired state.
	OSPolicyResourceConfigStep_DESIRED_STATE_ENFORCEMENT OSPolicyResourceConfigStep_Type = 3
	// Re-check desired state status for a resource after enforcement of all
	// resources in the current configuration run.
	//
	// This step is used to determine the final desired state status for the
	// resource. It accounts for any resources that might have drifted from
	// their desired state due to side effects from configuring other resources
	// during the current configuration run.
	OSPolicyResourceConfigStep_DESIRED_STATE_CHECK_POST_ENFORCEMENT OSPolicyResourceConfigStep_Type = 4
)

// Enum value maps for OSPolicyResourceConfigStep_Type.
var (
	OSPolicyResourceConfigStep_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "VALIDATION",
		2: "DESIRED_STATE_CHECK",
		3: "DESIRED_STATE_ENFORCEMENT",
		4: "DESIRED_STATE_CHECK_POST_ENFORCEMENT",
	}
	OSPolicyResourceConfigStep_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":                     0,
		"VALIDATION":                           1,
		"DESIRED_STATE_CHECK":                  2,
		"DESIRED_STATE_ENFORCEMENT":            3,
		"DESIRED_STATE_CHECK_POST_ENFORCEMENT": 4,
	}
)

func (x OSPolicyResourceConfigStep_Type) Enum() *OSPolicyResourceConfigStep_Type {
	p := new(OSPolicyResourceConfigStep_Type)
	*p = x
	return p
}

func (x OSPolicyResourceConfigStep_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OSPolicyResourceConfigStep_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[1].Descriptor()
}

func (OSPolicyResourceConfigStep_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[1]
}

func (x OSPolicyResourceConfigStep_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OSPolicyResourceConfigStep_Type.Descriptor instead.
func (OSPolicyResourceConfigStep_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{0, 0}
}

// Supported outcomes for a configuration step.
type OSPolicyResourceConfigStep_Outcome int32

const (
	// Default value. This value is unused.
	OSPolicyResourceConfigStep_OUTCOME_UNSPECIFIED OSPolicyResourceConfigStep_Outcome = 0
	// The step succeeded.
	OSPolicyResourceConfigStep_SUCCEEDED OSPolicyResourceConfigStep_Outcome = 1
	// The step failed.
	OSPolicyResourceConfigStep_FAILED OSPolicyResourceConfigStep_Outcome = 2
)

// Enum value maps for OSPolicyResourceConfigStep_Outcome.
var (
	OSPolicyResourceConfigStep_Outcome_name = map[int32]string{
		0: "OUTCOME_UNSPECIFIED",
		1: "SUCCEEDED",
		2: "FAILED",
	}
	OSPolicyResourceConfigStep_Outcome_value = map[string]int32{
		"OUTCOME_UNSPECIFIED": 0,
		"SUCCEEDED":           1,
		"FAILED":              2,
	}
)

func (x OSPolicyResourceConfigStep_Outcome) Enum() *OSPolicyResourceConfigStep_Outcome {
	p := new(OSPolicyResourceConfigStep_Outcome)
	*p = x
	return p
}

func (x OSPolicyResourceConfigStep_Outcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OSPolicyResourceConfigStep_Outcome) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[2].Descriptor()
}

func (OSPolicyResourceConfigStep_Outcome) Type() protoreflect.EnumType {
	return &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes[2]
}

func (x OSPolicyResourceConfigStep_Outcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OSPolicyResourceConfigStep_Outcome.Descriptor instead.
func (OSPolicyResourceConfigStep_Outcome) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{0, 1}
}

// Step performed by the OS Config agent for configuring an `OSPolicyResource`
// to its desired state.
type OSPolicyResourceConfigStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration step type.
	Type OSPolicyResourceConfigStep_Type `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep_Type" json:"type,omitempty"`
	// Outcome of the configuration step.
	Outcome OSPolicyResourceConfigStep_Outcome `protobuf:"varint,2,opt,name=outcome,proto3,enum=google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep_Outcome" json:"outcome,omitempty"`
	// An error message recorded during the execution of this step.
	// Only populated when outcome is FAILED.
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *OSPolicyResourceConfigStep) Reset() {
	*x = OSPolicyResourceConfigStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSPolicyResourceConfigStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSPolicyResourceConfigStep) ProtoMessage() {}

func (x *OSPolicyResourceConfigStep) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSPolicyResourceConfigStep.ProtoReflect.Descriptor instead.
func (*OSPolicyResourceConfigStep) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{0}
}

func (x *OSPolicyResourceConfigStep) GetType() OSPolicyResourceConfigStep_Type {
	if x != nil {
		return x.Type
	}
	return OSPolicyResourceConfigStep_TYPE_UNSPECIFIED
}

func (x *OSPolicyResourceConfigStep) GetOutcome() OSPolicyResourceConfigStep_Outcome {
	if x != nil {
		return x.Outcome
	}
	return OSPolicyResourceConfigStep_OUTCOME_UNSPECIFIED
}

func (x *OSPolicyResourceConfigStep) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Compliance data for an OS policy resource.
type OSPolicyResourceCompliance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the OS policy resource.
	OsPolicyResourceId string `protobuf:"bytes,1,opt,name=os_policy_resource_id,json=osPolicyResourceId,proto3" json:"os_policy_resource_id,omitempty"`
	// Ordered list of configuration steps taken by the agent for the OS policy
	// resource.
	ConfigSteps []*OSPolicyResourceConfigStep `protobuf:"bytes,2,rep,name=config_steps,json=configSteps,proto3" json:"config_steps,omitempty"`
	// Compliance state of the OS policy resource.
	State OSPolicyComplianceState `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.osconfig.agentendpoint.v1.OSPolicyComplianceState" json:"state,omitempty"`
	// Resource specific output.
	//
	// Types that are assignable to Output:
	//	*OSPolicyResourceCompliance_ExecResourceOutput_
	Output isOSPolicyResourceCompliance_Output `protobuf_oneof:"output"`
}

func (x *OSPolicyResourceCompliance) Reset() {
	*x = OSPolicyResourceCompliance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSPolicyResourceCompliance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSPolicyResourceCompliance) ProtoMessage() {}

func (x *OSPolicyResourceCompliance) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSPolicyResourceCompliance.ProtoReflect.Descriptor instead.
func (*OSPolicyResourceCompliance) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{1}
}

func (x *OSPolicyResourceCompliance) GetOsPolicyResourceId() string {
	if x != nil {
		return x.OsPolicyResourceId
	}
	return ""
}

func (x *OSPolicyResourceCompliance) GetConfigSteps() []*OSPolicyResourceConfigStep {
	if x != nil {
		return x.ConfigSteps
	}
	return nil
}

func (x *OSPolicyResourceCompliance) GetState() OSPolicyComplianceState {
	if x != nil {
		return x.State
	}
	return OSPolicyComplianceState_OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED
}

func (m *OSPolicyResourceCompliance) GetOutput() isOSPolicyResourceCompliance_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (x *OSPolicyResourceCompliance) GetExecResourceOutput() *OSPolicyResourceCompliance_ExecResourceOutput {
	if x, ok := x.GetOutput().(*OSPolicyResourceCompliance_ExecResourceOutput_); ok {
		return x.ExecResourceOutput
	}
	return nil
}

type isOSPolicyResourceCompliance_Output interface {
	isOSPolicyResourceCompliance_Output()
}

type OSPolicyResourceCompliance_ExecResourceOutput_ struct {
	// ExecResource specific output.
	ExecResourceOutput *OSPolicyResourceCompliance_ExecResourceOutput `protobuf:"bytes,4,opt,name=exec_resource_output,json=execResourceOutput,proto3,oneof"`
}

func (*OSPolicyResourceCompliance_ExecResourceOutput_) isOSPolicyResourceCompliance_Output() {}

// ExecResource specific output.
type OSPolicyResourceCompliance_ExecResourceOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output from Enforcement phase output file (if run).
	// Output size is limited to 100K bytes.
	EnforcementOutput []byte `protobuf:"bytes,2,opt,name=enforcement_output,json=enforcementOutput,proto3" json:"enforcement_output,omitempty"`
}

func (x *OSPolicyResourceCompliance_ExecResourceOutput) Reset() {
	*x = OSPolicyResourceCompliance_ExecResourceOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSPolicyResourceCompliance_ExecResourceOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSPolicyResourceCompliance_ExecResourceOutput) ProtoMessage() {}

func (x *OSPolicyResourceCompliance_ExecResourceOutput) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSPolicyResourceCompliance_ExecResourceOutput.ProtoReflect.Descriptor instead.
func (*OSPolicyResourceCompliance_ExecResourceOutput) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP(), []int{1, 0}
}

func (x *OSPolicyResourceCompliance_ExecResourceOutput) GetEnforcementOutput() []byte {
	if x != nil {
		return x.EnforcementOutput
	}
	return nil
}

var File_google_cloud_osconfig_agentendpoint_v1_config_common_proto protoreflect.FileDescriptor

var file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0xd4, 0x03, 0x0a, 0x1a, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x5b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x64, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x45,
	0x53, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43,
	0x4b, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x53, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x45, 0x53, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x45,
	0x4e, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x22, 0x3d, 0x0a, 0x07,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x55, 0x54, 0x43, 0x4f,
	0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0xe8, 0x03, 0x0a, 0x1a,
	0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x6f, 0x73,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x65, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x74, 0x65, 0x70, 0x73, 0x12, 0x55, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x14,
	0x65, 0x78, 0x65, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x48, 0x00, 0x52, 0x12, 0x65, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1a, 0x43, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2d, 0x0a,
	0x12, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x08, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2a, 0x93, 0x01, 0x0a, 0x17, 0x4f, 0x53, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x4f, 0x53, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x4e, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x12, 0x1d, 0x0a,
	0x19, 0x4e, 0x4f, 0x5f, 0x4f, 0x53, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x49, 0x45, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x42, 0x96, 0x01, 0x0a,
	0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescOnce sync.Once
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescData = file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDesc
)

func file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescGZIP() []byte {
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescData)
	})
	return file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDescData
}

var file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_goTypes = []interface{}{
	(OSPolicyComplianceState)(0),                          // 0: google.cloud.osconfig.agentendpoint.v1.OSPolicyComplianceState
	(OSPolicyResourceConfigStep_Type)(0),                  // 1: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.Type
	(OSPolicyResourceConfigStep_Outcome)(0),               // 2: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.Outcome
	(*OSPolicyResourceConfigStep)(nil),                    // 3: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep
	(*OSPolicyResourceCompliance)(nil),                    // 4: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance
	(*OSPolicyResourceCompliance_ExecResourceOutput)(nil), // 5: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance.ExecResourceOutput
}
var file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_depIdxs = []int32{
	1, // 0: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.type:type_name -> google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.Type
	2, // 1: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.outcome:type_name -> google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep.Outcome
	3, // 2: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance.config_steps:type_name -> google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceConfigStep
	0, // 3: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance.state:type_name -> google.cloud.osconfig.agentendpoint.v1.OSPolicyComplianceState
	5, // 4: google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance.exec_resource_output:type_name -> google.cloud.osconfig.agentendpoint.v1.OSPolicyResourceCompliance.ExecResourceOutput
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_init() }
func file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_init() {
	if File_google_cloud_osconfig_agentendpoint_v1_config_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSPolicyResourceConfigStep); i {
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
		file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSPolicyResourceCompliance); i {
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
		file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSPolicyResourceCompliance_ExecResourceOutput); i {
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
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*OSPolicyResourceCompliance_ExecResourceOutput_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_msgTypes,
	}.Build()
	File_google_cloud_osconfig_agentendpoint_v1_config_common_proto = out.File
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_rawDesc = nil
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_goTypes = nil
	file_google_cloud_osconfig_agentendpoint_v1_config_common_proto_depIdxs = nil
}
