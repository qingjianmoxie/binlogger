// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.0
// source: event.proto

package pbmysql

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EventType int32

const (
	EventType_UnknownEvent EventType = 0
	EventType_InsertEvent  EventType = 1
	EventType_UpdateEvent  EventType = 2
	EventType_DeleteEvent  EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "UnknownEvent",
		1: "InsertEvent",
		2: "UpdateEvent",
		3: "DeleteEvent",
	}
	EventType_value = map[string]int32{
		"UnknownEvent": 0,
		"InsertEvent":  1,
		"UpdateEvent":  2,
		"DeleteEvent":  3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

// one column values from old to new
type ColumnValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of one column
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value of the column,converted to bytes
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// value before
	Before     []byte `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	ValueNull  bool   `protobuf:"varint,4,opt,name=value_null,json=valueNull,proto3" json:"value_null,omitempty"`
	BeforeNull bool   `protobuf:"varint,5,opt,name=before_null,json=beforeNull,proto3" json:"before_null,omitempty"`
}

func (x *ColumnValue) Reset() {
	*x = ColumnValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnValue) ProtoMessage() {}

func (x *ColumnValue) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnValue.ProtoReflect.Descriptor instead.
func (*ColumnValue) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *ColumnValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ColumnValue) GetBefore() []byte {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *ColumnValue) GetValueNull() bool {
	if x != nil {
		return x.ValueNull
	}
	return false
}

func (x *ColumnValue) GetBeforeNull() bool {
	if x != nil {
		return x.BeforeNull
	}
	return false
}

// one event include
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schema name
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// table name
	Table string `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	// event type
	Et EventType `protobuf:"varint,3,opt,name=et,proto3,enum=pbmysql.EventType" json:"et,omitempty"`
	// all columns names and values exclude id
	Columns []*ColumnValue `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	// used as global unique identifier
	Gtid string `protobuf:"bytes,5,opt,name=gtid,proto3" json:"gtid,omitempty"`
	// timestamp for nano seconds
	NanoTimestamp int64 `protobuf:"varint,6,opt,name=nano_timestamp,json=nanoTimestamp,proto3" json:"nano_timestamp,omitempty"`
	// primary key index
	PkColumns []int32 `protobuf:"varint,7,rep,packed,name=pk_columns,json=pkColumns,proto3" json:"pk_columns,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Event) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Event) GetEt() EventType {
	if x != nil {
		return x.Et
	}
	return EventType_UnknownEvent
}

func (x *Event) GetColumns() []*ColumnValue {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Event) GetGtid() string {
	if x != nil {
		return x.Gtid
	}
	return ""
}

func (x *Event) GetNanoTimestamp() int64 {
	if x != nil {
		return x.NanoTimestamp
	}
	return 0
}

func (x *Event) GetPkColumns() []int32 {
	if x != nil {
		return x.PkColumns
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x62, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x4e, 0x75, 0x6c, 0x6c, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x22, 0x0a, 0x02, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70,
	0x62, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x02, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x74, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x67, 0x74, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6e, 0x6f,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6e, 0x61, 0x6e, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6b, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6b, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x2a, 0x50,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x03,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x62, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_proto_goTypes = []interface{}{
	(EventType)(0),      // 0: pbmysql.EventType
	(*ColumnValue)(nil), // 1: pbmysql.ColumnValue
	(*Event)(nil),       // 2: pbmysql.Event
}
var file_event_proto_depIdxs = []int32{
	0, // 0: pbmysql.Event.et:type_name -> pbmysql.EventType
	1, // 1: pbmysql.Event.columns:type_name -> pbmysql.ColumnValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnValue); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		EnumInfos:         file_event_proto_enumTypes,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
