// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: communication.proto

package communication

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

type NodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *NodeReq) Reset() {
	*x = NodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeReq) ProtoMessage() {}

func (x *NodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeReq.ProtoReflect.Descriptor instead.
func (*NodeReq) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{0}
}

func (x *NodeReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *NodeRes) Reset() {
	*x = NodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRes) ProtoMessage() {}

func (x *NodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRes.ProtoReflect.Descriptor instead.
func (*NodeRes) Descriptor() ([]byte, []int) {
	return file_communication_proto_rawDescGZIP(), []int{1}
}

func (x *NodeRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_communication_proto protoreflect.FileDescriptor

var file_communication_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x07, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x2a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x08, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_communication_proto_rawDescOnce sync.Once
	file_communication_proto_rawDescData = file_communication_proto_rawDesc
)

func file_communication_proto_rawDescGZIP() []byte {
	file_communication_proto_rawDescOnce.Do(func() {
		file_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_communication_proto_rawDescData)
	})
	return file_communication_proto_rawDescData
}

var file_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_communication_proto_goTypes = []interface{}{
	(*NodeReq)(nil), // 0: NodeReq
	(*NodeRes)(nil), // 1: NodeRes
}
var file_communication_proto_depIdxs = []int32{
	0, // 0: Node.InitNode:input_type -> NodeReq
	1, // 1: Node.InitNode:output_type -> NodeRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_communication_proto_init() }
func file_communication_proto_init() {
	if File_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeReq); i {
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
		file_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRes); i {
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
			RawDescriptor: file_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_communication_proto_goTypes,
		DependencyIndexes: file_communication_proto_depIdxs,
		MessageInfos:      file_communication_proto_msgTypes,
	}.Build()
	File_communication_proto = out.File
	file_communication_proto_rawDesc = nil
	file_communication_proto_goTypes = nil
	file_communication_proto_depIdxs = nil
}
