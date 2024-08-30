// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: externaldata/externaldata.proto

package externaldata

import (
	_ "github.com/taomics/pramanapi"
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

type ExternalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Since  int64  `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
	Until  int64  `protobuf:"varint,3,opt,name=until,proto3" json:"until,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExternalData) Reset() {
	*x = ExternalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalData) ProtoMessage() {}

func (x *ExternalData) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalData.ProtoReflect.Descriptor instead.
func (*ExternalData) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalData) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ExternalData) GetSince() int64 {
	if x != nil {
		return x.Since
	}
	return 0
}

func (x *ExternalData) GetUntil() int64 {
	if x != nil {
		return x.Until
	}
	return 0
}

func (x *ExternalData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExternalDataSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Latest int64  `protobuf:"varint,2,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (x *ExternalDataSummary) Reset() {
	*x = ExternalDataSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalDataSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalDataSummary) ProtoMessage() {}

func (x *ExternalDataSummary) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalDataSummary.ProtoReflect.Descriptor instead.
func (*ExternalDataSummary) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalDataSummary) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ExternalDataSummary) GetLatest() int64 {
	if x != nil {
		return x.Latest
	}
	return 0
}

type StoreExternalDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The external data to store.
	ExternalData *ExternalData `protobuf:"bytes,1,opt,name=external_data,json=externalData,proto3" json:"external_data,omitempty"`
}

func (x *StoreExternalDataRequest) Reset() {
	*x = StoreExternalDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreExternalDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreExternalDataRequest) ProtoMessage() {}

func (x *StoreExternalDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreExternalDataRequest.ProtoReflect.Descriptor instead.
func (*StoreExternalDataRequest) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{2}
}

func (x *StoreExternalDataRequest) GetExternalData() *ExternalData {
	if x != nil {
		return x.ExternalData
	}
	return nil
}

type StoreExternalDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreExternalDataResponse) Reset() {
	*x = StoreExternalDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreExternalDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreExternalDataResponse) ProtoMessage() {}

func (x *StoreExternalDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreExternalDataResponse.ProtoReflect.Descriptor instead.
func (*StoreExternalDataResponse) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{3}
}

type ExternalDataSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sources specified the summaries of external data. If sources is empty, return all summaries.
	Sources []string `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
}

func (x *ExternalDataSummaryRequest) Reset() {
	*x = ExternalDataSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalDataSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalDataSummaryRequest) ProtoMessage() {}

func (x *ExternalDataSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalDataSummaryRequest.ProtoReflect.Descriptor instead.
func (*ExternalDataSummaryRequest) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{4}
}

func (x *ExternalDataSummaryRequest) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

type ExternalDataSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summaries []*ExternalDataSummary `protobuf:"bytes,1,rep,name=summaries,proto3" json:"summaries,omitempty"`
}

func (x *ExternalDataSummaryResponse) Reset() {
	*x = ExternalDataSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_externaldata_externaldata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalDataSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalDataSummaryResponse) ProtoMessage() {}

func (x *ExternalDataSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_externaldata_externaldata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalDataSummaryResponse.ProtoReflect.Descriptor instead.
func (*ExternalDataSummaryResponse) Descriptor() ([]byte, []int) {
	return file_externaldata_externaldata_proto_rawDescGZIP(), []int{5}
}

func (x *ExternalDataSummaryResponse) GetSummaries() []*ExternalDataSummary {
	if x != nil {
		return x.Summaries
	}
	return nil
}

var File_externaldata_externaldata_proto protoreflect.FileDescriptor

var file_externaldata_externaldata_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61,
	0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0c,
	0x70, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0c,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x18, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x1a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x1b,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x09, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x32, 0xa8, 0x02, 0x0a, 0x13,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x2e, 0x74, 0x61, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61,
	0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x37, 0x2e, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x61, 0x6d, 0x61, 0x6e, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74,
	0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x61, 0x6d, 0x61, 0x6e, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x61,
	0x6d, 0x61, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_externaldata_externaldata_proto_rawDescOnce sync.Once
	file_externaldata_externaldata_proto_rawDescData = file_externaldata_externaldata_proto_rawDesc
)

func file_externaldata_externaldata_proto_rawDescGZIP() []byte {
	file_externaldata_externaldata_proto_rawDescOnce.Do(func() {
		file_externaldata_externaldata_proto_rawDescData = protoimpl.X.CompressGZIP(file_externaldata_externaldata_proto_rawDescData)
	})
	return file_externaldata_externaldata_proto_rawDescData
}

var file_externaldata_externaldata_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_externaldata_externaldata_proto_goTypes = []any{
	(*ExternalData)(nil),                // 0: taomics.praman.externaldata.ExternalData
	(*ExternalDataSummary)(nil),         // 1: taomics.praman.externaldata.ExternalDataSummary
	(*StoreExternalDataRequest)(nil),    // 2: taomics.praman.externaldata.StoreExternalDataRequest
	(*StoreExternalDataResponse)(nil),   // 3: taomics.praman.externaldata.StoreExternalDataResponse
	(*ExternalDataSummaryRequest)(nil),  // 4: taomics.praman.externaldata.ExternalDataSummaryRequest
	(*ExternalDataSummaryResponse)(nil), // 5: taomics.praman.externaldata.ExternalDataSummaryResponse
}
var file_externaldata_externaldata_proto_depIdxs = []int32{
	0, // 0: taomics.praman.externaldata.StoreExternalDataRequest.external_data:type_name -> taomics.praman.externaldata.ExternalData
	1, // 1: taomics.praman.externaldata.ExternalDataSummaryResponse.summaries:type_name -> taomics.praman.externaldata.ExternalDataSummary
	2, // 2: taomics.praman.externaldata.ExternalDataService.StoreExternalData:input_type -> taomics.praman.externaldata.StoreExternalDataRequest
	4, // 3: taomics.praman.externaldata.ExternalDataService.GetExternalDataSummary:input_type -> taomics.praman.externaldata.ExternalDataSummaryRequest
	3, // 4: taomics.praman.externaldata.ExternalDataService.StoreExternalData:output_type -> taomics.praman.externaldata.StoreExternalDataResponse
	5, // 5: taomics.praman.externaldata.ExternalDataService.GetExternalDataSummary:output_type -> taomics.praman.externaldata.ExternalDataSummaryResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_externaldata_externaldata_proto_init() }
func file_externaldata_externaldata_proto_init() {
	if File_externaldata_externaldata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_externaldata_externaldata_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExternalData); i {
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
		file_externaldata_externaldata_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExternalDataSummary); i {
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
		file_externaldata_externaldata_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StoreExternalDataRequest); i {
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
		file_externaldata_externaldata_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StoreExternalDataResponse); i {
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
		file_externaldata_externaldata_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ExternalDataSummaryRequest); i {
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
		file_externaldata_externaldata_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ExternalDataSummaryResponse); i {
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
			RawDescriptor: file_externaldata_externaldata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_externaldata_externaldata_proto_goTypes,
		DependencyIndexes: file_externaldata_externaldata_proto_depIdxs,
		MessageInfos:      file_externaldata_externaldata_proto_msgTypes,
	}.Build()
	File_externaldata_externaldata_proto = out.File
	file_externaldata_externaldata_proto_rawDesc = nil
	file_externaldata_externaldata_proto_goTypes = nil
	file_externaldata_externaldata_proto_depIdxs = nil
}
