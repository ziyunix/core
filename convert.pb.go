// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: ziyunix/core/convert.proto

package core

import (
	media "github.com/storezhang/media"
	transfer "github.com/storezhang/transfer"
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

// 转码
type Convert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 类型
	Type media.Type `protobuf:"varint,3,opt,name=type,proto3,enum=storezhang.media.Type" json:"type,omitempty"`
	// 目标
	Target *transfer.Target `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	// 视频
	Video *Video `protobuf:"bytes,5,opt,name=video,proto3" json:"video,omitempty"`
	// 音频
	Audio *Audio `protobuf:"bytes,6,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *Convert) Reset() {
	*x = Convert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ziyunix_core_convert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Convert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Convert) ProtoMessage() {}

func (x *Convert) ProtoReflect() protoreflect.Message {
	mi := &file_ziyunix_core_convert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Convert.ProtoReflect.Descriptor instead.
func (*Convert) Descriptor() ([]byte, []int) {
	return file_ziyunix_core_convert_proto_rawDescGZIP(), []int{0}
}

func (x *Convert) GetType() media.Type {
	if x != nil {
		return x.Type
	}
	return media.Type_TYPE_UNSPECIFIED
}

func (x *Convert) GetTarget() *transfer.Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Convert) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *Convert) GetAudio() *Audio {
	if x != nil {
		return x.Audio
	}
	return nil
}

var File_ziyunix_core_convert_proto protoreflect.FileDescriptor

var file_ziyunix_core_convert_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x7a, 0x69,
	0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1b, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a,
	0x68, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61,
	0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x7a, 0x69, 0x79, 0x75,
	0x6e, 0x69, 0x78, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x42, 0x32, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x79, 0x75, 0x6e, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ziyunix_core_convert_proto_rawDescOnce sync.Once
	file_ziyunix_core_convert_proto_rawDescData = file_ziyunix_core_convert_proto_rawDesc
)

func file_ziyunix_core_convert_proto_rawDescGZIP() []byte {
	file_ziyunix_core_convert_proto_rawDescOnce.Do(func() {
		file_ziyunix_core_convert_proto_rawDescData = protoimpl.X.CompressGZIP(file_ziyunix_core_convert_proto_rawDescData)
	})
	return file_ziyunix_core_convert_proto_rawDescData
}

var file_ziyunix_core_convert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ziyunix_core_convert_proto_goTypes = []interface{}{
	(*Convert)(nil),         // 0: ziyunix.core.Convert
	(media.Type)(0),         // 1: storezhang.media.Type
	(*transfer.Target)(nil), // 2: storezhang.transfer.Target
	(*Video)(nil),           // 3: ziyunix.core.Video
	(*Audio)(nil),           // 4: ziyunix.core.Audio
}
var file_ziyunix_core_convert_proto_depIdxs = []int32{
	1, // 0: ziyunix.core.Convert.type:type_name -> storezhang.media.Type
	2, // 1: ziyunix.core.Convert.target:type_name -> storezhang.transfer.Target
	3, // 2: ziyunix.core.Convert.video:type_name -> ziyunix.core.Video
	4, // 3: ziyunix.core.Convert.audio:type_name -> ziyunix.core.Audio
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ziyunix_core_convert_proto_init() }
func file_ziyunix_core_convert_proto_init() {
	if File_ziyunix_core_convert_proto != nil {
		return
	}
	file_ziyunix_core_video_proto_init()
	file_ziyunix_core_audio_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ziyunix_core_convert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Convert); i {
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
			RawDescriptor: file_ziyunix_core_convert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ziyunix_core_convert_proto_goTypes,
		DependencyIndexes: file_ziyunix_core_convert_proto_depIdxs,
		MessageInfos:      file_ziyunix_core_convert_proto_msgTypes,
	}.Build()
	File_ziyunix_core_convert_proto = out.File
	file_ziyunix_core_convert_proto_rawDesc = nil
	file_ziyunix_core_convert_proto_goTypes = nil
	file_ziyunix_core_convert_proto_depIdxs = nil
}