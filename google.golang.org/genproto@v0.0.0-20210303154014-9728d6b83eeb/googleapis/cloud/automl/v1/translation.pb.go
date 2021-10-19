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
// source: google/cloud/automl/v1/translation.proto

package automl

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Dataset metadata that is specific to translation.
type TranslationDatasetMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The BCP-47 language code of the source language.
	SourceLanguageCode string `protobuf:"bytes,1,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Required. The BCP-47 language code of the target language.
	TargetLanguageCode string `protobuf:"bytes,2,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
}

func (x *TranslationDatasetMetadata) Reset() {
	*x = TranslationDatasetMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationDatasetMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationDatasetMetadata) ProtoMessage() {}

func (x *TranslationDatasetMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationDatasetMetadata.ProtoReflect.Descriptor instead.
func (*TranslationDatasetMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_translation_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationDatasetMetadata) GetSourceLanguageCode() string {
	if x != nil {
		return x.SourceLanguageCode
	}
	return ""
}

func (x *TranslationDatasetMetadata) GetTargetLanguageCode() string {
	if x != nil {
		return x.TargetLanguageCode
	}
	return ""
}

// Evaluation metrics for the dataset.
type TranslationEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. BLEU score.
	BleuScore float64 `protobuf:"fixed64,1,opt,name=bleu_score,json=bleuScore,proto3" json:"bleu_score,omitempty"`
	// Output only. BLEU score for base model.
	BaseBleuScore float64 `protobuf:"fixed64,2,opt,name=base_bleu_score,json=baseBleuScore,proto3" json:"base_bleu_score,omitempty"`
}

func (x *TranslationEvaluationMetrics) Reset() {
	*x = TranslationEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationEvaluationMetrics) ProtoMessage() {}

func (x *TranslationEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*TranslationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_translation_proto_rawDescGZIP(), []int{1}
}

func (x *TranslationEvaluationMetrics) GetBleuScore() float64 {
	if x != nil {
		return x.BleuScore
	}
	return 0
}

func (x *TranslationEvaluationMetrics) GetBaseBleuScore() float64 {
	if x != nil {
		return x.BaseBleuScore
	}
	return 0
}

// Model metadata that is specific to translation.
type TranslationModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the model to use as a baseline to train the custom
	// model. If unset, we use the default base model provided by Google
	// Translate. Format:
	// `projects/{project_id}/locations/{location_id}/models/{model_id}`
	BaseModel string `protobuf:"bytes,1,opt,name=base_model,json=baseModel,proto3" json:"base_model,omitempty"`
	// Output only. Inferred from the dataset.
	// The source language (The BCP-47 language code) that is used for training.
	SourceLanguageCode string `protobuf:"bytes,2,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// Output only. The target language (The BCP-47 language code) that is used
	// for training.
	TargetLanguageCode string `protobuf:"bytes,3,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
}

func (x *TranslationModelMetadata) Reset() {
	*x = TranslationModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationModelMetadata) ProtoMessage() {}

func (x *TranslationModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationModelMetadata.ProtoReflect.Descriptor instead.
func (*TranslationModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_translation_proto_rawDescGZIP(), []int{2}
}

func (x *TranslationModelMetadata) GetBaseModel() string {
	if x != nil {
		return x.BaseModel
	}
	return ""
}

func (x *TranslationModelMetadata) GetSourceLanguageCode() string {
	if x != nil {
		return x.SourceLanguageCode
	}
	return ""
}

func (x *TranslationModelMetadata) GetTargetLanguageCode() string {
	if x != nil {
		return x.TargetLanguageCode
	}
	return ""
}

// Annotation details specific to translation.
type TranslationAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only . The translated content.
	TranslatedContent *TextSnippet `protobuf:"bytes,1,opt,name=translated_content,json=translatedContent,proto3" json:"translated_content,omitempty"`
}

func (x *TranslationAnnotation) Reset() {
	*x = TranslationAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationAnnotation) ProtoMessage() {}

func (x *TranslationAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_translation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationAnnotation.ProtoReflect.Descriptor instead.
func (*TranslationAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_translation_proto_rawDescGZIP(), []int{3}
}

func (x *TranslationAnnotation) GetTranslatedContent() *TextSnippet {
	if x != nil {
		return x.TranslatedContent
	}
	return nil
}

var File_google_cloud_automl_v1_translation_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1_translation_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x1a, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x14, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x12, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x35, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x65, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x65, 0x75, 0x5f,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x62, 0x6c, 0x65,
	0x75, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x62,
	0x6c, 0x65, 0x75, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x62, 0x61, 0x73, 0x65, 0x42, 0x6c, 0x65, 0x75, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x9d,
	0x01, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x6b,
	0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0xbc, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0xaa, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f,
	0x4d, 0x4c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1_translation_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1_translation_proto_rawDescData = file_google_cloud_automl_v1_translation_proto_rawDesc
)

func file_google_cloud_automl_v1_translation_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1_translation_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1_translation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1_translation_proto_rawDescData)
	})
	return file_google_cloud_automl_v1_translation_proto_rawDescData
}

var file_google_cloud_automl_v1_translation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_automl_v1_translation_proto_goTypes = []interface{}{
	(*TranslationDatasetMetadata)(nil),   // 0: google.cloud.automl.v1.TranslationDatasetMetadata
	(*TranslationEvaluationMetrics)(nil), // 1: google.cloud.automl.v1.TranslationEvaluationMetrics
	(*TranslationModelMetadata)(nil),     // 2: google.cloud.automl.v1.TranslationModelMetadata
	(*TranslationAnnotation)(nil),        // 3: google.cloud.automl.v1.TranslationAnnotation
	(*TextSnippet)(nil),                  // 4: google.cloud.automl.v1.TextSnippet
}
var file_google_cloud_automl_v1_translation_proto_depIdxs = []int32{
	4, // 0: google.cloud.automl.v1.TranslationAnnotation.translated_content:type_name -> google.cloud.automl.v1.TextSnippet
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1_translation_proto_init() }
func file_google_cloud_automl_v1_translation_proto_init() {
	if File_google_cloud_automl_v1_translation_proto != nil {
		return
	}
	file_google_cloud_automl_v1_data_items_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1_translation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationDatasetMetadata); i {
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
		file_google_cloud_automl_v1_translation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationEvaluationMetrics); i {
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
		file_google_cloud_automl_v1_translation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationModelMetadata); i {
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
		file_google_cloud_automl_v1_translation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationAnnotation); i {
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
			RawDescriptor: file_google_cloud_automl_v1_translation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1_translation_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1_translation_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1_translation_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1_translation_proto = out.File
	file_google_cloud_automl_v1_translation_proto_rawDesc = nil
	file_google_cloud_automl_v1_translation_proto_goTypes = nil
	file_google_cloud_automl_v1_translation_proto_depIdxs = nil
}
