// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User_EmailType int32

const (
	User_HOME   User_EmailType = 0
	User_WORK   User_EmailType = 1
	User_MOBILE User_EmailType = 2
)

var User_EmailType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
	2: "MOBILE",
}
var User_EmailType_value = map[string]int32{
	"HOME":   0,
	"WORK":   1,
	"MOBILE": 2,
}

func (x User_EmailType) String() string {
	return proto.EnumName(User_EmailType_name, int32(x))
}
func (User_EmailType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_7e392db06f9e127a, []int{0, 0}
}

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Emails               []*User_Email        `protobuf:"bytes,4,rep,name=emails,proto3" json:"emails,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7e392db06f9e127a, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmails() []*User_Email {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type User_Email struct {
	Address              string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Type                 User_EmailType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.User_EmailType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *User_Email) Reset()         { *m = User_Email{} }
func (m *User_Email) String() string { return proto.CompactTextString(m) }
func (*User_Email) ProtoMessage()    {}
func (*User_Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7e392db06f9e127a, []int{0, 0}
}
func (m *User_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_Email.Unmarshal(m, b)
}
func (m *User_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_Email.Marshal(b, m, deterministic)
}
func (dst *User_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_Email.Merge(dst, src)
}
func (m *User_Email) XXX_Size() int {
	return xxx_messageInfo_User_Email.Size(m)
}
func (m *User_Email) XXX_DiscardUnknown() {
	xxx_messageInfo_User_Email.DiscardUnknown(m)
}

var xxx_messageInfo_User_Email proto.InternalMessageInfo

func (m *User_Email) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User_Email) GetType() User_EmailType {
	if m != nil {
		return m.Type
	}
	return User_HOME
}

func init() {
	proto.RegisterType((*User)(nil), "protobuf.User")
	proto.RegisterType((*User_Email)(nil), "protobuf.User.Email")
	proto.RegisterEnum("protobuf.User_EmailType", User_EmailType_name, User_EmailType_value)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_7e392db06f9e127a) }

var fileDescriptor_user_7e392db06f9e127a = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xbf, 0xfd, 0xb1, 0x7e, 0x9b, 0x77, 0x30, 0x4a, 0xf0, 0x10, 0x2a, 0x62, 0xd9, 0xa9,
	0xe0, 0xe8, 0xa0, 0x9e, 0x3c, 0x4e, 0x28, 0x28, 0x3a, 0x0b, 0x61, 0xe2, 0x71, 0x64, 0x26, 0x1b,
	0x85, 0xd6, 0x96, 0x24, 0x3d, 0xec, 0xff, 0xf6, 0x0f, 0x90, 0xa6, 0x8d, 0x5e, 0x04, 0x4f, 0x49,
	0x9e, 0xe7, 0xf3, 0xe4, 0x7d, 0x1f, 0x80, 0x5e, 0x09, 0x99, 0x75, 0xb2, 0xd5, 0x2d, 0x0e, 0xcd,
	0x71, 0xe8, 0x8f, 0xf1, 0xf5, 0xa9, 0x6d, 0x4f, 0xb5, 0x58, 0x5b, 0x61, 0xad, 0xab, 0x46, 0x28,
	0xcd, 0x9a, 0x6e, 0x44, 0x97, 0x9f, 0x2e, 0xf8, 0xaf, 0x4a, 0x48, 0xbc, 0x00, 0xb7, 0xe2, 0xc4,
	0x49, 0x9c, 0xd4, 0xa3, 0x6e, 0xc5, 0xf1, 0x15, 0xc0, 0xb1, 0x92, 0x4a, 0xef, 0x3f, 0x58, 0x23,
	0x88, 0x9b, 0x38, 0x29, 0xa2, 0xc8, 0x28, 0x2f, 0xac, 0x11, 0xf8, 0x12, 0x50, 0xcd, 0xac, 0xeb,
	0x19, 0x37, 0x1c, 0x04, 0x63, 0xae, 0x20, 0x10, 0x0d, 0xab, 0x6a, 0x45, 0xfc, 0xc4, 0x4b, 0xe7,
	0xf9, 0x45, 0x66, 0xe7, 0x67, 0xc3, 0xac, 0xac, 0x18, 0x4c, 0x3a, 0x31, 0xf8, 0x0e, 0xe0, 0x5d,
	0x0a, 0xa6, 0x05, 0xdf, 0x33, 0x4d, 0x66, 0x89, 0x93, 0xce, 0xf3, 0x38, 0x1b, 0x17, 0xff, 0x09,
	0xee, 0xec, 0xe2, 0x14, 0x4d, 0xf4, 0x46, 0x0f, 0xd1, 0xbe, 0xe3, 0x36, 0x1a, 0xfc, 0x1d, 0x9d,
	0xe8, 0x8d, 0x8e, 0x4b, 0x98, 0x99, 0x35, 0x30, 0x81, 0xff, 0x8c, 0x73, 0x29, 0x94, 0x32, 0xed,
	0x11, 0xb5, 0x4f, 0xbc, 0x02, 0x5f, 0x9f, 0xbb, 0xb1, 0xfc, 0x22, 0x27, 0xbf, 0x95, 0xd8, 0x9d,
	0x3b, 0x41, 0x0d, 0xb5, 0xbc, 0x01, 0xf4, 0x2d, 0xe1, 0x10, 0xfc, 0x87, 0x72, 0x5b, 0x44, 0xff,
	0x86, 0xdb, 0x5b, 0x49, 0x9f, 0x22, 0x07, 0x03, 0x04, 0xdb, 0xf2, 0xfe, 0xf1, 0xb9, 0x88, 0xdc,
	0x43, 0x60, 0xfe, 0xba, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xfd, 0xf5, 0x00, 0xb6, 0x01,
	0x00, 0x00,
}