// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: timeseries.proto

package pb

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

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start       int64  `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End         int64  `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Window      string `protobuf:"bytes,3,opt,name=window,proto3" json:"window,omitempty"`
	Aggregation string `protobuf:"bytes,4,opt,name=aggregation,proto3" json:"aggregation,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	mi := &file_timeseries_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_timeseries_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_timeseries_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *QueryRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *QueryRequest) GetWindow() string {
	if x != nil {
		return x.Window
	}
	return ""
}

func (x *QueryRequest) GetAggregation() string {
	if x != nil {
		return x.Aggregation
	}
	return ""
}

type TimeSeriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  int64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TimeSeriesData) Reset() {
	*x = TimeSeriesData{}
	mi := &file_timeseries_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSeriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesData) ProtoMessage() {}

func (x *TimeSeriesData) ProtoReflect() protoreflect.Message {
	mi := &file_timeseries_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesData.ProtoReflect.Descriptor instead.
func (*TimeSeriesData) Descriptor() ([]byte, []int) {
	return file_timeseries_proto_rawDescGZIP(), []int{1}
}

func (x *TimeSeriesData) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *TimeSeriesData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TimeSeriesData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	mi := &file_timeseries_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_timeseries_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_timeseries_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResponse) GetData() []*TimeSeriesData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_timeseries_proto protoreflect.FileDescriptor

var file_timeseries_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x70,
	0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3a, 0x0a, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f, 0x0a, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x57, 0x0a,
	0x11, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timeseries_proto_rawDescOnce sync.Once
	file_timeseries_proto_rawDescData = file_timeseries_proto_rawDesc
)

func file_timeseries_proto_rawDescGZIP() []byte {
	file_timeseries_proto_rawDescOnce.Do(func() {
		file_timeseries_proto_rawDescData = protoimpl.X.CompressGZIP(file_timeseries_proto_rawDescData)
	})
	return file_timeseries_proto_rawDescData
}

var file_timeseries_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_timeseries_proto_goTypes = []any{
	(*QueryRequest)(nil),   // 0: timeseries.QueryRequest
	(*TimeSeriesData)(nil), // 1: timeseries.TimeSeriesData
	(*QueryResponse)(nil),  // 2: timeseries.QueryResponse
}
var file_timeseries_proto_depIdxs = []int32{
	1, // 0: timeseries.QueryResponse.data:type_name -> timeseries.TimeSeriesData
	0, // 1: timeseries.TimeSeriesService.QueryData:input_type -> timeseries.QueryRequest
	2, // 2: timeseries.TimeSeriesService.QueryData:output_type -> timeseries.QueryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_timeseries_proto_init() }
func file_timeseries_proto_init() {
	if File_timeseries_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_timeseries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_timeseries_proto_goTypes,
		DependencyIndexes: file_timeseries_proto_depIdxs,
		MessageInfos:      file_timeseries_proto_msgTypes,
	}.Build()
	File_timeseries_proto = out.File
	file_timeseries_proto_rawDesc = nil
	file_timeseries_proto_goTypes = nil
	file_timeseries_proto_depIdxs = nil
}
