// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: admin_account.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	status "vectara.com/public/proto/status"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Please note that this is an expensive operation, and the requests can be throttled
// by the platform.
type ComputeAccountSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComputeAccountSizeRequest) Reset() {
	*x = ComputeAccountSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAccountSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAccountSizeRequest) ProtoMessage() {}

func (x *ComputeAccountSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAccountSizeRequest.ProtoReflect.Descriptor instead.
func (*ComputeAccountSizeRequest) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{0}
}

// A TextSize object represents the size of stored text.
type TextSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of actual characters in the text that will be searched.
	NumChars uint64 `protobuf:"varint,1,opt,name=num_chars,json=numChars,proto3" json:"num_chars,omitempty"`
	// Count of metadata characters such as URL, author, date of creation etc.
	NumMetadataChars uint64 `protobuf:"varint,2,opt,name=num_metadata_chars,json=numMetadataChars,proto3" json:"num_metadata_chars,omitempty"`
}

func (x *TextSize) Reset() {
	*x = TextSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSize) ProtoMessage() {}

func (x *TextSize) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSize.ProtoReflect.Descriptor instead.
func (*TextSize) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{1}
}

func (x *TextSize) GetNumChars() uint64 {
	if x != nil {
		return x.NumChars
	}
	return 0
}

func (x *TextSize) GetNumMetadataChars() uint64 {
	if x != nil {
		return x.NumMetadataChars
	}
	return 0
}

type ComputeAccountSizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One TextSize represents one cluster. The account size is a sum of all the sizes.
	// Generally, this will only have one value.
	Size []*TextSize `protobuf:"bytes,1,rep,name=size,proto3" json:"size,omitempty"`
	// Status response of the operation.
	Status *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ComputeAccountSizeResponse) Reset() {
	*x = ComputeAccountSizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAccountSizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAccountSizeResponse) ProtoMessage() {}

func (x *ComputeAccountSizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAccountSizeResponse.ProtoReflect.Descriptor instead.
func (*ComputeAccountSizeResponse) Descriptor() ([]byte, []int) {
	return file_admin_account_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeAccountSizeResponse) GetSize() []*TextSize {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *ComputeAccountSizeResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_admin_account_proto protoreflect.FileDescriptor

var file_admin_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x68, 0x61, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x72, 0x73, 0x22, 0x7a, 0x0a, 0x1a, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63,
	0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65,
	0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x12, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a,
	0x1e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_account_proto_rawDescOnce sync.Once
	file_admin_account_proto_rawDescData = file_admin_account_proto_rawDesc
)

func file_admin_account_proto_rawDescGZIP() []byte {
	file_admin_account_proto_rawDescOnce.Do(func() {
		file_admin_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_account_proto_rawDescData)
	})
	return file_admin_account_proto_rawDescData
}

var file_admin_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_admin_account_proto_goTypes = []interface{}{
	(*ComputeAccountSizeRequest)(nil),  // 0: com.vectara.admin.ComputeAccountSizeRequest
	(*TextSize)(nil),                   // 1: com.vectara.admin.TextSize
	(*ComputeAccountSizeResponse)(nil), // 2: com.vectara.admin.ComputeAccountSizeResponse
	(*status.Status)(nil),              // 3: com.vectara.Status
}
var file_admin_account_proto_depIdxs = []int32{
	1, // 0: com.vectara.admin.ComputeAccountSizeResponse.size:type_name -> com.vectara.admin.TextSize
	3, // 1: com.vectara.admin.ComputeAccountSizeResponse.status:type_name -> com.vectara.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_admin_account_proto_init() }
func file_admin_account_proto_init() {
	if File_admin_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAccountSizeRequest); i {
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
		file_admin_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSize); i {
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
		file_admin_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAccountSizeResponse); i {
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
			RawDescriptor: file_admin_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_account_proto_goTypes,
		DependencyIndexes: file_admin_account_proto_depIdxs,
		MessageInfos:      file_admin_account_proto_msgTypes,
	}.Build()
	File_admin_account_proto = out.File
	file_admin_account_proto_rawDesc = nil
	file_admin_account_proto_goTypes = nil
	file_admin_account_proto_depIdxs = nil
}
