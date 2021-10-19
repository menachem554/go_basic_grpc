// Copyright 2016 Google Inc.
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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/genomics/v1/cigar.proto

package genomics

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

// Describes the different types of CIGAR alignment operations that exist.
// Used wherever CIGAR alignments are used.
type CigarUnit_Operation int32

const (
	CigarUnit_OPERATION_UNSPECIFIED CigarUnit_Operation = 0
	// An alignment match indicates that a sequence can be aligned to the
	// reference without evidence of an INDEL. Unlike the
	// `SEQUENCE_MATCH` and `SEQUENCE_MISMATCH` operators,
	// the `ALIGNMENT_MATCH` operator does not indicate whether the
	// reference and read sequences are an exact match. This operator is
	// equivalent to SAM's `M`.
	CigarUnit_ALIGNMENT_MATCH CigarUnit_Operation = 1
	// The insert operator indicates that the read contains evidence of bases
	// being inserted into the reference. This operator is equivalent to SAM's
	// `I`.
	CigarUnit_INSERT CigarUnit_Operation = 2
	// The delete operator indicates that the read contains evidence of bases
	// being deleted from the reference. This operator is equivalent to SAM's
	// `D`.
	CigarUnit_DELETE CigarUnit_Operation = 3
	// The skip operator indicates that this read skips a long segment of the
	// reference, but the bases have not been deleted. This operator is commonly
	// used when working with RNA-seq data, where reads may skip long segments
	// of the reference between exons. This operator is equivalent to SAM's
	// `N`.
	CigarUnit_SKIP CigarUnit_Operation = 4
	// The soft clip operator indicates that bases at the start/end of a read
	// have not been considered during alignment. This may occur if the majority
	// of a read maps, except for low quality bases at the start/end of a read.
	// This operator is equivalent to SAM's `S`. Bases that are soft
	// clipped will still be stored in the read.
	CigarUnit_CLIP_SOFT CigarUnit_Operation = 5
	// The hard clip operator indicates that bases at the start/end of a read
	// have been omitted from this alignment. This may occur if this linear
	// alignment is part of a chimeric alignment, or if the read has been
	// trimmed (for example, during error correction or to trim poly-A tails for
	// RNA-seq). This operator is equivalent to SAM's `H`.
	CigarUnit_CLIP_HARD CigarUnit_Operation = 6
	// The pad operator indicates that there is padding in an alignment. This
	// operator is equivalent to SAM's `P`.
	CigarUnit_PAD CigarUnit_Operation = 7
	// This operator indicates that this portion of the aligned sequence exactly
	// matches the reference. This operator is equivalent to SAM's `=`.
	CigarUnit_SEQUENCE_MATCH CigarUnit_Operation = 8
	// This operator indicates that this portion of the aligned sequence is an
	// alignment match to the reference, but a sequence mismatch. This can
	// indicate a SNP or a read error. This operator is equivalent to SAM's
	// `X`.
	CigarUnit_SEQUENCE_MISMATCH CigarUnit_Operation = 9
)

// Enum value maps for CigarUnit_Operation.
var (
	CigarUnit_Operation_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "ALIGNMENT_MATCH",
		2: "INSERT",
		3: "DELETE",
		4: "SKIP",
		5: "CLIP_SOFT",
		6: "CLIP_HARD",
		7: "PAD",
		8: "SEQUENCE_MATCH",
		9: "SEQUENCE_MISMATCH",
	}
	CigarUnit_Operation_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"ALIGNMENT_MATCH":       1,
		"INSERT":                2,
		"DELETE":                3,
		"SKIP":                  4,
		"CLIP_SOFT":             5,
		"CLIP_HARD":             6,
		"PAD":                   7,
		"SEQUENCE_MATCH":        8,
		"SEQUENCE_MISMATCH":     9,
	}
)

func (x CigarUnit_Operation) Enum() *CigarUnit_Operation {
	p := new(CigarUnit_Operation)
	*p = x
	return p
}

func (x CigarUnit_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CigarUnit_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_google_genomics_v1_cigar_proto_enumTypes[0].Descriptor()
}

func (CigarUnit_Operation) Type() protoreflect.EnumType {
	return &file_google_genomics_v1_cigar_proto_enumTypes[0]
}

