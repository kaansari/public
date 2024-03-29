// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: common.proto

package common

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

type Language_Code int32

const (
	Language_UNKNOWN Language_Code = 0
	// English.
	Language_ENG Language_Code = 1
	// German.
	Language_DEU Language_Code = 2
	// French.
	Language_FRA Language_Code = 3
	// Chinese.
	Language_ZHO Language_Code = 4
	// Korean.
	Language_KOR Language_Code = 5
	// Arabic.
	Language_ARA Language_Code = 6
	// Russian.
	Language_RUS Language_Code = 7
	// Thai.
	Language_THA Language_Code = 8
	// Dutch.
	Language_NLD Language_Code = 9
	// Italian.
	Language_ITA Language_Code = 10
	// Portuguese.
	Language_POR Language_Code = 11
	// Spanish.
	Language_SPA Language_Code = 12
	// Japanese.
	Language_JPN Language_Code = 13
	// Polish.
	Language_POL Language_Code = 14
	// Turkish.
	Language_TUR Language_Code = 15
	// Vietnamese.
	Language_VIE Language_Code = 16
	// Indonesian.
	Language_IND Language_Code = 17
	// Czech.
	Language_CES Language_Code = 18
	// Ukrainian.
	Language_UKR Language_Code = 19
	// Greek.
	Language_ELL Language_Code = 20
	// Hebrew.
	Language_HEB Language_Code = 21
	// Farsi/Persian.
	Language_FAS Language_Code = 22
	// Hindi.
	Language_HIN Language_Code = 23
	// Urdu.
	Language_URD Language_Code = 24
	// Swedish.
	Language_SWE Language_Code = 25
	// Bengali.
	Language_BEN Language_Code = 26
	// Malay.
	Language_MSA Language_Code = 27
	// Romanian.
	Language_RON Language_Code = 28
	// Let the platform detect the language.
	Language_AUTO Language_Code = 5000
)

// Enum value maps for Language_Code.
var (
	Language_Code_name = map[int32]string{
		0:    "UNKNOWN",
		1:    "ENG",
		2:    "DEU",
		3:    "FRA",
		4:    "ZHO",
		5:    "KOR",
		6:    "ARA",
		7:    "RUS",
		8:    "THA",
		9:    "NLD",
		10:   "ITA",
		11:   "POR",
		12:   "SPA",
		13:   "JPN",
		14:   "POL",
		15:   "TUR",
		16:   "VIE",
		17:   "IND",
		18:   "CES",
		19:   "UKR",
		20:   "ELL",
		21:   "HEB",
		22:   "FAS",
		23:   "HIN",
		24:   "URD",
		25:   "SWE",
		26:   "BEN",
		27:   "MSA",
		28:   "RON",
		5000: "AUTO",
	}
	Language_Code_value = map[string]int32{
		"UNKNOWN": 0,
		"ENG":     1,
		"DEU":     2,
		"FRA":     3,
		"ZHO":     4,
		"KOR":     5,
		"ARA":     6,
		"RUS":     7,
		"THA":     8,
		"NLD":     9,
		"ITA":     10,
		"POR":     11,
		"SPA":     12,
		"JPN":     13,
		"POL":     14,
		"TUR":     15,
		"VIE":     16,
		"IND":     17,
		"CES":     18,
		"UKR":     19,
		"ELL":     20,
		"HEB":     21,
		"FAS":     22,
		"HIN":     23,
		"URD":     24,
		"SWE":     25,
		"BEN":     26,
		"MSA":     27,
		"RON":     28,
		"AUTO":    5000,
	}
)

func (x Language_Code) Enum() *Language_Code {
	p := new(Language_Code)
	*p = x
	return p
}

func (x Language_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Language_Code) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Language_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language_Code.Descriptor instead.
func (Language_Code) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0, 0}
}

