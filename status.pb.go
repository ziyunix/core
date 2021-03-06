// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ziyunix/core/status.proto

package core

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

// 状态
type Status int32

const (
	// 未知，不要使用
	Status_STATUS_UNSPECIFIED Status = 0
	// 以下是开始状态
	// 已创建
	Status_CREATED Status = 1
	// 以下是中间状态
	// 已开始
	Status_STARTED Status = 10
	// 下载中
	Status_DOWNLOADING Status = 11
	// 转码中
	Status_CONVERTING Status = 12
	// 上传中
	Status_UPLOADING Status = 13
	// 以下是结束状态
	// 完成
	Status_COMPLETED Status = 20
	// 失败
	Status_FAILED Status = 21
	// 以下是异常情况
	// 错误，指完全还没有开始就出错，比如数据验证错误、数据解析错误
	Status_ERROR Status = 30
	// 异常，在
	Status_EXCEPTION Status = 31
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "CREATED",
		10: "STARTED",
		11: "DOWNLOADING",
		12: "CONVERTING",
		13: "UPLOADING",
		20: "COMPLETED",
		21: "FAILED",
		30: "ERROR",
		31: "EXCEPTION",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"STARTED":            10,
		"DOWNLOADING":        11,
		"CONVERTING":         12,
		"UPLOADING":          13,
		"COMPLETED":          20,
		"FAILED":             21,
		"ERROR":              30,
		"EXCEPTION":          31,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_ziyunix_core_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_ziyunix_core_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_ziyunix_core_status_proto_rawDescGZIP(), []int{0}
}

var File_ziyunix_core_status_proto protoreflect.FileDescriptor

var file_ziyunix_core_status_proto_rawDesc = []byte{
	0x0a, 0x19, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x7a, 0x69, 0x79,
	0x75, 0x6e, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x9f, 0x01, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f,
	0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x50, 0x4c, 0x4f, 0x41,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x15, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1e, 0x12, 0x0d, 0x0a, 0x09,
	0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1f, 0x42, 0x32, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x50,
	0x01, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69,
	0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ziyunix_core_status_proto_rawDescOnce sync.Once
	file_ziyunix_core_status_proto_rawDescData = file_ziyunix_core_status_proto_rawDesc
)

func file_ziyunix_core_status_proto_rawDescGZIP() []byte {
	file_ziyunix_core_status_proto_rawDescOnce.Do(func() {
		file_ziyunix_core_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_ziyunix_core_status_proto_rawDescData)
	})
	return file_ziyunix_core_status_proto_rawDescData
}

var file_ziyunix_core_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ziyunix_core_status_proto_goTypes = []interface{}{
	(Status)(0), // 0: ziyunix.core.Status
}
var file_ziyunix_core_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ziyunix_core_status_proto_init() }
func file_ziyunix_core_status_proto_init() {
	if File_ziyunix_core_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ziyunix_core_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ziyunix_core_status_proto_goTypes,
		DependencyIndexes: file_ziyunix_core_status_proto_depIdxs,
		EnumInfos:         file_ziyunix_core_status_proto_enumTypes,
	}.Build()
	File_ziyunix_core_status_proto = out.File
	file_ziyunix_core_status_proto_rawDesc = nil
	file_ziyunix_core_status_proto_goTypes = nil
	file_ziyunix_core_status_proto_depIdxs = nil
}
