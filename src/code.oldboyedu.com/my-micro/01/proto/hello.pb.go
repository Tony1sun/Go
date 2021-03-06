// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

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

// 结构体
type InfoRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

func (m *InfoRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type InfoReponse struct {
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoReponse) Reset()         { *m = InfoReponse{} }
func (m *InfoReponse) String() string { return proto.CompactTextString(m) }
func (*InfoReponse) ProtoMessage()    {}
func (*InfoReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *InfoReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoReponse.Unmarshal(m, b)
}
func (m *InfoReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoReponse.Marshal(b, m, deterministic)
}
func (m *InfoReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoReponse.Merge(m, src)
}
func (m *InfoReponse) XXX_Size() int {
	return xxx_messageInfo_InfoReponse.Size(m)
}
func (m *InfoReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoReponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoReponse proto.InternalMessageInfo

func (m *InfoReponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "InfoRequest")
	proto.RegisterType((*InfoReponse)(nil), "InfoReponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe4, 0xe2, 0xf6, 0xcc, 0x4b, 0xcb, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0xe4, 0x61, 0x4a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x98, 0xc0, 0xaa, 0x40,
	0x4c, 0x23, 0x5d, 0x2e, 0x56, 0x0f, 0x90, 0xd1, 0x42, 0x2a, 0x5c, 0x2c, 0x20, 0x95, 0x42, 0x3c,
	0x7a, 0x48, 0x66, 0x4b, 0xc1, 0x78, 0x60, 0xed, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x17, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x50, 0x19, 0x92, 0x90, 0x00, 0x00, 0x00,
}
