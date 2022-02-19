// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: not_sunny_acitivites.proto

package proto

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

type NotSunnyActivitiesParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test string `protobuf:"bytes,1,opt,name=Test,proto3" json:"Test,omitempty"`
}

func (x *NotSunnyActivitiesParams) Reset() {
	*x = NotSunnyActivitiesParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_not_sunny_acitivites_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotSunnyActivitiesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotSunnyActivitiesParams) ProtoMessage() {}

func (x *NotSunnyActivitiesParams) ProtoReflect() protoreflect.Message {
	mi := &file_not_sunny_acitivites_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotSunnyActivitiesParams.ProtoReflect.Descriptor instead.
func (*NotSunnyActivitiesParams) Descriptor() ([]byte, []int) {
	return file_not_sunny_acitivites_proto_rawDescGZIP(), []int{0}
}

func (x *NotSunnyActivitiesParams) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

type ActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allweatheractivity string `protobuf:"bytes,1,opt,name=Allweatheractivity,proto3" json:"Allweatheractivity,omitempty"`
}

func (x *ActivityResponse) Reset() {
	*x = ActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_not_sunny_acitivites_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityResponse) ProtoMessage() {}

func (x *ActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_not_sunny_acitivites_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityResponse.ProtoReflect.Descriptor instead.
func (*ActivityResponse) Descriptor() ([]byte, []int) {
	return file_not_sunny_acitivites_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityResponse) GetAllweatheractivity() string {
	if x != nil {
		return x.Allweatheractivity
	}
	return ""
}

var File_not_sunny_acitivites_proto protoreflect.FileDescriptor

var file_not_sunny_acitivites_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x75, 0x6e, 0x6e, 0x79, 0x5f, 0x61, 0x63, 0x69, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x2e, 0x0a, 0x18, 0x4e, 0x6f, 0x74, 0x53, 0x75, 0x6e,
	0x6e, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x6c,
	0x6c, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x41, 0x6c, 0x6c, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x32, 0x71, 0x0a, 0x12, 0x4e, 0x6f,
	0x74, 0x53, 0x75, 0x6e, 0x6e, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x5b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x53, 0x75, 0x6e, 0x6e, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x1a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_not_sunny_acitivites_proto_rawDescOnce sync.Once
	file_not_sunny_acitivites_proto_rawDescData = file_not_sunny_acitivites_proto_rawDesc
)

func file_not_sunny_acitivites_proto_rawDescGZIP() []byte {
	file_not_sunny_acitivites_proto_rawDescOnce.Do(func() {
		file_not_sunny_acitivites_proto_rawDescData = protoimpl.X.CompressGZIP(file_not_sunny_acitivites_proto_rawDescData)
	})
	return file_not_sunny_acitivites_proto_rawDescData
}

var file_not_sunny_acitivites_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_not_sunny_acitivites_proto_goTypes = []interface{}{
	(*NotSunnyActivitiesParams)(nil), // 0: template.NotSunnyActivitiesParams
	(*ActivityResponse)(nil),         // 1: template.ActivityResponse
}
var file_not_sunny_acitivites_proto_depIdxs = []int32{
	0, // 0: template.NotSunnyActivities.GetAllWeatherActivities:input_type -> template.NotSunnyActivitiesParams
	1, // 1: template.NotSunnyActivities.GetAllWeatherActivities:output_type -> template.ActivityResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_not_sunny_acitivites_proto_init() }
func file_not_sunny_acitivites_proto_init() {
	if File_not_sunny_acitivites_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_not_sunny_acitivites_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotSunnyActivitiesParams); i {
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
		file_not_sunny_acitivites_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityResponse); i {
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
			RawDescriptor: file_not_sunny_acitivites_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_not_sunny_acitivites_proto_goTypes,
		DependencyIndexes: file_not_sunny_acitivites_proto_depIdxs,
		MessageInfos:      file_not_sunny_acitivites_proto_msgTypes,
	}.Build()
	File_not_sunny_acitivites_proto = out.File
	file_not_sunny_acitivites_proto_rawDesc = nil
	file_not_sunny_acitivites_proto_goTypes = nil
	file_not_sunny_acitivites_proto_depIdxs = nil
}