func (x CigarUnit_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CigarUnit_Operation.Descriptor instead.
func (CigarUnit_Operation) EnumDescriptor() ([]byte, []int) {
	return file_google_genomics_v1_cigar_proto_rawDescGZIP(), []int{0, 0}
}

// A single CIGAR operation.
type CigarUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation CigarUnit_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=google.genomics.v1.CigarUnit_Operation" json:"operation,omitempty"`
	// The number of genomic bases that the operation runs for. Required.
	OperationLength int64 `protobuf:"varint,2,opt,name=operation_length,json=operationLength,proto3" json:"operation_length,omitempty"`
	// `referenceSequence` is only used at mismatches
	// (`SEQUENCE_MISMATCH`) and deletions (`DELETE`).
	// Filling this field replaces SAM's MD tag. If the relevant information is
	// not available, this field is unset.
	ReferenceSequence string `protobuf:"bytes,3,opt,name=reference_sequence,json=referenceSequence,proto3" json:"reference_sequence,omitempty"`
}

func (x *CigarUnit) Reset() {
	*x = CigarUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_genomics_v1_cigar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CigarUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CigarUnit) ProtoMessage() {}

func (x *CigarUnit) ProtoReflect() protoreflect.Message {
	mi := &file_google_genomics_v1_cigar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CigarUnit.ProtoReflect.Descriptor instead.
func (*CigarUnit) Descriptor() ([]byte, []int) {
	return file_google_genomics_v1_cigar_proto_rawDescGZIP(), []int{0}
}

func (x *CigarUnit) GetOperation() CigarUnit_Operation {
	if x != nil {
		return x.Operation
	}
	return CigarUnit_OPERATION_UNSPECIFIED
}

func (x *CigarUnit) GetOperationLength() int64 {
	if x != nil {
		return x.OperationLength
	}
	return 0
}

func (x *CigarUnit) GetReferenceSequence() string {
	if x != nil {
		return x.ReferenceSequence
	}
	return ""
}

var File_google_genomics_v1_cigar_proto protoreflect.FileDescriptor

var file_google_genomics_v1_cigar_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x67, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xde, 0x02, 0x0a, 0x09, 0x43, 0x69, 0x67, 0x61, 0x72, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x45, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x67, 0x61, 0x72, 0x55, 0x6e,
	0x69, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x4c,
	0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x49, 0x50, 0x5f, 0x53, 0x4f, 0x46, 0x54, 0x10, 0x05,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x49, 0x50, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x10, 0x06, 0x12,
	0x07, 0x0a, 0x03, 0x50, 0x41, 0x44, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x51, 0x55,
	0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43,
	0x48, 0x10, 0x09, 0x42, 0x65, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x43,
	0x69, 0x67, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67,
	0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_genomics_v1_cigar_proto_rawDescOnce sync.Once
	file_google_genomics_v1_cigar_proto_rawDescData = file_google_genomics_v1_cigar_proto_rawDesc
)

func file_google_genomics_v1_cigar_proto_rawDescGZIP() []byte {
	file_google_genomics_v1_cigar_proto_rawDescOnce.Do(func() {
		file_google_genomics_v1_cigar_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_genomics_v1_cigar_proto_rawDescData)
	})
	return file_google_genomics_v1_cigar_proto_rawDescData
}

var file_google_genomics_v1_cigar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_genomics_v1_cigar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_genomics_v1_cigar_proto_goTypes = []interface{}{
	(CigarUnit_Operation)(0), // 0: google.genomics.v1.CigarUnit.Operation
	(*CigarUnit)(nil),        // 1: google.genomics.v1.CigarUnit
}
var file_google_genomics_v1_cigar_proto_depIdxs = []int32{
	0, // 0: google.genomics.v1.CigarUnit.operation:type_name -> google.genomics.v1.CigarUnit.Operation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_genomics_v1_cigar_proto_init() }
func file_google_genomics_v1_cigar_proto_init() {
	if File_google_genomics_v1_cigar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_genomics_v1_cigar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CigarUnit); i {
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
			RawDescriptor: file_google_genomics_v1_cigar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_genomics_v1_cigar_proto_goTypes,
		DependencyIndexes: file_google_genomics_v1_cigar_proto_depIdxs,
		EnumInfos:         file_google_genomics_v1_cigar_proto_enumTypes,
		MessageInfos:      file_google_genomics_v1_cigar_proto_msgTypes,
	}.Build()
	File_google_genomics_v1_cigar_proto = out.File
	file_google_genomics_v1_cigar_proto_rawDesc = nil
	file_google_genomics_v1_cigar_proto_goTypes = nil
	file_google_genomics_v1_cigar_proto_depIdxs = nil
}
