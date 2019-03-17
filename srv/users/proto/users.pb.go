// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users.proto

package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	// @inject_tag: bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	// @inject_tag: bson:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"username"`
	// @inject_tag: bson:"email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" bson:"email"`
	// @inject_tag: bson:"password"
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	// @inject_tag: bson:"-"
	Token                string   `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty" bson:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fb03113a765b0e47, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Empty Request
type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fb03113a765b0e47, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	Token                *Token   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fb03113a765b0e47, []int{2}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *UserResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fb03113a765b0e47, []int{3}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_fb03113a765b0e47, []int{4}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*Request)(nil), "users.Request")
	proto.RegisterType((*UserResponse)(nil), "users.UserResponse")
	proto.RegisterType((*Token)(nil), "users.Token")
	proto.RegisterType((*Error)(nil), "users.Error")
}

func init() { proto.RegisterFile("proto/users.proto", fileDescriptor_users_fb03113a765b0e47) }

var fileDescriptor_users_fb03113a765b0e47 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x6d, 0xfe, 0xfe, 0xda, 0x69, 0x7f, 0x05, 0x47, 0x0f, 0xa1, 0x17, 0x6b, 0xf4, 0xa0, 0xa0,
	0x2d, 0xd4, 0xb3, 0x87, 0x22, 0xd2, 0x7b, 0xb0, 0x82, 0xc7, 0xd8, 0x0c, 0x18, 0x4c, 0xb3, 0x71,
	0x77, 0xab, 0xe0, 0x67, 0xf1, 0xeb, 0xf9, 0x3d, 0x64, 0x67, 0x13, 0x4d, 0x0b, 0x42, 0x6f, 0xfb,
	0xde, 0xbc, 0xd9, 0x79, 0xf3, 0x76, 0xe1, 0xa0, 0x92, 0x42, 0x8b, 0xe9, 0x46, 0x91, 0x54, 0x13,
	0x3e, 0x63, 0xc0, 0x20, 0xfe, 0x00, 0x7f, 0xa9, 0x48, 0xe2, 0x10, 0xdc, 0x3c, 0x8b, 0x9c, 0xb1,
	0x73, 0xde, 0x4b, 0xdc, 0x3c, 0xc3, 0x11, 0x74, 0x8d, 0xa0, 0x4c, 0xd7, 0x14, 0xb9, 0xcc, 0xfe,
	0x60, 0x3c, 0x82, 0x80, 0xd6, 0x69, 0x5e, 0x44, 0x1e, 0x17, 0x2c, 0x30, 0x1d, 0x55, 0xaa, 0xd4,
	0xbb, 0x90, 0x59, 0x14, 0xd8, 0x8e, 0x06, 0x9b, 0x0e, 0x2d, 0x5e, 0xa8, 0x8c, 0x42, 0xdb, 0xc1,
	0x20, 0xee, 0xc1, 0xbf, 0x84, 0x5e, 0x37, 0xa4, 0x74, 0xfc, 0xe9, 0xc0, 0xc0, 0xf8, 0x48, 0x48,
	0x55, 0xa2, 0x54, 0x84, 0xc7, 0xe0, 0x9b, 0x79, 0xec, 0xa8, 0x3f, 0xeb, 0x4f, 0xac, 0x75, 0x96,
	0x70, 0x01, 0x4f, 0xc0, 0x6e, 0x10, 0xb9, 0x63, 0x6f, 0x57, 0x61, 0x2b, 0x78, 0x06, 0x21, 0x49,
	0x29, 0xa4, 0x8a, 0x3c, 0xd6, 0x0c, 0x6a, 0xcd, 0x9d, 0x21, 0x93, 0xba, 0x86, 0x71, 0xe3, 0xcd,
	0xe7, 0x51, 0x8d, 0xe8, 0xde, 0x70, 0x8d, 0xd3, 0x47, 0x08, 0x18, 0xff, 0x2e, 0xe2, 0xb4, 0x16,
	0x31, 0xec, 0x5b, 0x5a, 0xe4, 0x19, 0x27, 0xd5, 0x4d, 0x2c, 0xd8, 0x6f, 0x7c, 0x7c, 0x03, 0x01,
	0x13, 0x88, 0xe0, 0xaf, 0x44, 0x46, 0x7c, 0x73, 0x90, 0xf0, 0x19, 0xc7, 0xd0, 0xcf, 0x48, 0xad,
	0x64, 0x5e, 0xe9, 0x5c, 0x94, 0xf5, 0x43, 0xb4, 0xa9, 0xd9, 0x97, 0x03, 0xc1, 0x92, 0xb7, 0xbd,
	0x84, 0xf0, 0x56, 0x52, 0xaa, 0x09, 0xdb, 0x59, 0x8c, 0x0e, 0xdb, 0xc1, 0xd4, 0xe9, 0xc6, 0x1d,
	0xbc, 0x00, 0x6f, 0x41, 0x7a, 0x2f, 0xe9, 0x14, 0xc2, 0x05, 0xe9, 0x79, 0x51, 0xe0, 0xb0, 0x16,
	0xd4, 0xaf, 0xf6, 0x57, 0xc3, 0x29, 0xf8, 0xf3, 0x8d, 0x7e, 0xde, 0xbe, 0x7c, 0x2b, 0xd7, 0xb8,
	0x83, 0x57, 0xf0, 0xff, 0xc1, 0xc4, 0x94, 0x6a, 0xb2, 0xd1, 0x6e, 0x09, 0x76, 0xe5, 0x4f, 0x21,
	0xff, 0xda, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x24, 0x3f, 0x2a, 0xca, 0x02, 0x00,
	0x00,
}
