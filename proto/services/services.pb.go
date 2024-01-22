// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: services.proto

package services

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	admin "vectara.com/public/proto/admin"
	common "vectara.com/public/proto/common"
	indexing "vectara.com/public/proto/indexing"
	lists "vectara.com/public/proto/lists"
	serving "vectara.com/public/proto/serving"
	status "vectara.com/public/proto/status"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to index a document.
type IndexDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Customer ID to issue the request for.
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The Corpus ID to index the document into.
	CorpusId int64 `protobuf:"varint,2,opt,name=corpus_id,json=corpusId,proto3" json:"corpus_id,omitempty"`
	// The Document to index.
	Document *indexing.Document `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *IndexDocumentRequest) Reset() {
	*x = IndexDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDocumentRequest) ProtoMessage() {}

func (x *IndexDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDocumentRequest.ProtoReflect.Descriptor instead.
func (*IndexDocumentRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *IndexDocumentRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *IndexDocumentRequest) GetCorpusId() int64 {
	if x != nil {
		return x.CorpusId
	}
	return 0
}

func (x *IndexDocumentRequest) GetDocument() *indexing.Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type IndexDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If ALREADY_EXISTS, it means the document was already indexed, and no new quota was consumed.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The storage quota needed for the document indexed in the request.
	// If "status" is ALREADY_EXISTS, it means that the document was already in the index prior to
	// this request. In such cases, quota is not consumed again and the value in this field
	// represents the quota consumed when the document was indexed the first time.
	QuotaConsumed *common.StorageQuota `protobuf:"bytes,2,opt,name=quota_consumed,json=quotaConsumed,proto3" json:"quota_consumed,omitempty"`
}

func (x *IndexDocumentResponse) Reset() {
	*x = IndexDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDocumentResponse) ProtoMessage() {}

func (x *IndexDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDocumentResponse.ProtoReflect.Descriptor instead.
func (*IndexDocumentResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *IndexDocumentResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *IndexDocumentResponse) GetQuotaConsumed() *common.StorageQuota {
	if x != nil {
		return x.QuotaConsumed
	}
	return nil
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x1a, 0x13, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x72, 0x70, 0x75, 0x73, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x0d, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x32, 0xe2, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x05,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74,
	0x61, 0x72, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3a,
	0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x64, 0x6f, 0x63, 0x3a, 0x01, 0x2a,
	0x32, 0xe1, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6e, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x01,
	0x2a, 0x12, 0x61, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x32, 0xde, 0x11, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74,
	0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x75,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x72, 0x70, 0x75, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x79,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x12, 0x25, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x72,
	0x70, 0x6f, 0x72, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x63,
	0x6f, 0x72, 0x70, 0x6f, 0x72, 0x61, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74,
	0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x2d, 0x6a, 0x6f, 0x62, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63,
	0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x61, 0x64, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x92, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x70, 0x75,
	0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74,
	0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x6f,
	0x72, 0x70, 0x75, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x2d, 0x73, 0x69, 0x7a,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x75, 0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d,
	0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xa6, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x70, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x2d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0xaf, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12, 0x32, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x72,
	0x70, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x70,
	0x75, 0x73, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2d, 0x61, 0x74, 0x74, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63,
	0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22,
	0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65,
	0x79, 0x73, 0x3a, 0x01, 0x2a, 0x32, 0x95, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61,
	0x72, 0x61, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x2d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x3f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x65, 0x63, 0x74, 0x61, 0x72, 0x61, 0x42, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x21, 0x76, 0x65, 0x63,
	0x74, 0x61, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_proto_goTypes = []interface{}{
	(*IndexDocumentRequest)(nil),                   // 0: com.vectara.IndexDocumentRequest
	(*IndexDocumentResponse)(nil),                  // 1: com.vectara.IndexDocumentResponse
	(*indexing.Document)(nil),                      // 2: com.vectara.indexing.Document
	(*status.Status)(nil),                          // 3: com.vectara.Status
	(*common.StorageQuota)(nil),                    // 4: com.vectara.StorageQuota
	(*common.DeleteDocumentRequest)(nil),           // 5: com.vectara.DeleteDocumentRequest
	(*serving.BatchQueryRequest)(nil),              // 6: com.vectara.serving.BatchQueryRequest
	(*admin.CreateCorpusRequest)(nil),              // 7: com.vectara.admin.CreateCorpusRequest
	(*admin.DeleteCorpusRequest)(nil),              // 8: com.vectara.admin.DeleteCorpusRequest
	(*admin.ResetCorpusRequest)(nil),               // 9: com.vectara.admin.ResetCorpusRequest
	(*admin.ListCorporaRequest)(nil),               // 10: com.vectara.admin.ListCorporaRequest
	(*admin.ListJobsRequest)(nil),                  // 11: com.vectara.admin.ListJobsRequest
	(*admin.ReadCorpusRequest)(nil),                // 12: com.vectara.admin.ReadCorpusRequest
	(*admin.ComputeCorpusSizeRequest)(nil),         // 13: com.vectara.admin.ComputeCorpusSizeRequest
	(*admin.ComputeAccountSizeRequest)(nil),        // 14: com.vectara.admin.ComputeAccountSizeRequest
	(*admin.ListUsersRequest)(nil),                 // 15: com.vectara.admin.ListUsersRequest
	(*admin.ManageUserRequest)(nil),                // 16: com.vectara.admin.ManageUserRequest
	(*admin.UpdateCorpusEnablementRequest)(nil),    // 17: com.vectara.admin.UpdateCorpusEnablementRequest
	(*admin.ReplaceCorpusFilterAttrsRequest)(nil),  // 18: com.vectara.admin.ReplaceCorpusFilterAttrsRequest
	(*admin.UsageMetricsRequest)(nil),              // 19: com.vectara.admin.UsageMetricsRequest
	(*admin.CreateApiKeyRequest)(nil),              // 20: com.vectara.admin.CreateApiKeyRequest
	(*admin.EnableApiKeyRequest)(nil),              // 21: com.vectara.admin.EnableApiKeyRequest
	(*admin.DeleteApiKeyRequest)(nil),              // 22: com.vectara.admin.DeleteApiKeyRequest
	(*admin.ListApiKeysRequest)(nil),               // 23: com.vectara.admin.ListApiKeysRequest
	(*lists.ListDocumentsRequest)(nil),             // 24: com.vectara.lists.ListDocumentsRequest
	(*common.DeleteDocumentResponse)(nil),          // 25: com.vectara.DeleteDocumentResponse
	(*serving.BatchQueryResponse)(nil),             // 26: com.vectara.serving.BatchQueryResponse
	(*serving.QueryResponsePart)(nil),              // 27: com.vectara.serving.QueryResponsePart
	(*admin.CreateCorpusResponse)(nil),             // 28: com.vectara.admin.CreateCorpusResponse
	(*admin.DeleteCorpusResponse)(nil),             // 29: com.vectara.admin.DeleteCorpusResponse
	(*admin.ResetCorpusResponse)(nil),              // 30: com.vectara.admin.ResetCorpusResponse
	(*admin.ListCorporaResponse)(nil),              // 31: com.vectara.admin.ListCorporaResponse
	(*admin.ListJobsResponse)(nil),                 // 32: com.vectara.admin.ListJobsResponse
	(*admin.ReadCorpusResponse)(nil),               // 33: com.vectara.admin.ReadCorpusResponse
	(*admin.ComputeCorpusSizeResponse)(nil),        // 34: com.vectara.admin.ComputeCorpusSizeResponse
	(*admin.ComputeAccountSizeResponse)(nil),       // 35: com.vectara.admin.ComputeAccountSizeResponse
	(*admin.ListUsersResponse)(nil),                // 36: com.vectara.admin.ListUsersResponse
	(*admin.ManageUserResponse)(nil),               // 37: com.vectara.admin.ManageUserResponse
	(*admin.UpdateCorpusEnablementResponse)(nil),   // 38: com.vectara.admin.UpdateCorpusEnablementResponse
	(*admin.ReplaceCorpusFilterAttrsResponse)(nil), // 39: com.vectara.admin.ReplaceCorpusFilterAttrsResponse
	(*admin.UsageMetricsResponse)(nil),             // 40: com.vectara.admin.UsageMetricsResponse
	(*admin.CreateApiKeyResponse)(nil),             // 41: com.vectara.admin.CreateApiKeyResponse
	(*admin.EnableApiKeyResponse)(nil),             // 42: com.vectara.admin.EnableApiKeyResponse
	(*admin.DeleteApiKeyResponse)(nil),             // 43: com.vectara.admin.DeleteApiKeyResponse
	(*admin.ListApiKeysResponse)(nil),              // 44: com.vectara.admin.ListApiKeysResponse
	(*lists.ListDocumentsResponse)(nil),            // 45: com.vectara.lists.ListDocumentsResponse
}
var file_services_proto_depIdxs = []int32{
	2,  // 0: com.vectara.IndexDocumentRequest.document:type_name -> com.vectara.indexing.Document
	3,  // 1: com.vectara.IndexDocumentResponse.status:type_name -> com.vectara.Status
	4,  // 2: com.vectara.IndexDocumentResponse.quota_consumed:type_name -> com.vectara.StorageQuota
	0,  // 3: com.vectara.IndexService.Index:input_type -> com.vectara.IndexDocumentRequest
	5,  // 4: com.vectara.IndexService.Delete:input_type -> com.vectara.DeleteDocumentRequest
	6,  // 5: com.vectara.QueryService.Query:input_type -> com.vectara.serving.BatchQueryRequest
	6,  // 6: com.vectara.QueryService.StreamQuery:input_type -> com.vectara.serving.BatchQueryRequest
	7,  // 7: com.vectara.AdminService.CreateCorpus:input_type -> com.vectara.admin.CreateCorpusRequest
	8,  // 8: com.vectara.AdminService.DeleteCorpus:input_type -> com.vectara.admin.DeleteCorpusRequest
	9,  // 9: com.vectara.AdminService.ResetCorpus:input_type -> com.vectara.admin.ResetCorpusRequest
	10, // 10: com.vectara.AdminService.ListCorpora:input_type -> com.vectara.admin.ListCorporaRequest
	11, // 11: com.vectara.AdminService.ListJobs:input_type -> com.vectara.admin.ListJobsRequest
	12, // 12: com.vectara.AdminService.ReadCorpus:input_type -> com.vectara.admin.ReadCorpusRequest
	13, // 13: com.vectara.AdminService.ComputeCorpusSize:input_type -> com.vectara.admin.ComputeCorpusSizeRequest
	14, // 14: com.vectara.AdminService.ComputeAccountSize:input_type -> com.vectara.admin.ComputeAccountSizeRequest
	15, // 15: com.vectara.AdminService.ListUsers:input_type -> com.vectara.admin.ListUsersRequest
	16, // 16: com.vectara.AdminService.ManageUser:input_type -> com.vectara.admin.ManageUserRequest
	17, // 17: com.vectara.AdminService.UpdateCorpusEnablement:input_type -> com.vectara.admin.UpdateCorpusEnablementRequest
	18, // 18: com.vectara.AdminService.ReplaceCorpusFilterAttrs:input_type -> com.vectara.admin.ReplaceCorpusFilterAttrsRequest
	19, // 19: com.vectara.AdminService.GetUsageMetrics:input_type -> com.vectara.admin.UsageMetricsRequest
	20, // 20: com.vectara.AdminService.CreateApiKey:input_type -> com.vectara.admin.CreateApiKeyRequest
	21, // 21: com.vectara.AdminService.EnableApiKey:input_type -> com.vectara.admin.EnableApiKeyRequest
	22, // 22: com.vectara.AdminService.DeleteApiKey:input_type -> com.vectara.admin.DeleteApiKeyRequest
	23, // 23: com.vectara.AdminService.ListApiKeys:input_type -> com.vectara.admin.ListApiKeysRequest
	24, // 24: com.vectara.DocumentService.ListDocuments:input_type -> com.vectara.lists.ListDocumentsRequest
	1,  // 25: com.vectara.IndexService.Index:output_type -> com.vectara.IndexDocumentResponse
	25, // 26: com.vectara.IndexService.Delete:output_type -> com.vectara.DeleteDocumentResponse
	26, // 27: com.vectara.QueryService.Query:output_type -> com.vectara.serving.BatchQueryResponse
	27, // 28: com.vectara.QueryService.StreamQuery:output_type -> com.vectara.serving.QueryResponsePart
	28, // 29: com.vectara.AdminService.CreateCorpus:output_type -> com.vectara.admin.CreateCorpusResponse
	29, // 30: com.vectara.AdminService.DeleteCorpus:output_type -> com.vectara.admin.DeleteCorpusResponse
	30, // 31: com.vectara.AdminService.ResetCorpus:output_type -> com.vectara.admin.ResetCorpusResponse
	31, // 32: com.vectara.AdminService.ListCorpora:output_type -> com.vectara.admin.ListCorporaResponse
	32, // 33: com.vectara.AdminService.ListJobs:output_type -> com.vectara.admin.ListJobsResponse
	33, // 34: com.vectara.AdminService.ReadCorpus:output_type -> com.vectara.admin.ReadCorpusResponse
	34, // 35: com.vectara.AdminService.ComputeCorpusSize:output_type -> com.vectara.admin.ComputeCorpusSizeResponse
	35, // 36: com.vectara.AdminService.ComputeAccountSize:output_type -> com.vectara.admin.ComputeAccountSizeResponse
	36, // 37: com.vectara.AdminService.ListUsers:output_type -> com.vectara.admin.ListUsersResponse
	37, // 38: com.vectara.AdminService.ManageUser:output_type -> com.vectara.admin.ManageUserResponse
	38, // 39: com.vectara.AdminService.UpdateCorpusEnablement:output_type -> com.vectara.admin.UpdateCorpusEnablementResponse
	39, // 40: com.vectara.AdminService.ReplaceCorpusFilterAttrs:output_type -> com.vectara.admin.ReplaceCorpusFilterAttrsResponse
	40, // 41: com.vectara.AdminService.GetUsageMetrics:output_type -> com.vectara.admin.UsageMetricsResponse
	41, // 42: com.vectara.AdminService.CreateApiKey:output_type -> com.vectara.admin.CreateApiKeyResponse
	42, // 43: com.vectara.AdminService.EnableApiKey:output_type -> com.vectara.admin.EnableApiKeyResponse
	43, // 44: com.vectara.AdminService.DeleteApiKey:output_type -> com.vectara.admin.DeleteApiKeyResponse
	44, // 45: com.vectara.AdminService.ListApiKeys:output_type -> com.vectara.admin.ListApiKeysResponse
	45, // 46: com.vectara.DocumentService.ListDocuments:output_type -> com.vectara.lists.ListDocumentsResponse
	25, // [25:47] is the sub-list for method output_type
	3,  // [3:25] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDocumentRequest); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDocumentResponse); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
