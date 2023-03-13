// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/pb/user.proto

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

type AddNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId   int64  `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddNameRequest) Reset() {
	*x = AddNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNameRequest) ProtoMessage() {}

func (x *AddNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNameRequest.ProtoReflect.Descriptor instead.
func (*AddNameRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *AddNameRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *AddNameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddNameResponse) Reset() {
	*x = AddNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNameResponse) ProtoMessage() {}

func (x *AddNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNameResponse.ProtoReflect.Descriptor instead.
func (*AddNameResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *AddNameResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64  `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	Phone  string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *AddPhoneRequest) Reset() {
	*x = AddPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhoneRequest) ProtoMessage() {}

func (x *AddPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhoneRequest.ProtoReflect.Descriptor instead.
func (*AddPhoneRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *AddPhoneRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *AddPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type AddPhoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddPhoneResponse) Reset() {
	*x = AddPhoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhoneResponse) ProtoMessage() {}

func (x *AddPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhoneResponse.ProtoReflect.Descriptor instead.
func (*AddPhoneResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *AddPhoneResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IsAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int64 `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *IsAdminRequest) Reset() {
	*x = IsAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminRequest) ProtoMessage() {}

func (x *IsAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminRequest.ProtoReflect.Descriptor instead.
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *IsAdminRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type IsAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *IsAdminResponse) Reset() {
	*x = IsAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAdminResponse) ProtoMessage() {}

func (x *IsAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAdminResponse.ProtoReflect.Descriptor instead.
func (*IsAdminResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *IsAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type GetUsersByChatIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId []int64 `protobuf:"varint,1,rep,packed,name=chatId,proto3" json:"chatId,omitempty"`
}

func (x *GetUsersByChatIdRequest) Reset() {
	*x = GetUsersByChatIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersByChatIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersByChatIdRequest) ProtoMessage() {}

func (x *GetUsersByChatIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersByChatIdRequest.ProtoReflect.Descriptor instead.
func (*GetUsersByChatIdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersByChatIdRequest) GetChatId() []int64 {
	if x != nil {
		return x.ChatId
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetUsersByChatIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUsersByChatIdResponse) Reset() {
	*x = GetUsersByChatIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersByChatIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersByChatIdResponse) ProtoMessage() {}

func (x *GetUsersByChatIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersByChatIdResponse.ProtoReflect.Descriptor instead.
func (*GetUsersByChatIdResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUsersByChatIdResponse) GetData() []*UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pb_user_proto protoreflect.FileDescriptor

var file_pkg_pb_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x2c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x28, 0x0a, 0x0e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x49, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3e, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x8c, 0x02, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x49,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_user_proto_rawDescOnce sync.Once
	file_pkg_pb_user_proto_rawDescData = file_pkg_pb_user_proto_rawDesc
)

func file_pkg_pb_user_proto_rawDescGZIP() []byte {
	file_pkg_pb_user_proto_rawDescOnce.Do(func() {
		file_pkg_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_user_proto_rawDescData)
	})
	return file_pkg_pb_user_proto_rawDescData
}

var file_pkg_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_user_proto_goTypes = []interface{}{
	(*AddNameRequest)(nil),           // 0: user.AddNameRequest
	(*AddNameResponse)(nil),          // 1: user.AddNameResponse
	(*AddPhoneRequest)(nil),          // 2: user.AddPhoneRequest
	(*AddPhoneResponse)(nil),         // 3: user.AddPhoneResponse
	(*IsAdminRequest)(nil),           // 4: user.IsAdminRequest
	(*IsAdminResponse)(nil),          // 5: user.IsAdminResponse
	(*GetUsersByChatIdRequest)(nil),  // 6: user.GetUsersByChatIdRequest
	(*UserData)(nil),                 // 7: user.UserData
	(*GetUsersByChatIdResponse)(nil), // 8: user.GetUsersByChatIdResponse
}
var file_pkg_pb_user_proto_depIdxs = []int32{
	7, // 0: user.GetUsersByChatIdResponse.data:type_name -> user.UserData
	0, // 1: user.User.AddName:input_type -> user.AddNameRequest
	2, // 2: user.User.AddPhone:input_type -> user.AddPhoneRequest
	4, // 3: user.User.IsAdmin:input_type -> user.IsAdminRequest
	6, // 4: user.User.GetUsersByChatId:input_type -> user.GetUsersByChatIdRequest
	1, // 5: user.User.AddName:output_type -> user.AddNameResponse
	3, // 6: user.User.AddPhone:output_type -> user.AddPhoneResponse
	5, // 7: user.User.IsAdmin:output_type -> user.IsAdminResponse
	8, // 8: user.User.GetUsersByChatId:output_type -> user.GetUsersByChatIdResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_user_proto_init() }
func file_pkg_pb_user_proto_init() {
	if File_pkg_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNameRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNameResponse); i {
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
		file_pkg_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhoneRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhoneResponse); i {
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
		file_pkg_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAdminResponse); i {
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
		file_pkg_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersByChatIdRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_pkg_pb_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersByChatIdResponse); i {
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
			RawDescriptor: file_pkg_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_user_proto_goTypes,
		DependencyIndexes: file_pkg_pb_user_proto_depIdxs,
		MessageInfos:      file_pkg_pb_user_proto_msgTypes,
	}.Build()
	File_pkg_pb_user_proto = out.File
	file_pkg_pb_user_proto_rawDesc = nil
	file_pkg_pb_user_proto_goTypes = nil
	file_pkg_pb_user_proto_depIdxs = nil
}
