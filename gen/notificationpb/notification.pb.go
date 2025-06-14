// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: chat/notification.proto

package notificationpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body          string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	IsRead        bool                   `protobuf:"varint,5,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	Timestamp     int64                  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_chat_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_chat_notification_proto_msgTypes[0]
	if x != nil {
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
	return file_chat_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notification) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Notification) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *Notification) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetNotificationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsRequest) Reset() {
	*x = GetNotificationsRequest{}
	mi := &file_chat_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsRequest) ProtoMessage() {}

func (x *GetNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_chat_notification_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetNotificationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notifications []*Notification        `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotificationsResponse) Reset() {
	*x = GetNotificationsResponse{}
	mi := &file_chat_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsResponse) ProtoMessage() {}

func (x *GetNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_chat_notification_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type MarkReadRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	NotificationId string                 `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MarkReadRequest) Reset() {
	*x = MarkReadRequest{}
	mi := &file_chat_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkReadRequest) ProtoMessage() {}

func (x *MarkReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkReadRequest.ProtoReflect.Descriptor instead.
func (*MarkReadRequest) Descriptor() ([]byte, []int) {
	return file_chat_notification_proto_rawDescGZIP(), []int{3}
}

func (x *MarkReadRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type MarkReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MarkReadResponse) Reset() {
	*x = MarkReadResponse{}
	mi := &file_chat_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarkReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkReadResponse) ProtoMessage() {}

func (x *MarkReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkReadResponse.ProtoReflect.Descriptor instead.
func (*MarkReadResponse) Descriptor() ([]byte, []int) {
	return file_chat_notification_proto_rawDescGZIP(), []int{4}
}

func (x *MarkReadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_chat_notification_proto protoreflect.FileDescriptor

const file_chat_notification_proto_rawDesc = "" +
	"\n" +
	"\x17chat/notification.proto\x12\x0enotificationpb\"\x98\x01\n" +
	"\fNotification\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x12\n" +
	"\x04body\x18\x04 \x01(\tR\x04body\x12\x17\n" +
	"\ais_read\x18\x05 \x01(\bR\x06isRead\x12\x1c\n" +
	"\ttimestamp\x18\x06 \x01(\x03R\ttimestamp\"2\n" +
	"\x17GetNotificationsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"^\n" +
	"\x18GetNotificationsResponse\x12B\n" +
	"\rnotifications\x18\x01 \x03(\v2\x1c.notificationpb.NotificationR\rnotifications\":\n" +
	"\x0fMarkReadRequest\x12'\n" +
	"\x0fnotification_id\x18\x01 \x01(\tR\x0enotificationId\",\n" +
	"\x10MarkReadResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xd7\x01\n" +
	"\x13NotificationService\x12e\n" +
	"\x10GetNotifications\x12'.notificationpb.GetNotificationsRequest\x1a(.notificationpb.GetNotificationsResponse\x12Y\n" +
	"\x14MarkNotificationRead\x12\x1f.notificationpb.MarkReadRequest\x1a .notificationpb.MarkReadResponseB%Z#./gen/notificationpb;notificationpbb\x06proto3"

var (
	file_chat_notification_proto_rawDescOnce sync.Once
	file_chat_notification_proto_rawDescData []byte
)

func file_chat_notification_proto_rawDescGZIP() []byte {
	file_chat_notification_proto_rawDescOnce.Do(func() {
		file_chat_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_notification_proto_rawDesc), len(file_chat_notification_proto_rawDesc)))
	})
	return file_chat_notification_proto_rawDescData
}

var file_chat_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_notification_proto_goTypes = []any{
	(*Notification)(nil),             // 0: notificationpb.Notification
	(*GetNotificationsRequest)(nil),  // 1: notificationpb.GetNotificationsRequest
	(*GetNotificationsResponse)(nil), // 2: notificationpb.GetNotificationsResponse
	(*MarkReadRequest)(nil),          // 3: notificationpb.MarkReadRequest
	(*MarkReadResponse)(nil),         // 4: notificationpb.MarkReadResponse
}
var file_chat_notification_proto_depIdxs = []int32{
	0, // 0: notificationpb.GetNotificationsResponse.notifications:type_name -> notificationpb.Notification
	1, // 1: notificationpb.NotificationService.GetNotifications:input_type -> notificationpb.GetNotificationsRequest
	3, // 2: notificationpb.NotificationService.MarkNotificationRead:input_type -> notificationpb.MarkReadRequest
	2, // 3: notificationpb.NotificationService.GetNotifications:output_type -> notificationpb.GetNotificationsResponse
	4, // 4: notificationpb.NotificationService.MarkNotificationRead:output_type -> notificationpb.MarkReadResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_notification_proto_init() }
func file_chat_notification_proto_init() {
	if File_chat_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_notification_proto_rawDesc), len(file_chat_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_notification_proto_goTypes,
		DependencyIndexes: file_chat_notification_proto_depIdxs,
		MessageInfos:      file_chat_notification_proto_msgTypes,
	}.Build()
	File_chat_notification_proto = out.File
	file_chat_notification_proto_goTypes = nil
	file_chat_notification_proto_depIdxs = nil
}
