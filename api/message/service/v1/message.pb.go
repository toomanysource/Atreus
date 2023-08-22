// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.5
// source: api/message/service/v1/message.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MessageListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                                // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`       // 对方用户id
	PreMsgTime int64  `protobuf:"varint,3,opt,name=pre_msg_time,json=preMsgTime,proto3" json:"pre_msg_time,omitempty"` //上次最新消息的时间（新增字段-apk更新中）
}

func (x *MessageListRequest) Reset() {
	*x = MessageListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_service_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageListRequest) ProtoMessage() {}

func (x *MessageListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_service_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageListRequest.ProtoReflect.Descriptor instead.
func (*MessageListRequest) Descriptor() ([]byte, []int) {
	return file_api_message_service_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MessageListRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *MessageListRequest) GetPreMsgTime() int64 {
	if x != nil {
		return x.PreMsgTime
	}
	return 0
}

type MessageListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,proto3" json:"status_code,omitempty"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `protobuf:"bytes,2,opt,name=status_msg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	MessageList []*Message `protobuf:"bytes,3,rep,name=message_list,proto3" json:"message_list,omitempty"` // 消息列表
}

func (x *MessageListReply) Reset() {
	*x = MessageListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_service_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageListReply) ProtoMessage() {}

func (x *MessageListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_service_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageListReply.ProtoReflect.Descriptor instead.
func (*MessageListReply) Descriptor() ([]byte, []int) {
	return file_api_message_service_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageListReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *MessageListReply) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *MessageListReply) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type MessageActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                              // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // 对方用户id
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-发送消息
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                          // 消息内容
}

func (x *MessageActionRequest) Reset() {
	*x = MessageActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_service_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionRequest) ProtoMessage() {}

func (x *MessageActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_service_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionRequest.ProtoReflect.Descriptor instead.
func (*MessageActionRequest) Descriptor() ([]byte, []int) {
	return file_api_message_service_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *MessageActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MessageActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *MessageActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *MessageActionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessageActionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,proto3" json:"status_msg,omitempty"`    // 返回状态描述
}

func (x *MessageActionReply) Reset() {
	*x = MessageActionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_service_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionReply) ProtoMessage() {}

func (x *MessageActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_service_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionReply.ProtoReflect.Descriptor instead.
func (*MessageActionReply) Descriptor() ([]byte, []int) {
	return file_api_message_service_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *MessageActionReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *MessageActionReply) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                     // 消息id
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,proto3" json:"to_user_id,omitempty"`     // 该消息接收者的id
	FromUserId int64  `protobuf:"varint,3,opt,name=from_user_id,proto3" json:"from_user_id,omitempty"` // 该消息发送者的id
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`            // 消息内容
	CreateTime string `protobuf:"bytes,5,opt,name=create_time,proto3" json:"create_time,omitempty"`    // 消息创建时间
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_service_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_service_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_message_service_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *Message) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

var File_api_message_service_v1_message_proto protoreflect.FileDescriptor

var file_api_message_service_v1_message_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x4d, 0x73, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x12, 0x3f, 0x0a, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x85, 0x01, 0x0a,
	0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x99, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x95, 0x02, 0x0a, 0x0e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_message_service_v1_message_proto_rawDescOnce sync.Once
	file_api_message_service_v1_message_proto_rawDescData = file_api_message_service_v1_message_proto_rawDesc
)

func file_api_message_service_v1_message_proto_rawDescGZIP() []byte {
	file_api_message_service_v1_message_proto_rawDescOnce.Do(func() {
		file_api_message_service_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_message_service_v1_message_proto_rawDescData)
	})
	return file_api_message_service_v1_message_proto_rawDescData
}

var file_api_message_service_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_message_service_v1_message_proto_goTypes = []interface{}{
	(*MessageListRequest)(nil),   // 0: message.service.v1.MessageListRequest
	(*MessageListReply)(nil),     // 1: message.service.v1.MessageListReply
	(*MessageActionRequest)(nil), // 2: message.service.v1.MessageActionRequest
	(*MessageActionReply)(nil),   // 3: message.service.v1.MessageActionReply
	(*Message)(nil),              // 4: message.service.v1.Message
}
var file_api_message_service_v1_message_proto_depIdxs = []int32{
	4, // 0: message.service.v1.MessageListReply.message_list:type_name -> message.service.v1.Message
	0, // 1: message.service.v1.messageService.GetMessageList:input_type -> message.service.v1.MessageListRequest
	2, // 2: message.service.v1.messageService.MessageAction:input_type -> message.service.v1.MessageActionRequest
	1, // 3: message.service.v1.messageService.GetMessageList:output_type -> message.service.v1.MessageListReply
	3, // 4: message.service.v1.messageService.MessageAction:output_type -> message.service.v1.MessageActionReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_message_service_v1_message_proto_init() }
func file_api_message_service_v1_message_proto_init() {
	if File_api_message_service_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_message_service_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageListRequest); i {
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
		file_api_message_service_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageListReply); i {
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
		file_api_message_service_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionRequest); i {
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
		file_api_message_service_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionReply); i {
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
		file_api_message_service_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_api_message_service_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_message_service_v1_message_proto_goTypes,
		DependencyIndexes: file_api_message_service_v1_message_proto_depIdxs,
		MessageInfos:      file_api_message_service_v1_message_proto_msgTypes,
	}.Build()
	File_api_message_service_v1_message_proto = out.File
	file_api_message_service_v1_message_proto_rawDesc = nil
	file_api_message_service_v1_message_proto_goTypes = nil
	file_api_message_service_v1_message_proto_depIdxs = nil
}