// Languages supported by Vectara.
type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Request to delete a document from an index.
type DeleteDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Customer ID to issue the request for.
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The Corpus ID that contains the document.
	CorpusId int64 `protobuf:"varint,2,opt,name=corpus_id,json=corpusId,proto3" json:"corpus_id,omitempty"`
	// The Document ID to be deleted.
	DocumentId string `protobuf:"bytes,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *DeleteDocumentRequest) Reset() {
	*x = DeleteDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentRequest) ProtoMessage() {}

func (x *DeleteDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteDocumentRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DeleteDocumentRequest) GetCorpusId() int64 {
	if x != nil {
		return x.CorpusId
	}
	return 0
}

func (x *DeleteDocumentRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

type DeleteDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDocumentResponse) Reset() {
	*x = DeleteDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentResponse) ProtoMessage() {}

func (x *DeleteDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentResponse.ProtoReflect.Descriptor instead.
func (*DeleteDocumentResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

// Encapsulates storage quota consumed by indexed documents.
// NextId: 3
type StorageQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of chars from the document that consumed the storage quota.
	NumChars int64 `protobuf:"varint,1,opt,name=num_chars,json=numChars,proto3" json:"num_chars,omitempty"`
	// The number of chars in the metadata of the document that consumed the
	// storage quota.
	NumMetadataChars int64 `protobuf:"varint,2,opt,name=num_metadata_chars,json=numMetadataChars,proto3" json:"num_metadata_chars,omitempty"`
}

func (x *StorageQuota) Reset() {
	*x = StorageQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageQuota) ProtoMessage() {}

func (x *StorageQuota) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageQuota.ProtoReflect.Descriptor instead.
func (*StorageQuota) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *StorageQuota) GetNumChars() int64 {
	if x != nil {
		return x.NumChars
	}
	return 0
}

func (x *StorageQuota) GetNumMetadataChars() int64 {
	if x != nil {
		return x.NumMetadataChars
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x22, 0xa7, 0x02, 0x0a, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x9a, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x45, 0x55, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x46, 0x52, 0x41, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x5a, 0x48, 0x4f,
	0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x52, 0x41, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x55, 0x53, 0x10, 0x07, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x48, 0x41, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4c, 0x44, 0x10, 0x09, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x54, 0x41, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x52, 0x10,
	0x0b, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x50, 0x41, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50,
	0x4e, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x4c, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03,
	0x54, 0x55, 0x52, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x49, 0x45, 0x10, 0x10, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x4e, 0x44, 0x10, 0x11, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x45, 0x53, 0x10, 0x12,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x4b, 0x52, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4c, 0x4c,
	0x10, 0x14, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x45, 0x42, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x46,
	0x41, 0x53, 0x10, 0x16, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x49, 0x4e, 0x10, 0x17, 0x12, 0x07, 0x0a,
	0x03, 0x55, 0x52, 0x44, 0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x57, 0x45, 0x10, 0x19, 0x12,
	0x07, 0x0a, 0x03, 0x42, 0x45, 0x4e, 0x10, 0x1a, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x53, 0x41, 0x10,
	0x1b, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4f, 0x4e, 0x10, 0x1c, 0x12, 0x09, 0x0a, 0x04, 0x41, 0x55,
	0x54, 0x4f, 0x10, 0x88, 0x27, 0x22, 0x76, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x18, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43,
	0x68, 0x61, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61,
	0x72, 0x73, 0x42, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72,
	0x61, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a,
	0x1f, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(Language_Code)(0),             // 0: com.vectara.Language.Code
	(*Language)(nil),               // 1: com.vectara.Language
	(*DeleteDocumentRequest)(nil),  // 2: com.vectara.DeleteDocumentRequest
	(*DeleteDocumentResponse)(nil), // 3: com.vectara.DeleteDocumentResponse
	(*StorageQuota)(nil),           // 4: com.vectara.StorageQuota
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentRequest); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentResponse); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageQuota); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
