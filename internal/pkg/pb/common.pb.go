// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// app 类型
type App int32

const (
	App_APP_UNKNOWN App = 0
	App_APP_ANDROID App = 1
	App_APP_IOS     App = 2
	App_APP_WIN     App = 3
	App_APP_LINUX   App = 4
)

var App_name = map[int32]string{
	0: "APP_UNKNOWN",
	1: "APP_ANDROID",
	2: "APP_IOS",
	3: "APP_WIN",
	4: "APP_LINUX",
}

var App_value = map[string]int32{
	"APP_UNKNOWN": 0,
	"APP_ANDROID": 1,
	"APP_IOS":     2,
	"APP_WIN":     3,
	"APP_LINUX":   4,
}

func (x App) String() string {
	return proto.EnumName(App_name, int32(x))
}

func (App) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type ReceiverType int32

const (
	ReceiverType_RECEIVER_UNKNOWN ReceiverType = 0
	ReceiverType_RECEIVER_USER    ReceiverType = 1
	ReceiverType_RECEIVER_GROUP   ReceiverType = 2
)

var ReceiverType_name = map[int32]string{
	0: "RECEIVER_UNKNOWN",
	1: "RECEIVER_USER",
	2: "RECEIVER_GROUP",
}

var ReceiverType_value = map[string]int32{
	"RECEIVER_UNKNOWN": 0,
	"RECEIVER_USER":    1,
	"RECEIVER_GROUP":   2,
}

func (x ReceiverType) String() string {
	return proto.EnumName(ReceiverType_name, int32(x))
}

func (ReceiverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Image struct {
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.App", App_name, App_value)
	proto.RegisterEnum("pb.ReceiverType", ReceiverType_name, ReceiverType_value)
	proto.RegisterType((*Text)(nil), "pb.Text")
	proto.RegisterType((*Image)(nil), "pb.Image")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x05, 0x60, 0xd3, 0x46, 0x65, 0x67, 0x77, 0x75, 0x1c, 0x3c, 0xac, 0x9e, 0x64, 0x4f, 0xb2,
	0x07, 0x2f, 0xfe, 0x82, 0xe2, 0x06, 0x09, 0x95, 0x24, 0xa4, 0xad, 0xf5, 0x26, 0xb6, 0x04, 0x11,
	0xac, 0x09, 0x25, 0x4a, 0xfd, 0xf7, 0xd2, 0x42, 0xd5, 0xdb, 0x7b, 0xdf, 0xc0, 0xc0, 0x83, 0x55,
	0xeb, 0xbb, 0xce, 0x7f, 0xdc, 0x84, 0xde, 0x47, 0x4f, 0x49, 0x68, 0xb6, 0x97, 0xc0, 0x4b, 0x37,
	0x44, 0x22, 0xe0, 0xd1, 0x0d, 0x71, 0xc3, 0xae, 0xd8, 0xf5, 0xc2, 0x4e, 0x79, 0x7b, 0x01, 0x87,
	0xb2, 0x7b, 0x79, 0x75, 0x84, 0x90, 0x7e, 0xf6, 0xef, 0x1b, 0x3e, 0xdd, 0xc6, 0xb8, 0x33, 0x90,
	0x66, 0x21, 0xd0, 0x29, 0x2c, 0x33, 0x63, 0x9e, 0x2b, 0x95, 0x2b, 0x5d, 0x2b, 0x3c, 0x98, 0x21,
	0x53, 0x7b, 0xab, 0xe5, 0x1e, 0x19, 0x2d, 0xe1, 0x78, 0x04, 0xa9, 0x0b, 0x4c, 0xe6, 0x52, 0x4b,
	0x85, 0x29, 0xad, 0x61, 0x31, 0x96, 0x07, 0xa9, 0xaa, 0x27, 0xe4, 0xbb, 0x1c, 0x56, 0xd6, 0xb5,
	0xee, 0xed, 0xcb, 0xf5, 0xe5, 0x77, 0x70, 0x74, 0x0e, 0x68, 0xc5, 0x9d, 0x90, 0x8f, 0xc2, 0xfe,
	0xfb, 0x7f, 0x06, 0xeb, 0x3f, 0x2d, 0x84, 0x45, 0x46, 0x04, 0x27, 0xbf, 0x74, 0x6f, 0x75, 0x65,
	0x30, 0x69, 0x8e, 0xa6, 0x81, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x5e, 0x4a, 0x54,
	0xf0, 0x00, 0x00, 0x00,
}
