// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: t.proto

package t

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

type T1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A []int32          `protobuf:"varint,1,rep,packed,name=a,proto3" json:"a,omitempty"`
	B []*RepeatMessage `protobuf:"bytes,2,rep,name=b,proto3" json:"b,omitempty"`
	C string           `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	D bool             `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E *T1_Embed        `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *T1) Reset() {
	*x = T1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *T1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*T1) ProtoMessage() {}

func (x *T1) ProtoReflect() protoreflect.Message {
	mi := &file_t_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use T1.ProtoReflect.Descriptor instead.
func (*T1) Descriptor() ([]byte, []int) {
	return file_t_proto_rawDescGZIP(), []int{0}
}

func (x *T1) GetA() []int32 {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *T1) GetB() []*RepeatMessage {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *T1) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *T1) GetD() bool {
	if x != nil {
		return x.D
	}
	return false
}

func (x *T1) GetE() *T1_Embed {
	if x != nil {
		return x.E
	}
	return nil
}

type RepeatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M1 string `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
	M2 uint64 `protobuf:"varint,2,opt,name=m2,proto3" json:"m2,omitempty"`
}

func (x *RepeatMessage) Reset() {
	*x = RepeatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatMessage) ProtoMessage() {}

func (x *RepeatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_t_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatMessage.ProtoReflect.Descriptor instead.
func (*RepeatMessage) Descriptor() ([]byte, []int) {
	return file_t_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatMessage) GetM1() string {
	if x != nil {
		return x.M1
	}
	return ""
}

func (x *RepeatMessage) GetM2() uint64 {
	if x != nil {
		return x.M2
	}
	return 0
}

type T1_Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E1 int32  `protobuf:"varint,1,opt,name=e1,proto3" json:"e1,omitempty"`
	E2 string `protobuf:"bytes,2,opt,name=e2,proto3" json:"e2,omitempty"`
}

func (x *T1_Embed) Reset() {
	*x = T1_Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_t_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *T1_Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*T1_Embed) ProtoMessage() {}

func (x *T1_Embed) ProtoReflect() protoreflect.Message {
	mi := &file_t_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use T1_Embed.ProtoReflect.Descriptor instead.
func (*T1_Embed) Descriptor() ([]byte, []int) {
	return file_t_proto_rawDescGZIP(), []int{0, 0}
}

func (x *T1_Embed) GetE1() int32 {
	if x != nil {
		return x.E1
	}
	return 0
}

func (x *T1_Embed) GetE2() string {
	if x != nil {
		return x.E2
	}
	return ""
}

var File_t_proto protoreflect.FileDescriptor

var file_t_proto_rawDesc = []byte{
	0x0a, 0x07, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x01, 0x0a, 0x02, 0x54, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x26, 0x0a, 0x01, 0x62, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x01,
	0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x63, 0x12,
	0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01, 0x64, 0x12, 0x21, 0x0a,
	0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x01, 0x65,
	0x1a, 0x27, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x65, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x32, 0x22, 0x2f, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6d, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x6d, 0x32, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x3b,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_t_proto_rawDescOnce sync.Once
	file_t_proto_rawDescData = file_t_proto_rawDesc
)

func file_t_proto_rawDescGZIP() []byte {
	file_t_proto_rawDescOnce.Do(func() {
		file_t_proto_rawDescData = protoimpl.X.CompressGZIP(file_t_proto_rawDescData)
	})
	return file_t_proto_rawDescData
}

var file_t_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_t_proto_goTypes = []interface{}{
	(*T1)(nil),            // 0: debugdata.T1
	(*RepeatMessage)(nil), // 1: debugdata.RepeatMessage
	(*T1_Embed)(nil),      // 2: debugdata.T1.Embed
}
var file_t_proto_depIdxs = []int32{
	1, // 0: debugdata.T1.b:type_name -> debugdata.RepeatMessage
	2, // 1: debugdata.T1.e:type_name -> debugdata.T1.Embed
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_t_proto_init() }
func file_t_proto_init() {
	if File_t_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_t_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*T1); i {
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
		file_t_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatMessage); i {
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
		file_t_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*T1_Embed); i {
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
			RawDescriptor: file_t_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_t_proto_goTypes,
		DependencyIndexes: file_t_proto_depIdxs,
		MessageInfos:      file_t_proto_msgTypes,
	}.Build()
	File_t_proto = out.File
	file_t_proto_rawDesc = nil
	file_t_proto_goTypes = nil
	file_t_proto_depIdxs = nil
}
