// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subpkg/subproto.proto

package subpkg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SubObject struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubObject) Reset()         { *m = SubObject{} }
func (m *SubObject) String() string { return proto.CompactTextString(m) }
func (*SubObject) ProtoMessage()    {}
func (*SubObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_subproto_094c5f22e1aecb1e, []int{0}
}
func (m *SubObject) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_SubObject.Unmarshal(m, b)
}
func (m *SubObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubObject.Marshal(b, m, deterministic)
}
func (dst *SubObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubObject.Merge(dst, src)
}
func (m *SubObject) XXX_Size() int {
	return xxx_messageInfo_SubObject.Size(m)
}
func (m *SubObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SubObject.DiscardUnknown(m)
}

var xxx_messageInfo_SubObject proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SubObject)(nil), "subpkg.SubObject")
}

func init() { proto.RegisterFile("subpkg/subproto.proto", fileDescriptor_subproto_094c5f22e1aecb1e) }

var fileDescriptor_subproto_094c5f22e1aecb1e = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0x4d, 0x2a,
	0xc8, 0x4e, 0xd7, 0x07, 0x51, 0x45, 0xf9, 0x25, 0xf9, 0x7a, 0x60, 0x52, 0x88, 0x0d, 0x22, 0x2c,
	0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9e, 0x9f, 0x9e,
	0xaf, 0x0f, 0x96, 0x4e, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x4d, 0x89, 0x9b,
	0x8b, 0x33, 0xb8, 0x34, 0xc9, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0x04, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0x38, 0xf3, 0x28, 0x5b, 0x00, 0x00, 0x00,
}
