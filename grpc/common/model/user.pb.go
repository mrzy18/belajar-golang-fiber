// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: model/user.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

// Enum value maps for UserGender.
var (
	UserGender_name = map[int32]string{
		0: "UNDEFINED",
		1: "MALE",
		2: "FEMALE",
	}
	UserGender_value = map[string]int32{
		"UNDEFINED": 0,
		"MALE":      1,
		"FEMALE":    2,
	}
)

func (x UserGender) Enum() *UserGender {
	p := new(UserGender)
	*p = x
	return p
}

func (x UserGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGender) Descriptor() protoreflect.EnumDescriptor {
	return file_model_user_proto_enumTypes[0].Descriptor()
}

func (UserGender) Type() protoreflect.EnumType {
	return &file_model_user_proto_enumTypes[0]
}

func (x UserGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGender.Descriptor instead.
func (UserGender) EnumDescriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Gender        UserGender             `protobuf:"varint,4,opt,name=gender,proto3,enum=model.UserGender" json:"gender,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_model_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetGender() UserGender {
	if x != nil {
		return x.Gender
	}
	return UserGender_UNDEFINED
}

type UserList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*User                `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserList) Reset() {
	*x = UserList{}
	mi := &file_model_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_model_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserList) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

var File_model_user_proto protoreflect.FileDescriptor

const file_model_user_proto_rawDesc = "" +
	"\n" +
	"\x10model/user.proto\x12\x05model\x1a\x1bgoogle/protobuf/empty.proto\"q\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12)\n" +
	"\x06gender\x18\x04 \x01(\x0e2\x11.model.UserGenderR\x06gender\"+\n" +
	"\bUserList\x12\x1f\n" +
	"\x04List\x18\x01 \x03(\v2\v.model.UserR\x04List*1\n" +
	"\n" +
	"UserGender\x12\r\n" +
	"\tUNDEFINED\x10\x00\x12\b\n" +
	"\x04MALE\x10\x01\x12\n" +
	"\n" +
	"\x06FEMALE\x10\x022m\n" +
	"\x05Users\x121\n" +
	"\bRegister\x12\v.model.User\x1a\x16.google.protobuf.Empty\"\x00\x121\n" +
	"\x04List\x12\x16.google.protobuf.Empty\x1a\x0f.model.UserList\"\x00B\tZ\a./modelb\x06proto3"

var (
	file_model_user_proto_rawDescOnce sync.Once
	file_model_user_proto_rawDescData []byte
)

func file_model_user_proto_rawDescGZIP() []byte {
	file_model_user_proto_rawDescOnce.Do(func() {
		file_model_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_user_proto_rawDesc), len(file_model_user_proto_rawDesc)))
	})
	return file_model_user_proto_rawDescData
}

var file_model_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_user_proto_goTypes = []any{
	(UserGender)(0),       // 0: model.UserGender
	(*User)(nil),          // 1: model.User
	(*UserList)(nil),      // 2: model.UserList
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_model_user_proto_depIdxs = []int32{
	0, // 0: model.User.gender:type_name -> model.UserGender
	1, // 1: model.UserList.List:type_name -> model.User
	1, // 2: model.Users.Register:input_type -> model.User
	3, // 3: model.Users.List:input_type -> google.protobuf.Empty
	3, // 4: model.Users.Register:output_type -> google.protobuf.Empty
	2, // 5: model.Users.List:output_type -> model.UserList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_user_proto_init() }
func file_model_user_proto_init() {
	if File_model_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_user_proto_rawDesc), len(file_model_user_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_user_proto_goTypes,
		DependencyIndexes: file_model_user_proto_depIdxs,
		EnumInfos:         file_model_user_proto_enumTypes,
		MessageInfos:      file_model_user_proto_msgTypes,
	}.Build()
	File_model_user_proto = out.File
	file_model_user_proto_goTypes = nil
	file_model_user_proto_depIdxs = nil
}
