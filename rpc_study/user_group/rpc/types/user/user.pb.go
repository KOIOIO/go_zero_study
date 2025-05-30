// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.9.0
// source: user.proto

package user

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

type UserInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfoResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type UserCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPwd       string                 `protobuf:"bytes,2,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserCreateRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserCreateRequest) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Err           string                 `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserCreateResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x04user\"*\n" +
	"\x0fUserInfoRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\"H\n" +
	"\x10UserInfoResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\"K\n" +
	"\x11UserCreateRequest\x12\x1b\n" +
	"\tuser_name\x18\x01 \x01(\tR\buserName\x12\x19\n" +
	"\buser_pwd\x18\x02 \x01(\tR\auserPwd\"&\n" +
	"\x12UserCreateResponse\x12\x10\n" +
	"\x03err\x18\x01 \x01(\tR\x03err2E\n" +
	"\bUserInfo\x129\n" +
	"\bUserInfo\x12\x15.user.UserInfoRequest\x1a\x16.user.UserInfoResponse2M\n" +
	"\n" +
	"UserAction\x12?\n" +
	"\n" +
	"UserCreate\x12\x17.user.UserCreateRequest\x1a\x18.user.UserCreateResponseB\bZ\x06./userb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []any{
	(*UserInfoRequest)(nil),    // 0: user.UserInfoRequest
	(*UserInfoResponse)(nil),   // 1: user.UserInfoResponse
	(*UserCreateRequest)(nil),  // 2: user.UserCreateRequest
	(*UserCreateResponse)(nil), // 3: user.UserCreateResponse
}
var file_user_proto_depIdxs = []int32{
	0, // 0: user.UserInfo.UserInfo:input_type -> user.UserInfoRequest
	2, // 1: user.UserAction.UserCreate:input_type -> user.UserCreateRequest
	1, // 2: user.UserInfo.UserInfo:output_type -> user.UserInfoResponse
	3, // 3: user.UserAction.UserCreate:output_type -> user.UserCreateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
