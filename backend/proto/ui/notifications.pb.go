// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: ui/notifications.proto

package ui

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

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Notification:
	//	*Notification_ConnState
	//	*Notification_DataState
	//	*Notification_Status
	Notification isNotification_Notification `protobuf_oneof:"notification"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_ui_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_ui_notifications_proto_rawDescGZIP(), []int{0}
}

func (m *Notification) GetNotification() isNotification_Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (x *Notification) GetConnState() *ConnectionState {
	if x, ok := x.GetNotification().(*Notification_ConnState); ok {
		return x.ConnState
	}
	return nil
}

func (x *Notification) GetDataState() *DataState {
	if x, ok := x.GetNotification().(*Notification_DataState); ok {
		return x.DataState
	}
	return nil
}

func (x *Notification) GetStatus() *GetStatusResponse {
	if x, ok := x.GetNotification().(*Notification_Status); ok {
		return x.Status
	}
	return nil
}

type isNotification_Notification interface {
	isNotification_Notification()
}

type Notification_ConnState struct {
	ConnState *ConnectionState `protobuf:"bytes,1,opt,name=conn_state,json=connState,proto3,oneof"`
}

type Notification_DataState struct {
	DataState *DataState `protobuf:"bytes,2,opt,name=data_state,json=dataState,proto3,oneof"`
}

type Notification_Status struct {
	Status *GetStatusResponse `protobuf:"bytes,3,opt,name=status,proto3,oneof"`
}

func (*Notification_ConnState) isNotification_Notification() {}

func (*Notification_DataState) isNotification_Notification() {}

func (*Notification_Status) isNotification_Notification() {}

type ConnectionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Backend is successfully connected to hubble-relay
	Connected bool `protobuf:"varint,1,opt,name=connected,proto3" json:"connected,omitempty"`
	// Backend has lost the connection to hubble-relay and is reconnecting now
	Reconnecting bool `protobuf:"varint,2,opt,name=reconnecting,proto3" json:"reconnecting,omitempty"`
}

func (x *ConnectionState) Reset() {
	*x = ConnectionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionState) ProtoMessage() {}

func (x *ConnectionState) ProtoReflect() protoreflect.Message {
	mi := &file_ui_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionState.ProtoReflect.Descriptor instead.
func (*ConnectionState) Descriptor() ([]byte, []int) {
	return file_ui_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionState) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *ConnectionState) GetReconnecting() bool {
	if x != nil {
		return x.Reconnecting
	}
	return false
}

type DataState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// No pods in selected namespace
	NoActivity bool `protobuf:"varint,1,opt,name=no_activity,json=noActivity,proto3" json:"no_activity,omitempty"`
}

func (x *DataState) Reset() {
	*x = DataState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ui_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataState) ProtoMessage() {}

func (x *DataState) ProtoReflect() protoreflect.Message {
	mi := &file_ui_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataState.ProtoReflect.Descriptor instead.
func (*DataState) Descriptor() ([]byte, []int) {
	return file_ui_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *DataState) GetNoActivity() bool {
	if x != nil {
		return x.NoActivity
	}
	return false
}

var File_ui_notifications_proto protoreflect.FileDescriptor

var file_ui_notifications_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x75, 0x69, 0x1a, 0x0f, 0x75, 0x69,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x6f,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ui_notifications_proto_rawDescOnce sync.Once
	file_ui_notifications_proto_rawDescData = file_ui_notifications_proto_rawDesc
)

func file_ui_notifications_proto_rawDescGZIP() []byte {
	file_ui_notifications_proto_rawDescOnce.Do(func() {
		file_ui_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_ui_notifications_proto_rawDescData)
	})
	return file_ui_notifications_proto_rawDescData
}

var file_ui_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ui_notifications_proto_goTypes = []interface{}{
	(*Notification)(nil),      // 0: ui.Notification
	(*ConnectionState)(nil),   // 1: ui.ConnectionState
	(*DataState)(nil),         // 2: ui.DataState
	(*GetStatusResponse)(nil), // 3: ui.GetStatusResponse
}
var file_ui_notifications_proto_depIdxs = []int32{
	1, // 0: ui.Notification.conn_state:type_name -> ui.ConnectionState
	2, // 1: ui.Notification.data_state:type_name -> ui.DataState
	3, // 2: ui.Notification.status:type_name -> ui.GetStatusResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ui_notifications_proto_init() }
func file_ui_notifications_proto_init() {
	if File_ui_notifications_proto != nil {
		return
	}
	file_ui_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ui_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_ui_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionState); i {
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
		file_ui_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataState); i {
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
	file_ui_notifications_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Notification_ConnState)(nil),
		(*Notification_DataState)(nil),
		(*Notification_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ui_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ui_notifications_proto_goTypes,
		DependencyIndexes: file_ui_notifications_proto_depIdxs,
		MessageInfos:      file_ui_notifications_proto_msgTypes,
	}.Build()
	File_ui_notifications_proto = out.File
	file_ui_notifications_proto_rawDesc = nil
	file_ui_notifications_proto_goTypes = nil
	file_ui_notifications_proto_depIdxs = nil
}
